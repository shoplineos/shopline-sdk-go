package client

import (
	"context"
)

// GetResources
// eg:
//
//	func getResources(resp interface{}) []Order {
//		apiResp := resp.(*ListOrdersAPIResp)
//		return apiResp.Orders
//	}
type GetResources[T any] func(resp interface{}) []T

// ListAll List all resources
func ListAll[T any](cli *Client, ctx context.Context, req APIRequest, resp interface{}, getResources GetResources[T]) ([]T, error) {
	collector := []T{}
	// 1. API request
	shopLineReq := &ShopLineRequest{
		Query: req, // API request params
	}

	for {
		// 2. API endpoint
		endpoint := req.GetEndpoint()

		// 3. API response resource
		apiResp := resp

		// 4. Call the API
		shoplineResp, err := cli.Get(ctx, endpoint, shopLineReq, apiResp)

		if err != nil {
			return collector, err
		}

		resources := getResources(apiResp)
		collector = append(collector, resources...)

		if !shoplineResp.HasNext() {
			break
		}

		shopLineReq.Query = shoplineResp.Pagination.Next
	}

	return collector, nil
}
