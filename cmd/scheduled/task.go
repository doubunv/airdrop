package scheduled

import (
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/scheduled/event"
	"air-drop/pkg/alert"
	"air-drop/pkg/concurrency"
	"context"
	"fmt"
	"time"

	"github.com/robfig/cron"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/time/rate"
)

const BlockExecutePeriod = 1

func GetEthData(svc *svc.ServiceContext) {

	// deal block
	chainLogic := event.NewLogLogic(context.Background(), svc)
	SyncBlock(chainLogic)

	// deal tx
	syncTx(chainLogic)

	// deal listen event
	eventLogic := event.NewEventLogic(context.Background(), svc)
	tableGC(eventLogic)

	dealEvent(eventLogic)
}

func SyncBlock(chainLogic *event.LogLogic) {
	concurrency.Go(context.Background(), func(ctx context.Context) {
		limit := rate.Every(time.Duration(BlockExecutePeriod) * time.Second)
		limiter := rate.NewLimiter(limit, 1)
		for {
			err := limiter.Wait(ctx)
			if err != nil {
				return
			}
			err = chainLogic.SyncBlock()
			if err != nil {
				alert.SendMsg(fmt.Sprintf("can not run SyncBlock err:%+v ", err))
				logx.Errorf("can not run SyncBlock err:%+v ", err)
			}
		}
	})
}

func syncTx(chainLogic *event.LogLogic) {
	concurrency.Go(context.Background(), func(ctx context.Context) {
		limit := rate.Every(time.Duration(BlockExecutePeriod) * time.Second)
		limiter := rate.NewLimiter(limit, 1)
		for {
			err := limiter.Wait(ctx)
			if err != nil {
				return
			}
			err = chainLogic.SyncTx()
			if err != nil {
				alert.SendMsg(fmt.Sprintf("can not run SyncTx err:%+v ", err))
				logx.Errorf("can not run SyncTx err:%+v ", err)
			}
		}
	})
}

func tableGC(chainLogic *event.EventLogic) {
	chainLogic.TableGC()
	spec := "@every 24h"
	c := cron.New()
	err := c.AddFunc(spec, func() {
		chainLogic.TableGC()
	})
	if err != nil {
		logx.Errorf("spec err: %v", err)
	}
	c.Start()
}

func dealEvent(eventLogic *event.EventLogic) {
	concurrency.Go(context.Background(), func(ctx context.Context) {
		limit := rate.Every(time.Duration(BlockExecutePeriod) * time.Second)
		limiter := rate.NewLimiter(limit, 1)
		for {
			err := limiter.Wait(ctx)
			if err != nil {
				return
			}
			eventName, err := eventLogic.EventHandler()
			if err != nil {
				alert.SendMsg(fmt.Sprintf("%s: can not run SyncBlock err:%+v", eventName, err))
				logx.Errorf("%s: can not run SyncBlock err:%+v ", eventName, err)
			}
		}
	})
}
