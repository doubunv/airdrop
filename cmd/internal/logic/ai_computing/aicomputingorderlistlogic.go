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

func (l *AiComputingOrderListLogic) AiComputingOrderList(req *types.AiComputingOrderListReq) (resp *types.AiComputingOrderListResp, err error) {
	resp = &types.AiComputingOrderListResp{
		List:     make([]types.AiComputingOrderListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	rq := &schema.AiComputingOrder{
		UserId: utils.GetTokenUid(l.ctx),
	}
	list, total, err := l.svcCtx.AiComputingOrderModel.GetList(rq, int64(0), int64(0), int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total

	if len(list) == 0 {
		return
	}

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
		t := types.AiComputingOrderListItem{
			Id:                 v.ID,
			Name:               idsDic[v.AiComputingId].Name,
			Amount:             v.Amount,
			ComputingPowerUnit: idsDic[v.AiComputingId].ComputingPowerUnit,
			ServiceMonth:       v.ServiceMonth,
			LeftDay:            int64(utils.DiffDays(v.CreatedAt, v.EndTime)),
			WithdrawAmount:     v.WithdrawAmount,
			CreatedAt:          v.CreatedAt,
		}

		resp.List = append(resp.List, t)
	}

	return
}
