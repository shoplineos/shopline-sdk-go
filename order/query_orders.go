package order

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/manager"
)

type QueryOrdersAPIReq struct {
	ContractIDs       string // Contract IDs, Separate multiple with commas
	CreatedAtMin      string // Minimum order creation time（ISO 8601）
	Name              string // Order name
	SortCondition     string // Sort Condition（eg: "order_at:desc"）
	UpdatedAtMax      string // Max order update time（ISO 8601）
	FulfillmentStatus string // Fulfillment Status（unshipped/partial/shipped）
	HiddenOrder       string // Is Hidden Order（true/false）
	IDs               string // Order ids, Separate multiple with commas
	Limit             string // Limit（Max 100）
	Location          string // Location ids, Separate multiple with commas
	Status            string // Status（open/cancelled/any）
	BuyerID           string // Buyer ID
	Email             string // Email
	Fields            string // Fields, Separate multiple with commas
	FinancialStatus   string // Financial Status（unpaid/authorized）
	SearchContent     string // Search Content（Order id、Product title and so on）
	UpdatedAtMin      string // Minimum order update time（ISO 8601）
	CreatedAtMax      string // Max order creation time（ISO 8601）
	PageInfo          string // Page Info
	SinceID           string // Order Since ID
}

type QueryOrdersAPIResp struct {
	Orders []Order `json:"orders"`

	client.CommonAPIRespData
	Pagination *client.Pagination
}

// QueryOrdersV2
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/get-orders?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/get-orders?version=v20251201
func QueryOrdersV2(appkey, storeHandle string, apiReq *QueryOrdersAPIReq) (*QueryOrdersAPIResp, error) {

	// 1. API request
	shoplineReq := &client.ShopLineRequest{
		Query: apiReq,
	}

	// 2. API endpoint
	endpoint := "orders.json"

	// 3. API response
	apiResp := &QueryOrdersAPIResp{}

	// 4. Invoke API
	shoplineResp, err := manager.GetClient(appkey, storeHandle).Get(context.Background(), endpoint, shoplineReq, apiResp)
	if err != nil {
		fmt.Printf("Execute request failed: %v\n", err)
		return nil, err
	}

	apiResp.Pagination = shoplineResp.Pagination

	return apiResp, err
}

// QueryOrders
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/get-orders?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/get-orders?version=v20251201
func QueryOrders(apiReq *QueryOrdersAPIReq) (*QueryOrdersAPIResp, error) {
	return QueryOrdersV2(manager.GetDefaultClient().GetAppKey(), manager.GetDefaultClient().StoreHandle, apiReq)
}
