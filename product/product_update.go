package product

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// UpdateProductAPIReq
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/update-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/update-a-product?version=v20251201
type UpdateProductAPIReq struct {
	client.BaseAPIRequest
	Product Product `json:"product"`
}

func (req *UpdateProductAPIReq) GetMethod() string {
	return "PUT"
}

func (req *UpdateProductAPIReq) GetData() interface{} {
	return req
}

func (req *UpdateProductAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *UpdateProductAPIReq) GetEndpoint() string {
	endpoint := fmt.Sprintf("products/%s.json", req.Product.Id)
	return endpoint
}

type UpdateProductAPIResp struct {
	client.BaseAPIResponse
	Product ProductRespData `json:"product"`
}

// UpdateProduct
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/update-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/update-a-product?version=v20251201
// Deprecated
// see ProductService
func UpdateProduct(c *client.Client, req *UpdateProductAPIReq) (*UpdateProductAPIResp, error) {
	// 1. API response data
	apiResp := &UpdateProductAPIResp{}

	// 2. Call the API
	err := c.Call(context.Background(), req, apiResp)
	return apiResp, err
}
