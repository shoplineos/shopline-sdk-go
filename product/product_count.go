package product

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetProductCountAPIReq
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-product-count?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/batch-query-product-quantity?version=v20251201
type GetProductCountAPIReq struct {
	client.BaseAPIRequest
	Status       string `url:"status,omitempty"`
	CollectionId string `url:"collection_id,omitempty"`
	CreatedAtMin string `url:"created_at_min,omitempty"` // Minimum order creation time（ISO 8601）
	CreatedAtMax string `url:"created_at_max,omitempty"` // Max order creation time（ISO 8601）
	UpdatedAtMin string `url:"updated_at_min,omitempty"` // Minimum order update time（ISO 8601）
	UpdatedAtMax string `url:"updated_at_max,omitempty"` // Max order update time（ISO 8601）
}

func (req *GetProductCountAPIReq) GetMethod() string {
	return "GET"
}

func (req *GetProductCountAPIReq) GetQuery() interface{} {
	return req
}

func (req *GetProductCountAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *GetProductCountAPIReq) GetEndpoint() string {
	return "products/count.json"
}

type GetProductCountAPIResp struct {
	client.BaseAPIResponse

	Count int `json:"count"`
}

// GetProductsCount
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-product-count?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/batch-query-product-quantity?version=v20251201
// Deprecated
// see ProductService
func GetProductsCount(c *client.Client, apiReq *GetProductCountAPIReq) (*GetProductCountAPIResp, error) {

	// 1. API response resource
	apiResp := &GetProductCountAPIResp{}

	// 2. Call the API
	err := c.Call(context.Background(), apiReq, apiResp)

	return apiResp, err
}
