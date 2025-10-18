package order

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type IOrderRefundService interface {
	Refund(context.Context, *RefundAPIReq) (*RefundAPIResp, error)
	List(context.Context, *ListOrderRefundsAPIReq) (*ListOrderRefundsAPIResp, error)
	Get(context.Context, *GetOrderRefundAPIReq) (*GetOrderRefundAPIResp, error)
	Calculate(context.Context, *CalculateOrderRefundAPIReq) (*CalculateOrderRefundAPIResp, error)
}

var orderRefundServiceInst = &OrderRefundService{}

func GetOrderRefundService() *OrderRefundService {
	return orderRefundServiceInst
}

type OrderRefundService struct {
	client.BaseService
}

func (o *OrderRefundService) List(ctx context.Context, req *ListOrderRefundsAPIReq) (*ListOrderRefundsAPIResp, error) {
	// 1. API response resource
	apiResp := &ListOrderRefundsAPIResp{}

	// 2. Call the API
	err := o.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (o *OrderRefundService) Get(ctx context.Context, req *GetOrderRefundAPIReq) (*GetOrderRefundAPIResp, error) {

	// 1. API response data
	apiResp := &GetOrderRefundAPIResp{}

	// 2. Call API
	err := o.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (o *OrderRefundService) Calculate(ctx context.Context, req *CalculateOrderRefundAPIReq) (*CalculateOrderRefundAPIResp, error) {
	// 1. API response data
	apiResp := &CalculateOrderRefundAPIResp{}

	// 2. Call API
	err := o.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

// Refund
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/order-refund?version=v20251201
// en：https://developer.shopline.com/docs/admin-rest-api/order/order-management/order-refund?version=v20251201
func (o *OrderRefundService) Refund(ctx context.Context, req *RefundAPIReq) (*RefundAPIResp, error) {
	// 1. API response data
	apiResp := &RefundAPIResp{}

	// 2. Call API
	err := o.Client.Call(ctx, req, apiResp)
	return apiResp, err
}
