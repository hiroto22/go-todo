package todos

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "applicaiton/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://todo-22-front.herokuapp.com")
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
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	defer db.Close()

	id := r.URL.Query().Get("id")

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
