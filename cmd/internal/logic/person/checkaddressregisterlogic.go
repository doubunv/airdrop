package person

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAddressRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckAddressRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAddressRegisterLogic {
	return &CheckAddressRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckAddressRegisterLogic) CheckAddressRegister(req *types.UserAddressResgiterReq) (resp *types.UserAddressResgiterResp, err error) {
	resp = &types.UserAddressResgiterResp{}
	addressInfo, err := l.svcCtx.UserModel.GetUserByUAddress(req.UAddress)
	if err != nil {
		return nil, err
	}
	if addressInfo.ID != 0 {
		resp.IsRegister = true
	}
	resp.UAddress = req.UAddress

	return
}
