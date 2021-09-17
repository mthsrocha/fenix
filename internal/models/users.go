package models

import (


)

type Users struct {
	Id			int64
	Email 		string
	password	string
	Name		string
	Cpfcnpj		string
	Address		string
	Gender		string
	City		string
	State 		string
}


func InsertUser() {

}

func SelectAllUsers()