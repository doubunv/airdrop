package linkproject

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrUpdateLinkProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrUpdateLinkProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrUpdateLinkProjectLogic {
	return &AddOrUpdateLinkProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrUpdateLinkProjectLogic) AddOrUpdateLinkProject(req *types.AddOrUpdateLinkProjectReq) (resp *types.AddOrUpdateLinkProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
