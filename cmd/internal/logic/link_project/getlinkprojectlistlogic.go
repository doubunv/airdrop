package link_project

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLinkProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLinkProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLinkProjectListLogic {
	return &GetLinkProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLinkProjectListLogic) GetLinkProjectList(req *types.GetLinkProjectListReq) (resp *types.GetLinkProjectListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
