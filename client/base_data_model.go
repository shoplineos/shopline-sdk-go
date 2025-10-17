package client

type APIRequest interface {
	GetEndpoint() string                // API Endpoint (required)
	GetMethod() string                  // Http method  (required)
	GetQuery() interface{}              // Your own struct or an APIRequest, for http url query parameters
	GetData() interface{}               // Your own struct or an APIRequest, for http body parameters
	Verify() error                      // Verify API request parameters
	GetHeaders() map[string]string      // Http headers
	GetRequestOptions() *RequestOptions // Request options
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

//func (req BaseAPIRequest) GetEndpoint() string {
//	return ""
//}
//
//func (req BaseAPIRequest) GetMethod() string {
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
