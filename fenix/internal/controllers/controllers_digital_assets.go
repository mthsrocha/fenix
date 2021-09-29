package controllers

import "net/http"

func DigitalAssetsPricesTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeDigitalAssetsPrices", nil)
}

func DigitalAssetsOthersTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeDigitalAssetsOthers", nil)
}

func DigitalAssetsAboutTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeDigitalAssetsAbout", nil)
}

func DigitalAssetsMBCONS01Template(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeDigitalAssetsCONS01", nil)
}

func DigitalAssetsMBPRK01Template(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeDigitalAssetsPRK01", nil)
}

func DigitalAssetsMBPRK03Template(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeDigitalAssetsPRK03", nil)
}

func DigitalAssetsMBPRK04Template(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeDigitalAssetsPRK04", nil)
}