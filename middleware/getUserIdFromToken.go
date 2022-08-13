package middleware

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

//tokenからuseIdを特定
func GetUserIdFromToken(tokenString string) interface{} {
	secretKey := os.Getenv("SECURITY_KEY")
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(userid *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	userID := claims["userid"]

	return userID
}
