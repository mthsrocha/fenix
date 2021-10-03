package models

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/Fenix/internal/database"
)

type Asset struct {
	Id       int64
	WalletId int64
	Tag      string
	Name     string
	Balance  float64
}

func InsertAsset(walletId int64, tag string, name string, balance float64) {
	db := database.ConnectDB()

	insert_asset_script := "INSERT INTO assets(walletid, tag, name, balance) VALUES($1, $2, $3, $4)"
	insert_asset_db, err := db.Prepare(insert_asset_script)
	if err != nil {
		log.Println(err)
	}

	insert_asset_db.Exec(tag, name, balance)
}

func SelectAssetById(id string) Asset {
	db := database.ConnectDB()

	var tag, name string
	var balance float64
	asset := Asset{}

	select_asset_script := "SELECT * FROM assets WHERE id=$1"
	if err := db.QueryRow(select_asset_script, id).Scan(&id, &tag,
		&name, &balance); err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
		}
	}

	id_converted, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err)
	}

	asset.Id = id_converted
	asset.Tag = tag
	asset.Name = name
	asset.Balance = balance

	defer db.Close()

	return asset
}

func SelectAssetByTag(tag string) Asset {
	db := database.ConnectDB()

	var id int64
	var name string
	var balance float64
	asset := Asset{}

	select_asset_script := "SELECT * FROM assets WHERE tag=$1"
	if err := db.QueryRow(select_asset_script, tag).Scan(&id, &tag,
		&name, &balance); err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
		}
	}

	asset.Id = id
	asset.Tag = tag
	asset.Name = name
	asset.Balance = balance

	defer db.Close()
	return asset
}

func SelectAllAssetByWalletId(id string) []Asset {
	db := database.ConnectDB()

	assets := []Asset{}

	select_assets_script := "SELECT * FROM assets WHERE WalletId=$1"
	assets_queryset, err := db.Query(select_assets_script, id)
	if err != nil {
		log.Println(err)
	}

	for assets_queryset.Next() {
		var walletId int64
		var tag, name string
		var balance float64
		asset := Asset{}

		err := assets_queryset.Scan(&id, &walletId, &tag, &name, &balance)
		if err == sql.ErrNoRows {
			log.Println(err)
		}

		id_converted, err := strconv.ParseInt(id, 10, 64)
		asset.Id = id_converted
		asset.WalletId = walletId
		asset.Tag = tag
		asset.Name = name
		asset.Balance = balance

		assets = append(assets, asset)
	}

	defer db.Close()
	return assets
}

func DeleteAsset(walletId string) {
	db := database.ConnectDB()

	delete_script := "DELETE FROM assets WHERE WalletId=$1"
	delete_asset_db, err := db.Prepare(delete_script)
	if err != nil {
		log.Println(err)
	}

	delete_asset_db.Exec(walletId)

	defer db.Close()
}

func UpdateAsset(id string, tag string, name string, balance float64) {
	db := database.ConnectDB()

	update_asset_script := "UPDATE assets SET tag=$1, name=$2, balance=$3"
	update_asset_db, err := db.Prepare(update_asset_script)
	if err != nil {
		log.Println(err)
	}

	update_asset_db.Exec(tag, name, balance)

	defer db.Close()
}
