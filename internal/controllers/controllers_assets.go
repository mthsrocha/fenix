package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Fenix/internal/models"
)

func IndexAssetsTemplate(w http.ResponseWriter, r *http.Request) {
	wallet_id := r.URL.Query().Get("walletid")
	assets := models.SelectAllAssetByWalletId(wallet_id)
	temp.ExecuteTemplate(w, "IndexAssets", assets)
}

func NewAssetTemplate(w http.ResponseWriter, r *http.Request) {
	WalletId := r.URL.Query().Get("walletid")
	temp.ExecuteTemplate(w, "NewAsset", WalletId)
}

func EditAssetTemplate(w http.ResponseWriter, r *http.Request) {

}

func CreateAsset(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		walletId := r.URL.Query().Get("WalletId")
		symbol := r.FormValue("symbol")
		name := r.FormValue("name")
		balance := r.FormValue("balance")

		wallet_id_int, err := strconv.ParseInt(walletId, 10, 64)
		if err != nil {
			log.Println(err)
		}

		balance_float, err := strconv.ParseFloat(balance, 64)
		if err != nil {
			log.Println(err)
		}

		models.InsertAsset(wallet_id_int, symbol, name, balance_float)
	}

	http.Redirect(w, r, "/user/home", 301)
}

func UpdateAsset(w http.ResponseWriter, r *http.Request) {

}

func SearchAsset(w http.ResponseWriter, r *http.Request) {

}

func SearchAllAssets(w http.ResponseWriter, r *http.Request) {

}
