package todo

import (
	"encoding/json"
	"net/http"
	"strings"
	"todo-22-app/auth"
	"todo-22-app/db"
	"todo-22-app/middleware"
	model "todo-22-app/model/todo"
)

//todo変更に使うAPI
func GetTodoListWithUserId(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectedDb()
	//Tokenをリクエストのheaderから取得
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//Token認証
	_, err := auth.TokenVerify(tokenString)
	if err != nil {
		http.Error(w, "invalid_access_token", http.StatusUnauthorized)
		return
	}

	//指定されたisDoneの状態を取得
	isDone := r.URL.Query().Get("isdone")

	//tokenからuserIdを取得
	//取得したtokenからuseIdを特定
	userID := middleware.GetUserIdFromToken(tokenString)

	todoList := model.NewTodoList()
	todoList.GetTodoListWithUserId(isDone, userID, db)

	json.NewEncoder(w).Encode(todoList)

}
