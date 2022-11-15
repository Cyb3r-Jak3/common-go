package common

import (
	"net/http"
	"os"
	"strings"
	"testing"
)

type requestBody struct {
	Name string `json:"name"`
}

type responseBody struct {
	Test requestBody `json:"url"`
}

// Hello is a simple hello function
func StringTest(w http.ResponseWriter, _ *http.Request) { StringResponse(w, "Hello") }

func JSONTest(w http.ResponseWriter, _ *http.Request) {
	JSONResponse(w, []byte(`"hello": "world"`))
}

func JSONMarshalTest(w http.ResponseWriter, _ *http.Request) {
	JSONMarshalResponse(w, &requestBody{Name: "test"})
}

func JSONMarshalBadTest(w http.ResponseWriter, _ *http.Request) {
	JSONMarshalResponse(w, nil)
}

func ContentTest(w http.ResponseWriter, _ *http.Request) {
	ContentResponse(w, "test/content", []byte("Hello"))
}

func TestAllowedMethod(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, AllowedMethods(StringTest, "POST"))
	checkResponse(t, rr, http.StatusMethodNotAllowed)
	r, _ = http.NewRequest("POST", "/", nil)
	rr = executeRequest(r, AllowedMethods(StringTest, "GET,POST"))
	checkResponse(t, rr, http.StatusOK)
}

func TestJsonResponse(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, JSONTest)
	checkResponse(t, rr, http.StatusOK)
}

func TestJSONMarshall(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Wanted no error and got %s", err)
	}
	rr := executeRequest(r, JSONMarshalTest)
	resp := rr.Result()
	if resp.Header.Get("Content-Type") != JSONApplicationType {
		t.Errorf("Wanted JSON response and got %s", resp.Header.Get("Content-Type"))
	}
	checkResponse(t, rr, http.StatusOK)
}

func TestContentResponse(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, ContentTest)
	checkResponse(t, rr, http.StatusOK)
}

func TestEmptyResponse(t *testing.T) {
	resp, err := DoJSONRequest(
		"", "https://httpbin.org/anything", &requestBody{Name: "value"}, nil,
	)
	if err != nil {
		t.Errorf("Wanted no error and got %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Wanted status code and got %d", resp.StatusCode)
	}
}

func TestEmptyBody(t *testing.T) {
	resp, err := DoJSONRequest(
		"", "https://httpbin.org/anything", nil, nil,
	)
	if err != nil {
		t.Errorf("Wanted no error and got %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Wanted status code and got %d", resp.StatusCode)
	}
}

func TestErrors(t *testing.T) {
	resp, err := DoJSONRequest(
		"", "example.com", nil, nil,
	)
	if err == nil {
		t.Error("Got no error and wanted one")
	} else if !strings.Contains(err.Error(), "unsupported protocol scheme") {
		t.Errorf("Wanted bad protocol scheme and got %s", err)
	}
	if resp != nil {
		t.Error("Wanted empty response and it was not")
	}
}

func TestGoodRequest(t *testing.T) {
	var respBody responseBody
	resp, err := DoJSONRequest(
		"GET", "https://httpbin.org/anything", &requestBody{Name: "Hello"}, &respBody,
	)
	if err != nil {
		t.Errorf("Wanted no error and got %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Wanted status code and got %d", resp.StatusCode)
	}
}

func TestDownload(t *testing.T) {
	ok, err := DownloadFile(
		"https://raw.githubusercontent.com/Cyb3r-Jak3/Cyb3r-Jak3/main/README.md",
		"test.md",
	)
	if !ok || err != nil {
		t.Errorf("Download status: %t. Error Message: %s", ok, err)
	}
	os.Remove("test.md")
}

func TestFailedDownload(t *testing.T) {
	ok, err := DownloadFile(
		"",
		"test.md",
	)
	if ok || err == nil {
		t.Errorf("Download status: %t. Error Message: %s", ok, err)
	}
	os.Remove("test.md")
}

func TestWriteDownload(t *testing.T) {
	ok, err := DownloadFile(
		"https://raw.githubusercontent.com/Cyb3r-Jak3/Cyb3r-Jak3/main/README.md",
		"",
	)
	if ok || err == nil {
		t.Errorf("Download status: %t. Error Message: %s", ok, err)
	}
}
