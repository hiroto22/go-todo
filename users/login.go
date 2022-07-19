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
	w.Header().Set("Content-Type", "applicaiton/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://todo-22-front.herokuapp.com")
	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	e := godotenv.Load()
	if e != nil {
		http.Error(w, e.Error(), 500)
	}
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	log.Printf("request body=%s\n", r.Body)

	var data LoginState
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), 500)
	}

	email := data.Email

	var user User

	err = db.QueryRow("SELECT * FROM users WHERE Email=?", email).Scan(&user.ID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = auth.PasswordVerify(user.PassWord, data.PassWord)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		token, err := auth.CreateToken(user.ID)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		tokenData := tokenRes{token}

		json.NewEncoder(w).Encode(tokenData)
	}

}
