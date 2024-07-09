package package_project

import (
	"air-drop/cmd/internal/data/schema"
	"context"
	"strconv"
	"strings"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"
	"github.com/jinzhu/copier"
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

func (l *GetPackageProjectListLogic) GetPackageProjectList(req *types.GetPackageProjectListReq) (resp *types.GetPackageProjectListResp, err error) {
	resp = &types.GetPackageProjectListResp{
		List:     make([]types.GetPackageProjectList, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}
	airPackage := &schema.AirPackage{}
	list, total, err := l.svcCtx.PackageModel.GetList(airPackage, int64(0), int64(0), int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total

	for _, v := range list {
		ids := make([]int64, 0)
		for _, v := range strings.Split(v.ChildId, ",") {
			parseInt, _ := strconv.ParseInt(v, 10, 64)
			ids = append(ids, parseInt)
		}

		byIds, err := l.svcCtx.PackageChildModel.FindByIds(ids)
		if err != nil {
			return nil, err
		}
		childList := make([]types.PackageProjectChildListItem, 0)
		for _, v := range byIds {
			tt := types.PackageProjectChildListItem{
				Id:   v.ID,
				Name: v.Name,
			}
			childList = append(childList, tt)
		}

		t := types.GetPackageProjectList{
			Id:    v.ID,
			Name:  v.Name,
			Price: v.Price,
			Month: v.Month,
			List:  childList,
		}
		copier.Copy(&t, v)
		resp.List = append(resp.List, t)
	}

	return
}
