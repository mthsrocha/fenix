package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
)

type TickerResponse struct {
	Ticker struct {
		High json.Number `json:"high,omitempty"`
		Low  json.Number `json:"low,omitempty"`
		Vol  json.Number `json:"vol,omitempty"`
		Last json.Number `json:"last,omitempty"`
		Buy  json.Number `json:"buy,omitempty"`
		Sell json.Number `json:"sell,omitempty"`
		Date json.Number `json:"date,omitempty"`
	} `json:"ticker"`
}

func GetRealTimeAsset(asset string) TickerResponse {
	ticker := TickerResponse{}
	api_url := "https://www.mercadobitcoin.net/api/" + asset + "/ticker/"

	response, err := http.Get(api_url)
	CheckError(err)

	body, ok := ioutil.ReadAll(response.Body)
	CheckError(ok)
	if err := json.Unmarshal(body, &ticker); err != nil {
		log.Fatal(err)
	}

	ticker.Ticker.High, err = JsonRoundConversion(ticker.Ticker.High)
	CheckError(err)
	ticker.Ticker.Low, err = JsonRoundConversion(ticker.Ticker.Low)
	CheckError(err)
	ticker.Ticker.Last, err = JsonRoundConversion(ticker.Ticker.Last)
	CheckError(err)
	ticker.Ticker.Buy, err = JsonRoundConversion(ticker.Ticker.Buy)
	CheckError(err)
	ticker.Ticker.Sell, err = JsonRoundConversion(ticker.Ticker.Sell)
	CheckError(err)

	return ticker
}

func GetAllRealTimeValueAssets(asset_list []string) []TickerResponse {
	ticker := TickerResponse{}
	all_tickers := []TickerResponse{}

	for _, asset := range asset_list {
		api_url := "https://www.mercadobitcoin.net/api/" + asset + "/ticker/"
		response, err := http.Get(api_url)
		CheckError(err)

		body, ok := ioutil.ReadAll(response.Body)
		CheckError(ok)
		if err := json.Unmarshal(body, &ticker); err != nil {
			log.Fatal(err)
		}

		ticker.Ticker.Last, err = JsonRoundConversion(ticker.Ticker.Last)
		CheckError(err)

		all_tickers = append(all_tickers, ticker)
	}

	return all_tickers
}

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func JsonRoundConversion(ticker_value json.Number) (json.Number, error) {
	conv, err := ticker_value.Float64()
	if err != nil {
		log.Println(err)
		return ticker_value, err
	}

	conv_str := strconv.FormatFloat((math.Round(conv*100) / 100), 'f', 2, 64)
	conv_json := json.Number(conv_str)
	return conv_json, nil
}
