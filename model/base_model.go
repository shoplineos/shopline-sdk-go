package model

type APIRequest interface {
	Verify() error    // Verify API request params
	Endpoint() string // API Endpoint
}
