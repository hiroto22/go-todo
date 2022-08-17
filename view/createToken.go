package view

import (
	"encoding/json"
	"net/http"
)

type tokenRes struct {
	Token string `json:"token"`
}

func CreateToken(w http.ResponseWriter, data string) {
	tokenData := tokenRes{data}
	json.NewEncoder(w).Encode(tokenData)
}
