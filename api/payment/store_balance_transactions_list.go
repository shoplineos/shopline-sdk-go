package payment

import (
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListStoreBalanceTransactionsAPIReq
// 中文:https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-payments/balance-transactions?version=v20251201
// En:https://developer.shopline.com/docs/admin-rest-api/shopline-payments/balance-transactions?version=v20251201
type ListStoreBalanceTransactionsAPIReq struct {
	client.BaseAPIRequest
	PageInfo string `url:"page_info,omitempty"`
	SinceId  string `url:"since_id,omitempty"`

	Limit    string `url:"limit,omitempty"`
	PayoutId string `url:"payout_id,omitempty"`

	PayoutStatus string `url:"payout_status,omitempty"`

	Status string `url:"status,omitempty"`
}

func (req *ListStoreBalanceTransactionsAPIReq) GetMethod() string {
	return "GET"
}

func (req *ListStoreBalanceTransactionsAPIReq) GetQuery() interface{} {
	return req
}

func (req *ListStoreBalanceTransactionsAPIReq) Verify() error {
	// Verify the api request params
	//if req.Limit == "" {
	//	return errors.New("limit can't be empty")
	//}
	return nil
}

func (req *ListStoreBalanceTransactionsAPIReq) GetEndpoint() string {
	return "payments/store/balance_transactions.json"
}

type ListStoreBalanceTransactionsAPIResp struct {
	client.BaseAPIResponse
	BalanceTransactions []BalanceTransaction `json:"transactions,omitempty"`
}
