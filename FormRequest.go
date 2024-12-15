package simple_http

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type FormRequest struct {
	Url             string
	Method          string
	Headers         map[string]string
	MultipartWriter *multipart.Writer
	Body            bytes.Buffer
}

// NewFormRequest creates a new form request
// Parameters:
// - method: the method to use for the request
// - url: the URL to send the request to
// Returns:
// - a new form request
// - an error if there was an issue creating the form request
func NewFormRequest(method, url string) (*FormRequest, error) {
	sr := &FormRequest{
		Headers: make(map[string]string),
		Body:    bytes.Buffer{},
	}

	// Initialize the Writer
	sr.MultipartWriter = multipart.NewWriter(&sr.Body)

	// Validate the URL and assign it to the struct
	sr.Url = url
	if !strings.Contains(sr.Url, "http") {
		return nil, errors.New("Invalid URL: " + sr.Url)
	}

	// Validate the method and assign it to the struct
	// Upper case the method
	sr.Method = strings.ToUpper(method)
	if ValidateRequestMethod(method) != nil {
		return nil, errors.New("Invalid method: " + method)
	}

	return sr, nil
}

// AddFileData adds a file to the form request
// Parameters:
// - key: the key to use for the file
// - filename: the filename to use for the file
// - value: the file contents
// Returns:
// - an error if there was an issue adding the file
func (fr *FormRequest) AddFileData(key, filename string, value []byte) error {
	// Create the file writer
	fileWriter, err := fr.MultipartWriter.CreateFormFile(key, filename)
	if err != nil {
		return err
	}

	// Write the file to the writer
	_, err = io.CopyBuffer(fileWriter, bytes.NewReader(value), make([]byte, 64))
	if err != nil {
		return err
	}
	return nil
}

// AddFile adds a file to the form request
// Parameters:
// - osFile: the file to add
// Returns:
// - the number of bytes written
// - an error if there was an issue adding the file
func (fr *FormRequest) AddFile(key string, osFile *os.File) (int64, error) {
	fileWriter, _ := fr.MultipartWriter.CreateFormFile(key, osFile.Name())
	return io.Copy(fileWriter, osFile)
}

// AddField adds a field to the form request
// Parameters:
// - key: the key to use for the field
// - value: the value to use for the field
// Returns:
// - an error if there was an issue adding the field
func (fr *FormRequest) AddField(key, value string) error {
	err := fr.MultipartWriter.WriteField(key, value)
	if err != nil {
		return err
	}
	return err
}

// AddHeader adds a header to the form request
// Parameters:
// - key: the key to use for the header
// - value: the value to use for the header
func (fr *FormRequest) AddHeader(key, value string) {
	fr.Headers[key] = value
}

// Send sends the form request
// Returns:
// - a map of the JSON response
// - an error if there was an issue sending the request
func (fr *FormRequest) Send() (map[string]interface{}, error) {

	// Close all the open files
	err := fr.MultipartWriter.Close()
	if err != nil {
		return nil, err
	}

	// Create the request
	formRequest, err := http.NewRequest(fr.Method, fr.Url, &fr.Body)
	if err != nil {
		return nil, err
	}

	// Set the headers
	for key, value := range fr.Headers {
		formRequest.Header.Set(key, value)
	}

	// Set the content type
	formRequest.Header.Set("Content-Type", fr.MultipartWriter.FormDataContentType())

	// Set the content length
	if fr.Body.Len() > 0 {
		formRequest.Header.Set("Content-Length", strconv.Itoa(fr.Body.Len()))
	}

	return SendHttpRequest(formRequest)
}
