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

var assets = []string{
	"BITCOIN",
    "REAL",
    "LITECOIN",
    "STELLAR",
    "BCASH",
    "ETHEREUM",
    "BGOLD",
    "RIPPLE",
    "BSV",
    "USDC",
    "WIBX",
    "CHILIZ",
    "PAXGOLD",
    "LINK",
    "MCO2",
    "AAVE",
    "BAL",
    "COMP",
    "CRV",
    "DAI",
    "KNC",
    "MKR",
    "REN",
    "SNX",
    "UMA",
    "UNI",
    "YFI",
    "ZRX",
    "AXS",
    "BAT",
    "ENJ",
    "GRT",
    "MANA",
    "SUSHI",
    "MATIC",
    "OMG",
    "WBTC",
    "ANKR",
    "BAND",
    "BNT",
    "ADA",
    "PSGFT",
    "BARFT",
    "JUVFT",
    "ASRFT",
    "ATMFT",
    "GALFT",
    "CAIFT",
    "ACMFT",
    "OGFT",
    "ALLFT",
    "NAVIFT",
    "THFT",
    "PFLFT",
    "CITYFT",
    "ARGFT",
    "AMFT",
    "SAUBERFT",
    "SCCPFT",
    "YBOFT",
    "GALOFT",
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

func GetAllRealTimeValueAssets() []TickerResponse {
	ticker := TickerResponse{}
	all_tickers := []TickerResponse{}

	for _, asset := range(assets) {
		api_url := "https://www.mercadobitcoin.net/api/" + asset + "/ticker/"
		response, err := http.Get(api_url)
		CheckError(err)

		body, ok := ioutil.ReadAll(response.Body)
		CheckError(ok)

		if err := json.Unmarshal(body, &ticker); err != nil {
			log.Fatal(err)
		}
		all_tickers = append(all_tickers, ticker)
		fmt.Println(ticker.Ticker.Last)
	}

	return all_tickers
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}