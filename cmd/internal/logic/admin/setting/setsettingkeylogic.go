package setting

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetSettingKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetSettingKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSettingKeyLogic {
	return &SetSettingKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetSettingKeyLogic) SetSettingKey(req *types.SetSettingKeyReq) (resp *types.SetSettingKeyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
