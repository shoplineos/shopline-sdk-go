package product

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CreateProductAPIReq Create Product Request Params
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/create-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201
type CreateProductAPIReq struct {
	Product Product `json:"product"`
}

func (req *CreateProductAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *CreateProductAPIReq) Endpoint() string {
	endpoint := "products/products.json"
	return endpoint
}

type CreateProductAPIResp struct {
	client.BaseAPIResponse
	Product ProductRespData `json:"product"`
}

type Product struct {
	Id              string    `json:"id"`                         // Product id
	Title           string    `json:"title,omitempty"`            // Product Title
	BodyHTML        string    `json:"body_html,omitempty"`        // Product Get HTML
	Handle          string    `json:"shopHandle,omitempty"`       // shopHandle
	Subtitle        string    `json:"subtitle,omitempty"`         // Product Subtitle
	Vendor          string    `json:"vendor,omitempty"`           // Vendor
	ProductCategory string    `json:"product_category,omitempty"` // Product Category
	Tags            []string  `json:"tags,omitempty"`             // Product Tags
	SPU             string    `json:"spu,omitempty"`              // Product SPU
	Status          string    `json:"status,omitempty"`           // Status（active/draft）
	PublishedScope  string    `json:"published_scope,omitempty"`  // Published Scope
	Options         []Option  `json:"options,omitempty"`          // Options（eg:Color、Size）
	Images          []Image   `json:"images,omitempty"`           // Product Images
	Variants        []Variant `json:"variants,omitempty"`         // Product Variants
	Path            string    `json:"path,omitempty"`             // Path
}

// ProductRespData Product Response Body
type ProductRespData struct {
	Id              string    `json:"id"`                         // Product id
	Title           string    `json:"title,omitempty"`            // Product Title
	BodyHTML        string    `json:"body_html,omitempty"`        // Product Get HTML
	Handle          string    `json:"shopHandle,omitempty"`       // shopHandle
	Subtitle        string    `json:"subtitle,omitempty"`         // Product Subtitle
	Vendor          string    `json:"vendor,omitempty"`           // Vendor
	ProductCategory string    `json:"product_category,omitempty"` // Product Category
	Tags            string    `json:"tags,omitempty"`             // Tags
	SPU             string    `json:"spu,omitempty"`              // Product SPU
	Status          string    `json:"status,omitempty"`           // Status（active/draft）
	PublishedScope  string    `json:"published_scope,omitempty"`  // Published Scope
	Options         []Option  `json:"options,omitempty"`          // Options（eg:Color、Size）
	Images          []Image   `json:"images,omitempty"`           // Product Images
	Variants        []Variant `json:"variants,omitempty"`         // Product Variants
	Path            string    `json:"path,omitempty"`             // Path
}

type Image struct {
	ID  string `json:"id,omitempty"`
	Src string `json:"src,omitempty"` // Image url
	Alt string `json:"alt,omitempty"` // Image description
}

// Option Product Option（eg:Color、Size）
type Option struct {
	Name         string            `json:"name,omitempty"`   // Option name（eg "Color"）
	Values       []string          `json:"values,omitempty"` // Option value（eg ["red", "blue"]）
	ValuesImages map[string]string `json:"values_images,omitempty"`
}

// Variant Product Variant
type Variant struct {
	ID               string `json:"id,omitempty"`
	Image            Image  `json:"image,omitempty"`
	SKU              string `json:"sku,omitempty"`              // SKU
	Price            string `json:"price,omitempty"`            // Price
	CompareAtPrice   string `json:"compare_at_price,omitempty"` // Compare At Price
	Barcode          string `json:"barcode,omitempty"`          // Barcode
	Weight           string `json:"weight,omitempty"`           // Weight
	WeightUnit       string `json:"weight_unit,omitempty"`      // Weight Unit（kg/g/lb/oz）
	Taxable          bool   `json:"taxable,omitempty"`          // Taxable
	Option1          string `json:"option1,omitempty"`
	Option2          string `json:"option2,omitempty"`
	Option3          string `json:"option3,omitempty"`
	Option4          string `json:"option4,omitempty"`
	Option5          string `json:"option5,omitempty"`
	InventoryPolicy  string `json:"inventory_policy,omitempty"`  // Inventory Policy
	RequiredShipping bool   `json:"required_shipping,omitempty"` // Required Shipping
	InventoryTracker bool   `json:"inventory_tracker,omitempty"`
}

// CreateProduct
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/create-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201
// Deprecated
// see ProductService
func CreateProduct(c *client.Client, apiReq *CreateProductAPIReq) (*CreateProductAPIResp, error) {

	// 1. API request
	request := &client.ShopLineRequest{ // client request
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &CreateProductAPIResp{}

	// 4. Call API
	_, err := c.Post(context.Background(), endpoint, request, apiResp)
	//if err != nil {
	//	log.Printf("CreateProduct request failed: %v\n", err)
	//	return nil, err
	//}
	return apiResp, err
}
