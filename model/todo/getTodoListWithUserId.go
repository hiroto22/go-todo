package todo

import (
	"database/sql"
	"time"
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
func (todoList *TodoList) GetTodoListWithUserId(isDone string, userID interface{}, db *sql.DB) error {

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
			todoList.AddTodo(todo)
		}
	}

	defer rows.Close()

	return nil
}

func (todoList *TodoList) AddTodo(a TodoForList) {
	*todoList = append(*todoList, a)
}
