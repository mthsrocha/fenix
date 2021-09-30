package controllers

import (
	"net/http"

	"github.com/Fenix/internal/api"
)

func CriptocoinsPricesTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeCriptocoinsPrices", nil)
}

func CriptocoinsOthersTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeCriptocoinsOthers", nil)
}

func CriptocoinsAboutTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeCriptocoinsAbout", nil)
}

func CriptocoinsBTCTemplate(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("BTC")
	temp.ExecuteTemplate(w, "HomeCriptocoinsBTC", value)
}

func CriptocoinsBCHTemplate(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("BCH")
	temp.ExecuteTemplate(w, "HomeCriptocoinsBCH", value)
}

func CriptocoinsETHTemplate(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("ETH")
	temp.ExecuteTemplate(w, "HomeCriptocoinsETH", value)
}

func CriptocoinsLTCTemplate(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("LTC")
	temp.ExecuteTemplate(w, "HomeCriptocoinsLTC", value)
}

func CriptocoinsPAXGTemplate(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("PAXG")
	temp.ExecuteTemplate(w, "HomeCriptocoinsPAXG", value)
}

func CriptocoinsXRPTemplate(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("XRP")
	temp.ExecuteTemplate(w, "HomeCriptocoinsXRP", value)
}
