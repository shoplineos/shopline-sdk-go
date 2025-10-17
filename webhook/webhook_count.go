package webhook

import (
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CountWebhooksAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/webhook/get-the-subscribed-webhook-count?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/webhook/get-the-subscribed-webhook-count?version=v20251201
type CountWebhooksAPIReq struct {
	client.BaseAPIRequest
}

func (r *CountWebhooksAPIReq) GetMethod() string {
	return "GET"
}

func (r *CountWebhooksAPIReq) GetQuery() interface{} {
	return r
}

func (r *CountWebhooksAPIReq) Verify() error {
	return nil
}

func (r *CountWebhooksAPIReq) GetEndpoint() string {
	return "webhooks/count.json"
}

type CountWebhooksAPIResp struct {
	client.BaseAPIResponse
	Count int `json:"count,omitempty"`
}
