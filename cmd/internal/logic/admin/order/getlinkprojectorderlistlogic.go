package order

import (
	"air-drop/cmd/internal/data/schema"
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
	resp = &types.GetLinkProjectListOrderResp{
		List:     make([]types.GetLinkProjectOrderListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	rq := &schema.LinkOrder{
		UAddress: req.UAddress,
	}
	l.svcCtx.LinkOrderModel.GetList(rq, req.STime, req.ETime, int(req.Page), int(req.PageSize))

	return
}
