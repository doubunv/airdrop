package ai_computing

import (
	"air-drop/cmd/internal/data/schema"
	"air-drop/pkg/utils"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AiComputingOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAiComputingOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AiComputingOrderListLogic {
	return &AiComputingOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AiComputingOrderListLogic) AiComputingOrderList(req *types.AdminAiComputingOrderListReq) (resp *types.AdminAiComputingOrderListResp, err error) {
	resp = &types.AdminAiComputingOrderListResp{
		List:     make([]types.AdminAiComputingOrderListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	rd := &schema.AiComputingOrder{}
	list, total, err := l.svcCtx.AiComputingOrderModel.GetList(rd, req.STime, req.ETime, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total

	var ids = make([]int64, 0)
	for _, v := range list {
		ids = append(ids, v.AiComputingId)
	}

	byIdsInfo, err := l.svcCtx.AiComputingModel.FindByIds(ids)
	if err != nil {
		return nil, err
	}

	var idsDic = make(map[int64]schema.AiComputing, 0)
	for _, v := range byIdsInfo {
		idsDic[v.ID] = v
	}

	for _, v := range list {
		t := types.AdminAiComputingOrderListItem{
			Id:                 v.ID,
			Name:               idsDic[v.AiComputingId].Name,
			Amount:             v.Amount,
			ComputingPowerUnit: idsDic[v.AiComputingId].Name,
			ServiceMonth:       v.ServiceMonth,
			LeftDay:            int64(utils.DiffDays(v.CreatedAt, v.EndTime)),
			SendAmount:         v.SendAmount,
			CreatedAt:          v.CreatedAt,
		}
		resp.List = append(resp.List, t)
	}

	return
}
