package order

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPackageOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPackageOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPackageOrderListLogic {
	return &GetPackageOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPackageOrderListLogic) GetPackageOrderList(req *types.GetPackageOrderListReq) (resp *types.GetPackageOrderListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
