package buda

import (
	"net/http"
)

const (
	name    = "buda.com"
	baseURL = "https://www.buda.com/api/v2"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Buda ...
type Buda struct {
	key        string
	secret     string
	baseURL    string
	shouldStop bool

	// http values
	client HTTPClient
}

// New ...
func New(key, secret string, client HTTPClient) *Buda {
	return &Buda{
		key:     key,
		secret:  secret,
		baseURL: baseURL,
		client:  client,
	}
}
