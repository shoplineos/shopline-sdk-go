package webhook

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/webhooks.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"webhook":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "address":"test desc", "topic":"key_test", "apiVersion":"123"}}`))

	req := &CreateWebhookAPIReq{
		Webhook: CreateWebhook{
			Address:    "test desc",
			Topic:      "key_test",
			ApiVersion: "133",
		},
	}

	apiResp, err := GetWebhookService().Create(context.Background(), req)
	if err != nil {
		t.Errorf("Webhook.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, uint64(1), apiResp.Webhook.Id)
	assert.Equal(t, "test desc", apiResp.Webhook.Address)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", apiResp.Webhook.CreatedAt)
}

func TestUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/1/webhooks.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"webhook":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-10-10T14:48:44-04:00", "address":"test1 desc", "topic":"key_test", "apiVersion":"123"}}`))

	req := &UpdateWebhookAPIReq{
		ID: 1,
		Webhook: UpdateWebhook{
			Address: "test1 desc",
		},
	}

	apiResp, err := GetWebhookService().Update(context.Background(), req)
	if err != nil {
		t.Errorf("Webhook.Update returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, uint64(1), apiResp.Webhook.Id)
	assert.Equal(t, "test1 desc", apiResp.Webhook.Address)
	assert.Equal(t, "2025-10-10T14:48:44-04:00", apiResp.Webhook.UpdatedAt)
}

func TestGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/1/webhooks.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"webhook":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "address":"test1 desc", "topic":"key_test", "apiVersion":"123"}}`))

	req := &GetWebhookAPIReq{
		ID: 1,
	}

	apiResp, err := GetWebhookService().Get(context.Background(), req)
	if err != nil {
		t.Errorf("Webhook.Get returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, uint64(1), apiResp.Webhook.Id)
	assert.Equal(t, "test1 desc", apiResp.Webhook.Address)
}

func TestList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/webhooks.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"webhooks":[{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "address":"test1 desc", "topic":"key_test", "apiVersion":"123"}]}`))

	req := &ListWebhooksAPIReq{}

	apiResp, err := GetWebhookService().List(context.Background(), req)
	if err != nil {
		t.Errorf("Webhook.List returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, len(apiResp.Webhooks))

	webhook := apiResp.Webhooks[0]
	assert.Equal(t, uint64(1), webhook.Id)
	assert.Equal(t, "test1 desc", webhook.Address)
}

func TestDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/1/webhooks.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	req := &DeleteWebhookAPIReq{
		ID: 1,
	}

	apiResp, err := GetWebhookService().Delete(context.Background(), req)
	if err != nil {
		t.Errorf("Webhook.Delete returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
}

func TestCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/webhooks/count.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"count":1}`))

	req := &CountWebhooksAPIReq{}

	apiResp, err := GetWebhookService().Count(context.Background(), req)
	if err != nil {
		t.Errorf("Webhook.Count returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, apiResp.Count)
}
