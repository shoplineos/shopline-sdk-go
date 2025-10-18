package order

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// UpdateOrderAPIReq
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/update-an-order?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/update-an-order?version=v20251201
type UpdateOrderAPIReq struct {
	client.BaseAPIRequest
	Order Order `json:"order"`
}

func (req *UpdateOrderAPIReq) GetMethod() string {
	return "PUT"
}

func (r *UpdateOrderAPIReq) GetData() interface{} {
	return r
}

func (req *UpdateOrderAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *UpdateOrderAPIReq) GetEndpoint() string {
	endpoint := fmt.Sprintf("orders/%s.json", req.Order.Id)
	return endpoint
}

// UpdateOrderAPIResp Define the request structure for upate an order (corresponding to the API request body)
type UpdateOrderAPIResp struct {
	client.BaseAPIResponse
	Order Order `json:"order"`
}

// UpdateOrder
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/update-an-order?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/update-an-order?version=v20251201
// Deprecated
// see OrderService
func UpdateOrder(c *client.Client, req *UpdateOrderAPIReq) (*UpdateOrderAPIResp, error) {
	// 1. API response data
	apiResp := &UpdateOrderAPIResp{}

	// 2. Call API
	err := c.Call(context.Background(), req, apiResp)
	return apiResp, err
}
