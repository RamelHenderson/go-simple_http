package main

import (
	_ "embed"
	"encoding/json"
	"io"
	"net/http"
)

const (
	POST   = "POST"
	GET    = "GET"
	PATCH  = "PATCH"
	DELETE = "DELETE"
	PUT    = "PUT"
)

// SendHttpRequest sends the specified request and returns the JSON response
// Parameters:
// - request: the request to send
// Returns:
// - a map of the JSON response
// - an error if there was an issue sending the request
func SendHttpRequest(request *http.Request) (map[string]interface{}, error) {
	// Create the client to send the request
	client := &http.Client{}

	// Send the request
	clientResponse, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer clientResponse.Body.Close()

	// Read the jsonResponse data
	data, err := io.ReadAll(clientResponse.Body)
	if err != nil {
		panic(err)
	}

	jsonResponse := make(map[string]interface{})
	jsonResponse["status"] = clientResponse.StatusCode
	jsonResponse["data"] = string(data)
	return jsonResponse, err
}

// PrettyPrintJSON pretty prints the specified JSON data
func PrettyPrintJSON(data map[string]interface{}) string {
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(prettyJSON)
}
