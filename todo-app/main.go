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
	r.HandleFunc("/createuser", user.CreateUser)
	r.HandleFunc("/login", users.Login)
	r.HandleFunc("/createtodo", todos.CreateTodo)
	r.HandleFunc("/deletetodo", todos.DeleteTodo)
	r.HandleFunc("/edittodo", todos.EditTodo)
	r.HandleFunc("/completetodo", todos.DoneTodo)
	r.HandleFunc("/gettodoList", todos.GetTodoList)
	r.HandleFunc("/gettodo", todos.GetTodo)
	http.ListenAndServe(":8080", r)

}
