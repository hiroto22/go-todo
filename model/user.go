package model

import (
	"time"
	"todo-22-app/db"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PassWord  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type LoginState struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

//loginのdb操作
func Login(body []byte, email string) User {
	db := db.ConnectDb()
	defer db.Close()

	var user User
	db.QueryRow("SELECT * FROM users WHERE Email=?", email).Scan(&user.ID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt, &user.UpdatedAt)

	return user

}
