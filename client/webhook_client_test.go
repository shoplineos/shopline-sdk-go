package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var webhookClient *WebhookClient

func setupWebhook() {
	setup()
	webhookClient = NewWebhookClient(app)
}

type TestProductCreatedEvent struct {
	WebhookEvent
	Id       string `json:"id,omitempty"`
	BodyHtml string `json:"body_html,omitempty"`
}

func (p TestProductCreatedEvent) GetSupportedTopic() string {
	return "products/create"
}

func TestVerifyAndDecode(t *testing.T) {

	setupWebhook()
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

	event := &TestProductCreatedEvent{}
	err = webhookClient.Decode(req, event)
	if err != nil {
		t.Fatalf("Decode(%v) err = %v, expected nil", inURL, err)
	}

	assert.Equal(t, "32398389389389", event.Id)
	assert.Equal(t, "BodyHtml", event.BodyHtml)
	assert.Equal(t, "testMerchantId", event.Header.MerchantId)
}
