package order

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLinkOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLinkOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLinkOrderListLogic {
	return &GetLinkOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLinkOrderListLogic) GetLinkOrderList(req *types.GetLinkOrderListReq) (resp *types.GetLinkOrderListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
