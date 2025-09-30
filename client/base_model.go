package client

type APIRequest interface {
	Verify() error    // Verify API request params
	Endpoint() string // API Endpoint
}

type APIResponse interface {
	SetTraceId(traceId string)
	SetPagination(pagination Pagination)
}

type Aware interface {
	SetClient(*Client)
}
