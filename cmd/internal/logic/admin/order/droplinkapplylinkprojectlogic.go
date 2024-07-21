package order

import (
	"air-drop/cmd/internal/data/schema"
	"context"
	"errors"
	"time"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropLinkApplyLinkProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropLinkApplyLinkProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropLinkApplyLinkProjectLogic {
	return &DropLinkApplyLinkProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropLinkApplyLinkProjectLogic) DropLinkApplyLinkProject(req *types.DropLinkApplyLinkProjectReq) (resp *types.DropApplyLinkProjectResp, err error) {
	resp = &types.DropApplyLinkProjectResp{}

	orderInfo, err := l.svcCtx.LinkReceiveModel.FindByOrderId(req.Id)
	if err != nil {
		return nil, err
	}

	if orderInfo.ID == 0 {
		return nil, errors.New("order not find")
	}

	if orderInfo.Status != 1 {
		return nil, errors.New("order status not wait drop ")
	}

	lo := &schema.ArLinkReceive{
		Status:   2,
		DropTime: time.Now().Unix(),
	}
	err = l.svcCtx.LinkReceiveModel.UpdateSchema(lo)
	if err != nil {
		return nil, err
	}

	lk := &schema.LinkOrder{
		ID:     req.Id,
		Status: 2,
	}
	err = l.svcCtx.LinkOrderModel.UpdateSchema(lk)
	if err != nil {
		return nil, err
	}
	return
}
