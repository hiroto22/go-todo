package todos

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"todo-22-app/auth"

	"github.com/joho/godotenv"
)

type TodoList struct {
	ID        int           `json:"id"`
	UserID    sql.NullInt64 `json:"userid"`
	Todo      string        `json:"todo"`
	CreatedAt time.Time     `json:"createdat"`
	UpdatedAt time.Time     `json:"updatedat"`
	IsDone    bool          `json:"isdone"`
}

func GetTodoList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

	isDone := r.URL.Query().Get("isdone")

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := auth.TokenVerify(tokenString)
	log.Printf("request token=%s\n", token)
	if err != nil {
		log.Println("")
	} else {
		rows, err := db.Query("SELECT * FROM todos WHERE IsDone=?", isDone)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		defer rows.Close()

		var data []TodoList

		for rows.Next() {
			var todoList TodoList

			err := rows.Scan(
				&todoList.ID,
				&todoList.UserID,
				&todoList.Todo,
				&todoList.CreatedAt,
				&todoList.UpdatedAt,
				&todoList.IsDone)

			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				data = append(data, TodoList{
					ID:        todoList.ID,
					UserID:    todoList.UserID,
					Todo:      todoList.Todo,
					CreatedAt: todoList.CreatedAt,
					UpdatedAt: todoList.UpdatedAt,
					IsDone:    todoList.IsDone,
				})
			}
		}

		// jsonData, _ := json.Marshal(data)

		json.NewEncoder(w).Encode(data)
	}

}
