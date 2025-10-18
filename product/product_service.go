package product

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// IProductService
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/create-a-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201
type IProductService interface {
	List(context.Context, *ListProductsAPIReq) (*ListProductsAPIResp, error)
	ListAll(context.Context, *ListProductsAPIReq) ([]ProductRespData, error)
	ListWithPagination(context.Context, *ListProductsAPIReq) (*ListProductsAPIResp, error)
	Count(context.Context, *GetProductCountAPIReq) (*GetProductCountAPIResp, error)
	Get(context.Context, *GetProductDetailAPIReq) (*GetProductDetailAPIResp, error)
	Create(context.Context, *CreateProductAPIReq) (*CreateProductAPIResp, error)
	Update(context.Context, *UpdateProductAPIReq) (*UpdateProductAPIResp, error)
	Delete(context.Context, *DeleteProductAPIReq) (*DeleteProductAPIResp, error)
}

var productServiceInst = &ProductService{}

func GetProductService() *ProductService {
	return productServiceInst
}

type ProductService struct {
	client.BaseService
}

func (p *ProductService) List(ctx context.Context, req *ListProductsAPIReq) (*ListProductsAPIResp, error) {
	// 1. API response data
	apiResp := &ListProductsAPIResp{}

	// 2. Call the API
	err := p.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (p *ProductService) ListAll(ctx context.Context, apiReq *ListProductsAPIReq) ([]ProductRespData, error) {
	collector := []ProductRespData{}
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	for {
		// 2. API endpoint
		endpoint := apiReq.GetEndpoint()

		// 3. API response data
		apiResp := &ListProductsAPIResp{}

		// 4. Call the API
		shoplineResp, err := p.Client.Get(ctx, endpoint, shopLineReq, apiResp)

		if err != nil {
			return collector, err
		}

		collector = append(collector, apiResp.Products...)

		if !shoplineResp.HasNext() {
			break
		}

		shopLineReq.Query = shoplineResp.Pagination.Next
	}

	return collector, nil
}

func (p *ProductService) ListWithPagination(ctx context.Context, apiReq *ListProductsAPIReq) (*ListProductsAPIResp, error) {
	return p.List(ctx, apiReq)
}

func (p *ProductService) Count(ctx context.Context, apiReq *GetProductCountAPIReq) (*GetProductCountAPIResp, error) {

	// 1. API response data
	apiResp := &GetProductCountAPIResp{}

	// 2. Call the API
	err := p.Client.Call(context.Background(), apiReq, apiResp)

	return apiResp, err
}
func (p *ProductService) Get(ctx context.Context, req *GetProductDetailAPIReq) (*GetProductDetailAPIResp, error) {
	// 1. API response data
	apiResp := &GetProductDetailAPIResp{}

	// 2. Call the API
	err := p.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (p *ProductService) Create(ctx context.Context, req *CreateProductAPIReq) (*CreateProductAPIResp, error) {
	// 1. API response data
	apiResp := &CreateProductAPIResp{}

	// 2. Call the API
	err := p.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (p *ProductService) Update(ctx context.Context, req *UpdateProductAPIReq) (*UpdateProductAPIResp, error) {
	// 1. API response data
	apiResp := &UpdateProductAPIResp{}

	// 2. Call the API
	err := p.Client.Call(ctx, req, apiResp)
	return apiResp, err
}

func (p *ProductService) Delete(ctx context.Context, req *DeleteProductAPIReq) (*DeleteProductAPIResp, error) {
	// 1. API response data
	apiResp := &DeleteProductAPIResp{}

	// 2. Call the API
	err := p.Client.Call(ctx, req, apiResp)
	return apiResp, err
}
