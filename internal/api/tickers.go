package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	response_obj := TickerResponse{}
	api_url := "https://www.mercadobitcoin.net/api/" + asset + "/ticker/"

	response, err := http.Get(api_url)
	CheckError(err)

	body, ok := ioutil.ReadAll(response.Body)
	CheckError(ok)
	if err := json.Unmarshal(body, &response_obj); err != nil {
		log.Fatal(err)
	}
	return response_obj
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
		all_tickers = append(all_tickers, ticker)
        fmt.Println(asset)
		fmt.Println(ticker.Ticker.Last)
	}

	return all_tickers
}

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}
