package todos

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
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
	w.Header().Set("Access-Control-Allow-Origin", "https://todo-22-front.herokuapp.com")
	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	e := godotenv.Load()
	if e != nil {
		http.Error(w, e.Error(), 500)
	}

	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	defer db.Close()

	id := r.URL.Query().Get("id")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	var data EditTodoBody
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), 500)
	}

	editData := EditTodoBody{data.Todo, time.Now()}

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err2 := auth.TokenVerify(tokenString)
	if err2 != nil {
		http.Error(w, err.Error(), 500)
	} else {

		stmt, err := db.Prepare("UPDATE todos set Todo=?, UpdatedAt=? WHERE ID=?")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		_, err = stmt.Exec(editData.Todo, editData.UpdatedAt, id)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		json.NewEncoder(w).Encode(editData)
	}

}
