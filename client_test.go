package kdb_client

import (
	"io/ioutil"
	"testing"
)

func Test_Save(t *testing.T) {
	client := Client{
		Host: "127.0.0.1",
		Port: 1337,
	}

	bytes, _ := ioutil.ReadFile("client_test.go")

	file := File{
		name: "client_test.go",
		mime: "text/plain",
		data: bytes,
	}
	res := client.Save("/test.go", file)

	if res == false {
		t.Error("Error while saving file")
	}
}

func Test_Get(t *testing.T) {
	client := Client{
		Host: "127.0.0.1",
		Port: 1337,
	}

	f, err := client.Get("/test.go")

	if f == nil || err != nil {
		t.Error("Error while reading file")
	}
}

func Test_Exists(t *testing.T) {
	client := Client{
		Host: "127.0.0.1",
		Port: 1337,
	}

	res := client.Exists("/test.go")

	if res == false {
		t.Error("Error while checking if file exists")
	}
}

func Test_Delete(t *testing.T) {
	client := Client{
		Host: "127.0.0.1",
		Port: 1337,
	}

	res := client.Delete("/test.go")

	if res == false {
		t.Error("Error while removing file")
	}
}
