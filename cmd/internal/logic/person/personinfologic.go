package person

import (
	"air-drop/pkg/utils"
	"context"
	"errors"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PersonInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPersonInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersonInfoLogic {
	return &PersonInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PersonInfoLogic) PersonInfo(req *types.PersonInfoReq) (resp *types.PersonInfoResp, err error) {
	resp = &types.PersonInfoResp{
		UAddress:      "",
		ParentAddress: "",
		InviteCode:    "",
		CreateAt:      0,
	}

	addr := utils.GetTokenAddress(l.ctx)
	userInfo, err := l.svcCtx.UserModel.GetUserByUAddress(addr)
	if err != nil {
		return nil, err
	}

	if userInfo.ID == 0 {
		return nil, errors.New("user not find")
	}

	resp.CreateAt = userInfo.CreatedAt
	resp.ParentAddress = userInfo.ParentAddress
	resp.InviteCode = userInfo.InviteCode
	resp.UAddress = userInfo.UAddress
	return
}
