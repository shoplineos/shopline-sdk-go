package order

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CancelOrderAPIReq
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/cancel-order?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/cancel-order?version=v20251201
type CancelOrderAPIReq struct {
	OrderId      string
	Amount       string `json:"amount,omitempty"`
	CancelReason string `json:"cancel_reason,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Email        string `json:"email,omitempty"`
	ProcessedAt  string `json:"processed_at,omitempty"`
	RefundType   string `json:"refund_type,omitempty"`
	Restock      string `json:"restock,omitempty"`
}

func (req *CancelOrderAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *CancelOrderAPIReq) Endpoint() string {
	endpoint := fmt.Sprintf("orders/%s/cancel.json", req.OrderId)
	return endpoint
}

// CancelOrderAPIResp Define the request structure for cancel an order (corresponding to the API request body)
type CancelOrderAPIResp struct {
	Order Order `json:"order"`
	client.BaseAPIResponse
}

// CancelOrder
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/cancel-order?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/cancel-order?version=v20251201
// Deprecated
// see OrderService
func CancelOrder(c *client.Client, apiReq *CancelOrderAPIReq) (*CancelOrderAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &CancelOrderAPIResp{}

	// 4. Call API
	_, err := c.Post(context.Background(), endpoint, shopLineReq, apiResp)
	if err != nil {
		fmt.Printf("Execute Request failed，endpoint: %s, shopLineReq: %v, err: %v\n", endpoint, shopLineReq, err)
		return nil, err
	}

	return apiResp, nil
}
