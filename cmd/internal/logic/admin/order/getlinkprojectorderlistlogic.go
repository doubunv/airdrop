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
	list, total, err := l.svcCtx.LinkOrderModel.GetList(rq, req.STime, req.ETime, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total

	for _, v := range list {
		linkDetail, _ := l.svcCtx.LinkModel.FindDetailById(v.LinkId)
		//appInfo, _ := l.svcCtx.LinkReceiveModel.FindByOrderId(v.ID)
		t := types.GetLinkProjectOrderListItem{
			Id:        v.ID,
			UserId:    v.UserId,
			Name:      linkDetail.ProjectName,
			UAddress:  v.UAddress,
			Amount:    v.BuyAmount,
			CreatedAt: v.CreatedAt,
			Status:    int32(v.Status),
			Icno:      linkDetail.ProjectIcon,
		}
		resp.List = append(resp.List, t)
	}

	return
}
