package order

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CancelOrderAPIRequest
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/cancel-order?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/cancel-order?version=v20251201
// Deprecated
// See CancelOrderAPIReq
type CancelOrderAPIRequest struct {
	client.BaseAPIRequest

	OrderId      string
	Amount       string `json:"amount,omitempty"`
	CancelReason string `json:"cancel_reason,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Email        string `json:"email,omitempty"`
	ProcessedAt  string `json:"processed_at,omitempty"`
	RefundType   string `json:"refund_type,omitempty"`
	Restock      string `json:"restock,omitempty"`
}

func (req *CancelOrderAPIRequest) GetMethod() string {
	return "POST"
}

func (req *CancelOrderAPIRequest) GetData() interface{} {
	return req
}

func (req *CancelOrderAPIRequest) Verify() error {
	// Verify the api request params
	return nil
}

func (req *CancelOrderAPIRequest) GetEndpoint() string {
	endpoint := fmt.Sprintf("orders/%s/cancel.json", req.OrderId)
	return endpoint
}

// CancelOrderAPIResponse Define the request structure for cancel an order (corresponding to the API request body)
type CancelOrderAPIResponse struct {
	client.BaseAPIResponse
	Order Order `json:"order"`
}

// CancelOrder
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/cancel-order?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/cancel-order?version=v20251201
// Deprecated
// see OrderService
func CancelOrder(c *client.Client, req *CancelOrderAPIRequest) (*CancelOrderAPIResponse, error) {

	// 1. API response data
	apiResp := &CancelOrderAPIResponse{}

	// 2. Call the API
	err := c.Call(context.Background(), req, apiResp)
	return apiResp, err
}
