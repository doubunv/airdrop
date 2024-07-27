package ai_computing

import (
	"air-drop/cmd/internal/data/schema"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropAiComputingApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropAiComputingApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropAiComputingApplyLogic {
	return &DropAiComputingApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropAiComputingApplyLogic) DropAiComputingApply(req *types.DropAiComputingApplyReq) (resp *types.DropAiComputingApplyResp, err error) {
	resp = &types.DropAiComputingApplyResp{}

	rd := &schema.AiComputingOrder{
		ID:             req.Id,
		WithdrawAmount: req.Amount,
	}

	err = l.svcCtx.AiComputingOrderModel.UpdateSchema(rd)
	if err != nil {
		return nil, err
	}

	return
}
