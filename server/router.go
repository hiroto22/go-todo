package server

import (
	todoController "todo-22-app/controller/todo"
	controller "todo-22-app/controller/user"
	"todo-22-app/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	db.CreateDb()

	r := mux.NewRouter()
	r.HandleFunc("/signup", controller.SingUp)
	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/createtodo", todoController.CreateTodo)
	r.HandleFunc("/deletetodo", todoController.DeleteTodo)
	r.HandleFunc("/edittodo", todoController.EditTodo)
	r.HandleFunc("/completetodo", todoController.DoneTodo)
	r.HandleFunc("/getusertodoList", todoController.GetTodoListWithUserId)

	return r

}
