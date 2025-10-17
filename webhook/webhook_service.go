package webhook

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// IWebhookService
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/webhook/subscribe-to-a-webhook?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/webhook/subscribe-to-a-webhook?version=v20251201
type IWebhookService interface {
	List(context.Context, *ListWebhooksAPIReq) (*ListWebhooksAPIResp, error)
	Count(context.Context, *CountWebhooksAPIReq) (*CountWebhooksAPIResp, error)
	Get(context.Context, *GetWebhookAPIReq) (*GetWebhookAPIResp, error)
	Create(context.Context, *CreateWebhookAPIReq) (*CreateWebhookAPIResp, error)
	Update(context.Context, *UpdateWebhookAPIReq) (*UpdateWebhookAPIResp, error)
	Delete(context.Context, *DeleteWebhookAPIReq) (*DeleteWebhookAPIResp, error)
}

type WebhookService struct {
	client.BaseService
}

var serviceInst = &WebhookService{}

func GetWebhookService() *WebhookService {
	return serviceInst
}

func (w WebhookService) List(ctx context.Context, req *ListWebhooksAPIReq) (*ListWebhooksAPIResp, error) {
	// 1. API response resource
	apiResp := &ListWebhooksAPIResp{}

	// 2. Call the API
	err := w.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (w WebhookService) Count(ctx context.Context, req *CountWebhooksAPIReq) (*CountWebhooksAPIResp, error) {
	// 1. API response resource
	apiResp := &CountWebhooksAPIResp{}

	// 2. Call the API
	err := w.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (w WebhookService) Get(ctx context.Context, req *GetWebhookAPIReq) (*GetWebhookAPIResp, error) {
	// 1. API response resource
	apiResp := &GetWebhookAPIResp{}

	// 2. Call the API
	err := w.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (w WebhookService) Create(ctx context.Context, req *CreateWebhookAPIReq) (*CreateWebhookAPIResp, error) {
	// 1. API response resource
	apiResp := &CreateWebhookAPIResp{}

	// 2. Call the API
	err := w.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (w WebhookService) Update(ctx context.Context, req *UpdateWebhookAPIReq) (*UpdateWebhookAPIResp, error) {
	// 1. API response resource
	apiResp := &UpdateWebhookAPIResp{}

	// 2. Call the API
	err := w.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (w WebhookService) Delete(ctx context.Context, req *DeleteWebhookAPIReq) (*DeleteWebhookAPIResp, error) {
	// 1. API response resource
	apiResp := &DeleteWebhookAPIResp{}

	// 2. Call the API
	err := w.Client.Call(ctx, req, apiResp)
	return apiResp, err
}
