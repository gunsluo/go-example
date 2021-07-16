package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	rate, err := QueryExchangeRate("GBP", "USD")
	if err != nil {
		panic(err)
	}

	fmt.Println(rate)
}

type exchangeRateRequest struct {
	Amount        float64    `json:"amount"`
	From          string     `json:"from"`
	To            string     `json:"to"`
	DecimalPoints int        `json:"decimalPoints"`
	CacheTime     string     `json:"cacheTime"`
	Exchanger     []exchange `json:"exchanger"`
}

type exchange struct {
	Name string `json:"name"`
}

type exchangeRateResposne struct {
	From            string  `json:"from"`
	To              string  `json:"to"`
	ExchangerName   string  `json:"exchangerName"`
	ExchangeValue   float64 `json:"exchangeValue"`
	OriginalAmount  float64 `json:"originalAmount"`
	ConvertedAmount float64 `json:"convertedAmount"`
	ConvertedText   string  `json:"convertedText"`
	RateDateTime    string  `json:"rateDateTime"`
	RateFromCache   bool    `json:"rateFromCache"`
}

func QueryExchangeRate(from, to string) (float64, error) {
	url := "https://go-swap-server.herokuapp.com/convert"

	req := &exchangeRateRequest{
		Amount:        1,
		From:          from,
		To:            to,
		DecimalPoints: 4,
		CacheTime:     "120s",
		Exchanger: []exchange{
			{Name: "yahoo"},
			{Name: "google"},
		},
	}

	query, err := json.Marshal(req)
	if err != nil {
		return 0, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(query))
	if err != nil {
		return 0, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return 0, err
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return 0, err
	}

	var resp exchangeRateResposne
	if err := json.Unmarshal(body, &resp); err != nil {
		return 0, err
	}

	return resp.ConvertedAmount, nil
}
