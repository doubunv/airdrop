package event

import (
	"air-drop/cmd/internal/data/model"
	"air-drop/cmd/internal/service"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
)

type EventLinkData struct {
	FormAddress string   `json:"_form"`
	ToAddress   string   `json:"_to"`
	OrderId     int64    `json:"_orderId"`
	Amount      *big.Int `json:"_spend"`
}

func (l *EventLogic) BuyLinkProject(chainTx *model.ChainTx) (err error) {
	var eventData EventLinkData
	err = json.Unmarshal([]byte(chainTx.Data), &eventData)
	if err != nil {
		return errors.New(fmt.Sprintf("data Unmarshal err %v", err))
	}

	order, _ := l.svcCtx.LinkOrderModel.FindById(eventData.OrderId)
	if order.ID == 0 {
		return errors.New("order not found")
	}

	usdtNum := big.NewInt(decimal.NewFromFloat(order.BuyAmount).Mul(decimal.NewFromFloat(float64(100))).IntPart())
	usdtNum.Mul(usdtNum, big.NewInt(10).Exp(big.NewInt(10), big.NewInt(16), nil))

	if eventData.Amount.Cmp(usdtNum) == -1 {
		return errors.New("pay amount error")
	}

	order.PayStatus = 2
	err = l.svcCtx.LinkOrderModel.UpdateSchema(&order)
	if err != nil {
		return err
	}

	userInfo, err := l.svcCtx.UserModel.GetUserById(order.UserId)
	if err != nil {
		return err
	}

	userInfo.PayAmount += order.BuyAmount
	err = l.svcCtx.UserModel.UpdateSchema(&userInfo)
	if err != nil {
		return err
	}
	service.NewInviteCommission(l.ctx, l.svcCtx).SendParentCommission(&userInfo, order.BuyAmount)

	return nil
}
