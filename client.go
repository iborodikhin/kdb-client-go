package kdb_client

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type File struct {
	name, mime string
	data       []byte
}

type Client struct {
	Host   string
	Port   int
	client *http.Client
}

type IClient interface {
	Get(filename string) (File, error)
	Save(filename string, file File) bool
	Delete(filename string) bool
	Exists(filename string) bool
}

// Read file
func (c Client) Get(filename string) (*File, error) {
	resp, err := http.Get(c.url(filename))
	defer resp.Body.Close()

	if err == nil && resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		return &File{
			name: filepath.Base(filename),
			mime: resp.Header.Get("Content-Type"),
			data: body,
		}, nil
	}

	return nil, errors.New("Failed to get file.")
}

// Save file
func (c Client) Save(filename string, file File) bool {
	resp, err := http.Post(c.url(filename), file.mime, bytes.NewReader(file.data))
	defer resp.Body.Close()

	if err == nil && resp.StatusCode == 200 {
		return true
	}

	return false
}

// Remove file
func (c Client) Delete(filename string) bool {
	req, err := http.NewRequest("DELETE", c.url(filename), nil)
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err == nil && resp.StatusCode == 200 {
		return true
	}

	return false
}

// Check if file exists
func (c Client) Exists(filename string) bool {
	resp, err := http.Head(c.url(filename))
	defer resp.Body.Close()

	if err == nil && resp.StatusCode == 200 {
		return true
	}

	return false
}

// Get url with host and port parts
func (c Client) url(filename string) string {
	return fmt.Sprintf("http://%s:%d%s", c.Host, c.Port, filename)
}
