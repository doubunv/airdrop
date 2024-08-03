package ai_computing

import (
	"air-drop/cmd/internal/data/schema"
	"air-drop/pkg/utils"
	"context"
	"errors"
	"time"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuyAiComputingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuyAiComputingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuyAiComputingLogic {
	return &BuyAiComputingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BuyAiComputingLogic) BuyAiComputing(req *types.BuyAiComputingReq) (resp *types.BuyAiComputingResp, err error) {
	resp = &types.BuyAiComputingResp{
		OrderId:   0,
		Amount:    0,
		Sign:      "",
		CreatedAt: 0,
		AmountStr: "",
	}

	aiInfo, err := l.svcCtx.AiComputingModel.FindById(req.Id)
	if err != nil {
		return nil, err
	}

	if aiInfo.ID == 0 {
		return nil, errors.New("not find")
	}

	ord := &schema.AiComputingOrder{
		UserId:        utils.GetTokenUid(l.ctx),
		UAddress:      utils.GetTokenAddress(l.ctx),
		AiComputingId: req.Id,
		Amount:        aiInfo.Price,
		ServiceMonth:  aiInfo.ServiceMonth,
		EndTime:       time.Now().AddDate(0, int(aiInfo.ServiceMonth), 0).Unix(),
	}

	err = l.svcCtx.AiComputingOrderModel.Insert(ord)
	if err != nil {
		return nil, err
	}

	resp.Amount = ord.Amount
	resp.Sign, resp.AmountStr = BuildAiComputingSign(l.svcCtx.Config, ord)
	resp.CreatedAt = ord.CreatedAt
	resp.OrderId = ord.ID
	resp.CreatedAt = ord.CreatedAt

	return
}
