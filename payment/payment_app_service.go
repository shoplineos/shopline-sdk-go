package payment

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// IPaymentAppService
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/payments-app-api/merchant-activation-successful-notification?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/payments-app-api/merchant-activation-successful-notification?version=v20251201
type IPaymentAppService interface {
	MerchantBindSuccess(context.Context, *MerchantBindSuccessAPIReq) (*MerchantBindSuccessAPIResp, error)
	MerchantPaySuccess(context.Context, *MerchantPaySuccessAPIReq) (*MerchantPaySuccessAPIResp, error)
	MerchantRefundSuccess(context.Context, *MerchantRefundSuccessAPIReq) (*MerchantRefundSuccessAPIResp, error)
	MerchantDeviceBindSuccess(context.Context, *MerchantDeviceBindSuccessAPIReq) (*MerchantDeviceBindSuccessAPIResp, error)
}

var serviceInst = &MerchantAppService{}

func GetMerchantAppService() *MerchantAppService {
	return serviceInst
}

type MerchantAppService struct {
	client.BaseService
}

func (m MerchantAppService) MerchantBindSuccess(ctx context.Context, req *MerchantBindSuccessAPIReq) (*MerchantBindSuccessAPIResp, error) {

	// 1. API response data
	apiResp := &MerchantBindSuccessAPIResp{}

	// 2. Call the API
	err := m.Client.Call(ctx, req, apiResp)

	return apiResp, err
}

func (m MerchantAppService) MerchantPaySuccess(ctx context.Context, req *MerchantPaySuccessAPIReq) (*MerchantPaySuccessAPIResp, error) {
	// 1. API response data
	apiResp := &MerchantPaySuccessAPIResp{}

	// 2. Call the API
	err := m.Client.Call(ctx, req, apiResp)

	return apiResp, err
}

func (m MerchantAppService) MerchantRefundSuccess(ctx context.Context, req *MerchantRefundSuccessAPIReq) (*MerchantRefundSuccessAPIResp, error) {
	// 1. API response data
	apiResp := &MerchantRefundSuccessAPIResp{}

	// 2. Call the API
	err := m.Client.Call(ctx, req, apiResp)

	return apiResp, err
}

func (m MerchantAppService) MerchantDeviceBindSuccess(ctx context.Context, req *MerchantDeviceBindSuccessAPIReq) (*MerchantDeviceBindSuccessAPIResp, error) {
	// 1. API response data
	apiResp := &MerchantDeviceBindSuccessAPIResp{}

	// 2. Call the API
	err := m.Client.Call(ctx, req, apiResp)

	return apiResp, err
}
