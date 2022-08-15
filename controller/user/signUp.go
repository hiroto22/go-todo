package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"todo-22-app/auth"
	"todo-22-app/db"
	model "todo-22-app/model/user"
	"todo-22-app/view"

	"golang.org/x/crypto/bcrypt"
)

type SignUpState struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

//signup
func SingUp(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectedDb()
	//requestされたデータの読み込み
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	var data SignUpState
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "400 Bad Request", 400)
		return
	}
	name := data.Name
	email := data.Email
	password := data.PassWord

	if name == "" || email == "" || password == "" {
		http.Error(w, "400 Bad Request", 400)
		return
	}

	//passwordをhash化する
	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//データベースに情報を登録
	signupUser := model.NewSignUpUser()
	signupUser.SingUp(name, email, string(hashPassWord), db)

	//データベースから情報を取得
	user := model.NewUser()
	user.Login(email, db)

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	view.CreateToken(w, token)

}
