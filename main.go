package main

import (
	"net/http"
	"os"
	"todo-22-app/db"
	"todo-22-app/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db.CreateDb()

	server := server.Router()
	http.ListenAndServe(":"+os.Getenv("PORT"), server)

}
