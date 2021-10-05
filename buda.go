package buda

import (
	"net/http"
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
func New(baseURL, key, secret string, client HTTPClient) *Buda {
	return &Buda{
		key:     key,
		secret:  secret,
		baseURL: baseURL,
		client:  client,
	}
}
