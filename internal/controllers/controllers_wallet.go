package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Fenix/internal/models"
)


func IndexWalletTemplate(w http.ResponseWriter, r *http.Request) {
	all_wallets := models.SelectAllWallets()
	temp.ExecuteTemplate(w, "IndexWallet", all_wallets)
}

func NewWalletTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewWallet", nil)
}

func UpdateWalletTemplate(w http.ResponseWriter, r *http.Request) {
	wallet_id := r.URL.Query().Get("id")
	get_wallet := models.SelectWallet(wallet_id)
	temp.ExecuteTemplate(w, "UpdateWallet", get_wallet)
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("user")
		crypto_tag := r.FormValue("asset")
		hash := r.FormValue("hash")
		balance := r.FormValue("balance")
		var crypto_name string
		if crypto_tag == "BTC" {
			crypto_name = "Bitcoin"
		}
		if crypto_tag == "ETH" {
			crypto_name = "Etherium"
		}
		if crypto_tag == "CHZ" {
			crypto_name = "Chiliz"
		}
		if crypto_tag == "XRP" {
			crypto_name = "Ripple"
		}

		balance_float, err := strconv.ParseFloat(balance, 64)
		if err != nil {
			log.Println(err)
		}

		models.InsertWallet(username, crypto_tag, crypto_name, hash, balance_float)
	}
	http.Redirect(w, r, "/", 301)
}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.DeleteWallet(id)
	http.Redirect(w, r, "/", 301)
	
}

func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	
}