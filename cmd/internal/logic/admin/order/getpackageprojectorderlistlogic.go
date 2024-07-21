package order

import (
	"air-drop/cmd/internal/data/schema"
	"air-drop/pkg/utils"
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
	resp = &types.GetPackageProjectListOrderResp{
		List:     make([]types.GetPackageProjectOrderListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	rq := &schema.PackageOrder{
		UserId:   req.UserId,
		UAddress: req.UAddress,
	}
	list, total, err := l.svcCtx.PackageOrderModel.GetList(rq, req.STime, req.ETime, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	resp.Total = total

	for _, v := range list {
		t := types.GetPackageProjectOrderListItem{
			Id:           v.ID,
			UserId:       v.UserId,
			UAddress:     v.UAddress,
			Amount:       v.Amount,
			BuyMonth:     int32(v.BuyMonth),
			LeftDay:      utils.DiffDays(v.CreatedAt, v.EndTime),
			SendEarnings: v.SendEarnings,
			CreatedAt:    v.CreatedAt,
		}
		resp.List = append(resp.List, t)
	}

	return
}
