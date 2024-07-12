package order

import (
	"air-drop/cmd/internal/data/schema"
	"air-drop/pkg/utils"
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
	resp = &types.GetPackageOrderListResp{
		List:     make([]types.GetPackageOrderListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	rq := &schema.PackageOrder{
		UserId: utils.GetTokenUid(l.ctx),
	}
	list, total, err := l.svcCtx.PackageOrderModel.GetList(rq, int64(0), int64(0), int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total

	ids := make([]int64, 0)
	for _, v := range list {
		ids = append(ids, v.PackageId)
	}

	for _, v := range list {
		t := types.GetPackageOrderListItem{
			Id:              v.ID,
			Created_at:      v.CreatedAt,
			Amount:          v.Amount,
			Month:           int32(v.BuyMonth),
			LeftDay:         utils.DiffDays(v.CreatedAt, v.EndTime),
			SendEarnings:    v.SendEarnings,
			MaxEarningsRate: v.MaxEarningsRate,
		}
		resp.List = append(resp.List, t)
	}

	return
}
