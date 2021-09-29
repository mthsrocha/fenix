package controllers

import (
	"net/http"

	"github.com/Fenix/internal/api"
)

func HomeTemplate(w http.ResponseWriter, r *http.Request) {
	value := api.GetAllRealTimeValueAssets()
	temp.ExecuteTemplate(w, "Home", value)
}

func DigitalAssetsTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Home", nil)
}

func FanTokensTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Home", nil)
}

func UtilityTokensTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Home", nil)
}

func AboutUsTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Home", nil)
}
