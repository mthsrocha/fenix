package controllers

import (
	"net/http"

	"github.com/Fenix/internal/models"
)

func IndexAssetsTemplate(w http.ResponseWriter, r *http.Request) {
	wallet_id := r.URL.Query().Get("walletid")
	assets := models.SelectAllAssetByWalletId(wallet_id)
	temp.ExecuteTemplate(w, "", assets)
}

func RegisterAssetTemplate(w http.ResponseWriter, r *http.Request) {

}

func CreateAsset(w http.ResponseWriter, r *http.Request) {

}

func UpdateAsset(w http.ResponseWriter, r *http.Request) {

}

func SearchAsset(w http.ResponseWriter, r *http.Request) {

}

func SearchAllAssets(w http.ResponseWriter, r *http.Request) {

}
