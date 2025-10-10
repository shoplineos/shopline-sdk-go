package order

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type ListOrdersAPIReq struct {
	ContractIDs       string `url:"contract_ids,omitempty"`       // Contract IDs, Separate multiple with commas
	FulfillmentStatus string `url:"fulfillment_status,omitempty"` // Fulfillment Status（unshipped/partial/shipped）

	Email         string `url:"email,omitempty"`          // Email
	Name          string `url:"name,omitempty"`           // Order name
	SortCondition string `url:"sort_condition,omitempty"` // Sort Condition（eg: "order_at:desc"）
	CreatedAtMin  string `url:"created_at_min,omitempty"` // Minimum order creation time（ISO 8601）
	CreatedAtMax  string `url:"created_at_max,omitempty"` // Max order creation time（ISO 8601）
	UpdatedAtMin  string `url:"updated_at_min,omitempty"` // Minimum order update time（ISO 8601）
	UpdatedAtMax  string `url:"updated_at_max,omitempty"` // Max order update time（ISO 8601）
	Location      string `url:"location,omitempty"`       // Location ids, Separate multiple with commas
	SearchContent string `url:"search_content,omitempty"` // Search Content（Order id、Product title and so on）

	BuyerID     string `url:"buyer_id,omitempty"`     // Buyer ID
	Fields      string `url:"fields,omitempty"`       // Fields, Separate multiple with commas
	HiddenOrder string `url:"hidden_order,omitempty"` // Is Hidden Order（true/false）
	IDs         string `url:"ids,omitempty"`          // Order ids, Separate multiple with commas

	Status          string `url:"status,omitempty"`           // Status（open/cancelled/any）
	FinancialStatus string `url:"financial_status,omitempty"` // Financial Status（unpaid/authorized）

	SinceID  string `url:"since_id,omitempty"`  // Order Since ID
	Limit    string `url:"limit,omitempty"`     // Limit
	PageInfo string `url:"page_info,omitempty"` // Page Info
}

func (req *ListOrdersAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *ListOrdersAPIReq) Endpoint() string {
	endpoint := "orders.json"
	return endpoint
}

type ListOrdersAPIResp struct {
	Orders []Order `json:"orders"`

	client.BaseAPIResponse
}

// ListOrders
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/get-orders?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/get-orders?version=v20251201
// Deprecated
// see OrderService
func ListOrders(c *client.Client, apiReq *ListOrdersAPIReq) (*ListOrdersAPIResp, error) {
	// 1. API request
	shoplineReq := &client.ShopLineRequest{
		Query: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &ListOrdersAPIResp{}

	// 4. Call API
	shoplineResp, err := c.Get(context.Background(), endpoint, shoplineReq, apiResp)
	if err != nil {
		fmt.Printf("Execute request failed: %v\n", err)
		return nil, err
	}

	apiResp.Pagination = shoplineResp.Pagination

	return apiResp, err
}

// ListOrdersAll
// Deprecated
// see OrderService
func ListOrdersAll(c *client.Client, apiReq *ListOrdersAPIReq) ([]Order, error) {
	collector := []Order{}
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	for {
		// 2. API endpoint
		endpoint := apiReq.Endpoint()

		// 3. API response data
		apiResp := &ListOrdersAPIResp{}

		// 4. Call API
		shoplineResp, err := c.Get(context.Background(), endpoint, shopLineReq, apiResp)

		if err != nil {
			return collector, err
		}

		collector = append(collector, apiResp.Orders...)

		if !shoplineResp.HasNext() {
			break
		}

		shopLineReq.Query = shoplineResp.Pagination.Next
	}

	return collector, nil
}
