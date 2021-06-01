package common

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// JSONApplicationType is the MIME type for JSON requests/responses
const JSONApplicationType = "application/json"

// DoJSONRequest sends a client JSON request and decodes the body back to what should be a struct.
// If a blank string is passed then it will default to a POST request
func DoJSONRequest(method string, url string, requestBody interface{}, responseBody interface{}) (interface{}, error) {
	if method == "" {
		method = "POST"
	}
	payloadBuf := new(bytes.Buffer)
	if requestBody != nil {
		json.NewEncoder(payloadBuf).Encode(requestBody)
	}
	req, err := http.NewRequest(method, url, payloadBuf)
	if err != nil {
		log.Printf("Got error encoding: '%s'", err)
		return nil, err
	}
	//Set headers to give best change at JSON response
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
		if decErr == io.EOF {
			err = nil
		}
	}

	return responseBody, err
}
