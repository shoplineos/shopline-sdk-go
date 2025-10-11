package metafield

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// IMetafieldService
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafields/resource-metafields/create-a-metafield-for-a-resource?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafields/resource-metafields/create-a-metafield-for-a-resource?version=v20251201
type IMetafieldService interface {
	List(context.Context, *ListMetafieldAPIReq) (*ListMetafieldAPIResp, error)
	ListAll(context.Context, *ListMetafieldAPIReq) ([]Metafield, error)
	ListWithPagination(context.Context, *ListMetafieldAPIReq) (*ListMetafieldAPIResp, error)
	Get(context.Context, *GetMetafieldAPIReq) (*GetMetafieldAPIResp, error)
	Count(context.Context, *CountMetafieldAPIReq) (*CountMetafieldAPIResp, error)
	Delete(context.Context, *DeleteMetafieldAPIReq) (*DeleteMetafieldAPIResp, error)
	Update(context.Context, *UpdateMetafieldAPIReq) (*UpdateMetafieldAPIResp, error)
	Create(context.Context, *CreateMetafieldAPIReq) (*CreateMetafieldAPIResp, error)
}

var metafieldServiceInst = &MetafieldService{}

func GetMetafieldService() *MetafieldService {
	return metafieldServiceInst
}

type MetafieldService struct {
	client.BaseService
}

func (m MetafieldService) ListWithPagination(ctx context.Context, req *ListMetafieldAPIReq) (*ListMetafieldAPIResp, error) {
	return m.List(ctx, req)
}

func (m MetafieldService) Get(ctx context.Context, apiReq *GetMetafieldAPIReq) (*GetMetafieldAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &GetMetafieldAPIResp{}

	// 4. Call API
	_, err := m.Client.Get(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (m MetafieldService) List(ctx context.Context, apiReq *ListMetafieldAPIReq) (*ListMetafieldAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &ListMetafieldAPIResp{}

	// 4. Call API
	_, err := m.Client.Get(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (m MetafieldService) ListAll(ctx context.Context, apiReq *ListMetafieldAPIReq) ([]Metafield, error) {
	collector := []Metafield{}
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	for {
		// 2. API endpoint
		endpoint := apiReq.Endpoint()

		// 3. API response data
		apiResp := &ListMetafieldAPIResp{}

		// 4. Call API
		shoplineResp, err := m.Client.Get(ctx, endpoint, shopLineReq, apiResp)

		if err != nil {
			return collector, err
		}

		collector = append(collector, apiResp.Metafields...)

		if !shoplineResp.HasNext() {
			break
		}

		shopLineReq.Query = shoplineResp.Pagination.Next
	}

	return collector, nil
}

func (m MetafieldService) Count(ctx context.Context, apiReq *CountMetafieldAPIReq) (*CountMetafieldAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &CountMetafieldAPIResp{}

	// 4. Call API
	_, err := m.Client.Get(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (m MetafieldService) Delete(ctx context.Context, apiReq *DeleteMetafieldAPIReq) (*DeleteMetafieldAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &DeleteMetafieldAPIResp{}

	// 4. Call API
	_, err := m.Client.Delete(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (m MetafieldService) Update(ctx context.Context, apiReq *UpdateMetafieldAPIReq) (*UpdateMetafieldAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request params
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &UpdateMetafieldAPIResp{}

	// 4. Call API
	_, err := m.Client.Put(ctx, endpoint, shopLineReq, apiResp)
	if err != nil {
		return nil, err
	}

	return apiResp, nil
}

func (m MetafieldService) Create(ctx context.Context, apiReq *CreateMetafieldAPIReq) (*CreateMetafieldAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request params
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &CreateMetafieldAPIResp{}

	// 4. Call API
	_, err := m.Client.Post(ctx, endpoint, shopLineReq, apiResp)
	if err != nil {
		return nil, err
	}

	return apiResp, nil
}
