package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"todo-22-app/auth"
	"todo-22-app/model"
)

type LoginState struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}
type tokenRes struct {
	Token string `json:"token"`
}

//loginの際に使うAPI
func Login(w http.ResponseWriter, r *http.Request) {
	//CORS
	w.Header().Set("Content-Type", "applicaiton/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	var data LoginState
	json.Unmarshal(body, &data)
	email := data.Email

	user := model.Login(body, email)

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

		tokenData := tokenRes{token}

		json.NewEncoder(w).Encode(tokenData)
	}

}
