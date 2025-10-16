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

func (c *DeleteWebhookAPIReq) Method() string {
	return "DELETE"
}

func (c *DeleteWebhookAPIReq) Verify() error {
	if c.Id == 0 {
		return errors.New("webhook.id is required")
	}
	return nil
}

func (c *DeleteWebhookAPIReq) Endpoint() string {
	return fmt.Sprintf("%d/webhooks.json", c.Id)
}

type DeleteWebhookAPIResp struct {
	client.BaseAPIResponse
}
