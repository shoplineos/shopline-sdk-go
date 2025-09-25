package product

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type GetProductDetailAPIReq struct {
	ProductId string
}

func (req *GetProductDetailAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *GetProductDetailAPIReq) Endpoint() string {
	endpoint := fmt.Sprintf("products/%s.json", req.ProductId)
	return endpoint
}

type GetProductDetailAPIResp struct {
	Product ProductRespData `json:"product"`

	client.CommonAPIRespData
}

// GetProductDetail
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/get-a-product?version=v20251201
func GetProductDetail(c *client.Client, apiReq *GetProductDetailAPIReq) (*GetProductDetailAPIResp, error) {
	// 1. API request
	shoplineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &GetProductDetailAPIResp{}

	// 4. Invoke API
	_, err := c.Get(context.Background(), endpoint, shoplineReq, apiResp)

	if err != nil {
		return apiResp, err
	}

	//apiResp.TraceId = shoplineResp.TraceId

	//fmt.Println(shoplineResp)
	return apiResp, nil
}
