package order

import (
	"context"
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
	cli *client.Client
}

func GetOrderService() *OrderService {
	return serviceInst
}

func (p *OrderService) SetClient(c *client.Client) {
	p.cli = c
}

func (p *OrderService) Refund(ctx context.Context, req *RefundAPIReq) (*RefundAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: req, // API request params
	}

	// 2. API response data
	apiResp := &RefundAPIResp{}

	// 3. Call API
	_, err := p.cli.Get(context.Background(), req.Endpoint(), shopLineReq, apiResp)

	return apiResp, err
}

func (p *OrderService) Cancel(ctx context.Context, apiReq *CancelOrderAPIReq) (*CancelOrderAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request params
	}

	// 2. API response data
	apiResp := &CancelOrderAPIResp{}

	// 3. Call API
	_, err := p.cli.Get(context.Background(), apiReq.Endpoint(), shopLineReq, apiResp)
	return apiResp, err
}

func (p *OrderService) List(ctx context.Context, apiReq *QueryOrdersAPIReq) (*QueryOrdersAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &QueryOrdersAPIResp{}

	// 4. Call API
	shopLineResp, err := p.cli.Get(context.Background(), endpoint, shopLineReq, apiResp)
	if err != nil {
		return nil, err
	}

	apiResp.Pagination = shopLineResp.Pagination

	return apiResp, nil
}

func (p *OrderService) ListAll(ctx context.Context, apiReq *QueryOrdersAPIReq) ([]Order, error) {
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
		shoplineResp, err := p.cli.Get(context.Background(), endpoint, shopLineReq, apiResp)

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

func (p *OrderService) ListWithPagination(ctx context.Context, apiReq *QueryOrdersAPIReq) (*QueryOrdersAPIResp, error) {
	return p.List(ctx, apiReq)
}

func (p *OrderService) Count(ctx context.Context, apiReq *GetOrdersCountAPIReq) (*GetOrdersCountAPIResp, error) {
	// 1. API request
	shoplineReq := &client.ShopLineRequest{
		Query: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &GetOrdersCountAPIResp{}

	// 4. Call API
	_, err := p.cli.Get(context.Background(), endpoint, shoplineReq, apiResp)

	return apiResp, err
}

func (p *OrderService) Get(ctx context.Context, apiReq *GetOrderDetailAPIReq) (*GetOrderDetailAPIResp, error) {
	// unsupported
	return nil, nil
}

func (p *OrderService) Create(ctx context.Context, apiReq *CreateOrderAPIReq) (*CreateOrderAPIResp, error) {
	// 1. API request
	request := &client.ShopLineRequest{ // client request
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &CreateOrderAPIResp{}

	// 4. Call API
	_, err := p.cli.Post(context.Background(), endpoint, request, apiResp)

	return apiResp, err
}

func (p *OrderService) Update(ctx context.Context, apiReq *UpdateOrderAPIReq) (*UpdateOrderAPIResp, error) {
	// 1. API request
	request := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &UpdateOrderAPIResp{}

	// 4. Call API
	_, err := p.cli.Put(context.Background(), endpoint, request, apiResp)

	return apiResp, err
}

func (p *OrderService) Delete(ctx context.Context, apiReq *DeleteOrderAPIReq) (*DeleteOrderAPIResp, error) {
	// 1. API request
	shoplineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &DeleteOrderAPIResp{}

	// 4. Call API
	_, err := p.cli.Delete(context.Background(), endpoint, shoplineReq, apiResp)

	return apiResp, err
}
