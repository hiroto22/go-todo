package users

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(u *User) {

}
