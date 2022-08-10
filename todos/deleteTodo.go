package todos

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"todo-22-app/auth"

	"github.com/joho/godotenv"
)

//todoを削除するAPI
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
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

	//requestのheaderからtokenを取得
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//Token認証
	_, err := auth.TokenVerify(tokenString)
	if err != nil {
		http.Error(w, "invalid_access_token", http.StatusUnauthorized)
		return
	}

	//MySQL接続
	e := godotenv.Load() //環境変数の読み込み
	if e != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//MySQL接続
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	defer db.Close()

	//queryからtodoのidを取得
	id := r.URL.Query().Get("id")

	//dbから特定のtodoを削除
	stmt, err := db.Prepare("DELETE FROM todos WHERE ID=?")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	json.NewEncoder(w).Encode(id)
}
