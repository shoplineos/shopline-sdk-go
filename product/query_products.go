package product

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type QueryProductsAPIResp struct {
	Products []ProductRespData `json:"products"`

	client.CommonAPIRespData
	Pagination *client.Pagination
}

type QueryProductsAPIReq struct {
	CollectionID    string // Collection ID
	CreatedAtMax    string // Latest creation time（ISO 8601）
	CreatedAtMin    string // Earliest creation time（ISO 8601）
	Fields          string // Fields（comma separated, eg "title,id,created_at"）
	Handle          string // Product Handle
	IDs             string // Ids（comma separated）
	Limit           int    // Limit（1-250, default 50）
	OrderBy         string // Sorting rules（created_at_asc/created_at_desc）
	PageInfo        string // Page Info（Get it from the response header 'link'）
	ProductCategory string // Product Category
	SinceID         string // Since ID（Start querying from this ID）
	Status          string // Product Status（active/draft/archived）
	Title           string // Product Title（fuzzy matching）
	UpdatedAtMax    string // Latest update time（ISO 8601）
	UpdatedAtMin    string // Earliest update time（ISO 8601）
	Vendor          string // Vendor
}

// QueryProducts
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-products?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/get-products?version=v20251201
func QueryProducts(c *client.Client, apiReq *QueryProductsAPIReq) (*QueryProductsAPIResp, error) {

	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := "products/products.json"

	// 3. API response data
	apiResp := &QueryProductsAPIResp{}

	// 4. Invoke API
	shopLineResp, err := c.Get(context.Background(), endpoint, shopLineReq, apiResp)
	if err != nil {
		return nil, err
	}

	//apiResp.TraceId = shopLineResp.TraceId
	apiResp.Pagination = shopLineResp.Pagination

	//log.Printf("product count:%v\n", len(apiResp.Products))
	//log.Printf("body:%v\n", apiResp)¬

	return apiResp, nil
}
