package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type CreateTokenState struct {
	Id int `json:"id"`
}

//tokenを発行
func CreateToken(userid int) (string, error) {

	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"userid": userid,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}

	var secretKey = os.Getenv("SECURITY_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//password認証
func PasswordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

//token認証
func TokenVerify(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("gotodo"), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
