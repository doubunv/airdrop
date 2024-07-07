package person

import (
	"air-drop/cmd/internal/data/schema"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"air-drop/pkg/utils"
	"air-drop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (resp *types.UserListResp, err error) {

	tokenAddress := utils.GetAdminAddress(l.ctx)
	if tokenAddress == "" {
		return nil, xerr.NewErrMsg("error")
	}

	resp = &types.UserListResp{}
	user := &schema.User{
		InviteCode:    req.InviteCode,
		UAddress:      req.UAddress,
		ParentAddress: req.PAddress}

	list, total, err := l.svcCtx.UserModel.GetUserList(user, req.StartTime, req.EndTime, req.PageNo, req.PageSize)
	if err != nil {
		return nil, err
	}
	// for list
	resp.List = make([]types.UserItem, 0)
	for _, item := range list {
		userT := types.UserItem{
			Id:            item.ID,
			UAddress:      item.UAddress,
			ParentAddress: item.ParentAddress,
			Amount:        item.Amount,
			InviteCode:    item.InviteCode,
			Path:          string(item.Path),
			CreateAt:      item.CreateAt.Format("2006-01-02 15:04:05"),
		}
		resp.List = append(resp.List, userT)
	}

	resp.PageNo = int64(req.PageNo)
	resp.PageSize = int64(req.PageSize)
	resp.Total = total
	return
}
