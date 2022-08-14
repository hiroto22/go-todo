package todo

import (
	"todo-22-app/db"
)

//login
func DoneTodo(id string, isComplete string) error {
	db := db.ConnectDb()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE todos set IsDone=? WHERE ID=?")
	if err != nil {
		return err
	}

	//現在のisCompleteにあわせて更新する
	if isComplete == "false" {
		_, err = stmt.Exec(true, id)
		if err != nil {
			return err
		}
	} else {
		_, err = stmt.Exec(false, id)
		if err != nil {
			return err
		}
	}
	return err
}
