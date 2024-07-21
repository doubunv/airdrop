package user

import (
	"air-drop/cmd/internal/data/schema"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.GetUserListResp, err error) {
	resp = &types.GetUserListResp{
		List:     make([]types.GetUserListitem, 0),
		Page:     0,
		PageSize: 0,
		Total:    0,
	}

	rq := &schema.User{
		UAddress:      req.UAddress,
		ParentAddress: req.PAddress,
		InviteCode:    req.InviteCode,
	}
	list, total, err := l.svcCtx.UserModel.GetUserList(rq, req.STime, req.ETime, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	resp.Total = total

	for _, v := range list {
		t := types.GetUserListitem{
			Id:         v.ID,
			UAddress:   v.UAddress,
			InviteCode: v.InviteCode,
			CreatedAt:  v.CreatedAt,
			PAddress:   v.ParentAddress,
		}
		resp.List = append(resp.List, t)
	}

	return
}
