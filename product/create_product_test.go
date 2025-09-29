package product

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/test"
	"testing"
)

// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/create-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201
//func TestCreateProduct(t *testing.T) {
//
//	apiReq := &CreateProductAPIReq{
//		Product: Product{
//			Title:          "Test product - " + time.Now().Format("20060102150405"),
//			BodyHTML:       "<p>This is a test product created via the API</p>",
//			Subtitle:       "Limited time offer",
//			Vendor:         "Test provider",
//			Status:         "active",
//			PublishedScope: "web",
//			Tags:           []string{"Test", "New", "API Create"},
//
//			// Options（Color、Size）
//			Options: []Option{
//				{Name: "color", Values: []string{"red", "blue"}},
//				{Name: "size", Values: []string{"S", "M", "L"}},
//			},
//
//			Images: []Image{
//				{Src: "https://example.com/product-main.jpg", Alt: "Main picture"},
//				{Src: "https://example.com/product-detail.jpg", Alt: "Detail picture"},
//			},
//
//			// Variants（red S、red M、blue S）
//			Variants: []Variant{
//				{
//					SKU:            "RED-S-001",
//					Price:          "99.99",
//					CompareAtPrice: "129.99",
//					Option1:        "red",
//					Option2:        "S",
//					Weight:         "0.5",
//					WeightUnit:     "kg",
//					Taxable:        true,
//				},
//				{
//					SKU:            "RED-M-002",
//					Price:          "109.99",
//					CompareAtPrice: "139.99",
//					Option1:        "red",
//					Option2:        "M",
//					Weight:         "0.6",
//					WeightUnit:     "kg",
//					Taxable:        true,
//				},
//				{
//					SKU:            "BLUE-S-003",
//					Price:          "99.99",
//					CompareAtPrice: "129.99",
//					Option1:        "blue",
//					Option2:        "S",
//					Weight:         "0.5",
//					WeightUnit:     "kg",
//					Taxable:        true,
//				},
//			},
//		},
//	}
//
//	cli := manager.GetDefaultClient()
//	apiResp, err := CreateProduct(cli, apiReq)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Printf("New product ID: %v\n", apiResp.Product.Id)
//
//	// option
//	log.Printf("Create product resp: %+v", apiResp)
//
//	//16071506956220681507143380
//	//16071531976175512823163380
//}

func TestProductCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("product/product.json")))

	product := Product{
		Title:    "Hello shopline Freestyle 111",
		BodyHTML: "<strong>Hello shopline!<\\/strong>",
		Vendor:   "shopline",
	}

	apiReq := &CreateProductAPIReq{
		Product: product,
	}

	apiResp, err := CreateProduct(cli, apiReq)
	if err != nil {
		t.Errorf("Product.Create returned error: %v", err)
	}

	createProductTests(t, *apiResp)
}

func createProductTests(t *testing.T, product CreateProductAPIResp) {
	var expectedProductId = "111"
	if product.Product.Id != expectedProductId {
		t.Errorf("Product.Id returned %+v, expected %+v", product.Product.Id, expectedProductId)
	}
}
