package access

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// IStorefrontAccessTokenService
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/access/storefront-api/create-an-access-token?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/access/storefront-api/create-an-access-token?version=v20251201
type IStorefrontAccessTokenService interface {
	List(context.Context, *ListStorefrontAccessTokensAPIReq) (*ListStorefrontAccessTokensAPIResp, error)
	Delete(context.Context, *DeleteStorefrontAccessTokenAPIReq) (*DeleteStorefrontAccessTokenAPIResp, error)
	Create(context.Context, *CreateStorefrontAccessTokenAPIReq) (*CreateStorefrontAccessTokenAPIResp, error)
}

var serviceInst = &StorefrontAccessTokenService{}

func GetStorefrontAccessTokenService() *StorefrontAccessTokenService {
	return serviceInst
}

type StorefrontAccessTokenService struct {
	client.BaseService
}

func (s StorefrontAccessTokenService) List(ctx context.Context, apiReq *ListStorefrontAccessTokensAPIReq) (*ListStorefrontAccessTokensAPIResp, error) {
	// 1. API response resource
	apiResp := &ListStorefrontAccessTokensAPIResp{}

	// 2. Call the API
	err := s.Client.Call(ctx, apiReq, apiResp)
	return apiResp, err
}

func (s StorefrontAccessTokenService) ListAll(ctx context.Context, apiReq *ListStorefrontAccessTokensAPIReq) ([]StorefrontAccessToken, error) {
	collector := []StorefrontAccessToken{}
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	for {
		// 2. API endpoint
		endpoint := apiReq.Endpoint()

		// 3. API response data
		apiResp := &ListStorefrontAccessTokensAPIResp{}

		// 4. Call API
		shoplineResp, err := s.Client.Get(ctx, endpoint, shopLineReq, apiResp)

		if err != nil {
			return collector, err
		}

		collector = append(collector, apiResp.StorefrontAccessTokens...)

		if !shoplineResp.HasNext() {
			break
		}

		shopLineReq.Query = shoplineResp.Pagination.Next
	}

	return collector, nil
}

func (s StorefrontAccessTokenService) ListWithPagination(ctx context.Context, req *ListStorefrontAccessTokensAPIReq) (*ListStorefrontAccessTokensAPIResp, error) {
	return s.List(ctx, req)
}

func (s StorefrontAccessTokenService) Delete(ctx context.Context, apiReq *DeleteStorefrontAccessTokenAPIReq) (*DeleteStorefrontAccessTokenAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request params
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &DeleteStorefrontAccessTokenAPIResp{}

	// 4. Call API
	_, err := s.Client.Delete(ctx, endpoint, shopLineReq, apiResp)
	return apiResp, err
}

func (s StorefrontAccessTokenService) Create(ctx context.Context, apiReq *CreateStorefrontAccessTokenAPIReq) (*CreateStorefrontAccessTokenAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request params
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &CreateStorefrontAccessTokenAPIResp{}

	// 4. Call API
	_, err := s.Client.Post(ctx, endpoint, shopLineReq, apiResp)
	return apiResp, err
}
