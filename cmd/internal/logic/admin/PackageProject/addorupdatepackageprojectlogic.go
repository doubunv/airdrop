package PackageProject

import (
	"air-drop/cmd/internal/data/schema"
	"context"
	"errors"
	"gorm.io/plugin/soft_delete"
	"strings"
	"time"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrUpdatePackageProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrUpdatePackageProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrUpdatePackageProjectLogic {
	return &AddOrUpdatePackageProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrUpdatePackageProjectLogic) AddOrUpdatePackageProject(req *types.AddOrUpdatePackageProjectReq) (resp *types.AddOrUpdatePackageProjectResp, err error) {
	resp = &types.AddOrUpdatePackageProjectResp{}

	if len(req.ProjectIds) == 0 {
		return nil, errors.New("ProjectIds id none")
	}

	var projectIds = ""
	for _, v := range req.ProjectIds {
		projectIds = v.Name + "," + projectIds
	}
	projectIds = strings.TrimRight(projectIds, ",")

	ar := &schema.AirPackage{
		Name:       "",
		ProjectIds: projectIds,
		Price:      req.Price,
		Month:      req.Month,
		Status:     req.Status,
	}

	if req.Id == 0 {
		err = l.svcCtx.PackageModel.Insert(ar)
		if err != nil {
			return nil, err
		}
	} else {
		ar.ID = req.Id
		if req.IsDeleted == 1 {
			ar.DeletedAt = soft_delete.DeletedAt(time.Now().Unix())
		}
		err = l.svcCtx.PackageModel.UpdateSchema(ar)
		if err != nil {
			return nil, err
		}
	}

	return
}
