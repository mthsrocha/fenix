package models

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/Fenix/internal/database"
)

type Wallet struct {
	Id         int64
	UserName   string
	Asset_tag  string
	Asset_name string
	Hash       string
	Balance    float64
	CreatedAt  string
	UpdatedAt  string
}

func SearchAllWallets() []Wallet {
	db := database.ConnectDB()

	select_all_wallets := "select * from wallets"

	queryset, err := db.Query(select_all_wallets)
	if err != nil {
		log.Println("Error: ", err)
	}

	wallet := Wallet{}
	wallets := []Wallet{}

	for queryset.Next() {

		var id int64
		var userName, asset_tag, asset_name, hash, createdAt, updatedAt string
		var Balance float64

		err = queryset.Scan(&id, &userName, &asset_tag, &asset_name, &hash, &Balance, &createdAt, &updatedAt)
		if err != nil {
			log.Println("Error: ", err)
		}

		wallet.Id = id
		wallet.UserName = userName
		wallet.Asset_tag = asset_tag
		wallet.Balance = Balance
		wallet.Hash = hash
		wallet.CreatedAt = createdAt
		wallet.UpdatedAt = updatedAt

		wallets = append(wallets, wallet)
	}
	defer db.Close()
	return wallets
}

func InsertWallet(username string, crypto_tag string, crypto_name string, hash string, balance float64) {
	db := database.ConnectDB()

	created_at := time.Now().Format("02/01/2006 15:04:05")
	updated_at := time.Now().Format("02/01/2006 15:04:05")

	insert_script := "insert into wallets(UserName, Asset_tag, Asset_name, Hash, Balance, CreatedAt, UpdatedAt) values ($1, $2, $3, $4, $5, $6, $7)"
	insert_db, err := db.Prepare(insert_script)
	if err != nil {
		panic(err.Error())
	}

	insert_db.Exec(username, crypto_tag, crypto_name, hash, balance, created_at, updated_at)

	defer db.Close()
}


func DeleteWallet(id string) {
	db := database.ConnectDB()

	delete_script := "DELETE FROM wallets WHERE id=$1"
	delete_db, err := db.Prepare(delete_script)
	if err != nil {
		panic(err.Error())
	}

	delete_db.Exec(id)

	defer db.Close()
}

func UpdateWallet(username string, crypto_tag string, crypto_name string, hash string, balance float64, id string) {
	updateAt := time.Now().Format("02/01/2006 15:04:05")

	db := database.ConnectDB()

	update_script := "UPDATE wallets SET username=$1, asset_tag=$2, asset_name=$3, hash=$4, balance=$5, updatedat=$6 WHERE id=$7"
	update_db, err := db.Prepare(update_script)
	if err != nil {
		panic(err.Error())
	}
	
	update_db.Exec(username, crypto_tag, crypto_name, hash, balance, updateAt, id)

	defer update_db.Close()
}

func SearchWallet(id string) Wallet {
	db := database.ConnectDB()

	var userName, asset_tag, asset_name, hash, createdAt, updatedAt string
	var balance float64
	wallet := Wallet{}

	select_script := "SELECT * FROM wallets WHERE id=$1;"
	if err := db.QueryRow(select_script, id).Scan(&id, &userName, &asset_tag, 
		&asset_name, &hash, &balance, &createdAt, &updatedAt); err != nil {
			
			if err == sql.ErrNoRows {
				panic(err.Error())
			}
		}
	

	id_convert, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err)
	}

	wallet.Id = id_convert
	wallet.UserName = userName
	wallet.Asset_tag = asset_tag
	wallet.Asset_name = asset_name
	wallet.Hash = hash
	wallet.Balance = balance
	wallet.CreatedAt = createdAt
	wallet.UpdatedAt = updatedAt

	defer db.Close()
	return wallet
}