package controllers

import (
	"net/http"

	"github.com/Fenix/internal/api"
)

func FanTokensPricesTemplate(w http.ResponseWriter, r *http.Request) {
	values := api.GetAllRealTimeValueAssets(api.GetFanTokens())
	temp.ExecuteTemplate(w, "HomeFanTokensPrices", values)
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
