package ai_computing

import (
	"air-drop/cmd/internal/data/schema"
	"context"
	"gorm.io/plugin/soft_delete"
	"time"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrUpdateAiComputingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrUpdateAiComputingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrUpdateAiComputingLogic {
	return &AddOrUpdateAiComputingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrUpdateAiComputingLogic) AddOrUpdateAiComputing(req *types.AddOrUpdateAiComputingReq) (resp *types.AddOrUpdateAiComputingResp, err error) {
	resp = &types.AddOrUpdateAiComputingResp{}

	if req.IsDeleted == 1 {
		err = l.svcCtx.AiComputingModel.UpdateSchema(&schema.AiComputing{DeletedAt: soft_delete.DeletedAt(time.Now().Unix())})
		if err != nil {
			return nil, err
		}
		return
	}

	ai := &schema.AiComputing{
		Icon:                req.Icon,
		Name:                req.Name,
		Content:             req.Content,
		Price:               req.Price,
		ComputingPowerValue: req.ComputingPowerValue,
		ComputingPowerUnit:  req.ComputingPowerUnit,
		ServiceMonth:        int64(req.ServiceMonth),
		Status:              int64(req.Status),
	}
	if req.Id == 0 {
		err = l.svcCtx.AiComputingModel.Insert(ai)
		if err != nil {
			return nil, err
		}
	}

	if req.Id != 0 {
		ai.ID = req.Id
		err = l.svcCtx.AiComputingModel.UpdateSchema(ai)
		if err != nil {
			return nil, err
		}
	}

	return
}
