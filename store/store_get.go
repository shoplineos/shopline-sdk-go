package store

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetStoreAPIReq
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/query-store-information?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/store/query-store-information?version=v20251201
type GetStoreAPIReq struct {
	client.BaseAPIRequest
}

func (req *GetStoreAPIReq) GetMethod() string {
	return "GET"
}

func (req *GetStoreAPIReq) GetQuery() interface{} {
	return req
}

func (req *GetStoreAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *GetStoreAPIReq) GetEndpoint() string {
	return "merchants/shop.json"
}

type GetStoreAPIResp struct {
	client.BaseAPIResponse
	Store Store `json:"data,omitempty"`
}

// GetStoreInfo
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/query-store-information?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/store/query-store-information?version=v20251201
// Deprecated
// see StoreService
func GetStoreInfo(c *client.Client, apiReq *GetStoreAPIReq) (*GetStoreAPIResp, error) {
	// 1. API response resource
	apiResp := &GetStoreAPIResp{}

	// 2. Call the API
	err := c.Call(context.Background(), apiReq, apiResp)
	return apiResp, err
}
