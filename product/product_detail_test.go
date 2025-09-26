package product

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestProductDetailDataNotExists(t *testing.T) {
//	//productId := "16070822412102455208483380"
//	productId := "16070828389930044390453380"
//	apiReq := &GetProductDetailAPIReq{
//		ProductId: productId,
//	}
//
//	c := manager.GetDefaultClient()
//
//	apiResp, err := GetProductDetail(c, apiReq)
//	log.Printf("apiResp: %v, err:%v", apiResp, err)
//	assert.NotNil(t, err, "err should be nil")
//}

// TestGetProductDetailError product detail
// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/query-single-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/query-single-product?version=v20251201
func TestGetProductDetailError(t *testing.T) {

	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/111.json", c.StoreHandle, c.PathPrefix, c.ApiVersion),
		httpmock.NewStringResponder(404, `{"errors":"DataNotExists"}`))

	//shopLineReq := &client.ShopLineRequest{}
	apiReq := &GetProductDetailAPIReq{
		ProductId: "111",
	}

	_, err := GetProductDetail(c, apiReq)

	a := assert.New(t)
	a.NotNil(err)
	a.Equal("DataNotExists", err.Error())
}

// product detail
// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/query-single-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/query-single-product?version=v20251201
func TestGetProductDetail(t *testing.T) {

	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/111.json", c.StoreHandle, c.PathPrefix, c.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("product/product.json")))

	//shopLineReq := &client.ShopLineRequest{}
	apiReq := &GetProductDetailAPIReq{
		ProductId: "111",
	}

	//responseData := &map[string]any{}
	//productId := "111"
	//endpoint := fmt.Sprintf("products/%s.json", productId)
	resp, err := GetProductDetail(c, apiReq)

	if err != nil {
		t.Fatal(err)
	}

	a := assert.New(t)
	a.Equal("111", resp.Product.Id)
}
