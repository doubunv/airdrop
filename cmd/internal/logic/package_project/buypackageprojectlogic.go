package package_project

import (
	"air-drop/cmd/internal/data/schema"
	"air-drop/pkg/utils"
	"context"
	"errors"
	"time"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuyPackageProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuyPackageProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuyPackageProjectLogic {
	return &BuyPackageProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BuyPackageProjectLogic) BuyPackageProject(req *types.BuyPackageProjectReq) (resp *types.BuyPackageProjectResp, err error) {
	resp = &types.BuyPackageProjectResp{
		OrderId:   0,
		AmountStr: "",
		Sign:      "",
		CreatedAt: 0,
		Amount:    0,
	}

	idInfo, err := l.svcCtx.PackageModel.FindById(req.Id)
	if err != nil {
		return nil, err
	}

	if idInfo.ID == 0 {
		return nil, errors.New("package project not find")
	}

	rd := &schema.PackageOrder{
		UserId:          utils.GetTokenUid(l.ctx),
		UAddress:        utils.GetTokenAddress(l.ctx),
		PackageId:       idInfo.ID,
		Amount:          idInfo.Price,
		BuyMonth:        idInfo.Month,
		EndTime:         time.Now().AddDate(0, int(idInfo.Month), 0).Unix(),
		SendEarnings:    0,
		MaxEarningsRate: 0,
		CreatedAt:       0,
		UpdatedAt:       0,
	}

	err = l.svcCtx.PackageOrderModel.Insert(rd)
	if err != nil {
		return nil, err
	}

	resp.Amount = rd.Amount
	resp.Sign, resp.AmountStr = BuildPackageSign(l.svcCtx.Config, rd)
	resp.CreatedAt = rd.CreatedAt
	resp.OrderId = rd.ID
	resp.CreatedAt = rd.CreatedAt

	return
}
