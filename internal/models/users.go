package models

import (
	"database/sql"
	"log"

	"github.com/Fenix/internal/database"
)

type User struct {
	Id         int64
	Email      string
	Passw      string
	UserName   string
	Cpfcnpj    string
	Address    string
	Complement string
	Gender     string
	City       string
	Country    string
	Cep        string
}

func InsertUser(email string, passw string, name string, cpfcnpj string, gender string, address string, address_complement string, city string, cep string, country string) {
	db := database.ConnectDB()

	insert_user_script := "INSERT INTO users(email, passw, userName, cpfcnpj, gender, address, address_complement, city, cep, country) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	insert_user_db, err := db.Prepare(insert_user_script)
	if err != nil {
		log.Println(err)
	}

	insert_user_db.Exec(email, passw, name, cpfcnpj, gender, address, address_complement, city, cep, country)

	defer db.Close()
}

func SelectUser(email string) User {
	db := database.ConnectDB()

	var passw, name, cpfcnpj, gender, address, address_complement, city, cep, country string
	user := User{}

	select_user_script := "SELECT * FROM users WHERE email=$1"
	if err := db.QueryRow(select_user_script, email).Scan(&email, &passw, &name,
		&cpfcnpj, &gender, &address, &address_complement, &city, &cep, &country); err != nil {

		if err == sql.ErrNoRows {
			log.Println(err)
		}
	}

	user.Email = email
	user.Passw = passw
	user.UserName = name
	user.Cpfcnpj = cpfcnpj
	user.Gender = gender
	user.Address = address
	user.Complement = address_complement
	user.City = city
	user.Cep = cep
	user.Country = country

	defer db.Close()
	return user
}

func DeleteUser(email string) {
	db := database.ConnectDB()

	delete_user_script := "DELETE FROM users WHERE email=$1"
	delete_user_db, err := db.Prepare(delete_user_script)
	if err != nil {
		log.Println(err)
	}

	delete_user_db.Exec(email)

	defer db.Close()
}

func SelectUserId(email string) int64 {
	db := database.ConnectDB()

	var id int64

	select_script := "SELECT id FROM users WHERE email=$1"
	if err := db.QueryRow(select_script, email).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
		}
	}
	
	defer db.Close()
	return id
}

func SelectAllUsers() []User {
	db := database.ConnectDB()
	users := []User{}

	defer db.Close()
	return users
}

func UpdateUser() {

}
