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

func GetTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "applicaiton/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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
		log.Println(e)
	}
	// dbConnectionInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go_todo?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	id := r.URL.Query().Get("id")

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err2 := auth.TokenVerify(tokenString)
	if err2 != nil {
		log.Println(err)
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
			log.Println(err)
		}

		json.NewEncoder(w).Encode(todo)
	}

}
