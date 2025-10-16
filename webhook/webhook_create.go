package webhook

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CreateWebhookAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/webhook/subscribe-to-a-webhook?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/webhook/subscribe-to-a-webhook?version=v20251201
type CreateWebhookAPIReq struct {
	client.BaseAPIRequest
	Webhook CreateWebhook `json:"webhook,omitempty"`
}

func (r *CreateWebhookAPIReq) Method() string {
	return "POST"
}

func (r *CreateWebhookAPIReq) GetData() interface{} {
	return r
}

func (r *CreateWebhookAPIReq) Verify() error {
	if r.Webhook.Address == "" {
		return errors.New("webhook.address is required")
	}
	if r.Webhook.Topic == "" {
		return errors.New("webhook.topic is required")
	}
	if r.Webhook.ApiVersion == "" {
		return errors.New("webhook.api_version is required")
	}
	return nil
}

func (r *CreateWebhookAPIReq) Endpoint() string {
	return "webhooks.json"
}

type CreateWebhook struct {
	Address    string `json:"address,omitempty"`
	Topic      string `json:"topic,omitempty"`
	ApiVersion string `json:"api_version,omitempty"`
}

type CreateWebhookAPIResp struct {
	client.BaseAPIResponse
	Webhook Webhook `json:"webhook,omitempty"`
}
