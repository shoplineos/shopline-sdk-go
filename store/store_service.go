package store

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// IStoreService
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/query-store-information?version=v20251201
// En: https://developer.shopline.com/docs/admin-rest-api/store/query-store-information?version=v20251201
type IStoreService interface {
	Get(context.Context, *GetStoreAPIReq) (*GetStoreAPIResp, error)
	ListCurrencies(context.Context, *ListStoreCurrenciesAPIReq) (*ListStoreCurrenciesAPIResp, error)
	GetStaff(context.Context, *GetStoreStaffAPIReq) (*GetStoreStaffAPIResp, error)
	ListStaffs(context.Context, *ListStoreStaffsAPIReq) (*ListStoreStaffsAPIResp, error)
	GetOperationLog(context.Context, *GetStoreOperationLogAPIReq) (*GetStoreOperationLogAPIResp, error)
	ListOperationLogs(context.Context, *ListStoreOperationLogsAPIReq) (*ListStoreOperationLogsAPIResp, error)
	CountOperationLogs(context.Context, *CountStoreOperationLogsAPIReq) (*CountStoreOperationLogsAPIResp, error)
	ListSubscriptions(context.Context, *ListStoreSubscriptionsAPIReq) (*ListStoreSubscriptionsAPIResp, error)
}

var serviceInst = &StoreService{}

func GetStoreService() *StoreService {
	return serviceInst
}

type StoreService struct {
	client.BaseService
}

func (s StoreService) Get(ctx context.Context, req *GetStoreAPIReq) (*GetStoreAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &GetStoreAPIResp{}

	// 4. Call API
	_, err := s.Client.Get(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (s StoreService) ListCurrencies(ctx context.Context, req *ListStoreCurrenciesAPIReq) (*ListStoreCurrenciesAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &ListStoreCurrenciesAPIResp{}

	// 4. Call API
	_, err := s.Client.Get(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (s StoreService) GetStaff(ctx context.Context, req *GetStoreStaffAPIReq) (*GetStoreStaffAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &GetStoreStaffAPIResp{}

	// 4. Call API
	_, err := s.Client.Get(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (s StoreService) ListStaffs(ctx context.Context, req *ListStoreStaffsAPIReq) (*ListStoreStaffsAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: req,
	}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &ListStoreStaffsAPIResp{}

	// 4. Call API
	_, err := s.Client.Get(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (s StoreService) GetOperationLog(ctx context.Context, req *GetStoreOperationLogAPIReq) (*GetStoreOperationLogAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &GetStoreOperationLogAPIResp{}

	// 4. Call API
	_, err := s.Client.Get(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (s StoreService) ListOperationLogs(ctx context.Context, req *ListStoreOperationLogsAPIReq) (*ListStoreOperationLogsAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: req,
	}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &ListStoreOperationLogsAPIResp{}

	// 4. Call API
	_, err := s.Client.Get(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (s StoreService) CountOperationLogs(ctx context.Context, req *CountStoreOperationLogsAPIReq) (*CountStoreOperationLogsAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: req,
	}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &CountStoreOperationLogsAPIResp{}

	// 4. Call API
	_, err := s.Client.Get(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (s StoreService) ListSubscriptions(ctx context.Context, req *ListStoreSubscriptionsAPIReq) (*ListStoreSubscriptionsAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: req,
	}

	// 2. API endpoint
	endpoint := req.Endpoint()

	// 3. API response data
	apiResp := &ListStoreSubscriptionsAPIResp{}

	// 4. Call API
	_, err := s.Client.Get(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}
