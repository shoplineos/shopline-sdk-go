package model

type APIRequest interface {
	Verify() error
	Endpoint() string
}
