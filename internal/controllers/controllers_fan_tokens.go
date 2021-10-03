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
	values := api.GetAllRealTimeValueAssets(api.GetFanTokensFutebol())
	temp.ExecuteTemplate(w, "HomeFanTokensFutebol", values)
}

func FanTokensGamingTemplate(w http.ResponseWriter, r *http.Request) {
	values := api.GetAllRealTimeValueAssets(api.GetFanTokensGaming())
	temp.ExecuteTemplate(w, "HomeFanTokensGaming", values)
}

func FanTokensMotorsportsTemplate(w http.ResponseWriter, r *http.Request) {
	values := api.GetAllRealTimeValueAssets(api.GetFanTokensMotorsports())
	temp.ExecuteTemplate(w, "HomeFanTokensMotorsports", values)
}
