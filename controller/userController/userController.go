package userController

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"todo-22-app/auth"
	usermodel "todo-22-app/model/userModel"
	"todo-22-app/view"

	"golang.org/x/crypto/bcrypt"
)

type LoginState struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

//loginの際に使うAPI
func Login(w http.ResponseWriter, r *http.Request) {
	//requestされたデータの読み込み
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	var data LoginState
	json.Unmarshal(body, &data)
	email := data.Email

	//データベースから情報を取得
	user := usermodel.NewUser()
	user.Login(email)

	//requestされたemail,passwordとDBの物が正しいか確認正しければtokenを返す
	err = auth.PasswordVerify(user.PassWord, data.PassWord)
	if err != nil {
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	} else {
		token, err := auth.CreateToken(user.ID)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

		view.CreateToken(w, token)
	}
}

type SignUpState struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

//signup
func SingUp(w http.ResponseWriter, r *http.Request) {
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
	signupUser := usermodel.NewSignUpUser()
	signupUser.SingUp(name, email, string(hashPassWord))

	//データベースから情報を取得
	user := usermodel.NewUser()
	user.Login(email)

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	view.CreateToken(w, token)

}
