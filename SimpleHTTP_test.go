package main

import (
	"os"
	"testing"
)

func TestParameterRequest_Send(t *testing.T) {
	pr, err := NewParameterRequest(GET, "https://pokeapi.co/api/v2/pokemon")
	if err != nil {
		panic(err)
	}
	pr.AddParameter("limit", "2")

	// Send the request
	response, err := pr.Send()
	if err != nil {
		panic(err)
	}
	t.Log(response)
}

func TestFormRequest_Send(t *testing.T) {
	TestFilePath := "test_image.png"
	request, err := NewFormRequest(POST, "https://postman-echo.com/post")
	if err != nil {
		panic(err)
	}

	// Add a file as a byte array
	err = request.AddFileData("bytes", "fake.txt", []byte("SET BY BYTE ARRAY"))
	if err != nil {
		panic(err)
	}

	// Add a file by path
	f, err := os.Open(TestFilePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()
	_, err = request.AddFile("file", f)
	if err != nil {
		panic(err)
	}

	// Send the request
	response, err := request.Send()
	if err != nil {
		panic(err)
	}
	t.Log(response)
}
