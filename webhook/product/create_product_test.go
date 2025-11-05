package product

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateProductEvent(t *testing.T) {
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

	slReq := &client.ShopLineRequest{
		Data: inBody,
	}

	req, err := cli.NewHttpRequest(context.Background(), http.MethodPost, inURL, slReq)
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

	e := &CreateProductEvent{}
	err = webhookClient.Decode(req, e)
	if err != nil {
		t.Fatalf("Decode(%v) err = %v, expected nil", inURL, err)
	}

	assert.Equal(t, "32398389389389", e.Id)
	assert.Equal(t, "BodyHtml", e.BodyHtml)
	assert.Equal(t, "testMerchantId", e.Header.MerchantId)

}
