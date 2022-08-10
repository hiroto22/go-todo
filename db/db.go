package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDb() *sql.DB {
	godotenv.Load()

	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, _ := sql.Open("mysql", dbConnectionInfo)
	return db
}

//DBをcreateする
func CreateDb() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}

	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	//users table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `users`(" +
		"ID int NOT NULL AUTO_INCREMENT," +
		"Name varchar(255) NOT NULL," +
		"Email varchar(255) NOT NULL," +
		"PassWord varchar(255) NOT NULL," +
		"CreatedAt datetime," +
		"UpdatedAt datetime," +
		"PRIMARY KEY(id))")

	if err != nil {
		log.Println(err)
	}

	//todos table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `todos`(" +
		"ID int NOT NULL AUTO_INCREMENT ," +
		"UserID int NOT NULL," +
		"Todo text NOT NULL," +
		"CreatedAt datetime," +
		"UpdatedAt datetime," +
		"isDone boolean," +
		"PRIMARY KEY(id))")

	if err != nil {
		log.Println(err)
	}
	defer db.Close()

}
