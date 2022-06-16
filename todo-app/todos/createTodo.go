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

type Todo struct {
	// UserID    int       `json:"userid"`
	Todo      string    `json:"todo"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type TodoBody struct {
	Todo string `json:"todo"`
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data TodoBody
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	// userId := 12
	todo := data.Todo

	todoData := Todo{todo, time.Now(), time.Now()}

	stmt, err := db.Prepare("INSERT INTO todos (Todo,CreatedAt,UpdatedAt) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(todoData.Todo, todoData.CreatedAt, todoData.UpdatedAt)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "applicaiton/json")

	json.NewEncoder(w).Encode(todoData)

}
