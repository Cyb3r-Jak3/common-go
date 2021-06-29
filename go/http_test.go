package common_test

import (
	"net/http"
	"testing"

	common "github.com/Cyb3r-Jak3/common/go"
)

// Hello is a simple hello function
func StringTest(w http.ResponseWriter, _ *http.Request) { common.StringResponse(w, "Hello") }

func JSONTest(w http.ResponseWriter, _ *http.Request) {
	common.JSONResponse(w, []byte(`"hello": "world"`))
}

func ContentTest(w http.ResponseWriter, _ *http.Request) {
	common.ContentResponse(w, "test/content", []byte("Hello"))
}

func TestAllowedMethod(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, common.AllowedMethod(StringTest, "POST"))
	checkResponse(t, rr, http.StatusMethodNotAllowed)
	r, _ = http.NewRequest("GET", "/", nil)
	rr = executeRequest(r, common.AllowedMethod(StringTest, "GET,POST"))
	checkResponse(t, rr, http.StatusOK)
}

func TestJsonResponse(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, JSONTest)
	checkResponse(t, rr, http.StatusOK)
}

func TestContentResponse(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, ContentTest)
	checkResponse(t, rr, http.StatusOK)
}
