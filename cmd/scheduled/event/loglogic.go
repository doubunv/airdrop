package event

import (
	_const "air-drop/cmd/const"
	"air-drop/cmd/internal/config"
	"air-drop/cmd/internal/data/dto"
	"air-drop/cmd/internal/data/model"
	"air-drop/cmd/internal/svc"
	"air-drop/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeromicro/go-zero/core/logx"
)

type LogLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	client    *ethclient.Client
	chainName string
}

func NewLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogLogic {
	client, err := ethclient.Dial(svcCtx.Config.ChainInfo.Rpc)
	if err != nil {
		logx.Errorf("init eth client err:%v", err)
	}
	return &LogLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		client:    client,
		chainName: svcCtx.Config.ChainInfo.Name,
	}
}

func (f *LogLogic) SyncBlock() error {
	chainName := f.svcCtx.Config.ChainInfo.Name
	chainNode := f.getChain()
	// per deal limit .
	limitBlockNum := chainNode.PerLimit
	chainLatestBlockNum, err := f.GetLatestBlockNumber()
	if err != nil {
		return err
	}
	localLatestBlockNum, err := f.svcCtx.ChainModel.GetLatestBlockNumber(chainNode.Name)
	if err != nil {
		return err
	}
	logx.Debugf("find chain:%s lastblock:%d local lastblock:%d", chainName, chainLatestBlockNum, localLatestBlockNum)
	configStartBlockNum := chainNode.InitBlock
	if localLatestBlockNum == 0 || localLatestBlockNum < configStartBlockNum {
		// from local set init blockNum start
		localLatestBlockNum = configStartBlockNum
	}
	if chainLatestBlockNum <= localLatestBlockNum {
		logx.Debugf("chain:%s chain block:%d is less than local block:%d finish this sync", chainName, chainLatestBlockNum, localLatestBlockNum)
		return nil
	}
	startBlockNum := localLatestBlockNum + 1
	var size = limitBlockNum
	if limitBlockNum > chainLatestBlockNum-localLatestBlockNum {
		size = chainLatestBlockNum - localLatestBlockNum
	}
	//size := min(limitBlockNum, chainLatestBlockNum-localLatestBlockNum)
	// generate need deal blocknumber list
	//logx.Infof("chain:%s generating block from:%d to %d", chainName, startBlockNum, startBlockNum+size)
	blockNumList, err := f.generateBlockNumList(startBlockNum, size)
	if err != nil {
		return err
	}
	if len(blockNumList) == 0 {
		//logx.Infof("chain:%s blockNumList is empty from:%d to %d", chainName, startBlockNum, startBlockNum+size)
		return nil
	}
	// get block basic data
	//logx.Infof("chain:%s getting block from:%d to %d", chainName, startBlockNum, startBlockNum+size)
	chainBlockList, err := f.getBlockList(chainNode.Rpc, blockNumList, chainName)
	if err != nil {
		return err
	}
	// save block basic data
	//logx.Infof("chain:%s saving block from:%d to %d", chainName, startBlockNum, startBlockNum+size)

	return f.svcCtx.ChainModel.CreateChainBlock(chainBlockList, chainName)
}

func (f *LogLogic) getBlockList(node string, blockList []int64, chainName string) ([]model.ChainBlock, error) {
	chainBlockList := make([]model.ChainBlock, 0, len(blockList))
	if len(blockList) == 0 {
		return nil, nil
	}
	for i := range blockList {
		blockNumberBig := big.NewInt(blockList[i])
		blockData, err := f.client.BlockByNumber(context.Background(), blockNumberBig)
		if err != nil {
			logx.Errorf("chain:%s getBlockList blockNumList is empty blockList: %v ", node, blockList)
			return nil, errors.WithStack(err)
		}
		blockData.Transactions()
		// todo ParentHash not equal last data block id ,need rollback
		chainBlockList = append(chainBlockList, model.ChainBlock{
			Chain:       chainName,
			Hash:        blockData.Hash().Hex(),
			Number:      blockList[i],
			IsSafe:      _const.ChainBlockSafe,
			IsSyncTx:    _const.ChainBlockNotSyncTx,
			IsConfirmTx: _const.ChainBlockNotConfirmTx,
			Previous:    blockData.ParentHash().Hex(),
			Size:        blockData.Size(),
			Timestamp:   int64(blockData.Time()),
		})
	}
	return chainBlockList, nil
}

func (f *LogLogic) SyncTx() error {
	// chain running rpc node
	chainNode := f.getChain()
	chainName := f.chainName
	// need listener contract map
	contractMap, err := f.GetListenerContractMap(chainName)
	if err != nil {
		return err
	}

	// per deal limit
	limitBlockNum := 5
	// get status is PENDING block list
	ChainBlockNotSyncTx := _const.ChainBlockNotSyncTx
	configStartBlockNum := chainNode.InitBlock
	chainBlockList, err := f.svcCtx.ChainModel.GetBlockList(chainName, &configStartBlockNum, nil, nil, &ChainBlockNotSyncTx, nil, int(limitBlockNum))
	if err != nil {
		return err
	}
	continueBlockList := f.FilterContinuedBlockNum(chainBlockList)
	if len(continueBlockList) == 0 {
		//logx.Infof("%s SyncTx no continue BlockList ", chainName)
		time.Sleep(1 * time.Second)
		return nil
	}
	from := continueBlockList[0].Number - 8 // Fetch the first 10 more blocks each time to prevent database rollback
	to := continueBlockList[len(continueBlockList)-1].Number
	syncTxList, err := f.getChainTxList(from, to, contractMap)
	logx.Infof("chain: %s sync Tx List: from %d to %d", chainName, from, to)

	if err != nil {
		logx.Errorf("getChainTxList err : blockNumber: %v %v %s", from, to, err)
		return err
	}

	for i := range continueBlockList {
		continueBlockList[i].IsSyncTx = _const.ChainBlockAlreadySyncTx
	}

	// set tx  block time
	blockNumMapTime := make(map[int64]int64)
	for i := range continueBlockList {
		blockNumMapTime[continueBlockList[i].Number] = continueBlockList[i].Timestamp
	}
	for i := range syncTxList {
		syncTxList[i].BlockTime = blockNumMapTime[syncTxList[i].BlockNumber]
	}

	err = f.svcCtx.ChainModel.CreateChainTx(syncTxList)
	if err != nil {
		logx.Errorf("CreateChainTx err  %s", err)
		return err
	}

	err = f.svcCtx.ChainModel.UpdateChainBlock(continueBlockList)
	if err != nil {
		logx.Errorf("CreateChainTx err  %s", err)
		return err
	}

	return err
}

func (f *LogLogic) ConfirmTx() error {
	// chain running rpc node
	chainNode := f.svcCtx.Config.ChainInfo
	chainName := f.chainName
	// need listener contract map
	contractMap, err := f.GetListenerContractMap(chainName)
	if err != nil {
		return err
	}
	// per deal limit
	limitBlockNum := chainNode.PerLimit
	// get status is PENDING block list
	blockIsSafe := _const.ChainBlockSafe
	txNotConfirm := _const.ChainBlockNotConfirmTx
	txAlreadySync := _const.ChainBlockAlreadySyncTx
	configStartBlockNum := chainNode.InitBlock

	chainBlockList, err := f.svcCtx.ChainModel.GetBlockList(chainName, &configStartBlockNum, nil, &blockIsSafe, &txAlreadySync, &txNotConfirm, int(limitBlockNum))
	if err != nil {
		return err
	}
	addTxList := make([]model.ChainTx, 0)
	dbTxMap := make(map[string]bool)
	chainTxMap := make(map[string]bool)

	continueBlockList := f.FilterContinuedBlockNum(chainBlockList)
	if len(continueBlockList) == 0 {
		//logx.Infof("%v ConfirmTx no continue BlockList ", chainName)
		time.Sleep(1 * time.Second)
		return nil
	}
	from := continueBlockList[0].Number
	to := from
	if len(continueBlockList) > 1 {
		to = continueBlockList[len(continueBlockList)-1].Number
	}
	confirmedTxList, err := f.getChainTxList(from, to, contractMap)
	if err != nil {
		logx.Errorf("getChainTxList err : blockNumber: %v %v %s", from, to, err.Error())
		return err
	}
	size := to - from + 1
	confirmedBlockList, err := f.generateBlockNumList(from, size)
	if err != nil {
		return err
	}
	dbTxList, err := f.svcCtx.ChainModel.FindTxListByBlockNums(confirmedBlockList)
	if err != nil {
		return err
	}
	for _, tx := range dbTxList {
		dbTxMap[tx.TxHash] = true
	}
	for _, tx := range confirmedTxList {
		chainTxMap[tx.TxHash] = true
	}
	for i, chainTx := range confirmedTxList {
		if _, ok := dbTxMap[chainTx.TxHash]; !ok {
			chainTx.Status = _const.ChainTxSuccess
			chainTx.ExecuteStatus = _const.ExecuteStatusPending
			addTxList = append(addTxList, confirmedTxList[i])
		}
	}
	for _, dbTx := range dbTxList {
		if _, ok := chainTxMap[dbTx.TxHash]; !ok {
			onChain, err := f.checkTxIsOnChain(f.ctx, chainName, dbTx)
			if err != nil {
				return err
			}
			if !onChain {
				dbTx.ExecuteStatus = _const.ExecuteStatusPending
				dbTx.Status = _const.ChainTxFailed
			}
		} else {
			dbTx.ExecuteStatus = _const.ExecuteStatusPending
			dbTx.Status = _const.ChainTxSuccess
		}
	}
	for _, block := range continueBlockList {
		block.IsConfirmTx = _const.ChainBlockAlreadyConfirmTx
	}
	logx.Infof("confirmed tx  start save and update in  , from %d, to%d", from, to)

	err = f.svcCtx.ChainModel.UpdateChainTx(dbTxList)
	if err != nil {
		return err
	}
	err = f.svcCtx.ChainModel.CreateChainTx(addTxList)
	if err != nil {
		return err
	}
	return f.svcCtx.ChainModel.UpdateChainBlock(continueBlockList)
}

func (f *LogLogic) ConfirmTxOld() error {
	// chain running rpc node
	chainNode := f.svcCtx.Config.ChainInfo
	chainName := f.chainName
	// need listener contract map
	contractMap, err := f.GetListenerContractMap(chainName)
	if err != nil {
		return err
	}
	// per deal limit
	limitBlockNum := chainNode.PerLimit
	// get status is PENDING block list
	blockIsSafe := _const.ChainBlockSafe
	txNotConfirm := _const.ChainBlockNotConfirmTx
	txAlreadySync := _const.ChainBlockAlreadySyncTx
	configStartBlockNum := chainNode.InitBlock

	chainBlockList, err := f.svcCtx.ChainModel.GetBlockList(chainName, &configStartBlockNum, nil, &blockIsSafe, &txAlreadySync, &txNotConfirm, int(limitBlockNum))
	if err != nil {
		return err
	}
	addTxList := make([]model.ChainTx, 0)
	dbTxMap := make(map[string]bool)
	chainTxMap := make(map[string]bool)

	continueBlockList := f.FilterContinuedBlockNum(chainBlockList)
	if len(continueBlockList) == 0 {
		logx.Debugf("%v ConfirmTx no continue BlockList ", chainName)
		return nil
	}
	from := continueBlockList[0].Number
	to := from
	if len(continueBlockList) > 1 {
		to = continueBlockList[len(continueBlockList)-1].Number
	}
	confirmedTxList, err := f.getChainTxList(from, to, contractMap)
	if err != nil {
		logx.Errorf("getChainTxList err : blockNumber: %v %v %s", from, to, err.Error())
		return err
	}
	size := to - from + 1
	confirmedBlockList, err := f.generateBlockNumList(from, size)
	if err != nil {
		return err
	}
	dbTxList, err := f.svcCtx.ChainModel.FindTxListByBlockNums(confirmedBlockList)
	if err != nil {
		return err
	}
	for _, tx := range dbTxList {
		dbTxMap[tx.TxHash] = true
	}
	for _, tx := range confirmedTxList {
		chainTxMap[tx.TxHash] = true
	}
	for i, chainTx := range confirmedTxList {
		if _, ok := dbTxMap[chainTx.TxHash]; !ok {
			chainTx.Status = _const.ChainTxSuccess
			chainTx.ExecuteStatus = _const.ExecuteStatusPending
			addTxList = append(addTxList, confirmedTxList[i])
		}
	}
	for _, dbTx := range dbTxList {
		if _, ok := chainTxMap[dbTx.TxHash]; !ok {
			onChain, err := f.checkTxIsOnChain(f.ctx, chainName, dbTx)
			if err != nil {
				return err
			}
			if !onChain {
				dbTx.ExecuteStatus = _const.ExecuteStatusPending
				dbTx.Status = _const.ChainTxFailed
			}
		} else {
			dbTx.ExecuteStatus = _const.ExecuteStatusPending
			dbTx.Status = _const.ChainTxSuccess
		}
	}
	for _, block := range continueBlockList {
		block.IsConfirmTx = _const.ChainBlockAlreadyConfirmTx
	}
	logx.Infof("confirmed tx  start save and update in  , from %d, to%d", from, to)

	err = f.svcCtx.ChainModel.UpdateChainTx(dbTxList)
	if err != nil {
		return err
	}
	err = f.svcCtx.ChainModel.CreateChainTx(addTxList)
	if err != nil {
		return err
	}
	return f.svcCtx.ChainModel.UpdateChainBlock(continueBlockList)
}

func (f *LogLogic) GetListenerContractMap(chainName string) (res map[string]dto.ListenerContract, err error) {
	res = make(map[string]dto.ListenerContract)
	list, err := f.svcCtx.ChainModel.GetChainListenerContractByName(chainName)
	if err != nil {
		return res, err
	}
	for _, row := range list {
		var events []string
		err = json.Unmarshal([]byte(row.Events), &events)
		if err != nil {
			return res, errors.WithStack(err)
		}
		res[row.ContractAddr] = dto.ListenerContract{
			Chain:        row.Chain,
			Name:         row.Name,
			ContractAddr: row.ContractAddr,
			Events:       events,
			AbiJSON:      row.AbiJSON,
		}
	}
	return res, nil
}

func (f *LogLogic) getChain() config.ChainInfo {
	return f.svcCtx.Config.ChainInfo
}

func (f *LogLogic) getChainTxList(from int64, to int64, contractMap map[string]dto.ListenerContract) ([]model.ChainTx, error) {
	var queryTopics [][]common.Hash
	var eventHasHex []string
	var events []string
	var addresses []common.Address

	chainTxList := make([]model.ChainTx, 0)
	contractAbiMap := make(map[string]abi.ABI)

	for _, contractRow := range contractMap {
		addresses = append(addresses, common.HexToAddress(contractRow.ContractAddr))
		for _, eventRow := range contractRow.Events {
			events = append(events, eventRow)
			eventSigHash := crypto.Keccak256Hash([]byte(eventRow))
			eventHasHex = append(eventHasHex, eventSigHash.Hex())
		}
		contractAbi, err := abi.JSON(strings.NewReader(contractRow.AbiJSON))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		contractAbiMap[strings.ToLower(contractRow.ContractAddr)] = contractAbi
	}

	eventHexNameMap := make(map[string]string, len(events))
	querySubTopics := make([]common.Hash, len(events))
	for k, eventHashHex := range eventHasHex {
		eventRow := events[k]
		eventName := utils.StringX{}.GetBetweenStr(eventRow, "", "(")
		eventHexNameMap[eventHashHex] = eventName
		querySubTopics[k] = common.HexToHash(eventHashHex)
	}
	// query block logs
	queryTopics = append(queryTopics, querySubTopics)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(from),
		ToBlock:   big.NewInt(to),
		Addresses: addresses,
		//Topics:    queryTopics,
	}
	logs, err := f.client.FilterLogs(f.ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(logs) == 0 {
		// in block not find need event data.
		return nil, nil
	}
	// read logs.

	for _, logRow := range logs {
		for _, Topic := range logRow.Topics {
			fmt.Println(Topic.Hex())
		}
		// no need event log exists
		if _, ok := eventHexNameMap[logRow.Topics[0].Hex()]; !ok {
			continue
		}
		var dataLog []interface{}
		// current contract address
		logContractAddress := strings.ToLower(logRow.Address.String())
		// current event name
		eventName := eventHexNameMap[logRow.Topics[0].Hex()]
		// abi
		contractAbi := abi.ABI{}
		if _, ok := contractAbiMap[logContractAddress]; ok {
			contractAbi = contractAbiMap[logContractAddress]
		}

		// if logs not null to continue
		if len(logRow.Data) > 0 {
			// unpack data to []interface
			dataLog, err = contractAbi.Unpack(eventName, logRow.Data)
			if err != nil {
				log.Error("Contract ABI not found EventName, please check ABI JSON")
				return nil, errors.WithStack(err)
			}
		}
		// get current event info
		event, _ := contractAbi.EventByID(logRow.Topics[0])
		allLog := make(map[int]interface{}, 0)
		insertNum := 0
		// keep need topic data , if contract function input parameters is indexed, is in the Topic
		for i, topic := range logRow.Topics {
			if i == 0 {
				continue
			}
			allLog[insertNum] = topic.Hex()
			insertNum++
		}

		// contract function input parameters, not is indexed, is in the Data
		for k := range dataLog {
			allLog[insertNum] = dataLog[k]
			insertNum++
		}

		// converts to a golang type based on the contract parameter type
		inputParams := make(map[string]interface{}, 0)
		for i, inputParam := range allLog {
			eventType := event.Inputs[i].Type.String()
			inputParamStr := utils.StringX{}.InterfaceToStr(inputParam)
			if strings.Contains(inputParamStr, "0x") {
				if eventType == "address" && len(inputParamStr) > 44 {
					inputParams[event.Inputs[i].Name] = utils.ChainTx{}.HexToString(inputParam)
				} else if strings.Contains(eventType, "int") {
					inputParams[event.Inputs[i].Name] = utils.ChainTx{}.Hex2Dec(inputParam)
				} else if strings.Contains(eventType, "bytes") {
					inputParams[event.Inputs[i].Name] = utils.ChainTx{}.HexToString(inputParam)
				} else {
					inputParams[event.Inputs[i].Name] = inputParam
				}
			} else if event.Inputs[i].Name == "cost" {
				inputParams[event.Inputs[i].Name] = utils.ChainTx{}.ToDecimal(inputParam, _const.Multiplier)
			} else if event.Inputs[i].Name == "odd" {
				inputParams[event.Inputs[i].Name] = utils.ChainTx{}.ToDecimal(inputParam, _const.Multiplier)
			} else {
				inputParams[event.Inputs[i].Name] = inputParam
			}
		}

		jsonStr, err := json.Marshal(inputParams)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		chainTxList = append(chainTxList, model.ChainTx{
			Chain:         f.chainName,
			ContractAddr:  logContractAddress,
			FromAddr:      "",
			BlockNumber:   int64(logRow.BlockNumber),
			BlockHash:     logRow.BlockHash.String(),
			TxHash:        logRow.TxHash.String(),
			Status:        _const.ChainTxPendingCheck,
			LogIndex:      int32(logRow.Index),
			EventHash:     logRow.Topics[0].Hex(),
			EventName:     eventName,
			Data:          string(jsonStr),
			ExecuteStatus: _const.ExecuteStatusPending,
		})
	}
	return chainTxList, nil
}

func (f *LogLogic) checkTxIsOnChain(ctx context.Context, chainName string, tx model.ChainTx) (bool, error) {

	transaction, err := f.client.TransactionReceipt(ctx, common.HexToHash(tx.TxHash))
	if transaction != nil {
		return true, nil
	}
	if err.Error() == "not found" {
		return false, nil
	}
	return false, err
}

func (f *LogLogic) FilterContinuedBlockNum(blockList []model.ChainBlock) []model.ChainBlock {
	if len(blockList) <= 1 {
		return blockList
	}
	flag := blockList[0].Number
	continueBlockList := make([]model.ChainBlock, 0, len(blockList))
	continueBlockList = append(continueBlockList, blockList[0])
	for _, block := range blockList[1:] {
		if block.Number == flag+1 {
			continueBlockList = append(continueBlockList, block)
			flag += 1
		} else {
			return continueBlockList
		}
	}
	return continueBlockList
}

func (f *LogLogic) generateBlockNumList(localBlockNum, size int64) ([]int64, error) {
	list := make([]int64, 0)
	for i := 0; i < int(size); i++ {
		list = append(list, localBlockNum+int64(i))
	}
	return list, nil
}

func (f *LogLogic) GetLatestBlockNumber() (int64, error) {
	blockHeader, err := f.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return blockHeader.Number.Int64(), nil
}
