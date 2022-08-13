package todo

import (
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

//login
func (todo *Todo) CreateTodo(todoText string, userID interface{}) error {
	db := db.ConnectDb()
	defer db.Close()

	//DBに送るデータ(user_id以外)
	nowTime := time.Now() //現在時刻の取得
	todo.Todo = todoText
	todo.CreatedAt = nowTime
	todo.UpdatedAt = nowTime

	//DBにTODOを追加
	stmt, err := db.Prepare("INSERT INTO todos (Todo,UserID,CreatedAt,UpdatedAt) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(todo.Todo, userID, todo.CreatedAt, todo.UpdatedAt)

	return err

}
