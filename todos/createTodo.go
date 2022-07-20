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
	"todo-22-app/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Todo struct {
	Todo string `json:"todo"`
	// UserID    int       `json:"userid"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type TodoBody struct {
	Todo   string `json:"todo"`
	UserID int    `json:"userid,string"`
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
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

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	log.Printf("request token=%s\n", tokenString)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	log.Printf("request body=%s\n", body)

	var data TodoBody

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), 500)
	}

	secretKey := os.Getenv("SECURITY_KEY")
	todo := data.Todo
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(userid *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	fmt.Println(claims["userid"])
	userID := claims["userid"]

	todoData := Todo{todo, time.Now(), time.Now()}

	_, err2 := auth.TokenVerify(tokenString)
	if err2 != nil {
		http.Error(w, err.Error(), 500)
	} else {

		stmt, err := db.Prepare("INSERT INTO todos (Todo,UserID,CreatedAt,UpdatedAt) VALUES(?,?,?,?)")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		_, err = stmt.Exec(todoData.Todo, userID, todoData.CreatedAt, todoData.UpdatedAt)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		json.NewEncoder(w).Encode(todoData)
	}
}
