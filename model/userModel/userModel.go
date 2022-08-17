package usermodel

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
func (user *User) Login(email string) error {
	if err := db.DB.QueryRow("SELECT * FROM users WHERE Email=?", email).Scan(&user.ID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return err
	}
	return nil
}

type SignUpUser struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PassWord  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

//Userの構造体を取得
func NewSignUpUser() *SignUpUser {
	return new(SignUpUser)
}

//signup
func (user *SignUpUser) SingUp(name string, email string, password string) error {
	nowTime := time.Now()
	//DBに送るuser情報
	user.Name = name
	user.Email = email
	user.PassWord = password
	user.CreatedAt = nowTime
	user.UpdatedAt = nowTime

	//DBにuser情報を登録
	_, err := db.DB.Exec("INSERT INTO users (Name,Email,PassWord,CreatedAt,UpdatedAt) VALUES(?,?,?,?,?)", user.Name, user.Email, user.PassWord, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return err

}
