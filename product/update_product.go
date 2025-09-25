package product

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
	"log"
)

type ProductUpdateAPIReq struct {
	Product Product `json:"product"`
}

type UpdateProductAPIResp struct {
	Product ProductRespData `json:"product"`

	client.CommonAPIRespData
}

//type Product struct {
//	BodyHTML        string    `json:"body_html,omitempty"`
//	Images          []Image   `json:"images,omitempty"`
//	Variants        []Variant `json:"variants,omitempty"`
//	DefaultStoreHandle          string    `json:"shopHandle,omitempty"`
//	Title           string    `json:"title,omitempty"`
//	Path            string    `json:"path,omitempty"`
//	ProductCategory string    `json:"product_category,omitempty"`
//	Tags            []string  `json:"tags,omitempty"`
//	Options         []Option  `json:"options,omitempty"`
//	PublishedScope  string    `json:"published_scope,omitempty"`
//	SPU             string    `json:"spu,omitempty"`
//	Status          string    `json:"status,omitempty"`
//	Subtitle        string    `json:"subtitle,omitempty"`
//	Vendor          string    `json:"vendor,omitempty"`
//}
//
//type Image struct {
//	Src string `json:"src,omitempty"`
//	Alt string `json:"alt,omitempty"`
//	ID  string `json:"id,omitempty"`
//}
//
//type Variant struct {
//	InventoryPolicy  string `json:"inventory_policy,omitempty"`
//	Barcode          string `json:"barcode,omitempty"`
//	Option4          string `json:"option4,omitempty"`
//	Image            Image  `json:"image,omitempty"`
//	Option5          string `json:"option5,omitempty"`
//	CompareAtPrice   string `json:"compare_at_price,omitempty"`
//	ID               string `json:"id,omitempty"`
//	Price            string `json:"price,omitempty"`
//	Option3          string `json:"option3,omitempty"`
//	RequiredShipping bool   `json:"required_shipping,omitempty"`
//	WeightUnit       string `json:"weight_unit,omitempty"`
//	InventoryTracker bool   `json:"inventory_tracker,omitempty"`
//	Option2          string `json:"option2,omitempty"`
//	SKU              string `json:"sku,omitempty"`
//	Taxable          bool   `json:"taxable,omitempty"`
//	Option1          string `json:"option1,omitempty"`
//	Weight           string `json:"weight,omitempty"`
//}
//
//type Option struct {
//	Name         string            `json:"name,omitempty"`
//	ValuesImages map[string]string `json:"values_images,omitempty"`
//}

// UpdateProduct
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/update-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/update-a-product?version=v20251201
func UpdateProduct(c *client.Client, apiReq *ProductUpdateAPIReq) (*UpdateProductAPIResp, error) {

	// 1. API request
	request := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	productID := apiReq.Product.Id
	endpoint := fmt.Sprintf("products/%s.json", productID)

	// 3. API response data
	apiResp := &UpdateProductAPIResp{}

	// 4. Invoke API
	shopLineResp, err := c.Put(context.Background(), endpoint, request, apiResp)

	if err != nil {
		log.Printf("Update product failed，shopLineResp: %v, err: %v\n", shopLineResp, err)
		return nil, err
	}

	//apiResp.TraceId = shopLineResp.TraceId

	return apiResp, nil
}
