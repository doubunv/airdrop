package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

const dbBatchSize = 1000

type ChainBlock struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Chain       string    `gorm:"column:chain;not null" json:"chain"`
	Hash        string    `gorm:"column:hash;not null" json:"hash"`
	Number      int64     `gorm:"column:number" json:"number"`
	IsSafe      bool      `gorm:"column:is_safe;not null" json:"is_safe"`
	IsConfirmTx bool      `gorm:"column:is_confirm_tx;not null" json:"is_confirm_tx"`
	IsSyncTx    bool      `gorm:"column:is_sync_tx;not null" json:"is_sync_tx"`
	Previous    string    `gorm:"column:previous;not null" json:"previous"`
	Size        uint64    `gorm:"column:size;not null" json:"size"`
	Timestamp   int64     `gorm:"column:timestamp;not null" json:"timestamp"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ChainBlock's table name
func (*ChainBlock) TableName() string {
	return "chain_block"
}

type ChainListenerContract struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Chain        string    `gorm:"column:chain;not null" json:"chain"`
	Name         string    `gorm:"column:name;not null" json:"name"`
	ContractAddr string    `gorm:"column:contract_addr;not null" json:"contract_addr"`
	Events       string    `gorm:"column:events" json:"events"`
	AbiJSON      string    `gorm:"column:abi_json" json:"abi_json"`
	IsEnable     bool      `gorm:"column:is_enable;not null" json:"is_enable"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ChainListenerContract's table name
func (*ChainListenerContract) TableName() string {
	return "chain_listener_contract"
}

// ChainTx mapped from table <chain_tx>
type ChainTx struct {
	ID               int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Chain            string    `gorm:"column:chain;not null" json:"chain"`
	ContractAddr     string    `gorm:"column:contract_addr;not null" json:"contract_addr"`
	FromAddr         string    `gorm:"column:from_addr" json:"from_addr"`
	BlockNumber      int64     `gorm:"column:block_number" json:"block_number"`
	BlockHash        string    `gorm:"column:block_hash" json:"block_hash"`
	TxHash           string    `gorm:"column:tx_hash" json:"tx_hash"`
	LogIndex         int32     `gorm:"column:log_index" json:"log_index"`
	BlockTime        int64     `gorm:"column:block_time" json:"block_time"`
	Status           string    `gorm:"column:status" json:"status"`
	EventHash        string    `gorm:"column:event_hash" json:"event_hash"`
	EventName        string    `gorm:"column:event_name" json:"event_name"`
	Data             string    `gorm:"column:data" json:"data"`
	ExecuteStatus    int64     `gorm:"column:execute_status" json:"execute_status"`
	RPCNodeIP        string    `gorm:"column:rpc_node_ip" json:"rpc_node_ip"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
	SignatureVersion string    `gorm:"column:signature_version;not null" json:"signature_version"`
	Signature        string    `gorm:"column:signature;not null" json:"signature"`
	FailReason       string    `gorm:"column:fail_reason;not null" json:"fail_reason"`
}

// TableName ChainTx's table name
func (*ChainTx) TableName() string {
	return "chain_tx"
}

type ChainModel struct {
	db *gorm.DB
}

func NewChainModel(db *gorm.DB) *ChainModel {
	return &ChainModel{db: db}
}

func (chainModel *ChainModel) GetChainListenerContractByName(name string) (res []ChainListenerContract, err error) {
	err = chainModel.db.Model(&ChainListenerContract{}).Where("chain = ? and is_enable = 1", name).Find(&res).Error
	return
}

func (chainModel *ChainModel) GetBlockList(chainName string, startBlockNum *int64, blockNums []int64, isSafe *bool, isSyncTx *bool, isConfirmTx *bool, limit int) (res []ChainBlock, err error) {
	res = make([]ChainBlock, 0)
	q := chainModel.db.Model(&ChainBlock{}).Where("chain = ? ", chainName)
	if blockNums != nil {
		q = q.Where("member in (?)", blockNums)
	}
	if startBlockNum != nil {
		q = q.Where("number > ?", startBlockNum)
	}
	if isSafe != nil {
		q = q.Where("is_safe = ?", isSafe)
	}
	if isSyncTx != nil {
		q = q.Where("is_sync_tx = ?", isSyncTx)
	}
	if isConfirmTx != nil {
		q = q.Where("is_confirm_tx = ?", isConfirmTx)
	}
	if limit != 0 {
		q = q.Limit(limit)
	}
	err = q.Order("number").Find(&res).Error
	return
}

func (chainModel *ChainModel) GetLatestBlockNumber(chainName string) (res int64, err error) {
	// err = chainModel.db.Model(&ChainBlock{}).Select("number").Where("chain = ? ", chainName).Order("number desc").First(&res).Error
	err = chainModel.db.Model(&ChainBlock{}).Select("number").Order("number desc").First(&res).Error // now only one chainName
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}
	return
}

func (chainModel *ChainModel) CreateChainBlock(blockList []ChainBlock, chainName string) error {
	err := chainModel.db.Model(&ChainBlock{}).CreateInBatches(blockList, 1000).Error
	return err
}

func (chainModel *ChainModel) FindTxListByBlockNums(blockNums []int64) (res []ChainTx, err error) {
	err = chainModel.db.Model(&ChainTx{}).Where("block_number in (?)", blockNums).Find(&res).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func (chainModel *ChainModel) UpdateChainBlock(pos []ChainBlock) error {
	for _, block := range pos {
		err := chainModel.db.Model(&ChainBlock{}).Where("id = ?", block.ID).Updates(block).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (chainModel *ChainModel) CreateChainTx(txList []ChainTx) error {

	if len(txList) == 0 {
		return nil
	}
	m := chainModel.db.Model(&ChainTx{})
	hashList := make([]string, 0, len(txList))
	checkRepeatedTxMap := make(map[string]bool)
	for _, tx := range txList {
		hashList = append(hashList, tx.TxHash)
	}
	// To prevent insertion failure caused by txid conflicts
	var repeatedTxList []string
	err := m.Where("tx_hash in (?)", hashList).Select("tx_hash").Find(&repeatedTxList).Error
	if err != nil {
		return err
	}
	for _, tx := range repeatedTxList {
		checkRepeatedTxMap[tx] = true
	}

	data := make([]ChainTx, 0, len(txList))
	for _, tx := range txList {
		if _, ok := checkRepeatedTxMap[tx.TxHash]; !ok {
			data = append(data, tx)
		}
	}

	return chainModel.db.CreateInBatches(data, dbBatchSize).Error
}

func (chainModel *ChainModel) UpdateChainTx(pos []ChainTx) error {
	m := chainModel.db.Model(&ChainTx{})
	for _, po := range pos {
		err := m.Where("id = ?", po.ID).Updates(po).Error
		if err != nil {
			return err
		}
	}
	return nil
}
func (chainModel *ChainModel) UpdateChainTxExecute(id int64) error {
	return chainModel.db.Model(&ChainTx{}).Where("id = ?", id).Update("execute_status", 1).Error
}

func (chainModel *ChainModel) UpdateChainTxError(id int64, kind string, reason string) error {
	data := ChainTx{
		Status:     kind,
		FailReason: reason,
	}
	return chainModel.db.Model(&ChainTx{}).Where("id = ?", id).Updates(data).Error
}

func (chainModel *ChainModel) GetLastChainTxNotExecute() (res ChainTx, err error) {
	err = chainModel.db.Model(&ChainTx{}).Where("execute_status = 0").Order("id").Limit(1).Find(&res).Error
	return
}

func (chainModel *ChainModel) TableGC() error {
	return chainModel.db.Delete(&ChainBlock{}, "timestamp < ? and is_sync_tx = 1", time.Now().Add(-3*24*time.Hour).Unix()).Error
}
