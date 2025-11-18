package client

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var retryAPIClient IClient

func setupRetryClient() {
	setup()

	retryAPIClient = NewRetryAPIClient(client)
}

func TestRetry(t *testing.T) {
	setupRetryClient()
	defer teardown()

	// test create product
	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", client.StoreHandle, client.PathPrefix, client.ApiVersion),
		httpmock.NewStringResponder(503, `{"errors":"service unavailable"}`))

	// 1. build request
	product := ProductStruct{
		Title:    "Hello shopline Freestyle 111",
		BodyHTML: "<strong>Hello shopline!<\\/strong>",
		Vendor:   "shopline",
	}

	apiReq := CreateProductAPIReqStruct{
		Product: product,
	}

	apiResp := &CreateProductAPIRespStruct{}

	// 2. Call the API
	err := retryAPIClient.Call(context.Background(), apiReq, apiResp)
	assert.NotNil(t, err)
	assert.Equal(t, "service unavailable", err.Error())
	//fmt.Printf("apiResp: %+v\n", apiResp.Product)

}

func TestNotRetry(t *testing.T) {
	setupRetryClient()
	defer teardown()

	// test create product
	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", client.StoreHandle, client.PathPrefix, client.ApiVersion),
		httpmock.NewStringResponder(504, `{"errors":"timeout"}`))

	// 1. build request
	product := ProductStruct{
		Title:    "Hello shopline Freestyle 111",
		BodyHTML: "<strong>Hello shopline!<\\/strong>",
		Vendor:   "shopline",
	}

	apiReq := CreateProductAPIReqStruct{
		Product: product,
	}

	apiResp := &CreateProductAPIRespStruct{}

	// 2. Call the API
	err := retryAPIClient.Call(context.Background(), apiReq, apiResp)
	assert.NotNil(t, err)
	assert.Equal(t, "timeout", err.Error())

	//fmt.Printf("apiResp: %+v\n", apiResp.Product)
	//assert.Equal(t, "111", apiResp.Product.Id)

}
