package todocontroller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	todomodel "todo-22-app/model/todoModel"
)

type TodoState struct {
	Todo   string  `json:"todo"`
	UserID float64 `json:"userid"`
}

//todo作成に使うAPI
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var data TodoState

	//取得したtokenからuseIdを特定
	userID := r.Context().Value("userID").(float64)

	data.UserID = userID

	//リクエストボディを取得
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "400 Bad Request", 400)
		return
	}

	json.Unmarshal(body, &data)

	//todoをデータベースに登録
	todo := todomodel.NewTodo()
	todo.CreateTodo(data.Todo, data.UserID)

	json.NewEncoder(w).Encode(todo)

}

//todo削除に使うAPI
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	//queryからtodoのidを取得
	id := r.URL.Query().Get("id")

	todomodel.DeleteTodo(id)

	json.NewEncoder(w).Encode(id)

}

//todoの状態変更に使うAPI
func DoneTodo(w http.ResponseWriter, r *http.Request) {
	//todoのidと状態(isComplete)を取得
	id := r.URL.Query().Get("id")
	isComplete := r.URL.Query().Get("isComplete")

	//todoの状態を変更
	todomodel.DoneTodo(id, isComplete)

	json.NewEncoder(w).Encode(id)

}

type EditTodoState struct {
	Todo string `json:"todo"`
}

//todo変更に使うAPI
func EditTodo(w http.ResponseWriter, r *http.Request) {

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
	EditTodo := todomodel.NewEditTodo()
	EditTodo.EditTodo(data.Todo, id)

	json.NewEncoder(w).Encode(EditTodo)

}

//todo一覧の取得に使うAPI
func GetTodoListWithUserId(w http.ResponseWriter, r *http.Request) {
	//指定されたisDoneの状態を取得
	isDone := r.URL.Query().Get("isdone")

	userID := r.Context().Value("userID").(float64)

	todoList := todomodel.NewTodoList()
	todoList.GetTodoListWithUserId(isDone, userID)

	json.NewEncoder(w).Encode(todoList)

}
