package simple_http

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

type ParameterRequest struct {
	Url        string
	Method     string
	Headers    map[string]string
	Parameters map[string]string
}

// NewParameterRequest creates a new ParameterRequest with the specified method and URL
// Parameters:
// - method: the HTTP method to use
// - url: the URL to send the request to
func NewParameterRequest(method, url string) (*ParameterRequest, error) {
	pr := &ParameterRequest{}

	// Validate the URL and assign it to the struct
	if !strings.Contains(url, "http") {
		return nil, errors.New("Invalid URL: " + url)
	}
	pr.Url = url

	// Validate the method and assign it to the struct
	if ValidateRequestMethod(method) != nil {
		return nil, errors.New("Invalid method: " + method)
	}

	pr.Headers = make(map[string]string)
	pr.Parameters = make(map[string]string)
	return pr, nil
}

// AddParameter AddHeader adds a header to the request
// Parameters:
// - key: the key of the header
// - value: the value of the header
func (pr *ParameterRequest) AddParameter(key, value string) {
	pr.Parameters[key] = value
}

// Send sends the request to the specified URL with the specified method, headers, and parameters
// Returns:
// - a map of the JSON response
func (pr *ParameterRequest) Send() (map[string]interface{}, error) {
	// Create the parameterizedRequest
	parameterizedRequest, err := http.NewRequest(pr.Method, pr.Url, nil)
	if err != nil {
		log.Printf("Error creating parameterizedRequest: %v", err)
	}

	// Set the parameters
	query := parameterizedRequest.URL.Query()
	for key, value := range pr.Parameters {
		query.Add(key, value)
	}

	// Encode the query
	parameterizedRequest.URL.RawQuery = query.Encode()

	// Set the headers
	for key, value := range pr.Headers {
		parameterizedRequest.Header.Set(key, value)
	}

	// Send the request
	return SendHttpRequest(parameterizedRequest)
}
