package project

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectListLogic {
	return &GetProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectListLogic) GetProjectList(req *types.GetProjectListReq) (resp *types.GetProjectListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
