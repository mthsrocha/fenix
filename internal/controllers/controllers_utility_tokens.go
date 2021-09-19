package controllers

import "net/http"

func UtilityTokensPricesTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeUtilityTokensPrices", nil)
}

func UtilityTokensOthersTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeUtilityTokensOthers", nil)
}

func UtilityTokensAboutTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeUtilityTokensAbout", nil)
}

func UtilityTokensCHZTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeUtilityTokensCHZ", nil)
}

func UtilityTokensWBXTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeUtilityTokensWBX", nil)
}

func UtilityTokensMCO2Template(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeUtilityTokensMCO2", nil)
}
