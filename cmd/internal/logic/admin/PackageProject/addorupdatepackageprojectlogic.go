package PackageProject

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrUpdatePackageProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrUpdatePackageProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrUpdatePackageProjectLogic {
	return &AddOrUpdatePackageProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrUpdatePackageProjectLogic) AddOrUpdatePackageProject(req *types.AddOrUpdatePackageProjectReq) (resp *types.AddOrUpdatePackageProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
