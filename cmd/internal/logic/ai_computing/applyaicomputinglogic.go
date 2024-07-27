package ai_computing

import (
	"air-drop/cmd/internal/data/schema"
	"air-drop/pkg/utils"
	"context"
	"errors"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyAiComputingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyAiComputingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyAiComputingLogic {
	return &ApplyAiComputingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyAiComputingLogic) ApplyAiComputing(req *types.ApplyAiComputingReq) (resp *types.ApplyAiComputingResp, err error) {
	resp = &types.ApplyAiComputingResp{}

	orderInfo, err := l.svcCtx.AiComputingOrderModel.FindById(req.Id)
	if err != nil {
		return nil, err
	}
	if orderInfo.ID == 0 {
		return nil, errors.New("not find order")
	}

	if orderInfo.WithdrawAmount <= 0 {
		return nil, errors.New("can withdraw amount is zero")
	}

	reD := &schema.AiApply{
		AiComputingOrderId: req.Id,
		UserId:             utils.GetTokenUid(l.ctx),
		UAddress:           utils.GetTokenAddress(l.ctx),
		Amount:             orderInfo.WithdrawAmount,
	}

	or := &schema.AiComputingOrder{
		WithdrawAmount: orderInfo.WithdrawAmount + orderInfo.SendAmount,
		SendAmount:     0,
	}
	err = l.svcCtx.AiComputingOrderModel.UpdateSchema(or)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.AiApplyModel.Insert(reD)
	if err != nil {
		return nil, err
	}
	return
}
