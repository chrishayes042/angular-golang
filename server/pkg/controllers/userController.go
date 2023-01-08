package controllers

import (
	"fmt"

	user "github.com/chrishayes042/angular-golang/pkg/model"
	_ "github.com/lib/pq"
)

// func to get all users from the sql table return them as a list
func GetAllUsers() []user.User {
	// create new list of Users
	var userList []user.User

	var (
		id    int
		name  string
		pass  string
		email string
	)
	// get the database info
	database := GetSQLInfo()
	// query string
	query := "SELECT * from users"
	row := QueryDB(database, query)
	for row.Next() {
		err := row.Scan(&id, &name, &pass, &email)
		if err != nil {
			panic(err)
		}

		var u user.User
		u.UserName = name
		u.UserId = id
		u.UserEmail = email
		fmt.Println("adding user" + u.UserName)
		// user.GetUserList(u)

		userList = append(userList, u)
	}

	return userList
}

// func addUser(string, string, string) {
// 	database := GetSQLInfo()

// 	query := "INSERT INTO users" +
// 		" VALUES( ?,?,?,? );"
// }
