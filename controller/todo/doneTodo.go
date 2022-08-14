package todo

import (
	"encoding/json"
	"net/http"
	"strings"
	"todo-22-app/auth"
	"todo-22-app/model/todo"
)

//todo作成に使うAPI
func DoneTodo(w http.ResponseWriter, r *http.Request) {
	//cors
	w.Header().Set("Content-Type", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")

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
	todo.DoneTodo(id, isComplete)

	json.NewEncoder(w).Encode(id)

}
