package person

import (
	"air-drop/cmd/internal/data/schema"
	"context"
	"time"

	"air-drop/pkg/utils"
	"air-drop/pkg/xerr"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	resp = &types.LoginResp{}
	//req.UAddress = strings.ToLower(req.UAddress)

	//ex, err := l.svcCtx.Redis.Get("mm:nonce:" + req.Nonce)
	//if ex == "" {
	//	return nil, xerr.NewErrCodeMsg(401, "nonce error")
	//}

	blVerify, err := utils.VerifyLoginAddress(req.ChainId, req.Timestamp, req.Nonce, req.UAddress, req.UAddress, req.Signature)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(401, "sign verify error")
	}
	if !blVerify {
		return nil, xerr.NewErrCodeMsg(401, "login address verify failed")
	}

	if req.UAddress == "" {
		return nil, xerr.NewErrCodeMsg(401, "uAddress empty")
	}

	user, err := l.svcCtx.UserModel.GetUserByUAddress(req.UAddress)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(401, "uAddress not find")
	}

	if req.InviteCode != "" {
		if user.InviteCode != req.InviteCode {
			return nil, xerr.NewErrCodeMsg(500, "InviteCode not find")
		}
	}

	if user.ID == 0 {
		user = schema.User{
			ParentAddress: user.ParentAddress,
			UAddress:      req.UAddress,
			CreateAt:      time.Now(),
		}
		err = l.svcCtx.UserModel.Insert(&user)
		if err != nil {
			logx.Errorf("Login InsertUser err:%v", err)
			return nil, xerr.NewErrCodeMsg(401, "insert uAddress error")
		}
	}

	token, err := utils.GenToken(l.svcCtx.Config.Auth.AccessSecret, user.ID, req.UAddress, l.svcCtx.Config.Auth.AccessExpire, false)
	if err != nil {
		return nil, xerr.NewErrMsg("token build error")
	}

	resp.Token = token.AccessToken
	return
}
