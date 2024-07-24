package person

import (
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"air-drop/pkg/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type NonceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNonceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NonceLogic {
	return &NonceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NonceLogic) Nonce() (resp *types.GetNonceResp, err error) {
	resp = &types.GetNonceResp{}
	str := utils.RandStr(8)
	resp.Nonce = str
	//ex, err := l.svcCtx.Redis.SetnxEx("mm:nonce:"+str, "1", 60)
	//if err != nil {
	//	alert.SendMsg(fmt.Sprintf("redis set error err:%+v", err))
	//	return nil, xerr.NewErrMsg("try later")
	//}
	//if !ex {
	//	alert.SendMsg("nonce repeat")
	//	return nil, xerr.NewErrMsg("try again later")
	//}
	return
}
