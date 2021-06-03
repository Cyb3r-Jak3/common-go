package common_test

import (
	"testing"

	common "github.com/Cyb3r-Jak3/common/go"
)

type requestBody struct {
	Name string `json:"name"`
}

type responseBody struct {
	Test requestBody `json:"url"`
}

func TestEmptyResponse(t *testing.T) {
	resp, err := common.DoJSONRequest(
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
	resp, err := common.DoJSONRequest(
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
	resp, err := common.DoJSONRequest(
		"", "example.com", nil, nil,
	)
	if err == nil {
		t.Error("Got no error and wanted one")
	} else if err.Error() != "Post \"example.com\": unsupported protocol scheme \"\"" {
		t.Errorf("Wanted bad protocol scheme and got %s", err)
	}
	if resp != nil {
		t.Error("Wanted empty response and it was not")
	}
}

func TestGoodRequest(t *testing.T) {
	var respBody responseBody
	resp, err := common.DoJSONRequest(
		"GET", "https://httpbin.org/anything", &requestBody{Name: "Hello"}, &respBody,
	)
	if err != nil {
		t.Errorf("Wanted no error and got %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Wanted status code and got %d", resp.StatusCode)
	}
}
