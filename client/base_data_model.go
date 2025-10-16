package client

type APIRequest interface {
	Endpoint() string      // API Endpoint
	Verify() error         // Verify API request params
	Method() string        // http method
	GetQuery() interface{} // your own struct or an APIRequest, for http url query params
	GetData() interface{}  // your own struct or an APIRequest, for http body params
	GetRequestOptions() *RequestOptions
}

type APIResponse interface {
	SetTraceId(traceId string)
	SetPagination(pagination *Pagination)
}

type BaseAPIRequest struct {
}

func (b BaseAPIRequest) GetQuery() interface{} {
	return b
}

func (b BaseAPIRequest) GetData() interface{} {
	return b
}

func (b BaseAPIRequest) GetRequestOptions() *RequestOptions {
	return nil
}

func (b BaseAPIRequest) Endpoint() string {
	return ""
}

func (b BaseAPIRequest) Verify() error {
	return nil
}

func (b BaseAPIRequest) Method() string {
	return ""
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
