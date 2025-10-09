package order

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// DeleteOrderAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/delete-an-order?version=v20251201
// en：https://developer.shopline.com/docs/admin-rest-api/order/order-management/delete-an-order?version=v20251201
type DeleteOrderAPIReq struct {
	OrderId string
}

func (req *DeleteOrderAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *DeleteOrderAPIReq) Endpoint() string {
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
func DeleteOrder(c *client.Client, apiReq *DeleteOrderAPIReq) (*DeleteOrderAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &DeleteOrderAPIResp{}

	// 4. Call API
	_, err := c.Delete(context.Background(), endpoint, shopLineReq, apiResp)
	if err != nil {
		fmt.Printf("Execute Request failed，endpoint: %s, shopLineReq: %v, err: %v\n", endpoint, shopLineReq, err)
		return nil, err
	}

	return apiResp, nil
}
