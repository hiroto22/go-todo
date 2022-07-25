package todos

import (
	"database/sql"
	"encoding/json"
	"log"
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

	//MySQL接続
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

	//todoのidと状態(isComplete)を取得
	id := r.URL.Query().Get("id")
	isComplete := r.URL.Query().Get("isComplete")

	//requestのheaderからtokenを取得
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//token認証を行い正しければDBのisCompleteを更新
	_, err2 := auth.TokenVerify(tokenString)
	if err2 != nil {
		log.Println("tokenはありません。")
	} else {

		stmt, err := db.Prepare("UPDATE todos set IsDone=? WHERE ID=?")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		//現在のisCompleteにあわせて更新する
		if isComplete == "false" {
			_, err = stmt.Exec(true, id)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
		} else {
			_, err = stmt.Exec(false, id)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
		}

		json.NewEncoder(w).Encode(id)
	}

}
