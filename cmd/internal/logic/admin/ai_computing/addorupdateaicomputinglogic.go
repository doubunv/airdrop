package ai_computing

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrUpdateAiComputingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrUpdateAiComputingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrUpdateAiComputingLogic {
	return &AddOrUpdateAiComputingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrUpdateAiComputingLogic) AddOrUpdateAiComputing(req *types.AddOrUpdateAiComputingReq) (resp *types.AddOrUpdateAiComputingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
