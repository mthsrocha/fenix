package handlers

import (
	"net/http"

	"github.com/Fenix/internal/controllers"
)

func UserHandlers() {

	http.HandleFunc("/user/home/", controllers.IndexUserTemplate)
	http.HandleFunc("/user/register/", controllers.RegisterUserTemplate)
	http.HandleFunc("/user/edit/", controllers.EditUserTemplate)
	http.HandleFunc("/user/insert/", controllers.CreateUser)
	http.HandleFunc("/user/update/", controllers.UpdateUser)
	http.HandleFunc("/user/delete/", controllers.DeleteUser)

	http.HandleFunc("/user/wallets/", controllers.IndexWalletTemplate)
	http.HandleFunc("/user/wallet/info/", controllers.WalletInfoTemplate)
	http.HandleFunc("/user/wallet/new/", controllers.NewWalletTemplate)
	http.HandleFunc("/user/wallet/edit/", controllers.EditWalletTemplate)
	http.HandleFunc("/user/wallet/update/", controllers.UpdateWallet)
	http.HandleFunc("/user/wallet/insert/", controllers.CreateWallet)
	http.HandleFunc("/user/wallet/delete/", controllers.DeleteWallet)


	http.HandleFunc("/user/wallet/assets/", controllers.IndexWalletTemplate)
	http.HandleFunc("/user/wallet/asset/info/", controllers.IndexWalletTemplate)
	http.HandleFunc("/user/wallet/asset/new/", controllers.NewAssetTemplate)
	http.HandleFunc("/user/wallet/asset/edit/", controllers.EditAssetTemplate)
	http.HandleFunc("/user/wallet/asset/update/", controllers.UpdateAsset)
	http.HandleFunc("/user/wallet/asset/insert/", controllers.CreateAsset)
	http.HandleFunc("/user/wallet/asset/delete/", controllers.DeleteAsset)

}
