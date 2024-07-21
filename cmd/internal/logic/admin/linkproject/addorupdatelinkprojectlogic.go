package linkproject

import (
	"air-drop/cmd/internal/data/schema"
	"context"
	"github.com/pkg/errors"
	"gorm.io/plugin/soft_delete"
	"strconv"
	"time"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrUpdateLinkProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrUpdateLinkProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrUpdateLinkProjectLogic {
	return &AddOrUpdateLinkProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrUpdateLinkProjectLogic) AddOrUpdateLinkProject(req *types.AddOrUpdateLinkProjectReq) (resp *types.AddOrUpdateLinkProjectResp, err error) {
	resp = &types.AddOrUpdateLinkProjectResp{}
	if req.ProjectId == 0 {
		return nil, errors.New("ProjectId is none")
	}

	lk := &schema.ArLink{
		ProjectIds:  strconv.FormatInt(req.ProjectId, 10),
		DropTime:    req.DropTime,
		Price:       req.Price,
		Status:      int64(req.Status),
		SellEndTime: req.SellEndTime,
		DropAmount:  req.DropAmount,
	}

	if req.Id == 0 {
		err = l.svcCtx.LinkModel.Insert(lk)
		if err != nil {
			return nil, err
		}
	} else {
		if req.IsDeleted == 1 {
			lk.DeletedAt = soft_delete.DeletedAt(time.Now().Unix())
		}
		lk.ID = req.Id
		err = l.svcCtx.LinkModel.UpdateSchema(lk)
		if err != nil {
			return nil, err
		}
	}

	return
}
