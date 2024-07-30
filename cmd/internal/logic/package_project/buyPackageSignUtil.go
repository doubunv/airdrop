package package_project

import (
	"air-drop/cmd/internal/config"
	"air-drop/cmd/internal/data/schema"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/shopspring/decimal"
	"github.com/storyicon/sigverify"
	"github.com/zeromicro/go-zero/core/logc"
	"math/big"
)

const buyPackageProjectSign = `{
    "types":{
        "EIP712Domain":[
          	{
                "name":"name",
                "type":"string"
            },
            {
                "name":"version",
                "type":"string"
            },
            {
                "name":"chainId",
                "type":"uint256"
            },
            {
                "name":"verifyingContract",
                "type":"address"
            }
        ], 
        "BuyPackageProjectSign":[
            {
                "name":"_orderId",
                "type":"uint256"
            },
			{
                "name":"_spentAmount",
                "type":"uint256"
            },
			{
                "name":"_orderTime",
                "type":"uint256"
            },
			{
                "name":"_from",
                "type":"address"
            }
        ]
    },
    "domain":{
        "name":"boot",
        "version":"1",
        "chainId":"%d",
        "verifyingContract":"%s"
    },
    "primaryType":"BuyPackageProjectSign",
    "message":{
		"_orderId":"%d",
 		"_spentAmount":"%d",
 		"_orderTime":"%d",
		"_from":"%s"
    }
}`

func BuildPackageSign(config config.Config, order *schema.PackageOrder) (string, string) {
	privateKeyStr := config.ChainInfo.PrivateKey
	chainID := config.ChainInfo.ChainID
	contractAddr := config.ChainInfo.ChainBootAddress
	if privateKeyStr == "" || chainID == 0 || contractAddr == "" {
		panic("Config.ChainInfo is not defined")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", ""
	}

	usdtNum := big.NewInt(decimal.NewFromFloat(order.Amount).Mul(decimal.NewFromFloat(float64(100))).IntPart())
	usdtNum.Mul(usdtNum, big.NewInt(10).Exp(big.NewInt(10), big.NewInt(16), nil))

	data := fmt.Sprintf(buyPackageProjectSign, chainID, contractAddr, order.ID, usdtNum, order.CreatedAt, order.UAddress)
	var typedData apitypes.TypedData
	if err := json.Unmarshal([]byte(data), &typedData); err != nil {
		logc.Error(context.Background(), err.Error())
		return "", ""
	}

	_, originHash, err := sigverify.HashTypedData(typedData)
	if err != nil {
		logc.Error(context.Background(), err.Error())
		return "", ""
	}
	sig, err := crypto.Sign(originHash, privateKey)
	if err != nil {
		logc.Error(context.Background(), err.Error())
		return "", ""
	}
	// 这里最容易出问题了
	if sig[64] == 0 || sig[64] == 1 {
		sig[64] = sig[64] + 27
	}
	return hexutil.Encode(sig), fmt.Sprintf("%d", usdtNum)
}
