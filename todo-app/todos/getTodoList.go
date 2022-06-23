package todos

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"todo-app/auth"

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
	w.Header().Set("Content-Type", "applicaiton/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}
	dbConnectionInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go_todo?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
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
				log.Fatal(err)
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
