package client

import (
	"context"
	"net/http"
)

const (
	maxRetries = DefaultMaxRetries
)

type RetryAPIClient struct {
	delegate *Client
	retries  int
}

func NewRetryAPIClient(cli *Client) IClient {
	if cli == nil {
		panic("cli is required")
	}
	return &RetryAPIClient{
		delegate: cli,
		retries:  maxRetries,
	}
}

func (cli *RetryAPIClient) Call(ctx context.Context, req APIRequest, resource interface{}) error {
	var err error
	retries := cli.retries

	for {
		err = cli.delegate.Call(ctx, req, resource)
		if err == nil {
			break
		}

		if retries <= 1 {
			return err
		}

		if !canRetry(err) {
			break
		}

		retries--
	}

	return err
}

// http status 503 can retry
func canRetry(err error) bool {
	if _, ok := err.(ResponseError); ok {
		e := err.(ResponseError)
		if http.StatusServiceUnavailable == e.Status {
			return true
		}
	}
	return false
}

func (cli *RetryAPIClient) NewHttpRequest(ctx context.Context, method HTTPMethod, path string, request *ShopLineRequest) (*http.Request, error) {
	return cli.delegate.NewHttpRequest(ctx, method, path, request)
}
