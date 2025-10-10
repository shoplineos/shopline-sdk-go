package webhook

import (
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetWebhookAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/webhook/get-a-subscribed-webhook?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/webhook/get-a-subscribed-webhook?version=v20251201
type GetWebhookAPIReq struct {
	ID uint64 // required
}

func (c GetWebhookAPIReq) Verify() error {
	if c.ID == 0 {
		return fmt.Errorf("webhook id is required")
	}
	return nil
}

func (c GetWebhookAPIReq) Endpoint() string {
	return fmt.Sprintf("%d/webhooks.json", c.ID)
}

type GetWebhookAPIResp struct {
	client.BaseAPIResponse
	Webhook Webhook `json:"webhook,omitempty"`
}
