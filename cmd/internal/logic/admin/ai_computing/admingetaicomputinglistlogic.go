package ai_computing

import (
	"air-drop/cmd/internal/data/schema"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetAiComputingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminGetAiComputingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetAiComputingListLogic {
	return &AdminGetAiComputingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetAiComputingListLogic) AdminGetAiComputingList(req *types.AdminGetAiComputingListReq) (resp *types.AdminGetAiComputingListResp, err error) {
	resp = &types.AdminGetAiComputingListResp{
		List:     make([]types.AdminGetAiComputingListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	rd := &schema.AiComputing{
		Name: req.Name,
	}
	list, total, err := l.svcCtx.AiComputingModel.GetList(rd, req.STime, req.ETime, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total

	for _, v := range list {
		t := types.AdminGetAiComputingListItem{
			Id:                  v.ID,
			Name:                v.Name,
			Price:               v.Price,
			Icon:                v.Icon,
			Content:             v.Content,
			ServiceMonth:        int32(v.ServiceMonth),
			ComputingPowerUnit:  v.ComputingPowerUnit,
			ComputingPowerValue: v.ComputingPowerValue,
			CreatedAt:           v.CreatedAt,
			Status:              int32(v.Status),
		}

		resp.List = append(resp.List, t)
	}

	return
}
