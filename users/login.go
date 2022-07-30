package users

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
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

//loginの際に使うAPI
func Login(w http.ResponseWriter, r *http.Request) {
	//CORS
	w.Header().Set("Content-Type", "applicaiton/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	//MySQLに接続
	e := godotenv.Load() //環境変数の読み込み
	if e != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//requestされたemailとpassword
	var data LoginState
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "400 Bad Request", 400)
		return
	}

	email := data.Email

	//DBからのuser情報
	var user User

	//emailをもとにDBからuser情報を取得
	err = db.QueryRow("SELECT * FROM users WHERE Email=?", email).Scan(&user.ID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//requestされたemail,passwordとDBの物が正しいか確認正しければtokenを返す
	err = auth.PasswordVerify(user.PassWord, data.PassWord)
	if err != nil {
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	} else {
		token, err := auth.CreateToken(user.ID)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

		tokenData := tokenRes{token}

		json.NewEncoder(w).Encode(tokenData)
	}

}
