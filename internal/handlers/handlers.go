package handlers

import (
	"net/http"

	"github.com/Fenix/internal/controllers"
)

func Handlers() {
	
	http.HandleFunc("/home", controllers.RegisterUserTemplate)


	http.HandleFunc("/user", controllers.IndexUserTemplate)
	http.HandleFunc("/user/login/", controllers.LoginUser)
	http.HandleFunc("/user/register/", controllers.RegisterUserTemplate)
	http.HandleFunc("/user/insert/", controllers.RegisterUserTemplate)
	
	http.HandleFunc("/", controllers.IndexWalletTemplate)
	http.HandleFunc("/wallet/new/", controllers.NewWalletTemplate)
	http.HandleFunc("/wallet/update/", controllers.UpdateWalletTemplate)
	http.HandleFunc("/wallet/insert/", controllers.CreateWallet)
	http.HandleFunc("/wallet/delete/", controllers.DeleteWallet)
}
