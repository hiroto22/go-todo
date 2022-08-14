package todo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"todo-22-app/auth"
	"todo-22-app/middleware"
	model "todo-22-app/model/todo"
)

type TodoState struct {
	Todo   string      `json:"todo"`
	UserID interface{} `json:"userid"`
}

//todo作成に使うAPI
func CreateTodo(w http.ResponseWriter, r *http.Request) {
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

	var data TodoState

	//取得したtokenからuseIdを特定
	userID := middleware.GetUserIdFromToken(tokenString)
	data.UserID = userID

	//リクエストボディを取得
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "400 Bad Request", 400)
		return
	}

	json.Unmarshal(body, &data)

	//todoをデータベースに登録
	todo := model.NewTodo()
	todo.CreateTodo(data.Todo, data.UserID)

	json.NewEncoder(w).Encode(todo)

}
