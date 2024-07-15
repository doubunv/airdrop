package order

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPackageProjectOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPackageProjectOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPackageProjectOrderListLogic {
	return &GetPackageProjectOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPackageProjectOrderListLogic) GetPackageProjectOrderList(req *types.GetPackageProjectOrderListReq) (resp *types.GetPackageProjectListOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
