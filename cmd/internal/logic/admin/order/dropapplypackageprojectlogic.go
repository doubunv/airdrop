package order

import (
	"air-drop/cmd/internal/data/schema"
	"context"
	"errors"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropApplyPackageProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDropApplyPackageProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropApplyPackageProjectLogic {
	return &DropApplyPackageProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropApplyPackageProjectLogic) DropApplyPackageProject(req *types.DropApplyPackageProjectReq) (resp *types.DropApplyPackageProjectResp, err error) {
	resp = &types.DropApplyPackageProjectResp{}
	if req.Id == 0 {
		return nil, errors.New("id is none")
	}

	orderInfo, err := l.svcCtx.PackageOrderModel.FindById(req.Id)
	if err != nil {
		return nil, err
	}
	if orderInfo.ID == 0 {
		return nil, errors.New("order not find")
	}

	am := &schema.AmountLog{
		UserId:   orderInfo.UserId,
		UAddress: orderInfo.UAddress,
		Balance:  orderInfo.Amount,
		TargetId: orderInfo.UserId,
		TypeId:   2, //1-佣金记录， 2-回报记录
		Mark:     "send by admin",
	}
	err = l.svcCtx.AmountLogModel.Insert(am)
	if err != nil {
		return nil, err
	}

	return
}
