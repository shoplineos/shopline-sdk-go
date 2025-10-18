package order

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type CreateOrderAPIReq struct {
	client.BaseAPIRequest
	Order Order `json:"order"`
}

func (req *CreateOrderAPIReq) GetMethod() string {
	return "POST"
}

func (req *CreateOrderAPIReq) GetData() interface{} {
	return req
}

func (req *CreateOrderAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *CreateOrderAPIReq) GetEndpoint() string {
	endpoint := "orders.json"
	return endpoint
}

// CreateOrderAPIResp Define the request structure for creating an order (corresponding to the API request body)
type CreateOrderAPIResp struct {
	client.BaseAPIResponse
	Order Order `json:"order"`
}

// CreateOrder
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/create-an-order?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/create-an-order?version=v20251201
// Deprecated
// see OrderService
func CreateOrder(c *client.Client, req *CreateOrderAPIReq) (*CreateOrderAPIResp, error) {

	// 1. API response data
	apiResp := &CreateOrderAPIResp{}

	// 2. Call the API
	err := c.Call(context.Background(), req, apiResp)
	return apiResp, err
}
