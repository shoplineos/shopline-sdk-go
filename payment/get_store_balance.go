package payment

import (
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetStoreBalanceAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-payments/balance?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-payments/balance?version=v20251201
type GetStoreBalanceAPIReq struct {
}

func (req *GetStoreBalanceAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *GetStoreBalanceAPIReq) Endpoint() string {
	return "payments/store/balance.json"
}

type GetStoreBalanceAPIResp struct {
	client.BaseAPIResponse
	Balance Balance `json:"balance,omitempty"`
}

type Balance struct {
	Amount   string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}
