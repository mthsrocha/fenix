package models

import (


)

type User struct {
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

func SelectUser() User {
	user := User{}

	
	return user
}

func SelectAllUsers() []User {
	
	users := []User{}

	return users
}

func UpdateUser() {

}

func DeleteUser() {
	
}