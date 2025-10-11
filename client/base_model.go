package client

type APIRequest interface {
	Endpoint() string // API Endpoint
	Verify() error    // Verify API request params
}

type APIResponse interface {
	SetTraceId(traceId string)
	SetPagination(pagination *Pagination)
}

type BaseAPIResponse struct {
	TraceId    string
	Pagination *Pagination
}

func (api BaseAPIResponse) SetTraceId(traceId string) {
	api.TraceId = traceId
}

func (api BaseAPIResponse) SetPagination(pagination *Pagination) {
	api.Pagination = pagination
}

type Aware interface {
	SetClient(*Client)
}

type BaseService struct {
	Client *Client
}

func (b *BaseService) SetClient(client *Client) {
	b.Client = client
}
