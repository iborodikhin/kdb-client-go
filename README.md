# kdb-client-go
Golang client for [KDB (Koraduba)](https://github.com/iborodikhin/kdb)

## Usage

Install in your `${GOPATH}` using `go get -u github.com/iborodikhin/kdb-client-go`

Then call it:
```go
package main

import (
  "github.com/iborodikhin.kdb-client-go"
)

func main() {
  // KDB connection parameters
  client := Client{
		Host: "127.0.0.1",
		Port: 1337,
	}

  // Read local file
	bytes, _ := ioutil.ReadFile("test_file.png")

  // Create file instance for upload
	file := File{
		name: "test_file.png",
		mime: "image/png",
		data: bytes,
	}
	
	// Save file to KDB
	res := client.Save("/test_file.png", file)
	
	// Check if file exists in KDB
	res := client.Exists("/test_file.png")
	
	// Read file from KDB
	f, err := client.Get("/test_file.png")
	// f is now of File type
	
	// Remove file from KDB
	res := client.Delete("/test_file.png")
}
```
