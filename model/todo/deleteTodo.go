package todo

import "database/sql"

//todoの削除を行う
func DeleteTodo(id string, db *sql.DB) error {
	//dbから特定のtodoを削除
	_, err := db.Query("DELETE FROM todos WHERE ID=?", id)
	if err != nil {
		return err
	}

	return nil
}
