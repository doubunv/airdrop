package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestBuildSignature(t *testing.T) {

	signature, err := BuildSignature("26343eafe1ef03f3f7e8749b91339a46876741656b09286ae221e8e764b061f6", []byte("123"))
	if err != nil {
		return
	}
	fmt.Println()
	fmt.Println(GetMsgHash("26343eafe1ef03f3f7e8749b91339a46876741656b09286ae221e8e764b061f6", []byte("123")))
	fmt.Println(signature)

	// 80be8bfb6648ca137326b3ba911473dd9e007d5dd4b65129cd7ccb2287edab635e546c0818f649810ab55c8d6ac7d6b0f82c101dffe9daf253ac20ce9ea4bf6800
	// 0xb42ca4636f721c7a331923e764587e98ec577cea1a185f60dfcc14dbb9bd900b
	// 0x390d704d7ab732ce034203599ee93dd5d3cb0d4d1d7c600ac11726659489773d559b12d220f99f41d17651b0c1c6a669d346a397f8541760d6b32a5725378b241c
}

func TestBuildSignature2(t *testing.T) {
	print(Verify2("26343eafe1ef03f3f7e8749b91339a46876741656b09286ae221e8e764b061f6", []byte("as10scasc55511da")))
}

func TestVerifyLoginAddress(t *testing.T) {
	chainId := 97
	timestamp := 1709619217
	address := "0xdEe465Ff2a3896F406b78a499DDC64196F2e0558"
	signer := "0xdEe465Ff2a3896F406b78a499DDC64196F2e0558"
	signature := "0x6dd0c82e727252445481cc0c2cb021b8ffc510b15cfab9a9d4d1d114b41e11ba41283760164c9a1918ba7f938cc1dcec279b524072451eebbef81ffb0de819091b"
	nonce := "123"

	loginAddress, err := VerifyLoginAddress(int64(chainId), int64(timestamp), nonce, address, signer, signature)
	fmt.Println("loginAddress", loginAddress)
	if err != nil {
		return
	}
}

func TestGenToken(t *testing.T) {
	//token, err := GenToken("1234abcd", "0x273b32e2ef27273597bd553e77ef0b05fd7e0e0f", 60*60*360, false)
	token, err := GenToken("1234abcd", "0x273b32e2ef27273597bd553e77ef0b05fd7e0e0f", 60*60*360, true)
	if err != nil {
		return
	}
	//token &{eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEwOTg5NzAsImlhdCI6MTcwOTgwMjk3MCwic2lnbiI6IlphZThOUnc4IiwidG9rZW5fYWRkcmVzcyI6IjB4MjczYjMyZTJlZjI3MjczNTk3YmQ1NTNlNzdlZjBiMDVmZDdlMGUwZiJ9.8ZmgK-zQVPGgM8x7Ma6NlOP2tFvTt2qFJk1CdrbpWKg 1711098970 Zae8NRw8}
	fmt.Println("token", token)
}

func TestAa(t *testing.T) {

	p3Price := 0.08234
	value := 100

	nftBoxP3Price := (1 / p3Price) * float64(value)
	PayP3 := decimal.NewFromFloat(0.8).Mul(decimal.NewFromFloat(nftBoxP3Price)).InexactFloat64()
	PayScore := decimal.NewFromFloat(0.2).Mul(decimal.NewFromFloat(float64(value))).IntPart()
	fmt.Println(nftBoxP3Price)
	fmt.Println(PayP3)
	fmt.Println(PayScore)
}
