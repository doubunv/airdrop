package event

import (
	"air-drop/cmd/internal/data/model"
	"air-drop/cmd/internal/svc"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	constx "air-drop/cmd/internal/const"
	"github.com/zeromicro/go-zero/core/logx"
)

type EventLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	events     map[string]func(chainTx *model.ChainTx) error
	NowChainTx *model.ChainTx
}

func NewEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EventLogic {
	logic := &EventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
	eventMap := map[string]func(res *model.ChainTx) error{
		"BuyPacketProject":      logic.BuyPacketProject,
		"BuyLinkProject":        logic.BuyLinkProject,
		"BuyAiComputingProject": logic.BuyAiComputingProject,
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
		if err := recover(); err != nil {
			fmt.Println(err)
			logc.Infof(context.Background(), "handler panic: %v", err)
		}
	}()

	fn, ok := l.events[chainTx.EventName]
	if !ok {
		logx.Errorf("event name: %s is no support", chainTx.EventName)
		err = l.svcCtx.ChainModel.UpdateChainTxExecute(chainTx.ID)
		return chainTx.EventName, err
	}
	err = fn(&chainTx)
	if err == nil {
		_ = l.svcCtx.ChainModel.UpdateChainTxError(chainTx.ID, constx.ChainTxSuccess, "success")
	} else {
		_ = l.svcCtx.ChainModel.UpdateChainTxError(chainTx.ID, constx.ChainTxFailed, err.Error())
	}
	_ = l.svcCtx.ChainModel.UpdateChainTxExecute(chainTx.ID)
	return chainTx.EventName, err
}

// TableGC Prevent deleting 30 days of data from the chain_block table being too large
func (l *EventLogic) TableGC() {
	_ = l.svcCtx.ChainModel.TableGC()
}
