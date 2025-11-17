package store

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListStoreStaffsAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/get-all-store-staff-members?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/store/get-all-store-staff-members?version=v20251201
type ListStoreStaffsAPIReq struct {
	client.BaseAPIRequest
	Limit string `url:"limit,omitempty"` // required
}

func (req *ListStoreStaffsAPIReq) GetMethod() string {
	return "GET"
}

func (req *ListStoreStaffsAPIReq) GetQuery() interface{} {
	return req
}

func (req *ListStoreStaffsAPIReq) Verify() error {
	// Verify the api request params
	if req.Limit == "" {
		return errors.New("limit is required")
	}
	return nil
}

func (req *ListStoreStaffsAPIReq) GetEndpoint() string {
	return "store/list/staff.json"
}

type ListStoreStaffsAPIResp struct {
	client.BaseAPIResponse
	StoreStaffs []StoreStaff `json:"data,omitempty"`
}
