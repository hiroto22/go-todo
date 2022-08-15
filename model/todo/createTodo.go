package todo

import (
	"database/sql"
	"time"
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
func (todo *Todo) CreateTodo(todoText string, userID interface{}, db *sql.DB) error {
	//DBに送るデータ(user_id以外)
	nowTime := time.Now() //現在時刻の取得
	todo.Todo = todoText
	todo.CreatedAt = nowTime
	todo.UpdatedAt = nowTime

	//DBにTODOを追加

	if _, err := db.Exec("INSERT INTO todos (Todo,UserID,CreatedAt,UpdatedAt) VALUES(?,?,?,?)", todo.Todo, userID, todo.CreatedAt, todo.UpdatedAt); err != nil {
		return err
	}

	return nil

}
