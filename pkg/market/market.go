package market

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Prices struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data struct {
		PORT3 struct {
			Id                int         `json:"id"`
			Name              string      `json:"name"`
			Symbol            string      `json:"symbol"`
			Slug              string      `json:"slug"`
			NumMarketPairs    int         `json:"num_market_pairs"`
			DateAdded         time.Time   `json:"date_added"`
			Tags              []string    `json:"tags"`
			MaxSupply         interface{} `json:"max_supply"`
			CirculatingSupply int         `json:"circulating_supply"`
			TotalSupply       int         `json:"total_supply"`
			Platform          struct {
				Id           int    `json:"id"`
				Name         string `json:"name"`
				Symbol       string `json:"symbol"`
				Slug         string `json:"slug"`
				TokenAddress string `json:"token_address"`
			} `json:"platform"`
			IsActive                      int         `json:"is_active"`
			InfiniteSupply                bool        `json:"infinite_supply"`
			CmcRank                       int         `json:"cmc_rank"`
			IsFiat                        int         `json:"is_fiat"`
			SelfReportedCirculatingSupply int         `json:"self_reported_circulating_supply"`
			SelfReportedMarketCap         float64     `json:"self_reported_market_cap"`
			TvlRatio                      interface{} `json:"tvl_ratio"`
			LastUpdated                   time.Time   `json:"last_updated"`
			Quote                         struct {
				USD struct {
					Price                 float64     `json:"price"`
					Volume24H             float64     `json:"volume_24h"`
					VolumeChange24H       float64     `json:"volume_change_24h"`
					PercentChange1H       float64     `json:"percent_change_1h"`
					PercentChange24H      float64     `json:"percent_change_24h"`
					PercentChange7D       float64     `json:"percent_change_7d"`
					PercentChange30D      float64     `json:"percent_change_30d"`
					PercentChange60D      float64     `json:"percent_change_60d"`
					PercentChange90D      float64     `json:"percent_change_90d"`
					MarketCap             int         `json:"market_cap"`
					MarketCapDominance    int         `json:"market_cap_dominance"`
					FullyDilutedMarketCap int         `json:"fully_diluted_market_cap"`
					Tvl                   interface{} `json:"tvl"`
					LastUpdated           time.Time   `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"PORT3"`
	} `json:"data"`
}

func Quotes() (res float64, err error) {
	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=PORT3&convert=USD"
	method := "GET"
	//client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-CMC_PRO_API_KEY", "6e49b979-78af-4d95-88f6-ac5edf98732b")
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
	var prices Prices
	err = json.Unmarshal(body, &prices)
	return prices.Data.PORT3.Quote.USD.Price, err
}
