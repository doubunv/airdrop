package person

import (
	"air-drop/cmd/internal/data/schema"
	"air-drop/pkg/utils"
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

func (l *WalletLogListLogic) WalletLogList(req *types.WalletLogListReq) (resp *types.WalletLogListResp, err error) {
	resp = &types.WalletLogListResp{
		List:     make([]types.WalletLogListItem, 0),
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    0,
	}

	tokenUid := utils.GetTokenUid(l.ctx)

	rq := &schema.AmountLog{
		UserId:   tokenUid,
		TypeId:   int64(req.TypeId),
		TargetId: req.TargetId,
	}

	list, total, err := l.svcCtx.AmountLogModel.GetList(rq, req.STime, req.ETime, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	resp.Total = total

	for _, v := range list {
		t := types.WalletLogListItem{
			CreatedAt:     v.CreatedAt,
			Amount:        v.Balance,
			TargetAddress: v.TargetAddress,
			UAddress:      v.UAddress,
		}
		resp.List = append(resp.List, t)
	}

	return
}
