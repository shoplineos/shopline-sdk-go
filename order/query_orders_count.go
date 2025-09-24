package order

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
	"log"
)

type GetOrdersCountAPIReq struct {
	Status            string // Status（open/cancelled/any）
	FinancialStatus   string // Financial Status（unpaid/authorized）
	FulfillmentStatus string // Fulfillment Status（unshipped/partial/shipped）
	CreatedAtMin      string // Minimum order creation time（ISO 8601）
	CreatedAtMax      string // Max order creation time（ISO 8601）
	UpdatedAtMin      string // Minimum order update time（ISO 8601）
	UpdatedAtMax      string // Max order update time（ISO 8601）
	OrderSource       string
}

type GetOrdersCountAPIResp struct {
	Count int `json:"count"`
	client.CommonAPIRespData
}

// QueryOrdersCount
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/order-quantity-query?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/order-quantity-query/?version=v20251201
func QueryOrdersCount(c *client.Client, apiReq *GetOrdersCountAPIReq) (*GetOrdersCountAPIResp, error) {

	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := "orders/count.json"

	// 3. API response data
	apiResp := &GetOrdersCountAPIResp{}

	// 4. Invoke API
	_, err := c.Get(context.Background(), endpoint, shopLineReq, apiResp)
	if err != nil {
		log.Printf("Failed to Get: %v\n", err)
		return apiResp, err
	}

	//fmt.Printf("body:%s\n", string(body))

	return apiResp, nil
}
