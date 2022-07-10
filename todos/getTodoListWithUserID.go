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
	"todo-22-app/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type TodoListWithUserID struct {
	ID        int           `json:"id"`
	UserID    sql.NullInt64 `json:"userid"`
	Todo      string        `json:"todo"`
	CreatedAt time.Time     `json:"createdat"`
	UpdatedAt time.Time     `json:"updatedat"`
	IsDone    bool          `json:"isdone"`
}

func GetTodoListWithUserId(w http.ResponseWriter, r *http.Request) {
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
		log.Println(e)
	}
	// dbConnectionInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go_todo?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	isDone := r.URL.Query().Get("isdone")

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// var secretKey = os.Getenv("SECURITY_KEY")

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(userid *jwt.Token) (interface{}, error) {
		return []byte("gotodo"), nil
	})
	if err != nil {
		log.Println(err)
	}
	// do something with decoded claims

	fmt.Println(claims["userid"])
	userID := claims["userid"]

	token, err := auth.TokenVerify(tokenString)
	log.Println(token)

	if err != nil {
		log.Fatal(err)
	} else {
		rows, err := db.Query("SELECT * FROM todos WHERE IsDone=? AND UserID=?", isDone, userID)
		if err != nil {
			log.Println(err)
		}

		defer rows.Close()

		var data []TodoListWithUserID

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
				log.Println(err)
			} else {
				data = append(data, TodoListWithUserID{
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
