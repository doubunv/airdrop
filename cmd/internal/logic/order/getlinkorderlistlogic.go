package order

import (
	"air-drop/cmd/internal/data/dto"
	"air-drop/cmd/internal/data/schema"
	"air-drop/pkg/utils"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"

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
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			logc.Infof(context.Background(), "handler panic: %v", err)
		}
	}()

	resp = &types.GetLinkOrderListResp{
		List:     make([]types.GetLinkOrderListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	rq := &schema.LinkOrder{
		UserId: utils.GetTokenUid(l.ctx),
		Status: int64(req.Ststus),
	}
	list, total, err := l.svcCtx.LinkOrderModel.GetList(rq, int64(0), int64(0), int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total

	dicLinkIds := make([]int64, 0)
	for _, v := range list {
		dicLinkIds = append(dicLinkIds, v.LinkId)
	}
	linkIds, err := l.svcCtx.LinkModel.FindByIds(dicLinkIds)
	if err != nil {
		return nil, err
	}
	dic := make(map[int64]dto.LinkDetail, 0)
	for _, v := range linkIds {
		dic[v.ID] = v
	}

	for _, v := range list {
		t := types.GetLinkOrderListItem{
			Id:        v.ID,
			CreatedAt: v.CreatedAt,
			Name:      "",
			BuyNumber: int32(v.BuyNumber),
			BuyAmount: v.BuyAmount,
			DropTime:  v.DropTime,
			Status:    int32(v.Status),
		}
		if data, ok := dic[v.LinkId]; ok {
			t.Name = data.ProjectName
		}
		resp.List = append(resp.List, t)
	}

	return
}
