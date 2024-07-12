package order

import (
	"air-drop/cmd/internal/data/model"
	"air-drop/cmd/internal/data/schema"
	"air-drop/pkg/utils"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderInfoLogic {
	return &GetOrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderInfoLogic) GetOrderInfo(req *types.GetOrderInfoReq) (resp *types.GetOrderInfoResp, err error) {
	resp = &types.GetOrderInfoResp{
		TotalBuyAmount:         0, //累计投资金额
		LinkEarnings:           0, //全链空投收益
		ReceivePackageEarnings: 0, //已发放私募空投
		WaitPackageEarnings:    0, //待领取的私募空投
		CommissionEarnings:     0, //佣金奖励
	}

	tokenUid := utils.GetTokenUid(l.ctx)
	userInfo, err := l.svcCtx.UserModel.GetUserById(tokenUid)
	if err != nil {
		return nil, err
	}
	resp.TotalBuyAmount = userInfo.PayAmount
	resp.CommissionEarnings = userInfo.TotalCommission

	sumSendEarnings, err := l.svcCtx.PackageOrderModel.SumSendEarnings(tokenUid)
	resp.ReceivePackageEarnings = sumSendEarnings.SendEarnings

	receiveSum, err := l.svcCtx.LinkReceiveModel.SumLinkReceive(&schema.ArLinkReceive{UserId: tokenUid, Status: model.LinkReceiveStatus3})
	resp.LinkEarnings = receiveSum.Amount

	receiveSum1, err := l.svcCtx.LinkReceiveModel.SumLinkReceive(&schema.ArLinkReceive{UserId: tokenUid, Status: model.LinkReceiveStatus2})
	resp.LinkEarnings = receiveSum1.Amount

	return
}
