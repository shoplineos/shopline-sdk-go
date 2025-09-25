package product

import (
	"github.com/shoplineos/shopline-sdk-go/manager"
	"log"
	"testing"
)

// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-products?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/get-products?version=v20251201
func TestQueryProducts(t *testing.T) {

	apiReq := &QueryProductsAPIReq{
		ProductCategory: "Clothes",
		Status:          "active",
		OrderBy:         "created_at_desc",
		Limit:           3,
		// Fields:         "id,title,status,vendor",
	}

	c := manager.GetDefaultClient()
	apiResp, err := QueryProducts(c, apiReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("apiResp:%v\n", apiResp)
}

// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-products?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/get-products?version=v20251201
func TestQueryProductsTitle(t *testing.T) {

	apiReq := &QueryProductsAPIReq{
		//ProductCategory: "Clothes",
		//Status:          "active",
		//OrderBy:         "created_at_desc",
		//Limit:           3,
		Fields: "title",
	}

	c := manager.GetDefaultClient()
	apiResp, err := QueryProducts(c, apiReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("apiResp:%v\n", apiResp)

}

// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-products?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/get-products?version=v20251201
func TestQueryProductsPagination(t *testing.T) {

	requestParams := &QueryProductsAPIReq{
		//ProductCategory: "Clothes",
		//Status:          "active",
		//OrderBy:         "created_at_desc",
		Limit:    4,
		Fields:   "title",
		PageInfo: "eyJzaW5jZUlkIjoiMTYwNzA4MjgzODk5MzAwNDQzOTA0NTMzODAiLCJkaXJlY3Rpb24iOiJuZXh0IiwibGltaXQiOjR9",
	}

	c := manager.GetDefaultClient()
	products, err := QueryProducts(c, requestParams)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("products:%v\n", products)

}
