# HTTP Client Utilities in Go

This Go module provides utilities for sending HTTP requests and processing JSON responses, including functionality for sending requests and pretty-printing JSON data.

## Features

- Send HTTP requests with various HTTP methods (GET, POST, PUT, PATCH, DELETE).
- Parse and return JSON responses as a map.
- Pretty-print JSON data for easier debugging and logging.

---

## Installation

Ensure you have [Go](https://golang.org/dl/) installed on your system.

Clone this repository and navigate to its directory:

```bash
git clone <repository-url>
cd <repository-name>
```

---

## Usage

### Sending HTTP Requests

You can use the `SendHttpRequest` function to send HTTP requests and parse JSON responses.

#### Example:

```go
package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	// Create a sample POST request
	url := "https://jsonplaceholder.typicode.com/posts"
	body := bytes.NewBuffer([]byte(`{"title": "foo", "body": "bar", "userId": 1}`))
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")

	// Send the request
	response, err := SendHttpRequest(request)
	if err != nil {
		panic(err)
	}

	// Pretty print the response
	fmt.Println(PrettyPrintJSON(response))
}
```

### Pretty Printing JSON

The `PrettyPrintJSON` function can be used to format JSON data for better readability.

#### Example:

```go
package main

import (
	"fmt"
)

func main() {
	jsonData := map[string]interface{}{
		"status": 200,
		"data":   "{"message":"Success"}",
	}

	// Pretty print the JSON
	fmt.Println(PrettyPrintJSON(jsonData))
}
```

---

## Functions

### `SendHttpRequest(request *http.Request) (map[string]interface{}, error)`

Sends an HTTP request and returns the JSON response.

- **Parameters**:
    - `request`: The `*http.Request` to send.
- **Returns**:
    - A map containing the JSON response.
    - An error if the request fails.

---

### `PrettyPrintJSON(data map[string]interface{}) string`

Pretty-prints a JSON object.

- **Parameters**:
    - `data`: A `map[string]interface{}` containing the JSON data.
- **Returns**:
    - A formatted JSON string.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
