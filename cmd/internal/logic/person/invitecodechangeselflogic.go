package person

import (
	"context"
	"regexp"
	"strings"

	errs2 "air-drop/cmd/errs"
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"air-drop/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteCodeChangeSelfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInviteCodeChangeSelfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteCodeChangeSelfLogic {
	return &InviteCodeChangeSelfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InviteCodeChangeSelfLogic) InviteCodeChangeSelf(req *types.InviteCodeReq) error {
	userAddress := utils.GetTokenAddress(l.ctx)

	regex := regexp.MustCompile(`^[a-zA-Z0-9]{6,18}$`) // 匹配 6 到 18 位数字和字母的组合
	match := regex.MatchString(req.InviteCode)
	if match {
		req.InviteCode = strings.ToLower(req.InviteCode)
	} else {
		return errs2.InviteCodeNowAllow
	}

	userInfo, err := l.svcCtx.UserModel.GetUserByUAddress(userAddress)
	if err != nil {
		return errs2.UserNotExists
	}
	if userInfo.InviteCode != "" {
		return errs2.UserInviteCodeNoChange
	}

	haveCode := l.svcCtx.UserModel.GetUnique(req.InviteCode)
	if haveCode.ID > 0 && haveCode.UAddress != userAddress {
		return errs2.InviteCodeExists
	}

	userInfo.InviteCode = req.InviteCode
	if err := l.svcCtx.UserModel.UpdateSchema(&userInfo); err != nil {
		logx.Errorf("UserModel.UpdateSchema err:%v", err)
		return errs2.DbUpdateErr
	}

	return nil
}
