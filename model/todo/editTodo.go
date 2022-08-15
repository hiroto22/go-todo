package todo

import (
	"database/sql"
	"time"
)

type EditTodo struct {
	Todo      string    `json:"todo"`
	UpdatedAt time.Time `json:"updatedat"`
}

func NewEditTodo() *EditTodo {
	return new(EditTodo)
}

//login
func (todo *EditTodo) EditTodo(todoText string, id string, db *sql.DB) error {
	todo.Todo = todoText
	todo.UpdatedAt = time.Now()

	//todoを更新
	_, err := db.Exec("UPDATE todos set Todo=?, UpdatedAt=? WHERE ID=?", todo.Todo, todo.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil

}
