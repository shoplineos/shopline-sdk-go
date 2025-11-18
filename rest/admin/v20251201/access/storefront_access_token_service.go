package access

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// IStorefrontAccessTokenService
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/access/storefront-api/create-an-access-token?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/access/storefront-api/create-an-access-token?version=v20251201
// Deprecated
// use the client.Call
type IStorefrontAccessTokenService interface {
	List(context.Context, *ListStorefrontAccessTokensAPIReq) (*ListStorefrontAccessTokensAPIResp, error)
	Delete(context.Context, *DeleteStorefrontAccessTokenAPIReq) (*DeleteStorefrontAccessTokenAPIResp, error)
	Create(context.Context, *CreateStorefrontAccessTokenAPIReq) (*CreateStorefrontAccessTokenAPIResp, error)
}

var serviceInst = &StorefrontAccessTokenService{}

// GetStorefrontAccessTokenService
// Deprecated
// use the client.Call
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

func getResources(resp interface{}) []StorefrontAccessToken {
	apiResp := resp.(*ListStorefrontAccessTokensAPIResp)
	return apiResp.StorefrontAccessTokens
}

func (s StorefrontAccessTokenService) ListAll(ctx context.Context, apiReq *ListStorefrontAccessTokensAPIReq) ([]StorefrontAccessToken, error) {

	apiResp := &ListStorefrontAccessTokensAPIResp{}
	return client.ListAll(s.Client, ctx, apiReq, apiResp, getResources)
}

func (s StorefrontAccessTokenService) ListWithPagination(ctx context.Context, req *ListStorefrontAccessTokensAPIReq) (*ListStorefrontAccessTokensAPIResp, error) {
	return s.List(ctx, req)
}

func (s StorefrontAccessTokenService) Delete(ctx context.Context, apiReq *DeleteStorefrontAccessTokenAPIReq) (*DeleteStorefrontAccessTokenAPIResp, error) {
	// 1. API response resource
	apiResp := &DeleteStorefrontAccessTokenAPIResp{}

	// 2. Call the API
	err := s.Client.Call(ctx, apiReq, apiResp)
	return apiResp, err
}

func (s StorefrontAccessTokenService) Create(ctx context.Context, apiReq *CreateStorefrontAccessTokenAPIReq) (*CreateStorefrontAccessTokenAPIResp, error) {
	// 1. API response resource
	apiResp := &CreateStorefrontAccessTokenAPIResp{}

	// 2. Call the API
	err := s.Client.Call(ctx, apiReq, apiResp)
	return apiResp, err
}
