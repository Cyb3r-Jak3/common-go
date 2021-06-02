package common

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// JSONApplicationType is the MIME type for JSON requests/responses
const JSONApplicationType = "application/json"

// DoJSONRequest sends a client JSON request. The responseBody should be a pointer to the address of a struct.
// If a blank string is passed then it will default to a POST request.
// Example:
//
// var response exampleStruct
//
// if err := DoJSONRequest("POST", "http://example.com", nil, &response),
func DoJSONRequest(method string, url string, requestBody interface{}, responseBody interface{}) error {
	if method == "" {
		method = "POST"
	}
	payloadBuf := new(bytes.Buffer)
	if requestBody != nil {
		json.NewEncoder(payloadBuf).Encode(requestBody)
	}
	req, err := http.NewRequest(method, url, payloadBuf)
	if err != nil {
		return err
	}
	//Set headers to give best chance at JSON response
	req.Header.Add("content-type", JSONApplicationType)
	req.Header.Add("Accept", JSONApplicationType)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
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

	return err
}
