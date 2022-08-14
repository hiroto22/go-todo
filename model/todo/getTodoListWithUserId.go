package todo

import (
	"database/sql"
	"time"
	"todo-22-app/db"
)

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

//login
func (todoList TodoList) GetTodoListWithUserId(isDone string, userID interface{}) error {
	db := db.ConnectDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todos WHERE IsDone=? AND UserID=?", isDone, userID)
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
			todoList = append(todoList, todo)
		}
	}

	defer rows.Close()

	return err
}
