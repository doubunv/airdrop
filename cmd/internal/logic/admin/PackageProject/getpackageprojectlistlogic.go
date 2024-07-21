package PackageProject

import (
	"air-drop/cmd/internal/data/schema"
	"context"
	"strconv"
	"strings"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPackageProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPackageProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPackageProjectListLogic {
	return &GetPackageProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPackageProjectListLogic) GetPackageProjectList(req *types.AdminGetPackageProjectListReq) (resp *types.AdminGetPackageProjectListResp, err error) {
	resp = &types.AdminGetPackageProjectListResp{
		List:     make([]types.GetPackageProjectListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	ap := &schema.AirPackage{}
	list, total, err := l.svcCtx.PackageModel.GetList(ap, req.STime, req.ETime, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total

	for _, v := range list {
		t := types.GetPackageProjectListItem{
			Id:     v.ID,
			Name:   v.Name,
			Price:  v.Price,
			Month:  v.Month,
			Status: v.Status,
			List:   make([]types.AdminPackageProjectChildListItem, 0),
		}

		childs := make([]int64, 0)
		for _, v := range strings.Split(v.ProjectIds, ",") {
			parseInt, _ := strconv.ParseInt(v, 10, 64)
			childs = append(childs, parseInt)
		}

		childsI, _ := l.svcCtx.PackageChildModel.FindByIds(childs)
		for _, v1 := range childsI {
			t1 := types.AdminPackageProjectChildListItem{
				Id:   v1.ID,
				Name: v1.Name,
				Icon: v1.Icon,
			}
			t.List = append(t.List, t1)
		}
		resp.List = append(resp.List, t)
	}

	return
}
