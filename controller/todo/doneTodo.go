package todo

import (
	"encoding/json"
	"net/http"
	"strings"
	"todo-22-app/auth"
	"todo-22-app/db"
	"todo-22-app/model/todo"
)

//todo作成に使うAPI
func DoneTodo(w http.ResponseWriter, r *http.Request) {
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

	//todoのidと状態(isComplete)を取得
	id := r.URL.Query().Get("id")
	isComplete := r.URL.Query().Get("isComplete")

	//todoの状態を変更
	todo.DoneTodo(id, isComplete, db)

	json.NewEncoder(w).Encode(id)

}
