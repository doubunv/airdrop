package person

import (
	"context"
	"strings"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"air-drop/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInviterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInviterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInviterLogic {
	return &GetInviterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInviterLogic) GetInviter() (resp *types.GetInviterResp, err error) {
	resp = &types.GetInviterResp{}
	value := utils.GetTokenAddress(l.ctx)
	addressInfo, err := l.svcCtx.UserModel.GetUserByUAddress(strings.ToLower(value))
	if err != nil {
		return nil, err
	}
	resp.PAddress = addressInfo.ParentAddress
	parentInfo, _ := l.svcCtx.UserModel.GetUserByUAddress(strings.ToLower(addressInfo.ParentAddress))

	resp.PInviteCode = parentInfo.InviteCode
	resp.InviteCode = addressInfo.InviteCode
	return
}
