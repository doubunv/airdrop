package person

import (
	constx "air-drop/cmd/internal/const"
	"context"
	"strings"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteLogic {
	return &InviteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InviteLogic) Invite(req *types.InviteReq) (resp *types.InviteResp, err error) {
	isOwner, err := l.svcCtx.UserModel.GetIsOwnerByParentAddress(strings.ToLower(req.UAddress))
	if err != nil {
		return nil, err
	}
	var tribe int
	var recommender int
	for i := range isOwner {
		recommender++
		if isOwner[i] == constx.TrueInt {
			tribe++
		}
	}
	return &types.InviteResp{
		Tribe:       tribe,
		Recommender: recommender,
	}, nil
}
