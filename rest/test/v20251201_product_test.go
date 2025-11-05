package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	product2 "github.com/shoplineos/shopline-sdk-go/rest/v20251201/product"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestProductCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "product/product.json")))

	product := product2.CreateAProductAPIReqProduct{
		Title:    "Hello shopline Freestyle 111",
		BodyHtml: "<strong>Hello shopline!<\\/strong>",
		Vendor:   "shopline",
	}

	apiReq := &product2.CreateAProductAPIReq{
		Product: product,
	}
	apiResp := &product2.CreateAProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)
	if err != nil {
		t.Errorf("Product.Create returned error: %v", err)
	}

	createProductTests(t, *apiResp)
}

func createProductTests(t *testing.T, product product2.CreateAProductAPIResp) {
	var expectedProductId = "111"
	if product.Product.Id != expectedProductId {
		t.Errorf("Product.Id returned %+v, expected %+v", product.Product.Id, expectedProductId)
	}
}

func TestProductUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/111.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "product/product.json")))

	product := product2.UpdateAProductAPIReqProduct{
		Title: "Test Product",
	}
	apiReq := &product2.UpdateAProductAPIReq{
		ProductId: "111",
		Product:   product,
	}
	apiResp := &product2.UpdateAProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("Product.Update returned error: %v", err)
	}

	productTests(t, apiResp)
}

func productTests(t *testing.T, product *product2.UpdateAProductAPIResp) {
	var expectedProductId = "111"
	if product.Product.Id != expectedProductId {
		t.Errorf("Product.Id returned %+v, expected %+v", product.Product.Id, expectedProductId)
	}
}

func TestProductCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/count.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"count":1}`))

	apiReq := &product2.GetProductCountAPIReq{}
	apiResp := &product2.GetProductCountAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("Product.Delete returned error: %v", err)
	} else {
		log.Printf("Product deleted, traceId: %s", apiResp.TraceId)
	}
	assert.Equal(t, 1, apiResp.Count)
}

func TestDeleteAProduct(t *testing.T) {

	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	productId := "1"
	apiReq := &product2.DeleteAProductAPIReq{
		ProductId: productId,
	}

	apiResp := &product2.DeleteAProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		log.Printf("delete product error, apiResp: %v, err:%v", apiResp, err)
	} else {
		log.Printf("delete product success")
	}

	assert.Nil(t, err, "err should be nil")
}

// 500 Internal Server Error
func TestDeleteProductError(t *testing.T) {

	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(500, `{"errors":"Internal Server Error"}`))

	apiReq := &product2.DeleteAProductAPIReq{
		ProductId: "1",
	}

	apiResp := &product2.DeleteAProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NotNil(t, err)
	assert.Equal(t, "Internal Server Error", err.Error())
}

// Unknown Error
func TestDeleteProductUnknowError(t *testing.T) {

	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(500, ""))

	apiReq := &product2.DeleteAProductAPIReq{
		ProductId: "1",
	}

	apiResp := &product2.DeleteAProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NotNil(t, err)
	assert.Equal(t, "Unknown Error", err.Error())
}

// ok
func TestDeleteProduct3(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, "{}"))

	apiReq := &product2.DeleteAProductAPIReq{
		ProductId: "1",
	}
	apiResp := &product2.DeleteAProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("Product.Delete returned error: %v", err)
	} else {
		log.Printf("Product deleted, traceId: %s", apiResp.TraceId)
	}
}

// TestGetProductDetailError product detail
// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/query-single-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/query-single-product?version=v20251201
func TestGetProductDetailError(t *testing.T) {

	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/111.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(404, `{"errors":"DataNotExists"}`))

	apiReq := &product2.GetAProductAPIReq{
		ProductId: "111",
	}
	apiResp := &product2.DeleteAProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

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

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/111.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "product/product.json")))

	apiReq := &product2.GetAProductAPIReq{
		ProductId: "111",
	}

	//responseData := &map[string]any{}
	//productId := "111"
	//endpoint := fmt.Sprintf("products/%s.json", productId)

	apiResp := &product2.GetAProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}

	a := assert.New(t)
	a.Equal("111", apiResp.Product.Id)
}
