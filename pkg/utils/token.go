package utils

import (
	"context"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type GenTokenResponse struct {
	AccessToken  string
	AccessExpire int64
	Sign         string
}

func GenToken(accessSecret, tokenAddress string, accessExpire int64, isAdmin bool) (*GenTokenResponse, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	claims["sign"] = RandStr(8)
	if isAdmin {
		claims["admin_address"] = tokenAddress
	} else {
		claims["token_address"] = tokenAddress
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	accessToken, _ := token.SignedString([]byte(accessSecret))

	return &GenTokenResponse{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		Sign:         claims["sign"].(string),
	}, nil
}

func checkSign(ctx context.Context) {
	// sign := ctx.Value("sign").(string)
	// todo check sign
}

// GetAdminAddress admin login address
func GetAdminAddress(ctx context.Context) string {
	if ok := ctx.Value("admin_address"); ok == nil {
		return ""
	}
	return ctx.Value("admin_address").(string)
}

// GetTokenAddress user login address
func GetTokenAddress(ctx context.Context) string {
	if ok := ctx.Value("token_address"); ok == nil {
		return ""
	}
	return strings.ToLower(ctx.Value("token_address").(string))
}
