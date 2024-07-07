package utils

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
	"github.com/storyicon/sigverify"
	"log"
	"math/big"
	"strings"
)

func BuildSignature(privateKeyStr string, typedDataBytes []byte) (string, error) {

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", err
	}

	hash := crypto.Keccak256Hash(typedDataBytes,
		common.HexToAddress("0xe085cbEA651bAAB36aFb338c65ae8744D2579A74").Bytes(),
		common.LeftPadBytes(big.NewInt(1).Bytes(), 32),
		common.LeftPadBytes(big.NewInt(1).Bytes(), 32),
		common.LeftPadBytes(big.NewInt(1).Bytes(), 32))
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureStr := hexutil.Encode(signature)
	return signatureStr, nil
}

func GetMsgHash(privateKeyStr string, typedDataBytes []byte) string {
	hash := crypto.Keccak256Hash(typedDataBytes)
	return hash.Hex() // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8
}

func Verify2(privateKeyStr string, typedDataBytes []byte) bool {

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Printf("publicKeyBytes:%s", hexutil.Encode(publicKeyBytes))

	hash := crypto.Keccak256Hash(typedDataBytes)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sigPublicKeyECDSA:%s\n", crypto.PubkeyToAddress(*sigPublicKeyECDSA).Hex())

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Printf("sigPublicKeyBytes:%s\n", string(sigPublicKeyBytes))

	signatureNoRecoverID := signature[:len(signature)-1]
	// remove recovery id
	fmt.Printf("signatureNoRecoverID:%s", hexutil.Encode(signatureNoRecoverID))
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	return verified
}

func VerifyLoginAddress(chainId int64, timestamp int64, nonce, address, signer, signature string) (bool, error) {
	if !strings.HasPrefix(signature, "0x") {
		signature = "0x" + signature
	}
	_signature, err := hexutil.Decode(signature)
	if err != nil {
		return false, errors.WithStack(err)
	}
	data := fmt.Sprintf(LoginTypedData, chainId, timestamp, address, nonce)
	var typedData apitypes.TypedData
	if err := json.Unmarshal([]byte(data), &typedData); err != nil {
		return false, errors.WithStack(err)
	}
	verify, err := sigverify.VerifyTypedDataSignatureEx(
		common.HexToAddress(signer),
		typedData,
		_signature,
	)
	if err != nil || !verify {
		fmt.Errorf("Verify signature error: %v\n", err)
		return false, errors.New("Verify signature failed")
	}
	return true, nil
}
