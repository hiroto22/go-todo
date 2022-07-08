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
		log.Println(e)
	}
	// dbConnectionInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go_todo", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	log.Printf("request token=%s\n", tokenString)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)

	}

	log.Printf("request body=%s\n", body)

	var data TodoBody

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	// userId := 12
	secretKey := os.Getenv("SECURITY_KEY")
	todo := data.Todo
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(userid *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Println(err)
	}
	// do something with decoded claims

	fmt.Println(claims["userid"])
	userID := claims["userid"]

	todoData := Todo{todo, time.Now(), time.Now()}

	_, err2 := auth.TokenVerify(tokenString)
	if err2 != nil {
		log.Println(err)
	} else {

		stmt, err := db.Prepare("INSERT INTO todos (Todo,UserID,CreatedAt,UpdatedAt) VALUES(?,?,?,?)")
		if err != nil {
			log.Println(err)
		}

		_, err = stmt.Exec(todoData.Todo, userID, todoData.CreatedAt, todoData.UpdatedAt)
		if err != nil {
			log.Println(err)
		}

		json.NewEncoder(w).Encode(todoData)
	}
}
