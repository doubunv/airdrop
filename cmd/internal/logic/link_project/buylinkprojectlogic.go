package link_project

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuyLinkProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuyLinkProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuyLinkProjectLogic {
	return &BuyLinkProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BuyLinkProjectLogic) BuyLinkProject(req *types.BuyLinkProjectReq) (resp *types.BuyLinkProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
