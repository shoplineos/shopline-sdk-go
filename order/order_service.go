package order

import (
	"context"
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type IOrderService interface {
	List(context.Context, *ListOrdersAPIReq) (*ListOrdersAPIResp, error)
	ListAll(context.Context, *ListOrdersAPIReq) ([]Order, error)
	ListWithPagination(context.Context, *ListOrdersAPIReq) (*ListOrdersAPIResp, error)
	Count(context.Context, *GetOrdersCountAPIReq) (*GetOrdersCountAPIResp, error)
	Get(context.Context, *GetOrderDetailAPIReq) (*GetOrderDetailAPIResp, error)
	Create(context.Context, *CreateOrderAPIReq) (*CreateOrderAPIResp, error)
	Update(context.Context, *UpdateOrderAPIReq) (*UpdateOrderAPIResp, error)
	Delete(context.Context, *DeleteOrderAPIReq) (*DeleteOrderAPIResp, error)
	Cancel(context.Context, *CancelOrderAPIReq) (*CancelOrderAPIResp, error)

	ListAttributionInfos(context.Context, *ListOrderAttributionInfosAPIReq) (*ListOrderAttributionInfosAPIResp, error)
}

var serviceInst = &OrderService{}

func GetOrderService() *OrderService {
	return serviceInst
}

type OrderService struct {
	client.BaseService
}

func (o *OrderService) ListAttributionInfos(ctx context.Context, req *ListOrderAttributionInfosAPIReq) (*ListOrderAttributionInfosAPIResp, error) {
	// 1. API response resource
	apiResp := &ListOrderAttributionInfosAPIResp{}

	// 2. Call the API
	err := o.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (o *OrderService) Cancel(ctx context.Context, req *CancelOrderAPIReq) (*CancelOrderAPIResp, error) {
	// 1. API response resource
	apiResp := &CancelOrderAPIResp{}

	// 2. Call the API
	err := o.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (o *OrderService) List(ctx context.Context, req *ListOrdersAPIReq) (*ListOrdersAPIResp, error) {
	// 1. API response resource
	apiResp := &ListOrdersAPIResp{}

	// 2. Call the API
	err := o.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func getResources(resp interface{}) []Order {
	apiResp := resp.(*ListOrdersAPIResp)
	return apiResp.Orders
}

func (o *OrderService) ListAll(ctx context.Context, apiReq *ListOrdersAPIReq) ([]Order, error) {

	return client.ListAll(o.Client, ctx, apiReq, &ListOrdersAPIResp{}, getResources)

	//
	//collector := []Order{}
	//// 1. API request
	//shopLineReq := &client.ShopLineRequest{
	//	Query: apiReq, // API request params
	//}
	//
	//for {
	//	// 2. API endpoint
	//	endpoint := apiReq.GetEndpoint()
	//
	//	// 3. API response resource
	//	apiResp := &ListOrdersAPIResp{}
	//
	//	// 4. Call the API
	//	shoplineResp, err := o.Client.Get(ctx, endpoint, shopLineReq, apiResp)
	//
	//	if err != nil {
	//		return collector, err
	//	}
	//
	//	collector = append(collector, apiResp.Orders...)
	//
	//	if !shoplineResp.HasNext() {
	//		break
	//	}
	//
	//	shopLineReq.Query = shoplineResp.Pagination.Next
	//}
	//
	//return collector, nil
}

func (o *OrderService) ListWithPagination(ctx context.Context, apiReq *ListOrdersAPIReq) (*ListOrdersAPIResp, error) {
	return o.List(ctx, apiReq)
}

func (o *OrderService) Count(ctx context.Context, req *GetOrdersCountAPIReq) (*GetOrdersCountAPIResp, error) {
	// 1. API response resource
	apiResp := &GetOrdersCountAPIResp{}

	// 2. Call the API
	err := o.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (o *OrderService) Get(ctx context.Context, apiReq *GetOrderDetailAPIReq) (*GetOrderDetailAPIResp, error) {
	// unsupported
	return nil, errors.New("Get method not implemented")
}

func (o *OrderService) Create(ctx context.Context, req *CreateOrderAPIReq) (*CreateOrderAPIResp, error) {
	// 1. API response resource
	apiResp := &CreateOrderAPIResp{}

	// 2. Call the API
	err := o.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (o *OrderService) Update(ctx context.Context, req *UpdateOrderAPIReq) (*UpdateOrderAPIResp, error) {
	// 1. API response resource
	apiResp := &UpdateOrderAPIResp{}

	// 2. Call the API
	err := o.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (o *OrderService) Delete(ctx context.Context, req *DeleteOrderAPIReq) (*DeleteOrderAPIResp, error) {
	// 1. API response resource
	apiResp := &DeleteOrderAPIResp{}

	// 2. Call the API
	err := o.Client.Call(ctx, req, apiResp)
	return apiResp, err
}
