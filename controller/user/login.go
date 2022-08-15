package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"todo-22-app/auth"
	"todo-22-app/db"
	model "todo-22-app/model/user"
	"todo-22-app/view"
)

type LoginState struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

//loginの際に使うAPI
func Login(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectedDb()
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
	user := model.NewUser()
	user.Login(email, db)

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
