package order

import (
	"context"
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type IOrderService interface {
	client.Aware
	List(context.Context, *QueryOrdersAPIReq) (*QueryOrdersAPIResp, error)
	ListAll(context.Context, *QueryOrdersAPIReq) ([]Order, error)
	ListWithPagination(context.Context, *QueryOrdersAPIReq) (*QueryOrdersAPIResp, error)
	Count(context.Context, *GetOrdersCountAPIReq) (*GetOrdersCountAPIResp, error)
	Get(context.Context, *GetOrderDetailAPIReq) (*GetOrderDetailAPIResp, error)
	Create(context.Context, *CreateOrderAPIReq) (*CreateOrderAPIResp, error)
	Update(context.Context, *UpdateOrderAPIReq) (*UpdateOrderAPIResp, error)
	Delete(context.Context, *DeleteOrderAPIReq) (*DeleteOrderAPIResp, error)
	Cancel(context.Context, *CancelOrderAPIReq) (*CancelOrderAPIResp, error)
	Refund(context.Context, *RefundAPIReq) (*RefundAPIResp, error)
}

var serviceInst = &OrderService{}

type OrderService struct {
	client.BaseService
}

func GetOrderService() *OrderService {
	return serviceInst
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
	_, err := o.Client.Post(context.Background(), req.Endpoint(), shopLineReq, apiResp)

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
	_, err := o.Client.Post(context.Background(), apiReq.Endpoint(), shopLineReq, apiResp)
	return apiResp, err
}

func (o *OrderService) List(ctx context.Context, apiReq *QueryOrdersAPIReq) (*QueryOrdersAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &QueryOrdersAPIResp{}

	// 4. Call API
	shopLineResp, err := o.Client.Get(context.Background(), endpoint, shopLineReq, apiResp)
	if err != nil {
		return nil, err
	}

	apiResp.Pagination = shopLineResp.Pagination

	return apiResp, nil
}

func (o *OrderService) ListAll(ctx context.Context, apiReq *QueryOrdersAPIReq) ([]Order, error) {
	collector := []Order{}
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	for {
		// 2. API endpoint
		endpoint := apiReq.Endpoint()

		// 3. API response data
		apiResp := &QueryOrdersAPIResp{}

		// 4. Call API
		shoplineResp, err := o.Client.Get(context.Background(), endpoint, shopLineReq, apiResp)

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

func (o *OrderService) ListWithPagination(ctx context.Context, apiReq *QueryOrdersAPIReq) (*QueryOrdersAPIResp, error) {
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
	_, err := o.Client.Get(context.Background(), endpoint, shoplineReq, apiResp)

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
	_, err := o.Client.Post(context.Background(), endpoint, request, apiResp)

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
	_, err := o.Client.Put(context.Background(), endpoint, request, apiResp)

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
	_, err := o.Client.Delete(context.Background(), endpoint, shoplineReq, apiResp)

	return apiResp, err
}
