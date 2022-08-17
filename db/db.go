package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDb() {
	var err error
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	DB, err = sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Println(err)
	}
	log.Println("connected db")
}
