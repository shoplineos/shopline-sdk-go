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

func (p *ProductService) List(ctx context.Context, apiReq *ListProductsAPIReq) (*ListProductsAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &ListProductsAPIResp{}

	// 4. Call API
	_, err := p.Client.Get(context.Background(), endpoint, shopLineReq, apiResp)
	if err != nil {
		return nil, err
	}

	//apiResp.Pagination = shopLineResp.Pagination

	return apiResp, nil
}

func (p *ProductService) ListAll(ctx context.Context, apiReq *ListProductsAPIReq) ([]ProductRespData, error) {
	collector := []ProductRespData{}
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Query: apiReq, // API request params
	}

	for {
		// 2. API endpoint
		endpoint := apiReq.Endpoint()

		// 3. API response data
		apiResp := &ListProductsAPIResp{}

		// 4. Call API
		shoplineResp, err := p.Client.Get(context.Background(), endpoint, shopLineReq, apiResp)

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
	// 1. API request
	shoplineReq := &client.ShopLineRequest{
		Query: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &GetProductCountAPIResp{}

	// 4. Call API
	_, err := p.Client.Get(context.Background(), endpoint, shoplineReq, apiResp)

	return apiResp, err
}
func (p *ProductService) Get(ctx context.Context, apiReq *GetProductDetailAPIReq) (*GetProductDetailAPIResp, error) {
	// 1. API request
	shoplineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &GetProductDetailAPIResp{}

	// 4. Call API
	_, err := p.Client.Get(context.Background(), endpoint, shoplineReq, apiResp)

	return apiResp, err
}

func (p *ProductService) Create(ctx context.Context, apiReq *CreateProductAPIReq) (*CreateProductAPIResp, error) {
	// 1. API request
	request := &client.ShopLineRequest{ // client request
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &CreateProductAPIResp{}

	// 4. Call API
	_, err := p.Client.Post(context.Background(), endpoint, request, apiResp)
	//if err != nil {
	//	log.Printf("CreateProduct request failed: %v\n", err)
	//	return nil, err
	//}

	return apiResp, err
}

func (p *ProductService) Update(ctx context.Context, apiReq *UpdateProductAPIReq) (*UpdateProductAPIResp, error) {
	// 1. API request
	request := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &UpdateProductAPIResp{}

	// 4. Call API
	_, err := p.Client.Put(context.Background(), endpoint, request, apiResp)

	//if err != nil {
	//	log.Printf("Update product failed，shopLineResp: %v, err: %v\n", shopLineResp, err)
	//	return nil, err
	//}

	return apiResp, err
}

func (p *ProductService) Delete(ctx context.Context, apiReq *DeleteProductAPIReq) (*DeleteProductAPIResp, error) {
	// 1. API request
	shoplineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response data
	apiResp := &DeleteProductAPIResp{}

	// 4. Call API
	_, err := p.Client.Delete(context.Background(), endpoint, shoplineReq, apiResp)

	return apiResp, err
}
