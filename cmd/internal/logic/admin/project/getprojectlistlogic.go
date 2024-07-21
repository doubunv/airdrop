package project

import (
	"air-drop/cmd/internal/data/schema"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectListLogic {
	return &GetProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectListLogic) GetProjectList(req *types.GetProjectListReq) (resp *types.GetProjectListResp, err error) {
	resp = &types.GetProjectListResp{
		List:     make([]types.GetProjectListItem, 0),
		Page:     0,
		PageSize: 0,
		Total:    0,
	}

	rq := &schema.AirPackageChild{}
	list, total, err := l.svcCtx.PackageChildModel.GetList(rq, req.STime, req.ETime, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total

	for _, v := range list {
		t := types.GetProjectListItem{
			Id:      v.ID,
			Icon:    v.Icon,
			Name:    v.Name,
			Content: v.Content,
		}
		resp.List = append(resp.List, t)
	}

	return
}
