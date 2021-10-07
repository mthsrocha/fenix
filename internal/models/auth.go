package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/Fenix/internal/database"
)

type Auth struct {
	id        int64
	UserId    int64
	SessionId string
	Expires   time.Time
}

func AuthenticationUser(email string, passw string) bool {
	db := database.ConnectDB()

	var email_db, pass_db string

	select_user_script := "SELECT email, passw FROM users WHERE email=$1"
	if err := db.QueryRow(select_user_script, email).Scan(&email_db, &pass_db); err != nil {
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

func InsertAuth(userid int64, sessionid string, expires time.Time) {
	db := database.ConnectDB()

	insert_auth_script := "INSERT INTO auth(userId, sessionId, expires) VALUES($1, $2, $3)"
	insert_auth_db, err := db.Prepare(insert_auth_script)
	if err != nil {
		log.Println(err)
	}

	insert_auth_db.Exec(userid, sessionid, expires)

	defer db.Close()
}

func ReadAuth(id int64) Auth {
	db := database.ConnectDB()

	auth_obj := Auth{}
	var Id, userid int64
	var sessionid string
	var expires time.Time

	select_auth_script := "SELECT * FROM auth WHERE id=$1"
	if err := db.QueryRow(select_auth_script, id).Scan(&Id, &userid, &sessionid, &expires); err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
		}
	}

	auth_obj.id = Id
	auth_obj.SessionId = sessionid
	auth_obj.UserId = userid
	auth_obj.Expires = expires

	defer db.Close()
	return auth_obj
}

func UpdateSessionIdAuth(sessionid string, id int64) {
	db := database.ConnectDB()

	update_auth_script := "UPDATE auth SET sessionid=$1 WHERE id=$2"
	update_auth_db, err := db.Prepare(update_auth_script)
	if err != nil {
		log.Println(err)
	}

	update_auth_db.Exec(sessionid, id)
	defer db.Close()
}

func UpdateExpiresAuth(expires string, id int64) {
	db := database.ConnectDB()

	update_auth_script := "UPDATE auth SET expires=$1 WHERE id=$2"
	update_auth_db, err := db.Prepare(update_auth_script)
	if err != nil {
		log.Println(err)
	}

	update_auth_db.Exec(expires, id)
	defer db.Close()
}

func DeleteAuth(id int64) {
	db := database.ConnectDB()

	delete_auth_script := "DELETE FROM auth WHERE id=$1"
	delete_auth_db, err := db.Prepare(delete_auth_script)
	if err != nil {
		log.Println(err)
	}

	delete_auth_db.Exec(id)
	defer db.Close()
}
