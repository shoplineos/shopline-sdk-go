package webhook

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// DeleteWebhookAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/webhook/unsubscribe-from-a-webhook?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/webhook/unsubscribe-from-a-webhook?version=v20251201
type DeleteWebhookAPIReq struct {
	client.BaseAPIRequest
	Id uint64
}

func (r *DeleteWebhookAPIReq) Method() string {
	return "DELETE"
}

func (r *DeleteWebhookAPIReq) GetData() interface{} {
	return r
}

func (r *DeleteWebhookAPIReq) Verify() error {
	if r.Id == 0 {
		return errors.New("webhook.id is required")
	}
	return nil
}

func (r *DeleteWebhookAPIReq) Endpoint() string {
	return fmt.Sprintf("%d/webhooks.json", r.Id)
}

type DeleteWebhookAPIResp struct {
	client.BaseAPIResponse
}
