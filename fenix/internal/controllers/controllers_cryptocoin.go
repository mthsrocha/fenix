package controllers

import (
	"net/http"
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
	temp.ExecuteTemplate(w, "HomeCriptocoinsBTC", nil)
}

func CriptocoinsBCHTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeCriptocoinsBCH", nil)
}

func CriptocoinsETHTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeCriptocoinsETH", nil)
}

func CriptocoinsLTCTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeCriptocoinsLTC", nil)
}

func CriptocoinsPAXGTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeCriptocoinsPAXG", nil)
}

func CriptocoinsXRPTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "HomeCriptocoinsXRP", nil)
}
