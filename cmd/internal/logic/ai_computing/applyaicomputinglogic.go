package ai_computing

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
