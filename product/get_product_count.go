package product

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type GetProductCountAPIReq struct {
	Status       string `url:"status,omitempty"`
	CollectionId string `url:"collection_id,omitempty"`
	CreatedAtMin string `url:"created_at_min,omitempty"` // Minimum order creation time（ISO 8601）
	CreatedAtMax string `url:"created_at_max,omitempty"` // Max order creation time（ISO 8601）
	UpdatedAtMin string `url:"updated_at_min,omitempty"` // Minimum order update time（ISO 8601）
	UpdatedAtMax string `url:"updated_at_max,omitempty"` // Max order update time（ISO 8601）
}

type GetProductCountAPIResp struct {
	Count int `json:"count"`

	client.CommonAPIRespData
}

// GetProductsCount
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-product-count?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/batch-query-product-quantity?version=v20251201
func GetProductsCount(c *client.Client, apiReq *GetProductCountAPIReq) (*GetProductCountAPIResp, error) {

	// 1. API request
	shoplineReq := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := "products/count.json"

	// 3. API response data
	apiResp := &GetProductCountAPIResp{}

	// 4. Invoke API
	_, err := c.Get(context.Background(), endpoint, shoplineReq, apiResp)

	//apiResp.TraceId = shoplineResp.TraceId

	return apiResp, err
}
