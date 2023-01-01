package controllers

import (
	user "github.com/chrishayes042/angular-golang/pkg/model"
	_ "github.com/lib/pq"
)

func GetAllUsers()(user.User){
	var  (
		id  	int
		name 	string
		pass 	string
		email 	string
	)
	database := GetSQLInfo()
	query := "SELECT * from users"
	row := QueryDB(database, query)
	for row.Next() {
		err := row.Scan(&id, &name, &pass, &email)
		if err != nil{
			panic(err)
		}

	}
	// This will only return the last row of the user table
	// TODO : Need to return a list of all the users in the table
	var u user.User
	u.UserName 	= name
	u.UserId 	= id
	u.UserEmail = email

	return u;
}