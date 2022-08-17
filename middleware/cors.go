package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func Cors() *cors.Cors {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	return cors

}
