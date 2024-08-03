package order

import (
	"air-drop/cmd/internal/data/model"
	"air-drop/cmd/internal/data/schema"
	"air-drop/pkg/utils"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyLinkEarningsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyLinkEarningsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyLinkEarningsLogic {
	return &ApplyLinkEarningsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyLinkEarningsLogic) ApplyLinkEarnings(req *types.ApplyLinkEarningsReq) (resp *types.ApplyLinkEarningsResp, err error) {
	resp = &types.ApplyLinkEarningsResp{}
	da := &schema.ArLinkReceive{
		UserId:   utils.GetTokenUid(l.ctx),
		UAddress: req.Address,
		Amount:   req.Amount,
		Status:   model.LinkReceiveStatus1,
		OrderId:  req.Id,
	}

	err = l.svcCtx.LinkReceiveModel.Insert(da)
	if err != nil {
		return nil, err
	}

	return
}
