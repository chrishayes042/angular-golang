package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host ="localhost"
	port =5432
	user ="postgres"
	password ="1234"
	dbname ="myDB"
)

var (
	id int
	userName string
	userPass string
	userEmail string
)




func main(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// connect to the DB
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// query
	rows, err := db.Query("select * from users")
	if err != nil{
		panic(err)
	}
	// scan each row
	for rows.Next(){
		err := rows.Scan(&id, &userName, &userPass, &userEmail)
		if err != nil{
			panic(err)
		}

		// print
		fmt.Println(id, userName, userEmail)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	fmt.Println("Established a successfull connection!")
}
