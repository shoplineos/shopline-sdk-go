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

func (req *GetProductCountAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *GetProductCountAPIReq) Endpoint() string {
	endpoint := "products/count.json"
	return endpoint
}

type GetProductCountAPIResp struct {
	Count int `json:"count"`

	client.BaseAPIResponse
}

// GetProductsCount
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-product-count?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/batch-query-product-quantity?version=v20251201
// Deprecated
// see ProductService
func GetProductsCount(c *client.Client, apiReq *GetProductCountAPIReq) (*GetProductCountAPIResp, error) {

	// 1. API request
	shoplineReq := &client.ShopLineRequest{
		Query: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &GetProductCountAPIResp{}

	// 4. Call API
	_, err := c.Get(context.Background(), endpoint, shoplineReq, apiResp)

	//apiResp.TraceId = shoplineResp.TraceId

	return apiResp, err
}
