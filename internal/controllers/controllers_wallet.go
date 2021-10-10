package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Fenix/internal/models"
)

type ResponseWallet struct {
	UserId    int64
	WalletId  int64
	AssetId   int64
	Symbol    string
	AssetName  string
	Balance   float64
	Avarage   float64
	Hash      string
	CreatedAt string
	UpdatedAt string
}

func IndexWalletTemplate(w http.ResponseWriter, r *http.Request) {
	all_wallets := models.SelectAllWallets()
	temp.ExecuteTemplate(w, "IndexWallets", all_wallets)
}

func WalletInfoTemplate(w http.ResponseWriter, r *http.Request) {
	response_obj := ResponseWallet{}
	response := []ResponseWallet{}
	walletid := r.URL.Query().Get("id")
	wallet := models.SelectWallet(walletid)
	assets := models.SelectAllAssetByWalletId(walletid)

	if len(assets) > 0 {
		for _, asset := range assets {
			response_obj.AssetId = asset.Id
			response_obj.WalletId = wallet.Id
			response_obj.Symbol = asset.Symbol
			response_obj.Balance = asset.Balance
			// Alterar avarage de acordo com a cotaçao do asset
			response_obj.Avarage = asset.Balance
			response_obj.Hash = wallet.Hash
			response_obj.CreatedAt = wallet.CreatedAt
			response_obj.UpdatedAt = wallet.UpdatedAt

			response = append(response, response_obj)
		}
	} else {
		response_obj.Symbol = "Empty wallet"
		response_obj.AssetName = ""
		response_obj.Balance = 0
		response_obj.Avarage = 0
		response_obj.WalletId = wallet.Id
		response_obj.Hash = wallet.Hash
		response = append(response, response_obj)
	}

	temp.ExecuteTemplate(w, "WalletInfo", response)
}

func NewWalletTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewWallet", nil)
}

func EditWalletTemplate(w http.ResponseWriter, r *http.Request) {
	wallet_id := r.URL.Query().Get("id")
	get_wallet := models.SelectWallet(wallet_id)
	fmt.Println(get_wallet)
	temp.ExecuteTemplate(w, "EditWallet", get_wallet)
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userId := r.FormValue("userId")
		hash := r.FormValue("hash")

		user_id_int, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			log.Println(err)
		}

		models.InsertWallet(user_id_int, hash)
	}
	http.Redirect(w, r, "/user/home/", 301)
}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.DeleteWallet(id)
	http.Redirect(w, r, "/user/home/", 301)

}

func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	hash := r.FormValue("hash")

	models.UpdateWallet(hash, id)
	http.Redirect(w, r, "/user/wallets/", 301)
}

func CreateWalletByCookie(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//		session := GetSession(w, r)
		//		email := session.Values["userEmail"]
		userId := models.SelectUserId("")
		hash := r.FormValue("hash")

		models.InsertWallet(userId, hash)
	}
	http.Redirect(w, r, "/user/home", 301)
}
