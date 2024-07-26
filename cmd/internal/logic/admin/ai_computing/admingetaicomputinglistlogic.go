package ai_computing

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetAiComputingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminGetAiComputingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetAiComputingListLogic {
	return &AdminGetAiComputingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetAiComputingListLogic) AdminGetAiComputingList(req *types.AdminGetAiComputingListReq) (resp *types.AdminGetAiComputingListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
