package webhook

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// UpdateWebhookAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/webhook/update-a-subscribed-webhook?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/webhook/update-a-subscribed-webhook?version=v20251201
type UpdateWebhookAPIReq struct {
	Id      uint64
	Webhook UpdateWebhook `json:"webhook,omitempty"`
}

func (c UpdateWebhookAPIReq) Verify() error {
	if c.Id == 0 {
		return errors.New("id is required")
	}
	if c.Webhook.Address == "" {
		return errors.New("webhook.address is required")
	}
	return nil
}

func (c UpdateWebhookAPIReq) Endpoint() string {
	return fmt.Sprintf("%d/webhooks.json", c.Id)
}

type UpdateWebhook struct {
	Address string `json:"address,omitempty"`
}

type UpdateWebhookAPIResp struct {
	client.BaseAPIResponse
	Webhook Webhook `json:"webhook,omitempty"`
}
