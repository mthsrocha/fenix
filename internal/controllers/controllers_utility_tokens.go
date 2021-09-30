package controllers

import (
	"net/http"

	"github.com/Fenix/internal/api"
)

func UtilityTokensPricesTemplate(w http.ResponseWriter, r *http.Request) {
	values := api.GetAllRealTimeValueAssets(api.GetUtilityTokens())
	temp.ExecuteTemplate(w, "HomeUtilityTokensPrices", values)
}

func UtilityTokensOthersTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeUtilityTokensOthers", nil)
}

func UtilityTokensAboutTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeUtilityTokensAbout", nil)
}

func UtilityTokensCHZTemplate(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("CHZ")
	temp.ExecuteTemplate(w, "HomeUtilityTokensCHZ", value)
}

func UtilityTokensWBXTemplate(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("WBX")
	temp.ExecuteTemplate(w, "HomeUtilityTokensWBX", value)
}

func UtilityTokensMCO2Template(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("MCO2")
	temp.ExecuteTemplate(w, "HomeUtilityTokensMCO2", value)
}
