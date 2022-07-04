package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"todo-app/auth"

	"github.com/joho/godotenv"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PassWord  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type LoginState struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}
type tokenRes struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}
	dbConnectionInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go_todo?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("request body=%s\n", r.Body)

	var data LoginState
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	email := data.Email
	// password := data.PassWord

	var user User

	err = db.QueryRow("SELECT * FROM users WHERE Email=?", email).Scan(&user.ID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Fatal(err)
	}

	err = auth.PasswordVerify(user.PassWord, data.PassWord)
	if err != nil {
		fmt.Println(err)
	} else {
		token, err := auth.CreateToken(user.ID)
		if err != nil {
			log.Fatal(err)
		}

		tokenData := tokenRes{token}

		w.Header().Set("Content-Type", "applicaiton/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		json.NewEncoder(w).Encode(tokenData)
	}

}
