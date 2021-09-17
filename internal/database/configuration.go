package database

import (
	"database/sql"
	"log"
)

func ConnectDB() *sql.DB {

	connection := "user=root dbname=mak password=root host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println("Error: ", err)
	}
	return db
}
