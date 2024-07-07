package order

import (
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
	// todo: add your logic here and delete this line

	return
}
