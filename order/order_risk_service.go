package order

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type IOrderRiskService interface {
	List(context.Context, *ListOrderRisksAPIReq) (*ListOrderRisksAPIResp, error)
	Get(context.Context, *GetOrderRiskAPIReq) (*GetOrderRiskAPIResp, error)
	Create(context.Context, *CreateOrderRiskAPIReq) (*CreateOrderRiskAPIResp, error)
	Update(context.Context, *UpdateOrderRiskAPIReq) (*UpdateOrderRiskAPIResp, error)
	DeleteAll(context.Context, *DeleteOrderRisksAPIReq) (*DeleteOrderRisksAPIResp, error)
	Delete(context.Context, *DeleteOrderRiskAPIReq) (*DeleteOrderRiskAPIResp, error)
}

var orderRiskServiceInst = &OrderRiskService{}

func GetOrderRiskService() *OrderRiskService {
	return orderRiskServiceInst
}

type OrderRiskService struct {
	client.BaseService
}

func (o OrderRiskService) List(ctx context.Context, req *ListOrderRisksAPIReq) (*ListOrderRisksAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		//Query: req, // API request params
	}
	// 2. API response data
	apiResp := &ListOrderRisksAPIResp{}
	// 3. Call API
	_, err := o.Client.Get(ctx, req.GetEndpoint(), shopLineReq, apiResp)
	return apiResp, err
}

func (o OrderRiskService) Get(ctx context.Context, req *GetOrderRiskAPIReq) (*GetOrderRiskAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		//Query: req, // API request params
	}
	// 2. API response data
	apiResp := &GetOrderRiskAPIResp{}
	// 3. Call API
	_, err := o.Client.Get(ctx, req.GetEndpoint(), shopLineReq, apiResp)
	return apiResp, err
}

func (o OrderRiskService) Create(ctx context.Context, req *CreateOrderRiskAPIReq) (*CreateOrderRiskAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: req, // API request params
	}
	// 2. API response data
	apiResp := &CreateOrderRiskAPIResp{}
	// 3. Call API
	_, err := o.Client.Post(ctx, req.GetEndpoint(), shopLineReq, apiResp)
	return apiResp, err
}

func (o OrderRiskService) Update(ctx context.Context, req *UpdateOrderRiskAPIReq) (*UpdateOrderRiskAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: req, // API request params
	}
	// 2. API response data
	apiResp := &UpdateOrderRiskAPIResp{}
	// 3. Call API
	_, err := o.Client.Put(ctx, req.GetEndpoint(), shopLineReq, apiResp)
	return apiResp, err
}

func (o OrderRiskService) DeleteAll(ctx context.Context, req *DeleteOrderRisksAPIReq) (*DeleteOrderRisksAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		//Data: req, // API request params
	}
	// 2. API response data
	apiResp := &DeleteOrderRisksAPIResp{}
	// 3. Call API
	_, err := o.Client.Delete(ctx, req.GetEndpoint(), shopLineReq, apiResp)
	return apiResp, err
}

func (o OrderRiskService) Delete(ctx context.Context, req *DeleteOrderRiskAPIReq) (*DeleteOrderRiskAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		//Query: req, // API request params
	}
	// 2. API response data
	apiResp := &DeleteOrderRiskAPIResp{}
	// 3. Call API
	_, err := o.Client.Delete(ctx, req.GetEndpoint(), shopLineReq, apiResp)
	return apiResp, err
}
