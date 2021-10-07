package controllers

import (
	"net/http"

	"github.com/Fenix/internal/models"
)

func IndexUserTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "IndexUserHome", nil)
}

func RegisterUserTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "RegisterUser", nil)

}
func LoginUserTemplate(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Login", nil)
}

func EditUserTemplate(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("inputRegisterEmail")
		passw := r.FormValue("inputRegisterPassword")
		name := r.FormValue("inputRegisterName")
		cpfcnpj := r.FormValue("inputRegisterCPFCNPJ")
		gender := r.FormValue("inputRegisterGender")
		address := r.FormValue("inputRegisterAddress")
		address_complement := r.FormValue("inputRegisterAddressComplement")
		city := r.FormValue("inputRegisterCity")
		cep := r.FormValue("inputRegisterCEP")
		country := r.FormValue("inputRegisterState")

		models.InsertUser(email, passw, name, cpfcnpj, gender, address, address_complement, city, cep, country)
	}

	http.Redirect(w, r, "/login/", 301)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func SearchUser(w http.ResponseWriter, r *http.Request) {

}

func SearchAllUsers(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
