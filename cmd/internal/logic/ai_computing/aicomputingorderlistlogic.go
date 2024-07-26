package ai_computing

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AiComputingOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAiComputingOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AiComputingOrderListLogic {
	return &AiComputingOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AiComputingOrderListLogic) AiComputingOrderList(req *types.AiComputingOrderListReq) (resp *types.AiComputingOrderListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
