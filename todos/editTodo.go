package todos

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"todo-22-app/auth"

	"github.com/joho/godotenv"
)

type EditTodoBody struct {
	Todo      string    `json:"todo"`
	UpdatedAt time.Time `json:"updatedat"`
}

func EditTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	// dbConnectionInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go_todo", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	id := r.URL.Query().Get("id")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var data EditTodoBody
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	editData := EditTodoBody{data.Todo, time.Now()}

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err2 := auth.TokenVerify(tokenString)
	if err2 != nil {
		log.Println(err)
	} else {

		stmt, err := db.Prepare("UPDATE todos set Todo=?, UpdatedAt=? WHERE ID=?")
		if err != nil {
			log.Println(err)
		}

		_, err = stmt.Exec(editData.Todo, editData.UpdatedAt, id)
		if err != nil {
			log.Println(err)
		}

		json.NewEncoder(w).Encode(editData)
	}

}
