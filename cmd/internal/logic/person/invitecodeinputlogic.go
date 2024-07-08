package person

import (
	"air-drop/cmd/errs"
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"air-drop/pkg/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteCodeInputLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInviteCodeInputLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteCodeInputLogic {
	return &InviteCodeInputLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InviteCodeInputLogic) InviteCodeInput(req *types.InviteCodeReq) error {
	//req.InviteCode = strings.ToLower(req.InviteCode)

	userAddress := utils.GetTokenAddress(l.ctx)
	userInfo, err := l.svcCtx.UserModel.GetUserByUAddress(userAddress)
	if err != nil {
		return errs.UserNotExists
	}
	if userInfo.ParentAddress != "" {
		return errs.UserInviteeExists
	}

	parentInfo := l.svcCtx.UserModel.GetUnique(req.InviteCode)
	if parentInfo.ID == 0 {
		return errs.UserInviteCodeNotExists
	}

	if parentInfo.UAddress == userAddress {
		return errs.UserInviteeSelf
	}

	userInfo.ParentAddress = parentInfo.UAddress
	userInfo.Path = parentInfo.Path.Append(parentInfo.ID)
	userInfo.InvitePathDistance = parentInfo.InvitePathDistance + 1

	err = l.svcCtx.UserModel.UpdateSchema(&userInfo)

	return nil
}
