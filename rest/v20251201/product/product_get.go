package product

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetProductDetailAPIReq
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/get-a-product?version=v20251201
// Deprecated
type GetProductDetailAPIReq struct {
	client.BaseAPIRequest
	ProductId string
}

func (req *GetProductDetailAPIReq) GetMethod() string {
	return "GET"
}

func (req *GetProductDetailAPIReq) GetQuery() interface{} {
	return req
}

func (req *GetProductDetailAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *GetProductDetailAPIReq) GetEndpoint() string {
	endpoint := fmt.Sprintf("products/%s.json", req.ProductId)
	return endpoint
}

type GetProductDetailAPIResp struct {
	client.BaseAPIResponse
	Product ProductRespData `json:"product"`
}

// GetProductDetail
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/get-a-product?version=v20251201
// Deprecated
// see ProductService
func GetProductDetail(c *client.Client, req *GetProductDetailAPIReq) (*GetProductDetailAPIResp, error) {
	// 1. API response data
	apiResp := &GetProductDetailAPIResp{}

	// 2. Call the API
	err := c.Call(context.Background(), req, apiResp)
	return apiResp, err
}
