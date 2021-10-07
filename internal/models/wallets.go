package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/Fenix/internal/database"
)

type Wallet struct {
	Id        int64
	UserId    int64
	Hash      string
	CreatedAt string
	UpdatedAt string
}

func SelectAllWallets() []Wallet {
	db := database.ConnectDB()

	select_all_wallets := "select * from wallets"

	queryset, err := db.Query(select_all_wallets)
	if err != nil {
		log.Println("Error: ", err)
	}

	wallet := Wallet{}
	wallets := []Wallet{}

	for queryset.Next() {

		var id, userId int64
		var hash string
		var createdAt, updatedAt time.Time

		err = queryset.Scan(&id, &userId, &hash, &createdAt, &updatedAt)
		if err != nil {
			log.Println("Error: ", err)
		}

		wallet.Id = id
		wallet.UserId = userId
		wallet.Hash = hash
		wallet.CreatedAt = createdAt.Format("02/01/2006 15:04:05")
		wallet.UpdatedAt = updatedAt.Format("02/01/2006 15:04:05")

		wallets = append(wallets, wallet)
	}
	defer db.Close()
	return wallets
}

func InsertWallet(userId int64, hash string) {
	db := database.ConnectDB()

	created_at := time.Now()
	updated_at := time.Now()

	insert_script := "INSERT INTO wallets(UserId, Hash, CreatedAt, UpdatedAt) VALUES ($1, $2, $3, $4)"
	insert_db, err := db.Prepare(insert_script)
	if err != nil {
		log.Println(err)
	}

	insert_db.Exec(userId, hash, created_at, updated_at)

	defer db.Close()
}

func DeleteWallet(id string) {
	db := database.ConnectDB()

	delete_script := "DELETE FROM wallets WHERE id=$1"
	delete_db, err := db.Prepare(delete_script)
	if err != nil {
		log.Println(err)
	}

	delete_db.Exec(id)

	defer db.Close()
}

func UpdateWallet(hash string, id string) {
	updateAt := time.Now()

	db := database.ConnectDB()

	update_script := "UPDATE wallets SET hash=$1, updatedat=$2 WHERE id=$3"
	update_db, err := db.Prepare(update_script)
	if err != nil {
		log.Println(err)
	}

	update_db.Exec(hash, updateAt, id)

	defer update_db.Close()
}

func SelectWallet(id string) Wallet {
	db := database.ConnectDB()

	var walletId, userId int64
	var hash string
	var createdAt, updatedAt time.Time

	wallet := Wallet{}

	select_script := "SELECT * FROM wallets WHERE id=$1;"
	if err := db.QueryRow(select_script, id).Scan(&walletId, &userId,
		&hash, &createdAt, &updatedAt); err != nil {

		if err == sql.ErrNoRows {
			log.Println(err)
		}
	}

	wallet.Id = walletId
	wallet.UserId = userId
	wallet.Hash = hash
	wallet.CreatedAt = createdAt.Format("02/01/2006 15:04:05")
	wallet.UpdatedAt = updatedAt.Format("02/01/2006 15:04:05")

	defer db.Close()
	return wallet
}
