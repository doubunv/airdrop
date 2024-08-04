package utils

import (
	"math/big"
)

func ChangeAmount(req *big.Int) float64 {
	// 执行除法运算
	result := new(big.Float).Quo(new(big.Float).SetInt(req), big.NewFloat(1000000000000000000))
	// 调用 Round 函数保留两位小数
	bigFloat, _ := result.Float64()
	return bigFloat
}
