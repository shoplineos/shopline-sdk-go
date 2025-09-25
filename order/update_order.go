package order

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type UpdateOrderAPIReq struct {
	Order Order `json:"order"`
}

// UpdateOrderAPIResp Define the request structure for upate an order (corresponding to the API request body)
type UpdateOrderAPIResp struct {
	Order Order `json:"order"`
	client.CommonAPIRespData
}

// UpdateOrder
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/update-an-order?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/update-an-order?version=v20251201
func UpdateOrder(c *client.Client, apiReq *UpdateOrderAPIReq) (*UpdateOrderAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := fmt.Sprintf("orders/%s.json", apiReq.Order.ID)

	// 3. API response data
	apiResp := &UpdateOrderAPIResp{}

	// 4. Invoke API
	_, err := c.Put(context.Background(), endpoint, shopLineReq, apiResp)
	if err != nil {
		fmt.Printf("Execute Request failed，endpoint:%s, shopLineReq: %v, err: %v\n", endpoint, shopLineReq, err)
		return nil, err
	}

	return apiResp, nil
}
