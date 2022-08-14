package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Todo struct {
	Todo      string    `json:"todo"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PassWord  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

var (
	//Client user_id database
	db  *sql.DB
	err error
)

func ConnectDb() *sql.DB {
	godotenv.Load()

	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, _ = sql.Open("mysql", dbConnectionInfo)
	return db
}

func init() {
	godotenv.Load()

	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, _ = sql.Open("mysql", dbConnectionInfo)
}

// func Login(email string, id int, name string, password string, createdAt time.Time, updatedAt time.Time) {
// 	db.QueryRow("SELECT * FROM users WHERE Email=?", email).Scan(&user.ID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt, &user.UpdatedAt)
// }

// func CreateTodo(todo string, userID interface{}, createdAt time.Time, updatedAt time.Time) error {
// 	stmt, err := db.Prepare("INSERT INTO todos (Todo,UserID,CreatedAt,UpdatedAt) VALUES(?,?,?,?)")
// 	if err != nil {
// 		return err
// 	}

// 	_, err = stmt.Exec(todo, userID, createdAt, updatedAt)
// 	return err
// }

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
