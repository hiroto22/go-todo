package todo

import (
	"encoding/json"
	"net/http"
	"strings"
	"todo-22-app/auth"
	"todo-22-app/model/todo"
)

//todo作成に使うAPI
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	//Tokenをリクエストのheaderから取得
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//Token認証
	_, err := auth.TokenVerify(tokenString)
	if err != nil {
		http.Error(w, "invalid_access_token", http.StatusUnauthorized)
		return
	}

	//queryからtodoのidを取得
	id := r.URL.Query().Get("id")

	todo.DeleteTodo(id)

	json.NewEncoder(w).Encode(id)

}
