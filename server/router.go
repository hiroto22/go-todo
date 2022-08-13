package server

import (
	controller "todo-22-app/controller/user"
	"todo-22-app/db"
	"todo-22-app/todos"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	db.CreateDb()

	r := mux.NewRouter()
	r.HandleFunc("/signup", controller.SingUp)
	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/createtodo", todos.CreateTodo)
	r.HandleFunc("/deletetodo", todos.DeleteTodo)
	r.HandleFunc("/edittodo", todos.EditTodo)
	r.HandleFunc("/completetodo", todos.DoneTodo)
	r.HandleFunc("/gettodoList", todos.GetTodoList)
	r.HandleFunc("/getusertodoList", todos.GetTodoListWithUserId)
	r.HandleFunc("/gettodo", todos.GetTodo)

	return r

}
