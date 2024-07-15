package order

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropApplyPackageListProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropApplyPackageListProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropApplyPackageListProjectLogic {
	return &DropApplyPackageListProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropApplyPackageListProjectLogic) DropApplyPackageListProject(req *types.DropApplyPackageListProjectReq) (resp *types.DropApplyPackageListProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
