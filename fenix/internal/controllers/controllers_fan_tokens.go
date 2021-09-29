package controllers

import "net/http"

func FanTokensPricesTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeFanTokensPrices", nil)
}

func FanTokensOthersTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeFanTokensOthers", nil)
}

func FanTokensAboutTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeFanTokensAbout", nil)
}

func FanTokensFutebolTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeFanTokensFutebol", nil)
}

func FanTokensGamingTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeFanTokensGaming", nil)
}

func FanTokensMotorsportsTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeFanTokensMotorsports", nil)
}
