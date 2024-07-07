package person

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WalletLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWalletLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WalletLogListLogic {
	return &WalletLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WalletLogListLogic) WalletLogList(req *types.WalletLogListReq) (resp *types.WalletLogListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
