package middleware

import (
	"context"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type key string

var userIDKey key

//tokenからuseIdを特定
func SetUserIdFromToken(tokenString string) interface{} {
	secretKey := os.Getenv("SECURITY_KEY")
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(userid *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	userID := claims["userid"]

	return userID
}

func GetUserIdFromToken(ctx context.Context) int {
	v := ctx.Value(userIDKey).(int)
	return v
}
