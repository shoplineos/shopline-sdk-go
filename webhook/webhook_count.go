package webhook

import (
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CountWebhooksAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/webhook/get-the-subscribed-webhook-count?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/webhook/get-the-subscribed-webhook-count?version=v20251201
type CountWebhooksAPIReq struct {
}

func (c CountWebhooksAPIReq) Verify() error {
	return nil
}

func (c CountWebhooksAPIReq) Endpoint() string {
	return "webhooks/count.json"
}

type CountWebhooksAPIResp struct {
	client.BaseAPIResponse
	Count int `json:"count,omitempty"`
}
