package controllers

import (
	_ "github.com/lib/pq"
	user "github.com/chrishayes042/angular-golang/pkg/model"
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
	row.Next(){
		err := row.Scan(&id, &name, &pass, &email)
		if err != nil{
			panic(err)
		}
	}
	var u user.User
	u.UserName 	= name
	u.UserId 	= id
	u.UserEmail = email

	return u;
}