package handlers

import (
	"net/http"

	"github.com/Fenix/internal/controllers"
)

func Handlers() {
		
	http.HandleFunc("/login/", controllers.LoginUserTemplate)
	http.HandleFunc("/logout/", controllers.UserLogout)
	http.HandleFunc("/login/auth/", controllers.AuteticationUser)

	http.HandleFunc("/user/home/", controllers.IndexUserTemplate)
	http.HandleFunc("/user/register/", controllers.RegisterUserTemplate)
	http.HandleFunc("/user/edit/", controllers.EditUserTemplate)
	http.HandleFunc("/user/insert/", controllers.CreateUser)
	http.HandleFunc("/user/update/", controllers.UpdateUser)
	http.HandleFunc("/user/delete/", controllers.DeleteUser)
	
	http.HandleFunc("/user/wallet/", controllers.IndexWalletTemplate)
	http.HandleFunc("/user/wallet/new/", controllers.NewWalletTemplate)
	http.HandleFunc("/user/wallet/update/", controllers.UpdateWalletTemplate)
	http.HandleFunc("/user/wallet/insert/", controllers.CreateWallet)
	http.HandleFunc("/user/wallet/delete/", controllers.DeleteWallet)

	http.HandleFunc("/user/wallet/assets/", controllers.IndexAssetsTemplate)

	http.HandleFunc("/asset/", controllers.IndexWalletTemplate)
	http.HandleFunc("/asset/new/", controllers.NewWalletTemplate)
	http.HandleFunc("/asset/update/", controllers.UpdateWalletTemplate)
	http.HandleFunc("/asset/insert/", controllers.CreateWallet)
	http.HandleFunc("/asset/delete/", controllers.DeleteWallet)

}
