package kdb_client

import (
	"io/ioutil"
	"testing"
)

func Test_Save(t *testing.T) {
	bytes, _ := ioutil.ReadFile("client_test.go")

	file := File{
		Name: "client_test.go",
		Mime: "text/plain",
		Data: bytes,
	}

	res := NewClient("127.0.0.1", 1337).Save("/test.go", file)

	if res == false {
		t.Error("Error while saving file")
	}
}

func Test_Get(t *testing.T) {
	f, err := NewClient("127.0.0.1", 1337).Get("/test.go")

	if f == nil || err != nil {
		t.Error("Error while reading file")
	}
}

func Test_Exists(t *testing.T) {
	res := NewClient("127.0.0.1", 1337).Exists("/test.go")

	if res == false {
		t.Error("Error while checking if file exists")
	}
}

func Test_Delete(t *testing.T) {
	res := NewClient("127.0.0.1", 1337).Delete("/test.go")

	if res == false {
		t.Error("Error while removing file")
	}
}
