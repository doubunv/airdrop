package setting

import (
	"air-drop/cmd/internal/data/schema"
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
	resp = &types.SetSettingKeyResp{}
	kd := &schema.Setting{
		Key:   req.Key,
		Value: req.Value,
	}
	err = l.svcCtx.SettingModel.Insert(kd)
	if err != nil {
		return nil, err
	}
	return
}
