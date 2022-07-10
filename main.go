package main

import (
	"net/http"
	"os"
	"todo-22-app/controlar"
	"todo-22-app/todos"
	"todo-22-app/users"

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
	r.HandleFunc("/getusertodoList", todos.GetTodoListWithUserId)
	r.HandleFunc("/gettodo", todos.GetTodo)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)

}
