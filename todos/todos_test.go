package todos

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestCreateTodo(t *testing.T) {
	expectedQueryString := "isdone=0"
	handler := func(w http.ResponseWriter, r *http.Request) {
		if expectedQueryString != r.URL.RawQuery {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	apitest.New(). // configuration
			HandlerFunc(handler).
			Get("/gettodo"). // request
			Query("isdone", "0").
			Header("Authorization", "5cCI6IkpXVCJ9.eyJleHAiOjE2NTkwMTI2MTQsInVzZXJpZCI6MTI0fQ.LwEAhIHZ3jZHAoRCVgA0fIlDJnUJ_ACCC8CBv_cGc60").
			Expect(t). // expectations
			Status(http.StatusOK).
			End()

}
