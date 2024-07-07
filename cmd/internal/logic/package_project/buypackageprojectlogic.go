package package_project

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuyPackageProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuyPackageProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuyPackageProjectLogic {
	return &BuyPackageProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BuyPackageProjectLogic) BuyPackageProject(req *types.BuyPackageProjectReq) (resp *types.BuyPackageProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
