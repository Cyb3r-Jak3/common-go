package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// JSONApplicationType is MIME type for JSON data.
const JSONApplicationType = "application/json; charset=utf-8"

// AllowedMethods is a decorator to only allow comma separated methods.
func AllowedMethods(handler http.HandlerFunc, methods string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if StringSearch(req.Method, strings.Split(methods, ",")) {
			handler(w, req)
			return
		}
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

// StringResponse writes a http response as a string.
func StringResponse(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(response)); err != nil {
		log.Printf("Error with writing string response: %s\n", err)
	}
}

// JSONResponse writes a http response as JSON. Takes a byte array as input.
func JSONResponse(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", JSONApplicationType)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		log.Printf("Error with writing JSON response: %s\n", err)
	}
}

// JSONMarshalResponse writes a http response as JSON. Takes an interface.
func JSONMarshalResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", JSONApplicationType)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("Error with writing JSON response: %s\n", err)
	}
}

// ContentResponse writes a http response with a given content type.
func ContentResponse(w http.ResponseWriter, contentType string, response []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		log.Printf("Error with writing content type '%s' response: %s\n", contentType, err)
	}
}

// DoJSONRequest sends a client JSON request. The responseBody should be a pointer to the address of a struct.
// If a blank string is passed then it will default to a POST request.
// Example:
//
// var response exampleStruct
//
// resp, err := DoJSONRequest("POST", "http://example.com", nil, &response)
//
// or
//
// response := new(exampleStruct)
//
// resp, err := DoJSONRequest("POST", "http://example.com", nil, response).
func DoJSONRequest(method, url string, requestBody, responseBody interface{}) (*http.Response, error) {
	if method == "" {
		method = "POST"
	}
	payloadBuf := new(bytes.Buffer)
	if requestBody != nil {
		if err := json.NewEncoder(payloadBuf).Encode(requestBody); err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, url, payloadBuf)
	if err != nil {
		return nil, err
	}
	// Set headers to give best chance at JSON response
	req.Header.Add("content-type", JSONApplicationType)
	req.Header.Add("Accept", JSONApplicationType)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	switch responseBody := responseBody.(type) {
	case nil:
	default:
		decErr := json.NewDecoder(resp.Body).Decode(responseBody)
		// ignore error if the JSON body is empty
		if errors.Is(decErr, io.EOF) {
			err = nil
		}
	}
	resp.Body.Close()
	return resp, err
}

// DownloadFile downloads a file and writes it to the file path. Overwrites any file at the path.
func DownloadFile(url, filename string) (ok bool, err error) {
	out, err := os.Create(filename)
	if err != nil {
		return false, err
	}
	resp, err := http.Get(url) //#nosec G107 User should validate URL
	if err != nil {
		return false, err
	}
	if _, err := io.Copy(out, resp.Body); err != nil {
		return false, err
	}
	if err := resp.Body.Close(); err != nil {
		return false, err
	}
	if err := out.Close(); err != nil {
		return false, err
	}
	return true, nil
}
