package main

import (
	"net/http"
	"todo-app/todos"
	"todo-app/users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	user := users.NewCreateUser()
	r.HandleFunc("/createuser", user.CreateUser).Methods("POST")
	r.HandleFunc("/login", users.Login).Methods("POST")
	r.HandleFunc("/createtodo", todos.CreateTodo).Methods("POST")
	r.HandleFunc("/deletetodo", todos.DeleteTodo).Methods("DELETE")
	r.HandleFunc("/edittodo", todos.EditTodo).Methods("POST")
	r.HandleFunc("/completetodo", todos.DoneTodo).Methods("POST")
	r.HandleFunc("/gettodoList", todos.GetTodoList).Methods("GET")
	r.HandleFunc("/gettodo", todos.GetTodo).Methods("GET")
	http.ListenAndServe(":8080", r)

}
