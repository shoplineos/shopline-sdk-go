package product

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
	"log"
)

type UpdateProductAPIReq struct {
	Product Product `json:"product"`
}

func (req *UpdateProductAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *UpdateProductAPIReq) Endpoint() string {
	endpoint := fmt.Sprintf("products/%s.json", req.Product.Id)
	return endpoint
}

type UpdateProductAPIResp struct {
	Product ProductRespData `json:"product"`

	client.CommonAPIRespData
}

// UpdateProduct
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/update-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/update-a-product?version=v20251201
func UpdateProduct(c *client.Client, apiReq *UpdateProductAPIReq) (*UpdateProductAPIResp, error) {

	// 1. API request
	request := &client.ShopLineRequest{
		Body: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &UpdateProductAPIResp{}

	// 4. Call API
	shopLineResp, err := c.Put(context.Background(), endpoint, request, apiResp)

	if err != nil {
		log.Printf("Update product failed，shopLineResp: %v, err: %v\n", shopLineResp, err)
		return nil, err
	}

	//apiResp.TraceId = shopLineResp.TraceId

	return apiResp, nil
}
