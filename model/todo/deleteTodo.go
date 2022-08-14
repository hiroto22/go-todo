package todo

import (
	"todo-22-app/db"
)

//login
func DeleteTodo(id string) error {
	db := db.ConnectDb()
	defer db.Close()

	//dbから特定のtodoを削除
	stmt, err := db.Prepare("DELETE FROM todos WHERE ID=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return err

}
