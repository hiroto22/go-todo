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

//現在使用していない
//1つ1つのtodoを取得するAPI
func GetTodo(w http.ResponseWriter, r *http.Request) {
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

	//tokenをrequestのheaderから取得
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//token認証
	_, err := auth.TokenVerify(tokenString)
	if err != nil {
		http.Error(w, "invalid_access_tokenr", http.StatusUnauthorized)
		return
	}

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

	//todoのidを取得
	id := r.URL.Query().Get("id")

	var todo TodoList

	//todoの取得
	err = db.QueryRow("SELECT * FROM todos WHERE ID=?", id).Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Todo,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.IsDone,
	)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	json.NewEncoder(w).Encode(todo)
}
