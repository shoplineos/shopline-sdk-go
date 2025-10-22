package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ProductCreatedEvent struct {
	WebhookEvent
	Id       string `json:"id,omitempty"`
	BodyHtml string `json:"body_html,omitempty"`
}

func (p ProductCreatedEvent) GetSupportedTopic() string {
	return "products/create"
}

func TestVerifyAndDecode(t *testing.T) {

	setup()
	defer teardown()

	inURL := "foo?page=1"

	inBody := struct {
		BodyHtml string `json:"body_html,omitempty"`
		Id       string `json:"id,omitempty"`
	}{
		Id:       "32398389389389",
		BodyHtml: "BodyHtml",
	}

	slReq := &ShopLineRequest{
		Data: inBody,
	}

	req, err := client.NewHttpRequest(context.Background(), MethodPost, inURL, slReq)
	req.Header.Set("X-Shopline-Hmac-Sha256", "edf64f4c46a9cb2c8aaa6c7034b77899aca1b63ad0c461901417a42b1cf139b0")
	req.Header.Set("X-Shopline-Shop-Domain", "testDomain")
	req.Header.Set("X-Shopline-Merchant-Id", "testMerchantId")
	req.Header.Set("X-Shopline-Api-Version", "v20251201")
	req.Header.Set("X-Shopline-Webhook-Id", "testWebhookId")
	req.Header.Set("X-Shopline-Topic", "testTopic")
	req.Header.Set("X-Shopline-Shop-Id", "testShopId")

	if err != nil {
		t.Fatalf("NewHttpRequest(%v) err = %v, expected nil", inURL, err)
	}

	webhookClient := NewWebhookClient(app)

	e := &ProductCreatedEvent{}
	err = webhookClient.VerifyAndDecode(req, e)
	if err != nil {
		t.Fatalf("VerifyAndDecode(%v) err = %v, expected nil", inURL, err)
	}

	assert.Equal(t, "32398389389389", e.Id)
	assert.Equal(t, "BodyHtml", e.BodyHtml)
	assert.Equal(t, "testMerchantId", e.Header.MerchantId)

}
