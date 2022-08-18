package server

import (
	todocontroller "todo-22-app/controller/todoController"
	usercontroller "todo-22-app/controller/userController"
	"todo-22-app/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()
	//cors処理
	r.Use(middleware.Cors().Handler)

	//token発行する
	r.HandleFunc("/signup", usercontroller.SingUp)
	r.HandleFunc("/login", usercontroller.Login)

	//token認証あり
	auth := r.PathPrefix("/").Subrouter()
	auth.Use(middleware.TokenVerify)
	auth.HandleFunc("/deletetodo", todocontroller.DeleteTodo)
	auth.HandleFunc("/edittodo", todocontroller.EditTodo)
	auth.HandleFunc("/completetodo", todocontroller.DoneTodo)
	auth.HandleFunc("/createtodo", todocontroller.CreateTodo)
	auth.HandleFunc("/getusertodoList", todocontroller.GetTodoListWithUserId)

	return r

}
