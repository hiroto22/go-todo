package main

import (
	"net/http"
	"os"
	"todo-app/controlar"
	"todo-app/todos"
	"todo-app/users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	controlar.CreateDb()
	r := mux.NewRouter()
	user := users.NewCreateUser()
	r.HandleFunc("/createuser", user.CreateUser)
	r.HandleFunc("/login", users.Login)
	r.HandleFunc("/createtodo", todos.CreateTodo)
	r.HandleFunc("/deletetodo", todos.DeleteTodo)
	r.HandleFunc("/edittodo", todos.EditTodo)
	r.HandleFunc("/completetodo", todos.DoneTodo)
	r.HandleFunc("/gettodoList", todos.GetTodoList)
	r.HandleFunc("/get-usertodoList", todos.GetTodoListWithUserId)
	r.HandleFunc("/gettodo", todos.GetTodo)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)

}
