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
	resp = &types.GetSettingKeyResp{
		Key:   req.Key,
		Value: "",
	}
	value := l.svcCtx.SettingModel.FindByKey(req.Key)
	resp.Value = value

	return
}
