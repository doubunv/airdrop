package market

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Prices2 struct {
	Coins struct {
		Ethereum0Xb4357054C3DA8D46ED642383F03139AC7F090343 struct {
			Decimals   int     `json:"decimals"`
			Symbol     string  `json:"symbol"`
			Price      float64 `json:"price"`
			Timestamp  int     `json:"timestamp"`
			Confidence float64 `json:"confidence"`
		} `json:"ethereum:0xb4357054c3dA8D46eD642383F03139aC7f090343"`
	} `json:"coins"`
}

func QuotesAll() (res float64, err error) {
	if val, err := Quotes2(); err == nil {
		return val, err
	}
	return Quotes()
}

func Quotes2() (res float64, err error) {
	url := "https://coins.llama.fi/prices/current/ethereum:0xb4357054c3dA8D46eD642383F03139aC7f090343"
	method := "GET"
	//client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resResp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resResp.Body.Close()

	body, err := io.ReadAll(resResp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var prices Prices2
	err = json.Unmarshal(body, &prices)
	return prices.Coins.Ethereum0Xb4357054C3DA8D46ED642383F03139AC7F090343.Price, err
}
