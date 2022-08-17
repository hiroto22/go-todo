package server

import (
	"todo-22-app/controller/todoController"
	"todo-22-app/controller/userController"
	"todo-22-app/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()
	//cors処理
	r.Use(middleware.Cors().Handler)

	r.HandleFunc("/signup", userController.SingUp)
	r.HandleFunc("/login", userController.Login)

	auth := r.PathPrefix("/").Subrouter()
	auth.Use(middleware.TokenVerify)
	auth.HandleFunc("/deletetodo", todoController.DeleteTodo)
	auth.HandleFunc("/edittodo", todoController.EditTodo)
	auth.HandleFunc("/completetodo", todoController.DoneTodo)
	auth.HandleFunc("/createtodo", todoController.CreateTodo)
	auth.HandleFunc("/getusertodoList", todoController.GetTodoListWithUserId)

	return r

}
