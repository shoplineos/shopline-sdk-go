package payment

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetStorePayoutAPIReq
// 中文:https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-payments/payout?version=v20251201
// En:https://developer.shopline.com/docs/admin-rest-api/shopline-payments/payout?version=v20251201
type GetStorePayoutAPIReq struct {
	client.BaseAPIRequest
	PayoutId string `url:"payout_id,omitempty"` // required
}

func (req *GetStorePayoutAPIReq) Method() string {
	return "GET"
}

func (req *GetStorePayoutAPIReq) GetQuery() interface{} {
	return req
}

func (req *GetStorePayoutAPIReq) Verify() error {
	// Verify the api request params
	if req.PayoutId == "" {
		return errors.New("PayoutId can't be empty")
	}
	return nil
}

func (req *GetStorePayoutAPIReq) Endpoint() string {
	return "payments/store/payout.json"
}

type GetStorePayoutAPIResp struct {
	client.BaseAPIResponse
	Payouts []Payout `json:"payouts,omitempty"`
}
