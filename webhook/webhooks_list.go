package webhook

import (
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListWebhooksAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/webhook/get-a-list-of-subscribed-webhooks?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/webhook/get-a-list-of-subscribed-webhooks?version=v20251201
type ListWebhooksAPIReq struct {
}

func (c ListWebhooksAPIReq) Verify() error {
	return nil
}

func (c ListWebhooksAPIReq) Endpoint() string {
	return "webhooks.json"
}

type ListWebhooksAPIResp struct {
	client.BaseAPIResponse
	Webhooks []Webhook `json:"webhooks,omitempty"`
}
