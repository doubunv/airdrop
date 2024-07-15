package setting

import (
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSettingListLogic {
	return &GetSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSettingListLogic) GetSettingList(req *types.GetSettingListReq) (resp *types.GetSettingListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
