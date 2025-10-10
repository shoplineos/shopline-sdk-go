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
	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &ListWebhooksAPIResp{}

	// 4. Call API
	_, err := w.Client.Get(context.Background(), endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (w WebhookService) Count(ctx context.Context, req *CountWebhooksAPIReq) (*CountWebhooksAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &CountWebhooksAPIResp{}

	// 4. Call API
	_, err := w.Client.Get(context.Background(), endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (w WebhookService) Get(ctx context.Context, req *GetWebhookAPIReq) (*GetWebhookAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &GetWebhookAPIResp{}

	// 4. Call API
	_, err := w.Client.Get(context.Background(), endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (w WebhookService) Create(ctx context.Context, req *CreateWebhookAPIReq) (*CreateWebhookAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: req,
	}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &CreateWebhookAPIResp{}

	// 4. Call API
	_, err := w.Client.Post(context.Background(), endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (w WebhookService) Update(ctx context.Context, req *UpdateWebhookAPIReq) (*UpdateWebhookAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: req,
	}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &UpdateWebhookAPIResp{}

	// 4. Call API
	_, err := w.Client.Put(context.Background(), endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (w WebhookService) Delete(ctx context.Context, req *DeleteWebhookAPIReq) (*DeleteWebhookAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &DeleteWebhookAPIResp{}

	// 4. Call API
	_, err := w.Client.Delete(context.Background(), endpoint, shopLineReq, apiResp)

	return apiResp, err
}
