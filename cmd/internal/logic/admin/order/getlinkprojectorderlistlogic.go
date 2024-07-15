package order

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLinkProjectOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLinkProjectOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLinkProjectOrderListLogic {
	return &GetLinkProjectOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLinkProjectOrderListLogic) GetLinkProjectOrderList(req *types.GetLinkProjectOrderListReq) (resp *types.GetLinkProjectListOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
