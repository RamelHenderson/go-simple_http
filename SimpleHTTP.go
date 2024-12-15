package main

import (
	_ "embed"
	"encoding/json"
	"io"
	"net/http"
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

	// Read the mappedResponse data
	data, err := io.ReadAll(clientResponse.Body)
	if err != nil {
		panic(err)
	}

	// Create the JSON response
	jsonData := make(map[string]interface{})
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		panic(err)
	}
	mappedResponse := make(map[string]interface{})
	mappedResponse["data"] = jsonData
	mappedResponse["status"] = clientResponse.StatusCode
	return mappedResponse, err
}

// PrettyPrintJSON pretty prints the specified JSON data
func PrettyPrintJSON(data map[string]interface{}) string {
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(prettyJSON)
}
