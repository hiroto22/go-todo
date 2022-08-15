package todo

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"todo-22-app/auth"
	model "todo-22-app/model/todo"

	"github.com/dgrijalva/jwt-go"
)

//todo変更に使うAPI
func GetTodoListWithUserId(w http.ResponseWriter, r *http.Request) {
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

	//指定されたisDoneの状態を取得
	isDone := r.URL.Query().Get("isdone")

	//tokenからuserIdを取得
	var secretKey = os.Getenv("SECURITY_KEY")
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(userid *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	userID := claims["userid"]

	todoList := model.NewTodoList()
	todoList.GetTodoListWithUserId(isDone, userID)

	json.NewEncoder(w).Encode(todoList)

}
