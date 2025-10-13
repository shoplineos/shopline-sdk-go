package payment

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListStorePayoutsAPIReq
// 中文:https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-payments/get-payouts?version=v20251201
// En:https://developer.shopline.com/docs/admin-rest-api/shopline-payments/get-payouts?version=v20251201
type ListStorePayoutsAPIReq struct {
	PageInfo string `url:"page_info,omitempty"`
	SinceId  string `url:"since_id,omitempty"`

	Limit               string `url:"limit,omitempty"` // required
	PayoutTransactionNo string `url:"payout_transaction_no,omitempty"`

	// 2025-04-30T00:00:00+08:00
	StartTime string `url:"start_time,omitempty"` // required

	// 2025-04-30T00:00:00+08:00
	EndTime string `url:"end_time,omitempty"` // required
	Status  string `url:"status,omitempty"`
}

func (req *ListStorePayoutsAPIReq) Verify() error {
	// Verify the api request params
	if req.Limit == "" {
		return errors.New("limit can't be empty")
	}
	if req.StartTime == "" {
		return errors.New("StartTime can't be empty")
	}
	if req.EndTime == "" {
		return errors.New("EndTime can't be empty")
	}
	return nil
}

func (req *ListStorePayoutsAPIReq) Endpoint() string {
	return "payments/store/payouts.json"
}

type ListStorePayoutsAPIResp struct {
	client.BaseAPIResponse
	Payouts []Payout `json:"payouts,omitempty"`
}
