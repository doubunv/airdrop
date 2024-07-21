package user

import (
	"air-drop/cmd/internal/data/schema"
	"context"

	"air-drop/cmd/internal/svc"
	"air-drop/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WalletLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWalletLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WalletLogListLogic {
	return &WalletLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WalletLogListLogic) WalletLogList(req *types.AdminWalletLogListReq) (resp *types.AdminWalletLogListResp, err error) {
	resp = &types.AdminWalletLogListResp{
		List:     make([]types.AdminWalletLogListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	am := &schema.AmountLog{
		TypeId:   int64(req.TypeId),
		UAddress: req.UAddress,
		UserId:   req.UserId,
	}

	list, total, err := l.svcCtx.AmountLogModel.GetList(am, req.STime, req.ETime, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	resp.Total = total
	for _, v := range list {
		t := types.AdminWalletLogListItem{
			CreatedAt:     v.CreatedAt,
			Amount:        v.Balance,
			TargetAddress: v.TargetAddress,
			UAddress:      v.UAddress,
		}

		resp.List = append(resp.List, t)
	}

	return
}
