package controllers

import (
	"net/http"

	"github.com/Fenix/internal/api"
)

func DigitalAssetsPricesTemplate(w http.ResponseWriter, r *http.Request) {
	values := api.GetAllRealTimeValueAssets(api.GetDigitalAssets())
	temp.ExecuteTemplate(w, "HomeDigitalAssetsPrices", values)
}

func DigitalAssetsOthersTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeDigitalAssetsOthers", nil)
}

func DigitalAssetsAboutTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeDigitalAssetsAbout", nil)
}

func DigitalAssetsMBCONS01Template(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("MBCONS01")
	temp.ExecuteTemplate(w, "HomeDigitalAssetsCONS01", value)
}

func DigitalAssetsMBPRK01Template(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("MBPRK01")
	temp.ExecuteTemplate(w, "HomeDigitalAssetsPRK01", value)
}

func DigitalAssetsMBPRK03Template(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("MBPRK03")
	temp.ExecuteTemplate(w, "HomeDigitalAssetsPRK03", value)
}

func DigitalAssetsMBPRK04Template(w http.ResponseWriter, r *http.Request) {
	value := api.GetRealTimeAsset("MBPRK04")
	temp.ExecuteTemplate(w, "HomeDigitalAssetsPRK04", value)
}
