package todos

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//todoを削除するAPI
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	//CORS
	CORS_URL := os.Getenv("CORS_URL") //呼び出しもとの情報
	w.Header().Set("Content-Type", "applicaiton/json")
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

	//queryからtodoのidを取得
	id := r.URL.Query().Get("id")

	//dbから特定のtodoを削除
	stmt, err := db.Prepare("DELETE FROM todos WHERE ID=?")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	json.NewEncoder(w).Encode(id)

}
