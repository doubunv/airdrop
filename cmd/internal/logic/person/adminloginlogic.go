package person

import (
	"air-drop/pkg/utils"
	"air-drop/pkg/xerr"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req *types.AdminLoginReq) (resp *types.AdminLoginResp, err error) {
	resp = &types.AdminLoginResp{}
	addressInfo, err := l.svcCtx.AdminUserModel.GetUserByUAddress(req.UAddress)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(401, "address get error")
	}

	if addressInfo.ID == 0 {
		return nil, xerr.NewErrCodeMsg(401, "address not find")
	}

	blVerify, err := utils.VerifyLoginAddress(req.ChainId, req.Timestamp, req.Nonce, req.UAddress, req.UAddress, req.Signature)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(401, "sign verify error")
	}
	if !blVerify {
		return nil, xerr.NewErrCodeMsg(401, "login address verify failed")
	}

	token, err := utils.GenToken(l.svcCtx.Config.Auth.AccessSecret, addressInfo.ID, req.UAddress, l.svcCtx.Config.Auth.AccessExpire, true)
	if err != nil {
		return nil, xerr.NewErrMsg("token build error")
	}

	resp.Token = token.AccessToken
	resp.Role = addressInfo.Role

	return
}
