package person

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"air-drop/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInviteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteCountLogic {
	return &InviteCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InviteCountLogic) InviteCount() (resp *types.InviteCountResp, err error) {
	userAddress := utils.GetTokenAddress(l.ctx)
	resp = &types.InviteCountResp{}

	addressInfo, err := l.svcCtx.UserModel.GetUserByUAddress(userAddress)
	if err != nil {
		return nil, err
	}
	resp.Count = l.svcCtx.UserModel.CountInvite(userAddress)
	resp.TeamLevel = addressInfo.TeamLevel
	return
}
