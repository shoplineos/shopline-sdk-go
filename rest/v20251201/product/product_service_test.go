package product

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/rest/test"
	"testing"
)

func TestProductServiceCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("product/product.json")))

	product := Product{
		Title:    "Hello shopline Freestyle 111",
		BodyHtml: "<strong>Hello shopline!<\\/strong>",
		Vendor:   "shopline",
	}

	apiReq := &CreateProductAPIReq{
		Product: product,
	}

	apiResp, err := GetProductService().Create(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Product.Create returned error: %v", err)
	}

	createProductTests(t, *apiResp)
}

func TestProductServiceUpdate(t *testing.T) {
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
