package handlers

import (
	"net/http"

	"github.com/Fenix/internal/controllers"
)

func AuthHandlers() {

	http.HandleFunc("/login/", controllers.LoginUserTemplate)
	http.HandleFunc("/logout/", controllers.UserLogout)
	http.HandleFunc("/login/auth/", controllers.AuteticationUser)

}