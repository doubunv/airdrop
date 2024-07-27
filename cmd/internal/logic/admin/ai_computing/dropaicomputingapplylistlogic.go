package ai_computing

import (
	"air-drop/cmd/internal/data/schema"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropAiComputingApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropAiComputingApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropAiComputingApplyListLogic {
	return &DropAiComputingApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropAiComputingApplyListLogic) DropAiComputingApplyList(req *types.DropAiComputingApplyListReq) (resp *types.DropAiComputingApplyListResp, err error) {
	resp = &types.DropAiComputingApplyListResp{
		List:     make([]types.DropAiComputingApplyListItem, 0),
		Page:     0,
		PageSize: 0,
		Total:    0,
	}

	rd := &schema.AiApply{
		AiComputingOrderId: req.Id,
	}
	list, total, err := l.svcCtx.AiApplyModel.GetList(rd, 0, 0, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total

	for _, v := range list {
		t := types.DropAiComputingApplyListItem{
			UAddress:  v.UAddress,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
		}
		resp.List = append(resp.List, t)
	}

	return
}
