package models

import (
	"database/sql"

	"github.com/Fenix/internal/database"
)

type Auth struct {
	id     int64
	UserId int64
	SessionId string
}

func AuthenticationUser(email string, passw string) bool {
	db := database.ConnectDB()
	var email_db, pass_db string
	select_user := "SELECT email, passw FROM users WHERE email=$1"
	if err := db.QueryRow(select_user, email).Scan(&email_db, &pass_db); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
	}

	if email == email_db && passw == pass_db {
		return true
	}
	defer db.Close()
	return false
}

func InsertAuthentication(userid int64, sessionid string) {
	
}

func ReadAuthentication() {
	
}

func UpdateAuthentication() {
	
}

func DeleteAuthentication() {
	
}