package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//token認証と同時にtokenからuserIDを取得
func TokenVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Tokenをリクエストのheaderから取得
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("gotodo"), nil
		})
		if err != nil {
			http.Error(w, "invalid_access_token", http.StatusUnauthorized)
		}

		//tokenからuseIDを取得
		secretKey := os.Getenv("SECURITY_KEY")
		claims := jwt.MapClaims{}
		jwt.ParseWithClaims(tokenString, claims, func(userid *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		userID := claims["userid"]

		ctx := context.WithValue(r.Context(), "userID", userID)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
