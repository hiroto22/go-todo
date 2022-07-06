package users

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"todo-22-app/auth"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PassWord  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type CreateUserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

func (user *CreateUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}
	// dbConnectionInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go_todo?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data CreateUserBody
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	name := data.Name
	email := data.Email
	password := data.PassWord

	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	userData := CreateUser{name, email, string(hashPassWord), time.Now(), time.Now()}

	stmt, err := db.Prepare("INSERT INTO users (Name,Email,PassWord,CreatedAt,UpdatedAt) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Println(err)
	}

	_, err2 := stmt.Exec(userData.Name, userData.Email, userData.PassWord, time.Now(), time.Now())
	if err != nil {
		log.Println(err2)
	}

	var userd User

	err = db.QueryRow("SELECT * FROM users WHERE Email=?", email).Scan(&userd.ID, &userd.Name, &userd.Email, &userd.PassWord, &userd.CreatedAt, &userd.UpdatedAt)
	if err != nil {
		log.Println(err)
	}

	token, err := auth.CreateToken(userd.ID)
	if err != nil {
		log.Println(err)
	}

	tokenData := tokenRes{token}

	w.Header().Set("Content-Type", "applicaiton/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	json.NewEncoder(w).Encode(tokenData)

}

func NewCreateUser() *CreateUser {
	return new(CreateUser)
}
