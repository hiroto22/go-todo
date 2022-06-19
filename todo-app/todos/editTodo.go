package todos

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type EditTodoBody struct {
	Todo      string    `json:"todo"`
	UpdatedAt time.Time `json:"updatedat"`
}

func EditTodo(w http.ResponseWriter, r *http.Request) {
	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}
	dbConnectionInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go_todo", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := r.URL.Query().Get("id")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data EditTodoBody
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	editData := EditTodoBody{data.Todo, time.Now()}

	stmt, err := db.Prepare("UPDATE todos set Todo=?, UpdatedAt=? WHERE ID=?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(editData.Todo, editData.UpdatedAt, id)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "applicaiton/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	json.NewEncoder(w).Encode(editData)

}
