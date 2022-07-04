package todos

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}
	// dbConnectionInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go_todo", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := r.URL.Query().Get("id")

	// tokenString := r.Header.Get("Authorization")
	// tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// _, err2 := auth.TokenVerify(tokenString)
	// if err2 != nil {
	// 	log.Fatal(err)
	// }

	stmt, err := db.Prepare("DELETE FROM todos WHERE ID=?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "applicaiton/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	json.NewEncoder(w).Encode(id)

}
