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

func GetTodo(w http.ResponseWriter, r *http.Request) {

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

	id := r.URL.Query().Get("id")

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err2 := auth.TokenVerify(tokenString)
	if err2 != nil {
		http.Error(w, err.Error(), 500)
	} else {

		var todo TodoList

		err = db.QueryRow("SELECT * FROM todos WHERE ID=?", id).Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Todo,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.IsDone,
		)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		json.NewEncoder(w).Encode(todo)
	}

}
