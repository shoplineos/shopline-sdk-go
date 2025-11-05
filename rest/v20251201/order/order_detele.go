package order

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// DeleteOrderAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/delete-an-order?version=v20251201
// en：https://developer.shopline.com/docs/admin-rest-api/order/order-management/delete-an-order?version=v20251201
// Deprecated
// See DeleteAnOrderAPIReq
type DeleteOrderAPIReq struct {
	client.BaseAPIRequest
	OrderId string
}

func (req *DeleteOrderAPIReq) GetMethod() string {
	return "DELETE"
}

func (req *DeleteOrderAPIReq) GetData() interface{} {
	return req
}

func (req *DeleteOrderAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *DeleteOrderAPIReq) GetEndpoint() string {
	endpoint := fmt.Sprintf("orders/%s.json", req.OrderId)
	return endpoint
}

// DeleteOrderAPIResp Define the request structure for delete an order (corresponding to the API request body)
type DeleteOrderAPIResp struct {
	client.BaseAPIResponse
}

// DeleteOrder
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/delete-an-order?version=v20251201
// en：https://developer.shopline.com/docs/admin-rest-api/order/order-management/delete-an-order?version=v20251201
// Deprecated
// see OrderService
func DeleteOrder(c *client.Client, apiReq *DeleteOrderAPIReq) (*DeleteOrderAPIResp, error) {
	// 1. API response resource
	apiResp := &DeleteOrderAPIResp{}

	// 2. Call the API
	err := c.Call(context.Background(), apiReq, apiResp)
	return apiResp, err
}
