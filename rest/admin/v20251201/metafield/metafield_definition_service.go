package metafield

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// IMetafieldDefinitionService Metafield Definition interface
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/create-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/create-a-metafield-definition?version=v20251201
// Deprecated
// use the client.Call
type IMetafieldDefinitionService interface {
	List(context.Context, *ListMetafieldDefinitionAPIReq) (*ListMetafieldDefinitionAPIResp, error)
	ListAll(context.Context, *ListMetafieldDefinitionAPIReq) ([]MetafieldDefinition, error)
	ListWithPagination(context.Context, *ListMetafieldDefinitionAPIReq) (*ListMetafieldDefinitionAPIResp, error)
	Get(context.Context, *GetMetafieldDefinitionAPIReq) (*GetMetafieldDefinitionAPIResp, error)
	Delete(context.Context, *DeleteMetafieldDefinitionAPIReq) (*DeleteMetafieldDefinitionAPIResp, error)
	Update(context.Context, *UpdateMetafieldDefinitionAPIReq) (*UpdateMetafieldDefinitionAPIResp, error)
	Create(context.Context, *CreateMetafieldDefinitionAPIReq) (*CreateMetafieldDefinitionAPIResp, error)
}

var metafieldDefinitionServiceInst = &MetafieldDefinitionService{}

// GetMetafieldDefinitionService
// Deprecated
// use the client.Call
func GetMetafieldDefinitionService() *MetafieldDefinitionService {
	return metafieldDefinitionServiceInst
}

type MetafieldDefinitionService struct {
	client.BaseService
}

func (m MetafieldDefinitionService) ListWithPagination(ctx context.Context, req *ListMetafieldDefinitionAPIReq) (*ListMetafieldDefinitionAPIResp, error) {
	return m.List(ctx, req)
}

func (m MetafieldDefinitionService) List(ctx context.Context, apiReq *ListMetafieldDefinitionAPIReq) (*ListMetafieldDefinitionAPIResp, error) {
	// 1. API response resource
	apiResp := &ListMetafieldDefinitionAPIResp{}

	// 2. Call the API
	err := m.Client.Call(ctx, apiReq, apiResp)
	return apiResp, err
}

func getMetafieldDefinitions(resp interface{}) []MetafieldDefinition {
	apiResp := resp.(*ListMetafieldDefinitionAPIResp)
	return apiResp.Data.MetafieldDefinitions
}

func (m MetafieldDefinitionService) ListAll(ctx context.Context, apiReq *ListMetafieldDefinitionAPIReq) ([]MetafieldDefinition, error) {
	apiResp := &ListMetafieldDefinitionAPIResp{}
	return client.ListAll(m.Client, ctx, apiReq, apiResp, getMetafieldDefinitions)
}

func (m MetafieldDefinitionService) Get(ctx context.Context, apiReq *GetMetafieldDefinitionAPIReq) (*GetMetafieldDefinitionAPIResp, error) {
	// 1. API response resource
	apiResp := &GetMetafieldDefinitionAPIResp{}

	// 2. Call the API
	err := m.Client.Call(ctx, apiReq, apiResp)
	return apiResp, err
}

func (m MetafieldDefinitionService) Delete(ctx context.Context, apiReq *DeleteMetafieldDefinitionAPIReq) (*DeleteMetafieldDefinitionAPIResp, error) {

	// 1. API response resource
	apiResp := &DeleteMetafieldDefinitionAPIResp{}

	// 2. Call the API
	err := m.Client.Call(ctx, apiReq, apiResp)
	return apiResp, err
}

func (m MetafieldDefinitionService) Update(ctx context.Context, apiReq *UpdateMetafieldDefinitionAPIReq) (*UpdateMetafieldDefinitionAPIResp, error) {
	// 1. API response resource
	apiResp := &UpdateMetafieldDefinitionAPIResp{}

	// 2. Call the API
	err := m.Client.Call(ctx, apiReq, apiResp)
	return apiResp, err
}

func (m MetafieldDefinitionService) Create(ctx context.Context, apiReq *CreateMetafieldDefinitionAPIReq) (*CreateMetafieldDefinitionAPIResp, error) {
	// 1. API response resource
	apiResp := &CreateMetafieldDefinitionAPIResp{}

	// 2. Call the API
	err := m.Client.Call(ctx, apiReq, apiResp)
	return apiResp, err
}
