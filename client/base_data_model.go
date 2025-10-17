package client

type APIRequest interface {
	Endpoint() string      // API Endpoint (required)
	Method() string        // http method  (required)
	GetQuery() interface{} // your own struct or an APIRequest, for http url query params
	GetData() interface{}  // your own struct or an APIRequest, for http body params
	Verify() error         // Verify API request params
	GetHeaders() map[string]string
	GetRequestOptions() *RequestOptions
}

type APIResponse interface {
	SetTraceId(traceId string)
	SetPagination(pagination *Pagination)
}

type BaseAPIRequest struct {
}

func (req BaseAPIRequest) GetHeaders() map[string]string {
	return nil
}

func (req BaseAPIRequest) GetQuery() interface{} {
	return nil
}

func (req BaseAPIRequest) GetData() interface{} {
	return nil
}

//func (req BaseAPIRequest) Endpoint() string {
//	return ""
//}
//
//func (req BaseAPIRequest) Method() string {
//	return ""
//}

func (req BaseAPIRequest) Verify() error {
	return nil
}

func (req BaseAPIRequest) GetRequestOptions() *RequestOptions {
	return nil
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
