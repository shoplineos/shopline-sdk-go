package store

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetStoreOperationLogAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/get-a-store-operation-log?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/store/get-a-store-operation-log?version=v20251201
type GetStoreOperationLogAPIReq struct {
	ID string // required
}

func (req *GetStoreOperationLogAPIReq) Verify() error {
	// Verify the api request params
	if req.ID == "" {
		return errors.New("id is required")
	}
	return nil
}

func (req *GetStoreOperationLogAPIReq) Endpoint() string {
	return fmt.Sprintf("store/operation_logs/%s.json", req.ID)
}

type GetStoreOperationLogAPIResp struct {
	client.BaseAPIResponse
	OperationLog OperationLog `json:"data,omitempty"`
}

type OperationLog struct {
	Author      string `json:"author,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	Detail      string `json:"detail,omitempty"`
	ID          string `json:"id,omitempty"`
	SubjectId   string `json:"subject_id,omitempty"`
	SubjectType string `json:"subject_type,omitempty"`
	Verb        string `json:"verb,omitempty"`
}
