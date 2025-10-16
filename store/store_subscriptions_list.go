package store

import (
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListStoreSubscriptionsAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/subscribe/get-active-store-plans?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/store/subscribe/get-active-store-plans?version=v20251201
type ListStoreSubscriptionsAPIReq struct {
	client.BaseAPIRequest
	IncludeTrial bool `url:"include_trial,omitempty"`
}

func (req *ListStoreSubscriptionsAPIReq) Method() string {
	return "GET"
}

func (req *ListStoreSubscriptionsAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *ListStoreSubscriptionsAPIReq) Endpoint() string {
	return "store/subscription"
}

type ListStoreSubscriptionsAPIResp struct {
	client.BaseAPIResponse
	Subscriptions []Subscription `json:"data,omitempty"`

	// Deprecated
	Message string `json:"message,omitempty"`
}
