package main

import (
	"net/http"
	"todo-app/users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	user := users.NewCreateUser()
	r.HandleFunc("/createuser", user.CreateUser)
	http.ListenAndServe(":8080", r)

}
