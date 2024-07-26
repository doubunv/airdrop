package ai_computing

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
