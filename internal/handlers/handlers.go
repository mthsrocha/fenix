package handlers

import (
	"net/http"

	"github.com/Fenix/internal/controllers"
)

func Handlers() {
		
	http.HandleFunc("/login/", controllers.LoginUser)
	http.HandleFunc("/user", controllers.IndexUserTemplate)
	http.HandleFunc("/user/register/", controllers.RegisterUserTemplate)
	http.HandleFunc("/user/edit/", controllers.EditUserTemplate)
	http.HandleFunc("/user/insert/", controllers.CreateUser)
	http.HandleFunc("/user/update/", controllers.UpdateUser)
	http.HandleFunc("/user/delete/", controllers.DeleteUser)
	
	http.HandleFunc("/", controllers.IndexWalletTemplate)
	http.HandleFunc("/wallet/new/", controllers.NewWalletTemplate)
	http.HandleFunc("/wallet/update/", controllers.UpdateWalletTemplate)
	http.HandleFunc("/wallet/insert/", controllers.CreateWallet)
	http.HandleFunc("/wallet/delete/", controllers.DeleteWallet)

	http.HandleFunc("/asset/", controllers.IndexWalletTemplate)
	http.HandleFunc("/asset/new/", controllers.NewWalletTemplate)
	http.HandleFunc("/asset/update/", controllers.UpdateWalletTemplate)
	http.HandleFunc("/asset/insert/", controllers.CreateWallet)
	http.HandleFunc("/asset/delete/", controllers.DeleteWallet)
}
