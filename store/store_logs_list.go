package store

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListStoreOperationLogsAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/get-store-operation-logs?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/store/get-store-operation-logs?version=v20251201
type ListStoreOperationLogsAPIReq struct {
	CreatedAtMin string `url:"created_at_min,omitempty"` // required
	CreatedAtMax string `url:"created_at_max,omitempty"` // required

	SinceId     string `url:"since_id,omitempty"`
	SubjectType string `url:"subject_type,omitempty"`
	Verb        string `url:"verb,omitempty"`
	Limit       string `url:"limit,omitempty"`
}

func (req *ListStoreOperationLogsAPIReq) Verify() error {
	// Verify the api request params
	if req.CreatedAtMin == "" {
		return errors.New("CreatedAtMin is required")
	}
	if req.CreatedAtMax == "" {
		return errors.New("CreatedAtMax is required")
	}
	return nil
}

func (req *ListStoreOperationLogsAPIReq) Endpoint() string {
	return "store/operation_logs.json"
}

type ListStoreOperationLogsAPIResp struct {
	client.BaseAPIResponse
	OperationLogs []OperationLog `json:"data,omitempty"`
}
