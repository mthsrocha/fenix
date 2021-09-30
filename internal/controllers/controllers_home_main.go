package controllers

import (
	"net/http"

	"github.com/Fenix/internal/api"
)

func HomeTemplate(w http.ResponseWriter, r *http.Request) {
	values := api.GetAllRealTimeValueAssets(api.GetHomeAssets())
	temp.ExecuteTemplate(w, "Home", values)
}

func DigitalAssetsTemplate(w http.ResponseWriter, r *http.Request) {
	values := api.GetAllRealTimeValueAssets(api.GetDigitalAssets())
	temp.ExecuteTemplate(w, "Home", values)
}

func FanTokensTemplate(w http.ResponseWriter, r *http.Request) {
	values := api.GetAllRealTimeValueAssets(api.GetFanTokens())
	temp.ExecuteTemplate(w, "Home", values)
}

func UtilityTokensTemplate(w http.ResponseWriter, r *http.Request) {
	values := api.GetAllRealTimeValueAssets(api.GetUtilityTokens())
	temp.ExecuteTemplate(w, "Home", values)
}

func AboutUsTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Home", nil)
}
