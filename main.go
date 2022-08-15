package main

import (
	"net/http"
	"os"
	"todo-22-app/middleware"
	"todo-22-app/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cors := middleware.Cors()
	server := server.Router()
	handler := cors.Handler(server)
	http.ListenAndServe(":"+os.Getenv("PORT"), handler)

}
