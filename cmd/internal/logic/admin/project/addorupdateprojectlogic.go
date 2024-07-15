package project

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrUpdateProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrUpdateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrUpdateProjectLogic {
	return &AddOrUpdateProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrUpdateProjectLogic) AddOrUpdateProject(req *types.AddOrUpdateProjectReq) (resp *types.AddOrUpdateProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
