package service

import (
	"air-drop/cmd/internal/data/schema"
	"air-drop/cmd/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"
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

func (l *InviteCommission) sendCommission(register *schema.User, parentUser *schema.User, commission float64) {
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

func (l *InviteCommission) SendParentCommission(newUser *schema.User, amount float64) {
	var LevelCommission = make(map[int]float64)

	atoi, err := strconv.Atoi(l.svcCtx.SettingModel.FindByKey(schema.Invite_one_level_commission))
	if err != nil {
		return
	}
	LevelCommission[1] = float64(atoi) / 100

	atoi, err = strconv.Atoi(l.svcCtx.SettingModel.FindByKey(schema.Invite_two_level_commission))
	if err != nil {
		return
	}
	LevelCommission[2] = float64(atoi) / 100

	atoi, err = strconv.Atoi(l.svcCtx.SettingModel.FindByKey(schema.Invite_three_level_commission))
	if err != nil {
		return
	}
	LevelCommission[3] = float64(atoi) / 100

	parentIds := newUser.Path.GetParentIdByNumber(string(newUser.Path), 3)
	for i, v := range parentIds {
		commission := amount * LevelCommission[i] //todo
		parentUser, _ := l.svcCtx.UserModel.GetUserById(v)
		if parentUser.ID == 0 {
			continue
		}
		l.sendCommission(newUser, &parentUser, commission)
	}
}
