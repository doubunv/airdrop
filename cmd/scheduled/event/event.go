package event

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	constx "xyz/cmd/internal/const"
	"xyz/cmd/internal/data/model"
	"xyz/cmd/internal/svc"
)

type EventLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	events     map[string]func(blockTime int64, reqJson string) error
	NowChainTx *model.ChainTx
}

func NewEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EventLogic {
	logic := &EventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
	eventMap := map[string]func(blockTime int64, reqJson string) error{
		"Staking":            logic.StakingEvent,
		"UnStaking":          logic.UnStakingEvent,
		"NewNode":            logic.NewNode,
		"AgentCommission":    logic.AgentCommissionEvent,
		"StakingCommission":  logic.StakingCommissionEvent,
		"Settle":             logic.SettleEvent,
		"SellNft":            logic.SellNftEvent,
		"CancelSell":         logic.CancelSellEvent,
		"StakingNft":         logic.StakingNftEvent,
		"UnStakingNft":       logic.UnStakingNftEvent,
		"WithdrawMany":       logic.WithdrawManyEvent,
		"CancelWithdrawMany": logic.CancelWithdrawManyEvent,
		"BuyScoreFromMany":   logic.BuyScoreFromManyEvent,
	}
	logic.events = eventMap
	return logic
}

func (l *EventLogic) EventHandler() (string, error) {
	chainTx, err := l.svcCtx.ChainModel.GetLastChainTxNotExecute()
	if err != nil {
		logx.Errorf("get last chain tx not execute error: %v", err)
		return chainTx.EventName, err
	}
	if chainTx.ID == 0 {
		time.Sleep(5 * time.Second)
		return chainTx.EventName, nil
	}

	// 保存到上下文中
	l.NowChainTx = &chainTx
	defer func() {
		l.NowChainTx = nil
	}()

	fn, ok := l.events[chainTx.EventName]
	if !ok {
		logx.Errorf("event name: %s is no support", chainTx.EventName)
		err = l.svcCtx.ChainModel.UpdateChainTxExecute(chainTx.ID)
		return chainTx.EventName, err
	}
	err = fn(chainTx.BlockTime, chainTx.Data)
	if err == nil {
		_ = l.svcCtx.ChainModel.UpdateChainTxError(chainTx.ID, constx.ChainTxSuccess)
	} else {
		_ = l.svcCtx.ChainModel.UpdateChainTxError(chainTx.ID, constx.ChainTxFailed)
	}
	_ = l.svcCtx.ChainModel.UpdateChainTxExecute(chainTx.ID)
	return chainTx.EventName, err
}

// TableGC Prevent deleting 30 days of data from the chain_block table being too large
func (l *EventLogic) TableGC() {
	_ = l.svcCtx.ChainModel.TableGC()
}
