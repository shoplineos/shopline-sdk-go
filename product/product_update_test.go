package product

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/test"
	"testing"
)

// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/update-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/update-product?version=v20251201
//func TestUpdateProduct(t *testing.T) {
//
//	requestBody := &UpdateProductAPIReq{
//		Product: Product{
//			Id:              "16071506529459141648923380",
//			BodyHTML:        "This is a description",
//			Title:           "A product title",
//			Handle:          "a-product-title",
//			Path:            "/my_product",
//			ProductCategory: "Electronic",
//			Tags:            []string{"tag1", "tag2"},
//			SPU:             "S000001",
//			Vendor:          "SHOPLine",
//			Status:          "active",
//			Subtitle:        "This is a subtitle",
//			PublishedScope:  "web",
//
//			Images: []Image{
//				{
//					Src: "https://img.myshopline.com/image/official/e46e6189dd5641a3b179444cacdcdd2a.png",
//					Alt: "main picture",
//					ID:  "7082838991712560845",
//				},
//			},
//
//			Options: []Option{
//				{
//					Name:         "Color",
//					ValuesImages: make(map[string]string),
//				},
//			},
//
//			Variants: []Variant{
//				{
//					InventoryPolicy: "deny",
//					Barcode:         "B000000001",
//					Option4:         "",
//					Image: Image{
//						Src: "https://img.myshopline.com/image/official/e46e6189dd5641a3b179444cacdcdd2a.png",
//						Alt: "This is a image alt",
//						ID:  "7082838991712560846",
//					},
//					Option5:          "",
//					CompareAtPrice:   "11.2",
//					ID:               "18070828389938432998533380",
//					Price:            "10.11",
//					Option3:          "",
//					RequiredShipping: true,
//					WeightUnit:       "kg",
//					InventoryTracker: true,
//					Option2:          "L",
//					SKU:              "T0000000001",
//					Taxable:          true,
//					Option1:          "Red",
//					Weight:           "1.2",
//				},
//			},
//		},
//	}
//
//	cli := manager.GetDefaultClient()
//
//	resp, err := UpdateProduct(cli, requestBody)
//	assert.Nil(t, err)
//	fmt.Printf("Update Product Resp: %v\n", resp)
//	// New Product Id: 16071506529459141648923380
//	//16071506956220681507143380
//	//16071531976175512823163380
//}

func TestProductUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/111.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("product/product.json")))

	product := Product{
		Id:    "111",
		Title: "Test Product",
	}

	apiReq := &UpdateProductAPIReq{
		Product: product,
	}

	apiResp, err := GetProductService().Update(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Product.Update returned error: %v", err)
	}

	productTests(t, apiResp)
}

func productTests(t *testing.T, product *UpdateProductAPIResp) {
	var expectedProductId = "111"
	if product.Product.Id != expectedProductId {
		t.Errorf("Product.Id returned %+v, expected %+v", product.Product.Id, expectedProductId)
	}
}
