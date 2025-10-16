package store

import "github.com/shoplineos/shopline-sdk-go/client"

// ListStoreCurrenciesAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/query-store-settlement-currency?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/store/query-store-settlement-currency?version=v20251201
type ListStoreCurrenciesAPIReq struct {
	client.BaseAPIRequest
}

func (req *ListStoreCurrenciesAPIReq) Method() string {
	return "GET"
}

func (req *ListStoreCurrenciesAPIReq) GetQuery() interface{} {
	return req
}

func (req *ListStoreCurrenciesAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *ListStoreCurrenciesAPIReq) Endpoint() string {
	return "currency/currencies.json"
}

type ListStoreCurrenciesAPIResp struct {
	client.BaseAPIResponse
	Currencies []Currency `json:"currencies,omitempty"`
}

type Currency struct {
	RateUpdateAt string `json:"rate_update_at,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Enabled      bool   `json:"enabled,omitempty"`
}
