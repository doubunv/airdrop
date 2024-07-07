package person

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
