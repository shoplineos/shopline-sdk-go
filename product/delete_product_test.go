package product

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

//
//func TestDeleteProduct(t *testing.T) {
//
//	productId := "16070828389930044390453380"
//	apiReq := &DeleteProductAPIReq{
//		ProductId: productId,
//	}
//
//	c := manager.GetDefaultClient()
//
//	apiResp, err := DeleteProduct(c, apiReq)
//	if err != nil {
//		log.Printf("delete product error, apiResp: %v, err:%v", apiResp, err)
//	} else {
//		log.Printf("delete product success")
//	}
//
//	assert.Nil(t, err, "err should be nil")
//}

// 500 Internal Server Error
func TestDeleteProductError(t *testing.T) {

	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/1.json", c.StoreHandle, c.PathPrefix, c.ApiVersion),
		httpmock.NewStringResponder(500, `{"errors":"Internal Server Error"}`))

	apiReq := &DeleteProductAPIReq{
		ProductId: "1",
	}
	_, err := DeleteProduct(c, apiReq)
	assert.NotNil(t, err)
	assert.Equal(t, "Internal Server Error", err.Error())
}

// Unknown Error
func TestDeleteProductUnknowError(t *testing.T) {

	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/1.json", c.StoreHandle, c.PathPrefix, c.ApiVersion),
		httpmock.NewStringResponder(500, ""))

	apiReq := &DeleteProductAPIReq{
		ProductId: "1",
	}
	_, err := DeleteProduct(c, apiReq)
	assert.NotNil(t, err)
	assert.Equal(t, "Unknown Error", err.Error())
}

// ok
func TestDeleteProduct3(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/1.json", c.StoreHandle, c.PathPrefix, c.ApiVersion),
		httpmock.NewStringResponder(200, "{}"))

	apiReq := &DeleteProductAPIReq{
		ProductId: "1",
	}
	apiResp, err := DeleteProduct(c, apiReq)
	if err != nil {
		t.Errorf("Product.Delete returned error: %v", err)
	} else {
		log.Printf("Product deleted, traceId: %s", apiResp.TraceId)
	}
}
