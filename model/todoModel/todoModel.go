package todomodel

import (
	"database/sql"
	"time"
	"todo-22-app/db"
)

type Todo struct {
	Todo      string    `json:"todo"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

func NewTodo() *Todo {
	return new(Todo)
}

//todoの登録を行う
func (todo *Todo) CreateTodo(todoText string, userID float64) error {
	//DBに送るデータ(user_id以外)
	nowTime := time.Now() //現在時刻の取得
	todo.Todo = todoText
	todo.CreatedAt = nowTime
	todo.UpdatedAt = nowTime

	//DBにTODOを追加

	if _, err := db.DB.Exec("INSERT INTO todos (Todo,UserID,CreatedAt,UpdatedAt) VALUES(?,?,?,?)", todo.Todo, userID, todo.CreatedAt, todo.UpdatedAt); err != nil {
		return err
	}

	return nil

}

//todoの削除を行う
func DeleteTodo(id string) error {
	//dbから特定のtodoを削除
	_, err := db.DB.Query("DELETE FROM todos WHERE ID=?", id)
	if err != nil {
		return err
	}

	return nil
}

//todoを完了または未完了にする
func DoneTodo(id string, isComplete string) error {
	//現在のisCompleteにあわせて更新する
	if isComplete == "false" {
		_, err := db.DB.Exec("UPDATE todos set IsDone=? WHERE ID=?", true, id)
		if err != nil {
			return err
		}
	} else {
		_, err := db.DB.Exec("UPDATE todos set IsDone=? WHERE ID=?", false, id)
		if err != nil {
			return err
		}
	}
	return nil
}

type EditTodo struct {
	Todo      string    `json:"todo"`
	UpdatedAt time.Time `json:"updatedat"`
}

func NewEditTodo() *EditTodo {
	return new(EditTodo)
}

//todoの内容を更新する
func (todo *EditTodo) EditTodo(todoText string, id string) error {
	todo.Todo = todoText
	todo.UpdatedAt = time.Now()

	//todoを更新
	_, err := db.DB.Exec("UPDATE todos set Todo=?, UpdatedAt=? WHERE ID=?", todo.Todo, todo.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil

}

type TodoForList struct {
	ID        int           `json:"id"`
	UserID    sql.NullInt64 `json:"userid"`
	Todo      string        `json:"todo"`
	CreatedAt time.Time     `json:"createdat"`
	UpdatedAt time.Time     `json:"updatedat"`
	IsDone    bool          `json:"isdone"`
}

type TodoList []TodoForList

func NewTodoList() *TodoList {
	return new(TodoList)
}

//todoの一覧を取得する
func (todoList *TodoList) GetTodoListWithUserId(isDone string, userID float64) error {

	rows, err := db.DB.Query("SELECT * FROM todos WHERE IsDone=? AND UserID=?", isDone, userID)
	if err != nil {
		return err
	}

	for rows.Next() {
		var todo TodoForList

		err := rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Todo,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.IsDone)

		if err != nil {
			return err
		} else {
			todoList.AddTodo(todo)
		}
	}

	defer rows.Close()

	return nil
}

func (todoList *TodoList) AddTodo(a TodoForList) {
	*todoList = append(*todoList, a)
}
