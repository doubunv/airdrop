package package_project

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPackageProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPackageProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPackageProjectListLogic {
	return &GetPackageProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPackageProjectListLogic) GetPackageProjectList(req *types.GetPackageProjectListReq) (resp *types.GetPackageProjectListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
