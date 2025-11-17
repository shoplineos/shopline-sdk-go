package product

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetProductCountAPIRequest
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-product-count?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/batch-query-product-quantity?version=v20251201
// Deprecated
type GetProductCountAPIRequest struct {
	client.BaseAPIRequest
	Status       string `url:"status,omitempty"`
	CollectionId string `url:"collection_id,omitempty"`
	CreatedAtMin string `url:"created_at_min,omitempty"` // Minimum order creation time（ISO 8601）
	CreatedAtMax string `url:"created_at_max,omitempty"` // Max order creation time（ISO 8601）
	UpdatedAtMin string `url:"updated_at_min,omitempty"` // Minimum order update time（ISO 8601）
	UpdatedAtMax string `url:"updated_at_max,omitempty"` // Max order update time（ISO 8601）
}

func (req *GetProductCountAPIRequest) GetMethod() string {
	return "GET"
}

func (req *GetProductCountAPIRequest) GetQuery() interface{} {
	return req
}

func (req *GetProductCountAPIRequest) Verify() error {
	// Verify the api request params
	return nil
}

func (req *GetProductCountAPIRequest) GetEndpoint() string {
	return "products/count.json"
}

type GetProductCountAPIResponse struct {
	client.BaseAPIResponse

	Count int `json:"count"`
}

// GetProductsCount
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-product-count?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/batch-query-product-quantity?version=v20251201
// Deprecated
// see ProductService
func GetProductsCount(c *client.Client, apiReq *GetProductCountAPIRequest) (*GetProductCountAPIResponse, error) {

	// 1. API response resource
	apiResp := &GetProductCountAPIResponse{}

	// 2. Call the API
	err := c.Call(context.Background(), apiReq, apiResp)

	return apiResp, err
}
