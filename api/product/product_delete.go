package product

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// DeleteProductAPIReq
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/delete-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/delete-a-product?version=v20251201
type DeleteProductAPIReq struct {
	client.BaseAPIRequest
	ProductId string
}

func (req *DeleteProductAPIReq) GetMethod() string {
	return "DELETE"
}

func (req *DeleteProductAPIReq) GetData() interface{} {
	return req
}

func (req *DeleteProductAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *DeleteProductAPIReq) GetEndpoint() string {
	endpoint := fmt.Sprintf("products/%s.json", req.ProductId)
	return endpoint
}

type DeleteProductAPIResp struct {
	client.BaseAPIResponse
}

// DeleteProduct
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/delete-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/delete-a-product?version=v20251201
// Deprecated
// see ProductService
func DeleteProduct(c *client.Client, req *DeleteProductAPIReq) (*DeleteProductAPIResp, error) {

	// 1. API response data
	apiResp := &DeleteProductAPIResp{}

	// 2. Call the API
	err := c.Call(context.Background(), req, apiResp)
	return apiResp, err
}
