package service

import (
	"air-drop/cmd/internal/data/schema"
	"air-drop/cmd/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logc"
)

type InviteCommission struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInviteCommission(ctx context.Context, svcCtx *svc.ServiceContext) *InviteCommission {
	return &InviteCommission{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InviteCommission) sendCommission(register *schema.User, parentUser *schema.User) {
	commission := float64(100)
	am := &schema.AmountLog{
		UserId:        parentUser.ID,
		UAddress:      parentUser.UAddress,
		Balance:       commission,
		TargetId:      register.ID,
		TargetUid:     register.ID,
		TargetAddress: register.UAddress,
		TypeId:        schema.AmountLogTypeId1,
		Mark:          "invite commission",
	}

	err := l.svcCtx.AmountLogModel.Insert(am)
	if err != nil {
		logc.Error(l.ctx, "invite commission insert log error:"+register.UAddress)
	}

	parentUser.TotalCommission += commission
	err = l.svcCtx.UserModel.UpdateSchema(parentUser)
	if err != nil {
		logc.Error(l.ctx, "invite add parentUser commission error:"+register.UAddress)
		return
	}
}

func (l *InviteCommission) SendParentCommission(parentIds []int64, newUser *schema.User) {

	for _, v := range parentIds {
		parentUser, _ := l.svcCtx.UserModel.GetUserById(v)
		if parentUser.ID == 0 {
			continue
		}
		l.sendCommission(newUser, &parentUser)
	}
}
