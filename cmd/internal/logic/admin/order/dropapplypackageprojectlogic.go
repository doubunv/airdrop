package order

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropApplyPackageProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropApplyPackageProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropApplyPackageProjectLogic {
	return &DropApplyPackageProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropApplyPackageProjectLogic) DropApplyPackageProject(req *types.DropApplyPackageProjectReq) (resp *types.DropApplyPackageProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
