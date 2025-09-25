package product

import (
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/manager"
	"github.com/stretchr/testify/assert"
	"testing"
)

// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/update-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/update-product?version=v20251201
func TestUpdateProduct(t *testing.T) {

	requestBody := &UpdateProductAPIReq{
		Product: Product{
			Id:              "16071506529459141648923380",
			BodyHTML:        "This is a description",
			Title:           "A product title",
			Handle:          "a-product-title",
			Path:            "/my_product",
			ProductCategory: "Electronic",
			Tags:            []string{"tag1", "tag2"},
			SPU:             "S000001",
			Vendor:          "SHOPLine",
			Status:          "active",
			Subtitle:        "This is a subtitle",
			PublishedScope:  "web",

			Images: []Image{
				{
					Src: "https://img.myshopline.com/image/official/e46e6189dd5641a3b179444cacdcdd2a.png",
					Alt: "main picture",
					ID:  "7082838991712560845",
				},
			},

			Options: []Option{
				{
					Name:         "Color",
					ValuesImages: make(map[string]string),
				},
			},

			Variants: []Variant{
				{
					InventoryPolicy: "deny",
					Barcode:         "B000000001",
					Option4:         "",
					Image: Image{
						Src: "https://img.myshopline.com/image/official/e46e6189dd5641a3b179444cacdcdd2a.png",
						Alt: "This is a image alt",
						ID:  "7082838991712560846",
					},
					Option5:          "",
					CompareAtPrice:   "11.2",
					ID:               "18070828389938432998533380",
					Price:            "10.11",
					Option3:          "",
					RequiredShipping: true,
					WeightUnit:       "kg",
					InventoryTracker: true,
					Option2:          "L",
					SKU:              "T0000000001",
					Taxable:          true,
					Option1:          "Red",
					Weight:           "1.2",
				},
			},
		},
	}

	c := manager.GetDefaultClient()

	resp, err := UpdateProduct(c, requestBody)
	assert.Nil(t, err)
	fmt.Printf("Update Product Resp: %v\n", resp)
	// New Product Id: 16071506529459141648923380
	//16071506956220681507143380
	//16071531976175512823163380
}
