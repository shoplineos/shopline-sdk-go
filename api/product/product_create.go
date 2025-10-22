package product

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CreateProductAPIReq Create Product Request Params
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/create-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201
type CreateProductAPIReq struct {
	client.BaseAPIRequest
	Product Product `json:"product"`
}

func (req *CreateProductAPIReq) GetMethod() string {
	return "POST"
}

func (req *CreateProductAPIReq) GetData() interface{} {
	return req
}

func (req *CreateProductAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *CreateProductAPIReq) GetEndpoint() string {
	endpoint := "products/products.json"
	return endpoint
}

type CreateProductAPIResp struct {
	client.BaseAPIResponse
	Product ProductRespData `json:"product"`
}

// CreateProduct
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/create-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201
// Deprecated
// see ProductService
func CreateProduct(c *client.Client, req *CreateProductAPIReq) (*CreateProductAPIResp, error) {
	// 1. API response data
	apiResp := &CreateProductAPIResp{}

	// 2. Call the API
	err := c.Call(context.Background(), req, apiResp)
	return apiResp, err
}
