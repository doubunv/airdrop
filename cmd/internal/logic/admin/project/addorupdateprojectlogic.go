package project

import (
	"air-drop/cmd/internal/data/schema"
	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"context"
	"gorm.io/plugin/soft_delete"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrUpdateProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrUpdateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrUpdateProjectLogic {
	return &AddOrUpdateProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrUpdateProjectLogic) AddOrUpdateProject(req *types.AddOrUpdateProjectReq) (resp *types.AddOrUpdateProjectResp, err error) {
	resp = &types.AddOrUpdateProjectResp{}

	ap := &schema.AirPackageChild{
		Icon:    req.Icon,
		Name:    req.Name,
		Content: req.Content,
	}
	if req.Id == 0 {
		err = l.svcCtx.PackageChildModel.Insert(ap)
		if err != nil {
			return nil, err
		}
	} else {
		ap.ID = req.Id
		if req.IsDeleted == 1 {
			ap.DeletedAt = soft_delete.DeletedAt(time.Now().Unix())
		}
		err = l.svcCtx.PackageChildModel.UpdateSchema(ap)
		if err != nil {
			return nil, err
		}
	}

	return
}
