package product

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/manager"
)

type GetProductCountAPIReq struct {
	CollectionId string `json:"collection_id"`
	CreatedAtMax string // Latest creation time（ISO 8601）
	CreatedAtMin string // Earliest creation time（ISO 8601）
	UpdatedAtMax string // Latest update time（ISO 8601）
	UpdatedAtMin string // Earliest update time（ISO 8601）
}

type GetProductCountAPIResp struct {
	Count int `json:"count"`

	client.CommonAPIRespData
}

// GetProductsCount
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-product-count?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/batch-query-product-quantity?version=v20251201
func GetProductsCount(apiReq *GetProductCountAPIReq) (*GetProductCountAPIResp, error) {
	return GetProductsCountV2(manager.GetDefaultClient().GetAppKey(), manager.GetDefaultClient().StoreHandle, apiReq)
}

// GetProductsCountV2
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-product-count?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/batch-query-product-quantity?version=v20251201
func GetProductsCountV2(appKey, storeHandle string, apiReq *GetProductCountAPIReq) (*GetProductCountAPIResp, error) {

	// 1. API request
	shoplineReq := &client.ShopLineRequest{
		Query: apiReq,
	}

	// 2. API endpoint
	endpoint := "products/count.json"

	// 3. API response
	apiResp := &GetProductCountAPIResp{}

	// 4. Invoke API
	_, err := manager.GetClient(appKey, storeHandle).Get(context.Background(), endpoint, shoplineReq, apiResp)

	//apiResp.TraceId = shoplineResp.TraceId

	return apiResp, err
}
