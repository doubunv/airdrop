package order

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropLinkApplyLinkProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropLinkApplyLinkProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropLinkApplyLinkProjectLogic {
	return &DropLinkApplyLinkProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropLinkApplyLinkProjectLogic) DropLinkApplyLinkProject(req *types.DropLinkApplyLinkProjectReq) (resp *types.DropApplyLinkProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
