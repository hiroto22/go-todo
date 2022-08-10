package todos

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"todo-22-app/auth"

	"github.com/joho/godotenv"
)

type EditTodoBody struct {
	Todo      string    `json:"todo"`
	UpdatedAt time.Time `json:"updatedat"`
}

//todoを編集できるAPI
func EditTodo(w http.ResponseWriter, r *http.Request) {
	//CORS
	w.Header().Set("Content-Type", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	e := godotenv.Load() //環境変数の読み込み
	if e != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	//tokenをrequestのheaderから取得
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//token認証
	_, err := auth.TokenVerify(tokenString)
	if err != nil {
		http.Error(w, "invalid_access_token", http.StatusUnauthorized)
		return
	}

	//MySQL接続
	dbConnectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConnectionInfo)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	defer db.Close()

	//todoのidを取得
	id := r.URL.Query().Get("id")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//todoの内容と変更日時を記録
	var data EditTodoBody
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "400 Bad Request", 400)
		return
	}

	editData := EditTodoBody{data.Todo, time.Now()}

	//todoを更新
	stmt, err := db.Prepare("UPDATE todos set Todo=?, UpdatedAt=? WHERE ID=?")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	_, err = stmt.Exec(editData.Todo, editData.UpdatedAt, id)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	json.NewEncoder(w).Encode(editData)
}
