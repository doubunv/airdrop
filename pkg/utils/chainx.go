package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math/big"
)

type ChainTx struct {
}

func (ChainTx) HexToString(str interface{}) string {
	return common.HexToAddress(StringX{}.InterfaceToStr(str)).String()
}

func (ChainTx) Hex2DecStr(str interface{}) string {
	return common.HexToAddress(StringX{}.Hex2DecStr(str)).String()
}

func (ChainTx) Hex2Dec(str interface{}) uint64 {
	return StringX{}.Hex2Dec(str)
}

func (ChainTx) ToDecimal(iValue interface{}, decimals uint8) decimal.Decimal {
	value := new(big.Int)
	switch v := iValue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	case int64:
		value.SetInt64(v)
	case uint64:
		value.SetUint64(v)
	}
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}
