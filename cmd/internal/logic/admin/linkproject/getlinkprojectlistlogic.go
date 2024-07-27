package linkproject

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

func (l *GetLinkProjectListLogic) GetLinkProjectList(req *types.AdminGetLinkProjectListReq) (resp *types.AdminGetLinkProjectListResp, err error) {
	resp = &types.AdminGetLinkProjectListResp{
		List:     make([]types.GetLinkProjectListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	rq := &schema.ArLink{}
	list, total, err := l.svcCtx.LinkModel.GetList(rq, req.STime, req.ETime, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total
	for _, v := range list {
		parseInt, _ := strconv.ParseInt(v.ProjectIds, 10, 64)
		childInfo, _ := l.svcCtx.PackageChildModel.FindById(parseInt)
		tt := types.GetLinkProjectListItem{
			Id:          v.ID,
			ProjectId:   parseInt,
			Name:        childInfo.Name,
			Icon:        childInfo.Icon,
			Price:       v.Price,
			DropAmount:  v.DropAmount,
			SellEndTime: v.SellEndTime,
			DropTime:    v.DropTime,
			CreatedAt:   v.CreatedAt,
			Status:      int32(v.Status),
		}
		resp.List = append(resp.List, tt)
	}

	return
}
