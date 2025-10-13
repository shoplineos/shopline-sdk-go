package product

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type ListProductsAPIResp struct {
	client.BaseAPIResponse
	Products []ProductRespData `json:"products"`
}

type ListProductsAPIReq struct {
	Status       string `url:"status,omitempty"`
	CollectionId string `url:"collection_id,omitempty"`
	CreatedAtMin string `url:"created_at_min,omitempty"` // Minimum order creation time（ISO 8601）
	CreatedAtMax string `url:"created_at_max,omitempty"` // Max order creation time（ISO 8601）
	UpdatedAtMin string `url:"updated_at_min,omitempty"` // Minimum order update time（ISO 8601）
	UpdatedAtMax string `url:"updated_at_max,omitempty"` // Max order update time（ISO 8601）

	Ids      string `url:"ids,omitempty"`       //  ids, Separate multiple with commas
	Limit    int32  `url:"limit,omitempty"`     // Limit（1-250, default 50）
	Fields   string `url:"fields,omitempty"`    // Fields, Separate multiple with commas
	PageInfo string `url:"page_info,omitempty"` // Page Info（Get it from the response header 'link'）

	Handle          string `url:"handle,omitempty"`           // Product Handle
	OrderBy         string `url:"order_by,omitempty"`         // Sorting rules（created_at_asc/created_at_desc）
	ProductCategory string `url:"product_category,omitempty"` // Product Category
	SinceId         string `url:"since_id,omitempty"`         // Since Id（Start querying from this Id）
	Title           string `url:"title,omitempty"`            // Product Title（fuzzy matching）

	Vendor string `url:"vendor,omitempty"` // Vendor
}

func (req *ListProductsAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *ListProductsAPIReq) Endpoint() string {
	return "products/products.json" // endpoint
}

// ListProducts
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-products?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/get-products?version=v20251201
// Deprecated
// see ProductService
func ListProducts(c *client.Client, apiReq *ListProductsAPIReq) (*ListProductsAPIResp, error) {

	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &ListProductsAPIResp{}

	// 4. Call API
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

// ListProductsAll
// Deprecated
// see ProductService
func ListProductsAll(c *client.Client, apiReq *ListProductsAPIReq) ([]ProductRespData, error) {
	collector := []ProductRespData{}
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	for {
		// 2. API endpoint
		endpoint := apiReq.Endpoint()

		// 3. API response data
		apiResp := &ListProductsAPIResp{}

		// 4. Call API
		shoplineResp, err := c.Get(context.Background(), endpoint, shopLineReq, apiResp)

		if err != nil {
			return collector, err
		}

		collector = append(collector, apiResp.Products...)

		if !shoplineResp.HasNext() {
			break
		}

		shopLineReq.Query = shoplineResp.Pagination.Next
	}

	return collector, nil
}
