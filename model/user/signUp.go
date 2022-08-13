package user

import (
	"time"
	"todo-22-app/db"
)

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

func (user *SignUpUser) SingUp(name string, email string, password string) error {
	db := db.ConnectDb()
	defer db.Close()

	nowTime := time.Now()
	//DBに送るuser情報
	userData := SignUpUser{name, email, password, nowTime, nowTime}

	//DBにuser情報を登録
	stmt, err := db.Prepare("INSERT INTO users (Name,Email,PassWord,CreatedAt,UpdatedAt) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userData.Name, userData.Email, userData.PassWord, userData.CreatedAt, userData.UpdatedAt)

	return err

}
