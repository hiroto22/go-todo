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
	w.Header().Set("Content-Type", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	//Tokenをリクエストのheaderから取得
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//Token認証
	_, err := auth.TokenVerify(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized error", http.StatusUnauthorized)
		return
	}

	//MySQL接続
	e := godotenv.Load() //環境変数の読み込み
	if e != nil {
		http.Error(w, "Internal Server Error", 500)
	}
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
	defer db.Close()

	//取得したtokenからuseIdを特定
	secretKey := os.Getenv("SECURITY_KEY")
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(userid *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
	userID := claims["userid"]

	//リクエストボディを取得
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	//userが入力したtodoとuserId
	var data TodoBody

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
	todo := data.Todo

	//DBに送るデータ(user_id以外)
	timeNow := time.Now() //現在時刻の取得
	todoData := Todo{todo, timeNow, timeNow}

	//DBにTODOを追加
	stmt, err := db.Prepare("INSERT INTO todos (Todo,UserID,CreatedAt,UpdatedAt) VALUES(?,?,?,?)")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	_, err = stmt.Exec(todoData.Todo, userID, todoData.CreatedAt, todoData.UpdatedAt)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	json.NewEncoder(w).Encode(todoData)
}
