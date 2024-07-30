package link_project

import (
	"air-drop/cmd/internal/data/schema"
	"air-drop/pkg/utils"
	"context"
	"errors"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuyLinkProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuyLinkProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuyLinkProjectLogic {
	return &BuyLinkProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BuyLinkProjectLogic) BuyLinkProject(req *types.BuyLinkProjectReq) (resp *types.BuyLinkProjectResp, err error) {
	resp = &types.BuyLinkProjectResp{
		OrderId:   0,
		Amount:    0,
		Sign:      "",
		CreatedAt: 0,
		AmountStr: "",
	}

	linkInfo, err := l.svcCtx.LinkModel.FindById(req.Id)
	if err != nil {
		return nil, err
	}

	if linkInfo.ID == 0 {
		return nil, errors.New("link project not find")
	}

	rd := &schema.LinkOrder{
		UserId:    utils.GetTokenUid(l.ctx),
		UAddress:  utils.GetTokenAddress(l.ctx),
		LinkId:    req.Id,
		BuyAmount: linkInfo.DropAmount,
		DropTime:  linkInfo.DropTime,
		BuyNumber: 1,
		Status:    1,
	}
	err = l.svcCtx.LinkOrderModel.Insert(rd)
	if err != nil {
		return nil, err
	}

	resp.Amount = rd.BuyAmount
	resp.Sign, resp.AmountStr = BuildLinkSign(l.svcCtx.Config, rd)
	resp.CreatedAt = rd.CreatedAt
	resp.OrderId = rd.ID
	resp.CreatedAt = rd.CreatedAt
	return
}
