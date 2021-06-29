package common

import (
	"log"
	"net/http"
	"strings"
)

// AllowedMethod is a decorator to get methods
func AllowedMethod(handler http.HandlerFunc, methods string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		for _, b := range strings.Split(methods, ",") {
			if b == req.Method {
				handler(w, req)
			}
		}
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

// StringResponse writes a http response as a string
func StringResponse(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(response)); err != nil {
		log.Printf("Error with writing string response: %s\n", err)
	}
}

// JSONResponse writes a http response as JSON
func JSONResponse(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		log.Printf("Error with writing JSON response: %s\n", err)
	}
}

// ContentResponse writes a http response with a given content type
func ContentResponse(w http.ResponseWriter, contentType string, response []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		log.Printf("Error with writing content type '%s' response: %s\n", contentType, err)
	}
}
