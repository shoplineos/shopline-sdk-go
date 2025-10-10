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
	UID string // required
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

type StoreStaff struct {
	AccountOwner bool   `json:"account_owner,omitempty"`
	Email        string `json:"email,omitempty"`
	Name         string `json:"name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	UID          string `json:"uid,omitempty"`
	UserType     string `json:"user_type,omitempty"`
}
