package person

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"air-drop/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInviteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteListLogic {
	return &InviteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InviteListLogic) InviteList(req *types.InviteUserReq) (resp *types.InviteUserResp, err error) {
	resp = &types.InviteUserResp{}
	userAddress := utils.GetTokenAddress(l.ctx)
	list, total, err := l.svcCtx.UserModel.GetInviteCountList(userAddress, int(req.PageNo), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.List = make([]types.UserInviteItem, 0)
	for _, user := range list {
		resp.List = append(resp.List, types.UserInviteItem{
			ID:        user.ID,
			UAddress:  user.UAddress,
			CreatedAt: user.CreatedAt.Unix(),
		})
	}
	resp.PageNo = req.PageNo
	resp.PageSize = req.PageSize
	resp.Total = total
	return
}
