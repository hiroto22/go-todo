package main

import (
	"net/http"
	"os"
	"todo-22-app/db"
	"todo-22-app/server"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db.ConnectDb()
	defer db.DB.Close()

	server := server.Router()
	http.ListenAndServe(":"+os.Getenv("PORT"), server)

}
