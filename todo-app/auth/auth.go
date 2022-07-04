package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type CreateTokenState struct {
	Id int `json:"id"`
}

func Test(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)

	}

	log.Printf("request body=%s\n", r.Body)

	var data CreateTokenState

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	id := data.Id

	token, err := CreateToken(id)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "applicaiton/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	json.NewEncoder(w).Encode(token)

}

func CreateToken(userid int) (string, error) {

	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"userid": userid,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}

	var secretKey = "gotodo"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func PasswordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

func TokenVerify(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("gotodo"), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
