package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// IWebhookEvent
// For more details, see:
// 中文: https://developer.shopline.com/zh-hans-cn/docs/webhook/product/create-product?version=v20251201
// En: https://developer.shopline.com/docs/webhook/product/create-product?version=v20251201
type IWebhookEvent interface {
	GetHeader() *WebhookEventHeader
	SetHeader(*WebhookEventHeader)
	GetSupportedTopic() string
}

type WebhookEventHeader struct {
	HmacSha256 string
	Topic      string
	ShopDomain string
	ShopId     string
	MerchantId string
	APIVersion string
	WebhookId  string
}
type WebhookEvent struct {
	Header *WebhookEventHeader
}

func (e *WebhookEvent) GetHeader() *WebhookEventHeader {
	return e.Header
}

func (e *WebhookEvent) SetHeader(h *WebhookEventHeader) {
	e.Header = h
}

type WebhookClient struct {
	App App
}

func NewWebhookClient(app App) *WebhookClient {
	return &WebhookClient{
		App: app,
	}
}

// VerifyWebhookRequest Verify a Webhook http request, sent by SHOPLINE.
// The body of the request is still readable after invoking the method.
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
// En: https://developer.shopline.com/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header

func (client *WebhookClient) VerifyWebhookRequest(r *http.Request) bool {
	return client.App.VerifyWebhookRequest(r)
}

// Decode a Webhook http request, sent by SHOPLINE.
// event To specify the return event resource, an WebhookEvent or your own event resource
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
// En: https://developer.shopline.com/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
func (client *WebhookClient) Decode(r *http.Request, event interface{}) error {
	if event == nil {
		return errors.New("event is required")
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	h, err := verifyHeader(r)

	actualMac := []byte(h.HmacSha256)

	mac := hmac.New(sha256.New, []byte(client.App.AppSecret))

	mac.Write(requestBody)
	macSum := mac.Sum(nil)
	expectedMac := []byte(hex.EncodeToString(macSum))

	success := hmac.Equal(actualMac, expectedMac)
	if !success {
		return errors.New(fmt.Sprintf("HMAC verification failed. Expected: %s, Actual: %s", expectedMac, actualMac))
	}

	if _, ok := event.(IWebhookEvent); ok {
		we := event.(IWebhookEvent)
		we.SetHeader(h)
	}

	err = json.Unmarshal(requestBody, event)
	return err
}

func verifyHeader(r *http.Request) (*WebhookEventHeader, error) {
	weh := &WebhookEventHeader{}
	shoplineSha256 := r.Header.Get("X-Shopline-Hmac-Sha256")
	weh.HmacSha256 = shoplineSha256
	weh.ShopDomain = r.Header.Get("X-Shopline-Shop-Domain")
	weh.MerchantId = r.Header.Get("X-Shopline-Merchant-Id")
	weh.APIVersion = r.Header.Get("X-Shopline-Api-Version")
	weh.WebhookId = r.Header.Get("X-Shopline-Webhook-Id")
	weh.Topic = r.Header.Get("X-Shopline-Topic")
	weh.ShopId = r.Header.Get("X-Shopline-Shop-Id")

	if shoplineSha256 == "" {
		return weh, errors.New("X-Shopline-Hmac-Sha256 is required")
	}

	if weh.ShopDomain == "" {
		return weh, errors.New("ShopDomain is required")
	}

	if weh.MerchantId == "" {
		return weh, errors.New("MerchantId is required")
	}
	if weh.APIVersion == "" {
		return weh, errors.New("ApiVersion is required")
	}

	if weh.WebhookId == "" {
		return weh, errors.New("WebhookId is required")
	}

	if weh.Topic == "" {
		return weh, errors.New("topic is required")
	}

	if weh.ShopId == "" {
		return weh, errors.New("ShopId is required")
	}

	return weh, nil
}
