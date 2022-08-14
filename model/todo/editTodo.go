package todo

import (
	"time"
	"todo-22-app/db"
)

type EditTodo struct {
	Todo      string    `json:"todo"`
	UpdatedAt time.Time `json:"updatedat"`
}

func NewEditTodo() *EditTodo {
	return new(EditTodo)
}

//login
func (todo *EditTodo) EditTodo(todoText string, id string) error {
	db := db.ConnectDb()
	defer db.Close()

	todo.Todo = todoText
	todo.UpdatedAt = time.Now()

	//todoを更新
	stmt, err := db.Prepare("UPDATE todos set Todo=?, UpdatedAt=? WHERE ID=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(todo.Todo, todo.UpdatedAt, id)
	if err != nil {
		return err
	}

	return err

}
