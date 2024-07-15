package setting

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSettingKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSettingKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSettingKeyLogic {
	return &GetSettingKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSettingKeyLogic) GetSettingKey(req *types.GetSettingKeyReq) (resp *types.GetSettingKeyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
