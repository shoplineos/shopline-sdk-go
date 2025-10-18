package payment

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// IPaymentStoreService
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-payments/balance?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-payments/balance?version=v20251201
type IPaymentStoreService interface {
	GetStoreBalance(context.Context, *GetStoreBalanceAPIReq) (*GetStoreBalanceAPIResp, error)
	ListStorePayouts(context.Context, *ListStorePayoutsAPIReq) (*ListStorePayoutsAPIResp, error)
	ListStoreBalanceTransactions(context.Context, *ListStoreBalanceTransactionsAPIReq) (*ListStoreBalanceTransactionsAPIResp, error)
	GetStorePayout(context.Context, *GetStorePayoutAPIReq) (*GetStorePayoutAPIResp, error)
	ListStoreTransactions(context.Context, *ListStoreTransactionsAPIReq) (*ListStoreTransactionsAPIResp, error)
}

var paymentStoreServiceInst = &PaymentStoreService{}

func GetPaymentStoreService() *PaymentStoreService {
	return paymentStoreServiceInst
}

type PaymentStoreService struct {
	client.BaseService
}

func (p PaymentStoreService) GetStoreBalance(ctx context.Context, req *GetStoreBalanceAPIReq) (*GetStoreBalanceAPIResp, error) {
	// 1. API response data
	apiResp := &GetStoreBalanceAPIResp{}

	// 2. Call the API
	err := p.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (p PaymentStoreService) ListStorePayouts(ctx context.Context, req *ListStorePayoutsAPIReq) (*ListStorePayoutsAPIResp, error) {
	// 1. API response data
	apiResp := &ListStorePayoutsAPIResp{}

	// 2. Call the API
	err := p.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (p PaymentStoreService) ListStoreBalanceTransactions(ctx context.Context, req *ListStoreBalanceTransactionsAPIReq) (*ListStoreBalanceTransactionsAPIResp, error) {
	// 1. API response data
	apiResp := &ListStoreBalanceTransactionsAPIResp{}

	// 2. Call the API
	err := p.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (p PaymentStoreService) GetStorePayout(ctx context.Context, req *GetStorePayoutAPIReq) (*GetStorePayoutAPIResp, error) {
	// 1. API response data
	apiResp := &GetStorePayoutAPIResp{}

	// 2. Call the API
	err := p.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (p PaymentStoreService) ListStoreTransactions(ctx context.Context, req *ListStoreTransactionsAPIReq) (*ListStoreTransactionsAPIResp, error) {
	// 1. API response data
	apiResp := &ListStoreTransactionsAPIResp{}

	// 2. Call the API
	err := p.Client.Call(ctx, req, apiResp)
	return apiResp, err
}
