package product

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/batch-query-product-quantity?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/batch-query-product-quantity?version=v20251201
func TestProductsCount(t *testing.T) {

	apiReq := &GetProductCountAPIReq{}
	apiResp, err := GetProductsCount(apiReq)
	fmt.Printf("count:%v\n", apiResp)

	assert.Nil(t, err)
	assert.Greater(t, apiResp.Count, -1)

}
