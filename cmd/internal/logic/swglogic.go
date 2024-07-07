package logic

import (
	"air-drop/cmd/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type SwgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwgLogic {
	return &SwgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwgLogic) Swg() error {
	return nil
}
