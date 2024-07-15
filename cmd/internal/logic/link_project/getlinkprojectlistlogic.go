package link_project

import (
	"air-drop/cmd/internal/data/schema"
	"context"
	"strconv"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLinkProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLinkProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLinkProjectListLogic {
	return &GetLinkProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLinkProjectListLogic) GetLinkProjectList(req *types.GetLinkProjectListReq) (resp *types.GetLinkProjectListResp, err error) {
	resp = &types.GetLinkProjectListResp{
		List:     make([]types.GetLinkProjectList, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	list, total, err := l.svcCtx.LinkModel.GetList(&schema.ArLink{}, int64(0), int64(0), int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	resp.Total = total
	for _, v := range list {
		parseInt, _ := strconv.ParseInt(v.ProjectIds, 10, 64)
		childInfo, _ := l.svcCtx.PackageChildModel.FindById(parseInt)
		tt := types.GetLinkProjectList{
			Id:       v.ID,
			Price:    v.Price,
			DropTime: v.DropTime,
			Name:     childInfo.Name,
		}
		resp.List = append(resp.List, tt)
	}

	return
}
