package product

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/batch-query-product-quantity?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/batch-query-product-quantity?version=v20251201
//func TestProductsCount(t *testing.T) {
//
//	cli := manager.GetDefaultClient()
//
//	apiReq := &GetProductCountAPIReq{}
//	apiResp, err := GetProductsCount(cli, apiReq)
//	fmt.Printf("count:%v\n", apiResp)
//
//	assert.Nil(t, err)
//	assert.Greater(t, apiResp.Count, -1)
//
//}

func TestProductCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/count.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"count":1}`))

	apiReq := &GetProductCountAPIReq{}
	apiResp, err := GetProductsCount(cli, apiReq)
	if err != nil {
		t.Errorf("Product.Delete returned error: %v", err)
	} else {
		log.Printf("Product deleted, traceId: %s", apiResp.TraceId)
	}
	assert.Equal(t, 1, apiResp.Count)
}
