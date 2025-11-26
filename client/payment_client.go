package client

import (
	"context"
	"errors"
	"github.com/fatih/structs"
	"net/http"
)

// PaymentClient
// Use NewPaymentClient to new a PaymentClient
// 中文:https://developer.shopline.com/zh-hans-cn/docs/apps/payment-application/overview?version=v20260301
// En:https://developer.shopline.com/docs/apps/payment-application/overview?version=v20260301
type PaymentClient struct {
	cli                       *Client
	paymentSignatureAlgorithm *PaymentSignatureAlgorithm
}

func NewPaymentClient(cli *Client, privateKey string, publicKey string) *PaymentClient {
	alg := NewPaymentSignatureAlgorithm(privateKey, publicKey)
	return &PaymentClient{
		cli:                       cli,
		paymentSignatureAlgorithm: alg,
	}
}

// Call an API
// resource : An API response resource, to specify the return type of the request, an APIResponse or your own resource
func (c *PaymentClient) Call(ctx context.Context, req APIRequest, resource interface{}) error {
	if req == nil {
		return errors.New("request is required")
	}
	if resource == nil {
		return errors.New("resource is required")
	}

	// 1. New a SHOPLINE API request
	shopLineReq := newShopLineRequest(req)

	// 2. Call an API
	_, err := c.executeInternal(ctx, HTTPMethod(req.GetMethod()), req.GetEndpoint(), shopLineReq, resource)
	return err
}

func (c *PaymentClient) executeInternal(ctx context.Context, method HTTPMethod, endpoint string, request *ShopLineRequest, resource interface{}) (*ShopLineResponse, error) {
	err := c.cli.Verify(endpoint, method, request)
	if err != nil {
		return nil, err
	}

	relPath := c.cli.resolveUrlPath(endpoint, request)

	httpReq, _, err := c.cli.NewHttpRequestWithoutHeaders(ctx, method, relPath, request)
	if err != nil {
		return nil, err
	}

	c.setPaymentHeaders(request, httpReq)

	return c.cli.executeHttpRequest(request, httpReq, resource)
}

// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/payment-application/payment-application-signature-logic?version=v20260301
// En: https://developer.shopline.com/docs/apps/payment-application/payment-application-signature-logic?version=v20260301
func (c *PaymentClient) setPaymentHeaders(request *ShopLineRequest, httpReq *http.Request) error {
	// Set common headers
	c.cli.setCommonHeaders(httpReq)

	timestamp := BuildTimestamp()
	httpReq.Header.Set("pay-api-timestamp", timestamp)

	var resourceMap map[string]interface{}
	if request.Data != nil {
		resourceMap = structs.Map(request.Data)
	} else if request.Query != nil {
		resourceMap = structs.Map(request.Query)
	}

	if resourceMap != nil {
		sign, err := c.paymentSignatureAlgorithm.Signature(resourceMap)
		if err != nil {
			return err
		}
		httpReq.Header.Set("pay-api-signature", sign)
	}

	setRequestHeadersIfNecessary(request, httpReq)
	return nil
}
