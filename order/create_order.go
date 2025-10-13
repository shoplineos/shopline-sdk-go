package order

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type CreateOrderAPIReq struct {
	Order Order `json:"order"`
}

func (req *CreateOrderAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *CreateOrderAPIReq) Endpoint() string {
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
func CreateOrder(c *client.Client, apiReq *CreateOrderAPIReq) (*CreateOrderAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &CreateOrderAPIResp{}

	// 4. Call API
	_, err := c.Post(context.Background(), endpoint, shopLineReq, apiResp)
	if err != nil {
		fmt.Printf("Execute Request failed，endpoint: %s, shopLineReq: %v, err: %v\n", endpoint, shopLineReq, err)
		return nil, err
	}

	return apiResp, nil
}
