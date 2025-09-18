package product

import (
	"log"
	"testing"
	"time"
)

// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/create-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201
func TestCreateProduct(t *testing.T) {

	apiReq := &CreateProductAPIReq{
		Product: Product{
			Title:          "Test product - " + time.Now().Format("20060102150405"),
			BodyHTML:       "<p>This is a test product created via the API</p>",
			Subtitle:       "Limited time offer",
			Vendor:         "Test provider",
			Status:         "active",
			PublishedScope: "web",
			Tags:           []string{"Test", "New", "API Create"},

			// Options（Color、Size）
			Options: []Option{
				{Name: "color", Values: []string{"red", "blue"}},
				{Name: "size", Values: []string{"S", "M", "L"}},
			},

			Images: []Image{
				{Src: "https://example.com/product-main.jpg", Alt: "Main picture"},
				{Src: "https://example.com/product-detail.jpg", Alt: "Detail picture"},
			},

			// Variants（red S、red M、blue S）
			Variants: []Variant{
				{
					SKU:            "RED-S-001",
					Price:          "99.99",
					CompareAtPrice: "129.99",
					Option1:        "red",
					Option2:        "S",
					Weight:         "0.5",
					WeightUnit:     "kg",
					Taxable:        true,
				},
				{
					SKU:            "RED-M-002",
					Price:          "109.99",
					CompareAtPrice: "139.99",
					Option1:        "red",
					Option2:        "M",
					Weight:         "0.6",
					WeightUnit:     "kg",
					Taxable:        true,
				},
				{
					SKU:            "BLUE-S-003",
					Price:          "99.99",
					CompareAtPrice: "129.99",
					Option1:        "blue",
					Option2:        "S",
					Weight:         "0.5",
					WeightUnit:     "kg",
					Taxable:        true,
				},
			},
		},
	}

	apiResp, err := CreateProduct(apiReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("New product ID: %v\n", apiResp.Product.Id)

	// option
	log.Printf("Create product resp: %+v", apiResp)

	//16071506956220681507143380
	//16071531976175512823163380
}

//func logInfo(response *client.ShopLineResponse, responseData *CreateProductAPIResp) {
//	if response.StatusCode == http.StatusOK {
//		log.Println("Create product success！")
//		productID := responseData.Product.Id
//		log.Printf("New product ID: %v\n", productID)
//	} else {
//		log.Printf("Create product failed, erros: %s\n", response.Errors)
//	}
//}
