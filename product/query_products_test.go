package product

import (
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

	apiResp, err := QueryProducts(apiReq)
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

	apiResp, err := QueryProducts(apiReq)
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

	products, err := QueryProducts(requestParams)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("products:%v\n", products)

}
