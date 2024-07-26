package ai_computing

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropAiComputingApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropAiComputingApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropAiComputingApplyListLogic {
	return &DropAiComputingApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropAiComputingApplyListLogic) DropAiComputingApplyList(req *types.DropAiComputingApplyListReq) (resp *types.DropAiComputingApplyListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
