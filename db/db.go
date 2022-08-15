package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func ConnectDb() *sql.DB {
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err = sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Println(err)
	}
	log.Println("connected db")

	return db
}

func ConnectedDb() *sql.DB {
	return db
}
