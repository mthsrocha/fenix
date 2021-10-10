package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Fenix/internal/models"
)

type ResponseAssets struct {
	UserId     int64
	WalletId   int64
	AssetId    int64
	Symbol     string
	AssetName  string
	Balance    float64
	Avarage    float64
	WalletHash string
}

func IndexAssetsTemplate(w http.ResponseWriter, r *http.Request) {
	response_obj := ResponseAssets{}
	response := []ResponseAssets{}
	walletid := r.URL.Query().Get("id")
	wallet := models.SelectWallet(walletid)
	assets := models.SelectAllAssetByWalletId(walletid)

	if len(assets) > 0 {
		for _, asset := range assets {
			response_obj.AssetId = asset.Id
			response_obj.WalletId = wallet.Id
			response_obj.Symbol = asset.Symbol
			response_obj.AssetName = asset.Name
			response_obj.Balance = asset.Balance
			response_obj.WalletHash = wallet.Hash

			response = append(response, response_obj)
		}
	} else {
		response_obj.AssetId = 0
		response_obj.WalletId = wallet.Id
		response_obj.Symbol = "Empty wallet"
		response_obj.AssetName = "Empty wallet"
		response_obj.Balance = 0
		response_obj.WalletHash = wallet.Hash
		response = append(response, response_obj)
	}

	temp.ExecuteTemplate(w, "IndexAssets", assets)
}

func NewAssetTemplate(w http.ResponseWriter, r *http.Request) {
	WalletId := r.URL.Query().Get("id")
	temp.ExecuteTemplate(w, "NewAsset", WalletId)
}

func EditAssetTemplate(w http.ResponseWriter, r *http.Request) {

}

func CreateAsset(w http.ResponseWriter, r *http.Request) {
	id_wallet := r.URL.Query().Get("id")

	if r.Method == "POST" {
		walletId := id_wallet
		symbol := r.FormValue("asset")
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

	http.Redirect(w, r, "/user/wallet/assets/", 301)
}

func UpdateAsset(w http.ResponseWriter, r *http.Request) {

}

func SearchAsset(w http.ResponseWriter, r *http.Request) {

}

func SearchAllAssets(w http.ResponseWriter, r *http.Request) {

}

func DeleteAsset(w http.ResponseWriter, r *http.Request) {

}
