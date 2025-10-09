package product

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type DeleteProductAPIReq struct {
	ProductId string
}

func (req *DeleteProductAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *DeleteProductAPIReq) Endpoint() string {
	endpoint := fmt.Sprintf("products/%s.json", req.ProductId)
	return endpoint
}

type DeleteProductAPIResp struct {
	client.BaseAPIResponse
}

// DeleteProduct
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/delete-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/delete-a-product?version=v20251201
func DeleteProduct(c *client.Client, apiReq *DeleteProductAPIReq) (*DeleteProductAPIResp, error) {

	// 1. API request
	shoplineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &DeleteProductAPIResp{}

	// 4. Call API
	_, err := c.Delete(context.Background(), endpoint, shoplineReq, apiResp)

	//apiResp.TraceId = shoplineResp.TraceId

	return apiResp, err
}
