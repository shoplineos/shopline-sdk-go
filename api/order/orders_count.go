package order

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetOrdersCountAPIReq
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/get-order-quantity?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/get-order-quantity?version=v20251201
type GetOrdersCountAPIReq struct {
	client.BaseAPIRequest

	Status            string `url:"status,omitempty"`           // Status（open/cancelled/any）
	FinancialStatus   string `url:"contract_ids,omitempty"`     // Financial Status（unpaid/authorized）
	FulfillmentStatus string `url:"financial_status,omitempty"` // Fulfillment Status（unshipped/partial/shipped）
	CreatedAtMin      string `url:"created_at_min,omitempty"`   // Minimum order creation time（ISO 8601）
	CreatedAtMax      string `url:"created_at_max,omitempty"`   // Max order creation time（ISO 8601）
	UpdatedAtMin      string `url:"updated_at_min,omitempty"`   // Minimum order update time（ISO 8601）
	UpdatedAtMax      string `url:"updated_at_max,omitempty"`   // Max order update time（ISO 8601）
	OrderSource       string `url:"order_source,omitempty"`
}

func (req *GetOrdersCountAPIReq) GetMethod() string {
	return "GET"
}

func (req *GetOrdersCountAPIReq) GetQuery() interface{} {
	return req
}

func (req *GetOrdersCountAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *GetOrdersCountAPIReq) GetEndpoint() string {
	endpoint := "orders/count.json"
	return endpoint
}

type GetOrdersCountAPIResp struct {
	client.BaseAPIResponse
	Count int `json:"count"`
}

// QueryOrdersCount
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/get-order-quantity?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/get-order-quantity?version=v20251201
// Deprecated
// see OrderService
func QueryOrdersCount(c *client.Client, req *GetOrdersCountAPIReq) (*GetOrdersCountAPIResp, error) {
	// 1. API response data
	apiResp := &GetOrdersCountAPIResp{}

	// 2. Call the API
	err := c.Call(context.Background(), req, apiResp)
	return apiResp, err
}
