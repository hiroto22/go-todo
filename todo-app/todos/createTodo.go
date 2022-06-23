package todos

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"todo-app/auth"

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
		log.Fatal(e)
	}
	dbConnectionInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go_todo", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	log.Printf("request token=%s\n", tokenString)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)

	}

	log.Printf("request body=%s\n", body)

	var data TodoBody

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	// userId := 12
	todo := data.Todo

	todoData := Todo{todo, time.Now(), time.Now()}

	_, err2 := auth.TokenVerify(tokenString)
	if err2 != nil {
		log.Fatal(err)
	} else {

		stmt, err := db.Prepare("INSERT INTO todos (Todo,CreatedAt,UpdatedAt) VALUES(?,?,?)")
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec(todoData.Todo, todoData.CreatedAt, todoData.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(todoData)
	}
}
