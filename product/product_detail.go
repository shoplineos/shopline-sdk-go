package product

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type GetProductDetailAPIReq struct {
	ProductId string
}

type GetProductDetailAPIResp struct {
	Product ProductRespData `json:"product"`

	client.CommonAPIRespData
}

// GetProductDetail
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/query-single-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/query-single-product?version=v20251201
func GetProductDetail(c *client.Client, apiReq *GetProductDetailAPIReq) (*GetProductDetailAPIResp, error) {
	// 1. API request
	shoplineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := fmt.Sprintf("products/%s.json", apiReq.ProductId)

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
