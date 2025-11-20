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
	return NewRetryAPIClientWithRetries(cli, maxRetries)
}

func NewRetryAPIClientWithRetries(cli *Client, retries int) IClient {
	if cli == nil {
		panic("cli is required")
	}
	return &RetryAPIClient{
		delegate: cli,
		retries:  retries,
	}
}

func (cli *RetryAPIClient) CreateAccessToken(ctx context.Context, code string) (*TokenResponse, error) {
	return cli.delegate.CreateAccessToken(ctx, code)
}

func (cli *RetryAPIClient) RefreshAccessToken(ctx context.Context, storeHandle string) (*TokenResponse, error) {
	return cli.delegate.RefreshAccessToken(ctx, storeHandle)
}

func (cli *RetryAPIClient) Get(ctx context.Context, endpoint string, req *ShopLineRequest, resource interface{}) (*ShopLineResponse, error) {
	var err error
	var resp *ShopLineResponse
	retries := cli.retries

	for {
		resp, err = cli.delegate.Get(ctx, endpoint, req, resource)
		if err == nil {
			break
		}

		if retries <= 1 {
			return resp, err
		}

		if !canRetry(err) {
			break
		}

		retries--
	}

	return resp, err
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
