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

//todoを完了または未完了にするAPI
func DoneTodo(w http.ResponseWriter, r *http.Request) {
	//CORS
	w.Header().Set("Content-Type", "*")
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

	//token認証
	_, err := auth.TokenVerify(tokenString)
	if err != nil {
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}

	//MySQL接続
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

	//todoのidと状態(isComplete)を取得
	id := r.URL.Query().Get("id")
	isComplete := r.URL.Query().Get("isComplete")

	stmt, err := db.Prepare("UPDATE todos set IsDone=? WHERE ID=?")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//現在のisCompleteにあわせて更新する
	if isComplete == "false" {
		_, err = stmt.Exec(true, id)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
	} else {
		_, err = stmt.Exec(false, id)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}

	json.NewEncoder(w).Encode(id)
}
