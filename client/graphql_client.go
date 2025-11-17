package client

import (
	"context"
)

// IGraphQLClient is an interface to interact with the graphql endpoint
// of the SHOPLINE API
// See https://developer.shopline.com/zh-hans-cn/docs/storefront-api/schema-documentation/?version=v20251201
// 中文：https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/storefront-api/overview?version=v20251201
// En：https://developer.shopline.com/docs/apps/api-instructions-for-use/storefront-api/overview?version=v20251201
type IGraphQLClient interface {
	Query(ctx context.Context, query string, vars interface{}, resp interface{}) error
	Mutation(ctx context.Context, mutation string, vars interface{}, resp interface{}) error
}

// GraphQLClient handles communication with the graphql endpoint of
// the SHOPLINE API.
type GraphQLClient struct {
	client *Client
}

func NewStorefrontClient(client *Client) *GraphQLClient {
	cli := &GraphQLClient{
		client: client,
	}
	client.PathPrefix = "storefront/graph"
	return cli
}

func NewAdminClient(client *Client) *GraphQLClient {
	cli := &GraphQLClient{
		client: client,
	}
	client.PathPrefix = "admin/graph"
	return cli
}

type graphqlMutationRequest struct {
	Mutation  string      `json:"mutation"`
	Variables interface{} `json:"variables"`
}

type graphqlQueryRequest struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

type graphQLResponse struct {
	Data       interface{}        `json:"data"`
	Errors     []graphQLError     `json:"errors"`
	Extensions *graphQLExtensions `json:"extensions"`
}

type graphQLExtensions struct {
	Cost GraphQLCost `json:"cost"`
}

// GraphQLCost represents the cost of the graphql query
type GraphQLCost struct {
	RequestedQueryCost int                   `json:"requestedQueryCost"`
	ActualQueryCost    *int                  `json:"actualQueryCost"`
	ThrottleStatus     GraphQLThrottleStatus `json:"throttleStatus"`
}

// GraphQLThrottleStatus represents the status of the shop's rate limit points
type GraphQLThrottleStatus struct {
	MaximumAvailable   float64 `json:"maximumAvailable"`
	CurrentlyAvailable float64 `json:"currentlyAvailable"`
	RestoreRate        float64 `json:"restoreRate"`
}

type graphQLError struct {
	Message    string                  `json:"message"`
	Extensions *graphQLErrorExtensions `json:"extensions"`
	Locations  []graphQLErrorLocation  `json:"locations"`
}

type graphQLErrorExtensions struct {
	Code          string
	Documentation string
}

type graphQLErrorLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// Query creates a graphql query against the SHOPLINE API
// the "data" portion of the response is unmarshalled into resp
// example:
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-graph-ql-api/apps/queries/publication?version=v20251201
// En：https://developer.shopline.com/docs/admin-graph-ql-api/apps/queries/publication?version=v20251201
func (cli *GraphQLClient) Query(ctx context.Context, q string, vars, resp interface{}) error {
	data := &graphqlQueryRequest{
		Query:     q,
		Variables: vars,
	}

	return cli.execute(ctx, data, resp)
}

func (cli *GraphQLClient) execute(ctx context.Context, data interface{}, resp interface{}) error {
	shoplineReq := &ShopLineRequest{
		Data: data,
	}

	gr := graphQLResponse{
		Data: resp,
	}

	_, err := cli.client.Post(ctx, "graphql.json", shoplineReq, &gr)

	var retryAfterSecs float64

	if gr.Extensions != nil {
		retryAfterSecs = gr.Extensions.Cost.RetryAfterSeconds()
		cli.client.RateLimits.GraphQLCost = &gr.Extensions.Cost
		cli.client.RateLimits.RetryAfterSeconds = retryAfterSecs
	}

	if len(gr.Errors) > 0 {
		responseError := ResponseError{Status: 200}

		for _, err := range gr.Errors {
			responseError.Errors = append(responseError.Errors, err.Message)
		}

		err = responseError
	}

	return err
}

// RetryAfterSeconds returns the estimated retry after seconds based on
// the requested query cost and throttle status
func (c GraphQLCost) RetryAfterSeconds() float64 {
	var diff float64

	if c.ActualQueryCost != nil {
		diff = c.ThrottleStatus.CurrentlyAvailable - float64(*c.ActualQueryCost)
	} else {
		diff = c.ThrottleStatus.CurrentlyAvailable - float64(c.RequestedQueryCost)
	}

	if diff < 0 {
		return -diff / c.ThrottleStatus.RestoreRate
	}

	return 0
}

// Mutation creates a graphql mutation against the SHOPLINE API
// the "data" portion of the response is unmarshalled into resp
// example:
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-graph-ql-api/apps/mutations/publication-update?version=v20251201
// En：https://developer.shopline.com/docs/admin-graph-ql-api/apps/mutations/publication-update?version=v20251201
func (cli *GraphQLClient) Mutation(ctx context.Context, mutation string, vars interface{}, resp interface{}) error {
	data := &graphqlMutationRequest{
		Mutation:  mutation,
		Variables: vars,
	}

	return cli.execute(ctx, data, resp)
}
