package main

import (
	"net/http"
	"os"
	"todo-22-app/db"
	"todo-22-app/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := db.ConnectDb()
	defer db.Close()

	server := server.Router()
	http.ListenAndServe(":"+os.Getenv("PORT"), server)

}
