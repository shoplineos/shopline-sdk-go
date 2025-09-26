package client

import (
	"net/http"
	"time"
)

// Options client options
type Options struct {
	EnableLogDetail bool // log detail switch
}

// Option is used to configure client with options
type Option func(c *Client)

func WithVersion(apiVersion string) Option {
	return func(c *Client) {
		c.ApiVersion = apiVersion
	}
}

func WithEnableSign(enableSign bool) Option {
	return func(c *Client) {
		c.EnableSign = enableSign
	}
}

// WithHTTPClient is used to set a custom http client
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.Client = client
	}
}

// WithTimeout http timeout
func WithTimeout(Timeout time.Duration) Option {
	return func(c *Client) {
		c.Client.Timeout = Timeout
	}
}

func WithEnableLogDetail(enableLogDetail bool) Option {
	return func(c *Client) {
		if c.Options == nil {
			c.Options = &Options{}
		}
		c.Options.EnableLogDetail = enableLogDetail
	}
}
