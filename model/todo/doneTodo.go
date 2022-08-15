package todo

import "database/sql"

//todoを完了または未完了にする
func DoneTodo(id string, isComplete string, db *sql.DB) error {
	//現在のisCompleteにあわせて更新する
	if isComplete == "false" {
		_, err := db.Exec("UPDATE todos set IsDone=? WHERE ID=?", true, id)
		if err != nil {
			return err
		}
	} else {
		_, err := db.Exec("UPDATE todos set IsDone=? WHERE ID=?", false, id)
		if err != nil {
			return err
		}
	}
	return nil
}
