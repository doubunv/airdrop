package ai_computing

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAiComputingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAiComputingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAiComputingListLogic {
	return &GetAiComputingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAiComputingListLogic) GetAiComputingList(req *types.GetAiComputingListReq) (resp *types.GetAiComputingListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
