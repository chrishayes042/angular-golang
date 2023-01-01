package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetSQLInfo()(*sql.DB){
	const (
		host ="localhost"
		port =5432
		user ="postgres"
		password ="1234"
		dbname ="myDB"
	)
	

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
	"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	
	// connect to the DB
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	panic(err)
	}
	return db;
}

func QueryDB(db *sql.DB, query string)(*sql.Rows){
	
	rows, err := db.Query(query)
	if err != nil{
		panic(err)
	}

	return rows
	
}

