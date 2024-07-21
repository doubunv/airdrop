package order

import (
	"air-drop/cmd/internal/data/schema"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropApplyPackageListProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropApplyPackageListProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropApplyPackageListProjectLogic {
	return &DropApplyPackageListProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropApplyPackageListProjectLogic) DropApplyPackageListProject(req *types.DropApplyPackageListProjectReq) (resp *types.DropApplyPackageListProjectResp, err error) {
	resp = &types.DropApplyPackageListProjectResp{
		List:     make([]types.DropApplyPackageListProjectItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	rq := &schema.AmountLog{
		TargetId: req.Id,
		TypeId:   2,
	}

	list, total, err := l.svcCtx.AmountLogModel.GetList(rq, req.STime, req.ETime, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	resp.Total = total

	for _, v := range list {
		t := types.DropApplyPackageListProjectItem{
			Id:        v.ID,
			UAddress:  v.UAddress,
			Amount:    v.Balance,
			CreatedAt: v.CreatedAt,
		}

		resp.List = append(resp.List, t)
	}

	return
}
