package order

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type UpdateOrderAPIReq struct {
	Order Order `json:"order"`
}

func (req *UpdateOrderAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *UpdateOrderAPIReq) Endpoint() string {
	endpoint := fmt.Sprintf("orders/%s.json", req.Order.ID)
	return endpoint
}

// UpdateOrderAPIResp Define the request structure for upate an order (corresponding to the API request body)
type UpdateOrderAPIResp struct {
	Order Order `json:"order"`
	client.BaseAPIResponse
}

// UpdateOrder
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/update-an-order?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/update-an-order?version=v20251201
// Deprecated
// see OrderService
func UpdateOrder(c *client.Client, apiReq *UpdateOrderAPIReq) (*UpdateOrderAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &UpdateOrderAPIResp{}

	// 4. Call API
	_, err := c.Put(context.Background(), endpoint, shopLineReq, apiResp)
	if err != nil {
		fmt.Printf("Execute Request failed，endpoint: %s, shopLineReq: %v, err: %v\n", endpoint, shopLineReq, err)
		return nil, err
	}

	return apiResp, nil
}
