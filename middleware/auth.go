package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//token認証
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
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
