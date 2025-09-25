package order

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
	"log"
)

type GetOrdersCountAPIReq struct {
	Status            string `url:"status,omitempty"`           // Status（open/cancelled/any）
	FinancialStatus   string `url:"contract_ids,omitempty"`     // Financial Status（unpaid/authorized）
	FulfillmentStatus string `url:"financial_status,omitempty"` // Fulfillment Status（unshipped/partial/shipped）
	CreatedAtMin      string `url:"created_at_min,omitempty"`   // Minimum order creation time（ISO 8601）
	CreatedAtMax      string `url:"created_at_max,omitempty"`   // Max order creation time（ISO 8601）
	UpdatedAtMin      string `url:"updated_at_min,omitempty"`   // Minimum order update time（ISO 8601）
	UpdatedAtMax      string `url:"updated_at_max,omitempty"`   // Max order update time（ISO 8601）
	OrderSource       string `url:"order_source,omitempty"`
}

type GetOrdersCountAPIResp struct {
	Count int `json:"count"`
	client.CommonAPIRespData
}

// QueryOrdersCount
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/get-order-quantity?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/get-order-quantity?version=v20251201
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
