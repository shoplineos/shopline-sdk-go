package store

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CountStoreOperationLogsAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/get-store-operation-log-count?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/store/get-store-operation-log-count?version=v20251201
type CountStoreOperationLogsAPIReq struct {
	client.BaseAPIRequest
	CreatedAtMin string `url:"created_at_min,omitempty"` // required
	CreatedAtMax string `url:"created_at_max,omitempty"` // required

	SubjectType string `url:"subject_type,omitempty"`
	Verb        string `url:"verb,omitempty"`
}

func (req *CountStoreOperationLogsAPIReq) Method() string {
	return "GET"
}

func (req *CountStoreOperationLogsAPIReq) GetQuery() interface{} {
	return req
}

func (req *CountStoreOperationLogsAPIReq) Verify() error {
	if req.CreatedAtMin == "" {
		return errors.New("CreatedAtMin is required")
	}
	if req.CreatedAtMax == "" {
		return errors.New("CreatedAtMax is required")
	}
	return nil
}

func (req *CountStoreOperationLogsAPIReq) Endpoint() string {
	return "store/operation_logs/count.json"
}

type CountStoreOperationLogsAPIResp struct {
	client.BaseAPIResponse
	CountOperationLogData CountOperationLogData `json:"data,omitempty"`
}

type CountOperationLogData struct {
	Count int `json:"count,omitempty"`
}
