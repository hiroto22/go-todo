package user

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

//Userの構造体を取得
func NewUser() *User {
	return new(User)
}

//login
func (user *User) Login(email string) {
	db := db.ConnectDb()
	defer db.Close()

	db.QueryRow("SELECT * FROM users WHERE Email=?", email).Scan(&user.ID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt, &user.UpdatedAt)
}
