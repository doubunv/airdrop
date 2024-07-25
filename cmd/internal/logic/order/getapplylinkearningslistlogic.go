package order

import (
	"air-drop/cmd/internal/data/schema"
	"air-drop/pkg/utils"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplyLinkEarningsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApplyLinkEarningsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplyLinkEarningsListLogic {
	return &GetApplyLinkEarningsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApplyLinkEarningsListLogic) GetApplyLinkEarningsList(req *types.GetApplyLinkLogListReq) (resp *types.GetApplyLinkLogListResp, err error) {
	resp = &types.GetApplyLinkLogListResp{
		List:     make([]types.GetApplyLinkLogListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	rq := &schema.ArLinkReceive{
		UserId: utils.GetTokenUid(l.ctx),
		Status: int64(req.Status),
	}
	list, total, err := l.svcCtx.LinkReceiveModel.GetList(rq, int64(0), int64(0), int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	resp.Total = total
	for _, v := range list {
		t := types.GetApplyLinkLogListItem{
			CreatedAt: v.CreatedAt,
			UAddress:  v.UAddress,
			Amount:    v.Amount,
		}
		resp.List = append(resp.List, t)
	}

	return
}
