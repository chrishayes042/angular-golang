package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnvInfo(key string) string {
	err := godotenv.Load("postgres.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}


func GetSQLInfo()(*sql.DB){
	var (
		host 		= 	getEnvInfo("HOST")
		user 		= 	getEnvInfo("USER")
		password 	= 	getEnvInfo("PASS")
		dbname 		= 	getEnvInfo("DBNAME")
		port 		= 	getEnvInfo("PORT")
	)
	

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s " +
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

