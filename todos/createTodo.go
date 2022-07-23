package todos

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"todo-22-app/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Todo struct {
	Todo      string    `json:"todo"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type TodoBody struct {
	Todo   string `json:"todo"`
	UserID int    `json:"userid,string"`
}

//新しいTODOを登録するAPI
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// CORS
	CORS_URL := os.Getenv("CORS_URL") //呼び出しもとの情報
	w.Header().Set("Content-Type", "*")
	w.Header().Set("Access-Control-Allow-Origin", CORS_URL)
	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	e := godotenv.Load()
	if e != nil {
		http.Error(w, e.Error(), 500)
	}

	//MySQL接続
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	defer db.Close()

	//Tokenをリクエストのheaderから取得
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//取得したtokenからuseIdを特定
	secretKey := os.Getenv("SECURITY_KEY")
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(userid *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	userID := claims["userid"]

	//リクエストボディを取得
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	//userが入力したtodoとuserId
	var data TodoBody

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), 500)
	}
	todo := data.Todo

	//DBに送るデータ(user_id以外)
	todoData := Todo{todo, time.Now(), time.Now()}

	//token認証を行い正しければDBにTODOを追加
	_, err2 := auth.TokenVerify(tokenString)
	if err2 != nil {
		http.Error(w, err.Error(), 500)
	} else {

		stmt, err := db.Prepare("INSERT INTO todos (Todo,UserID,CreatedAt,UpdatedAt) VALUES(?,?,?,?)")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		_, err = stmt.Exec(todoData.Todo, userID, todoData.CreatedAt, todoData.UpdatedAt)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		json.NewEncoder(w).Encode(todoData)
	}
}
