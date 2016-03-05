package kdbclient

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type File struct {
	Name, Mime string
	Data       []byte
}

type Client struct {
	host   string
	port   int
	client *http.Client
}

type IClient interface {
	NewClient(kdbHost string, kdbPort string) *Client
	Get(filename string) (File, error)
	Save(filename string, file File) bool
	Delete(filename string) bool
	Exists(filename string) bool
}

// Read file
func (c *Client) Get(filename string) (*File, error) {
	resp, err := http.Get(c.url(filename))
	defer resp.Body.Close()

	if err == nil && resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		return &File{
			Name: filepath.Base(filename),
			Mime: resp.Header.Get("Content-Type"),
			Data: body,
		}, nil
	}

	return nil, errors.New("Failed to get file.")
}

// Save file
func (c *Client) Save(filename string, file File) bool {
	resp, err := http.Post(c.url(filename), file.Mime, bytes.NewReader(file.Data))
	defer resp.Body.Close()

	if err == nil && resp.StatusCode == 200 {
		return true
	}

	return false
}

// Remove file
func (c *Client) Delete(filename string) bool {
	req, err := http.NewRequest("DELETE", c.url(filename), nil)
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err == nil && resp.StatusCode == 200 {
		return true
	}

	return false
}

// Check if file exists
func (c *Client) Exists(filename string) bool {
	resp, err := http.Head(c.url(filename))
	// log.Fatal(err)
	defer resp.Body.Close()

	if err == nil && resp.StatusCode == 200 {
		return true
	}

	return false
}

// Get url with host and port parts
func (c *Client) url(filename string) string {
	return fmt.Sprintf("http://%s:%d%s", c.host, c.port, filename)
}

// Create new instance of client
func NewClient(kdbHost string, kdbPort int) *Client {
	c := Client{
		host: kdbHost,
		port: kdbPort,
	}

	return &c
}
