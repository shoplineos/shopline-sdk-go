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
	Refund(context.Context, *RefundAPIReq) (*RefundAPIResp, error)
}

var serviceInst = &OrderService{}

func GetOrderService() *OrderService {
	return serviceInst
}

type OrderService struct {
	client.BaseService
}

// Refund
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/order-refund?version=v20251201
// en：https://developer.shopline.com/docs/admin-rest-api/order/order-management/order-refund?version=v20251201
func (o *OrderService) Refund(ctx context.Context, req *RefundAPIReq) (*RefundAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: req, // API request params
	}

	// 2. API response data
	apiResp := &RefundAPIResp{}

	// 3. Call API
	_, err := o.Client.Post(ctx, req.Endpoint(), shopLineReq, apiResp)

	return apiResp, err
}

func (o *OrderService) Cancel(ctx context.Context, apiReq *CancelOrderAPIReq) (*CancelOrderAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request params
	}

	// 2. API response data
	apiResp := &CancelOrderAPIResp{}

	// 3. Call API
	_, err := o.Client.Post(ctx, apiReq.Endpoint(), shopLineReq, apiResp)
	return apiResp, err
}

func (o *OrderService) List(ctx context.Context, apiReq *ListOrdersAPIReq) (*ListOrdersAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &ListOrdersAPIResp{}

	// 4. Call API
	_, err := o.Client.Get(ctx, endpoint, shopLineReq, apiResp)

	return apiResp, err
}

func (o *OrderService) ListAll(ctx context.Context, apiReq *ListOrdersAPIReq) ([]Order, error) {
	collector := []Order{}
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	for {
		// 2. API endpoint
		endpoint := apiReq.Endpoint()

		// 3. API response data
		apiResp := &ListOrdersAPIResp{}

		// 4. Call API
		shoplineResp, err := o.Client.Get(ctx, endpoint, shopLineReq, apiResp)

		if err != nil {
			return collector, err
		}

		collector = append(collector, apiResp.Orders...)

		if !shoplineResp.HasNext() {
			break
		}

		shopLineReq.Query = shoplineResp.Pagination.Next
	}

	return collector, nil
}

func (o *OrderService) ListWithPagination(ctx context.Context, apiReq *ListOrdersAPIReq) (*ListOrdersAPIResp, error) {
	return o.List(ctx, apiReq)
}

func (o *OrderService) Count(ctx context.Context, apiReq *GetOrdersCountAPIReq) (*GetOrdersCountAPIResp, error) {
	// 1. API request
	shoplineReq := &client.ShopLineRequest{
		Query: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &GetOrdersCountAPIResp{}

	// 4. Call API
	_, err := o.Client.Get(ctx, endpoint, shoplineReq, apiResp)

	return apiResp, err
}

func (o *OrderService) Get(ctx context.Context, apiReq *GetOrderDetailAPIReq) (*GetOrderDetailAPIResp, error) {
	// unsupported
	return nil, errors.New("Get method not implemented")
}

func (o *OrderService) Create(ctx context.Context, apiReq *CreateOrderAPIReq) (*CreateOrderAPIResp, error) {
	// 1. API request
	request := &client.ShopLineRequest{ // client request
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &CreateOrderAPIResp{}

	// 4. Call API
	_, err := o.Client.Post(ctx, endpoint, request, apiResp)

	return apiResp, err
}

func (o *OrderService) Update(ctx context.Context, apiReq *UpdateOrderAPIReq) (*UpdateOrderAPIResp, error) {
	// 1. API request
	request := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &UpdateOrderAPIResp{}

	// 4. Call API
	_, err := o.Client.Put(ctx, endpoint, request, apiResp)

	return apiResp, err
}

func (o *OrderService) Delete(ctx context.Context, apiReq *DeleteOrderAPIReq) (*DeleteOrderAPIResp, error) {
	// 1. API request
	shoplineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &DeleteOrderAPIResp{}

	// 4. Call API
	_, err := o.Client.Delete(ctx, endpoint, shoplineReq, apiResp)

	return apiResp, err
}
