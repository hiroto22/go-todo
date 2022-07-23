package todos

import (
	"database/sql"
	"encoding/json"
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

//userのidにあわせたtodoの一覧を取得するAPI(isDoneの値ごとに取得)
func GetTodoListWithUserId(w http.ResponseWriter, r *http.Request) {
	//CORS
	CORS_URL := os.Getenv("CORS_URL") //呼び出しもとの情報
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", CORS_URL)
	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	//MySQLに接続
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

	//指定されたisDoneの状態を取得
	isDone := r.URL.Query().Get("isdone")

	//tokenをrequestのheaderから取得
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//tokenからuserIdを取得
	var secretKey = os.Getenv("SECURITY_KEY")
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(userid *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	userID := claims["userid"]

	//token認証して正しければtodoの一覧を取得し値を返す
	_, err = auth.TokenVerify(tokenString)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		rows, err := db.Query("SELECT * FROM todos WHERE IsDone=? AND UserID=?", isDone, userID)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		defer rows.Close()

		//returnされる内容
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
				http.Error(w, err.Error(), 500)
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

		json.NewEncoder(w).Encode(data)
	}

}
