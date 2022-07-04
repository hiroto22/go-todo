package controlar

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func CreateDb() {
	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}
	// dbConnectionInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/track_test", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `users`(" +
		"ID int NOT NULL AUTO_INCREMENT," +
		"Name varchar(255)," +
		"Email varchar(255)," +
		"PassWord varchar(255)," +
		"CreatedAt datetime," +
		"UpdatedAt datetime," +
		"PRIMARY KEY(id))")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `todos`(" +
		"ID int NOT NULL AUTO_INCREMENT ," +
		"UserID int NOT NULL," +
		"Todo text NOT NULL," +
		"CreatedAt datetime," +
		"UpdatedAt datetime," +
		"PRIMARY KEY(id))")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
