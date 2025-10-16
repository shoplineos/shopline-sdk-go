package store

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetStoreStaffAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/get-a-staff-member?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/store/get-a-staff-member?version=v20251201
type GetStoreStaffAPIReq struct {
	client.BaseAPIRequest
	UID string // required
}

func (req *GetStoreStaffAPIReq) Method() string {
	return "GET"
}

func (req *GetStoreStaffAPIReq) GetQuery() interface{} {
	return req
}

func (req *GetStoreStaffAPIReq) Verify() error {
	// Verify the api request params
	if req.UID == "" {
		return errors.New("store staff uid is required")
	}
	return nil
}

func (req *GetStoreStaffAPIReq) Endpoint() string {
	return fmt.Sprintf("store/staff/%s.json", req.UID)
}

type GetStoreStaffAPIResp struct {
	client.BaseAPIResponse
	StoreStaff StoreStaff `json:"data,omitempty"`
}
