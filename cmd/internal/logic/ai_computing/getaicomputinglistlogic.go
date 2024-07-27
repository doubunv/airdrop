package ai_computing

import (
	"air-drop/cmd/internal/data/schema"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAiComputingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAiComputingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAiComputingListLogic {
	return &GetAiComputingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAiComputingListLogic) GetAiComputingList(req *types.GetAiComputingListReq) (resp *types.GetAiComputingListResp, err error) {
	resp = &types.GetAiComputingListResp{
		List:     make([]types.GetAiComputingListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}
	re := &schema.AiComputing{}
	list, total, err := l.svcCtx.AiComputingModel.GetList(re, 0, 0, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total

	for _, v := range list {
		t := types.GetAiComputingListItem{
			Id:                  v.ID,
			Name:                v.Name,
			Price:               v.Price,
			Icon:                v.Icon,
			Content:             v.Content,
			ServiceMonth:        v.ServiceMonth,
			ComputingPowerUnit:  v.ComputingPowerUnit,
			ComputingPowerValue: v.ComputingPowerValue,
		}
		resp.List = append(resp.List, t)
	}

	return
}
