package todo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"todo-22-app/auth"
	model "todo-22-app/model/todo"
)

type EditTodoState struct {
	Todo string `json:"todo"`
}

//todo変更に使うAPI
func EditTodo(w http.ResponseWriter, r *http.Request) {
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

	//todoのidと変更内容を取得
	id := r.URL.Query().Get("id")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//todoの内容と変更日時を記録
	var data EditTodoState
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "400 Bad Request", 400)
		return
	}

	//todoを更新
	EditTodo := model.NewEditTodo()
	EditTodo.EditTodo(data.Todo, id)

	json.NewEncoder(w).Encode(EditTodo)

}
