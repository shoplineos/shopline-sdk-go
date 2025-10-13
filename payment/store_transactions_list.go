package payment

import (
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListStoreTransactionsAPIReq
// 中文:https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-payments/query-store-transaction-records?version=v20251201
// En:https://developer.shopline.com/docs/admin-rest-api/shopline-payments/query-store-transaction-records?version=v20251201
type ListStoreTransactionsAPIReq struct {
	PageInfo string `url:"page_info,omitempty"`
	SinceId  string `url:"since_id,omitempty"`

	Limit string `url:"limit,omitempty"`

	// Example: 2024-12-10T00:00:00+08:00
	DateMin string `url:"date_min,omitempty"`
	DateMax string `url:"date_max,omitempty"`

	Status string `url:"status,omitempty"`

	TradeOrderId    string `url:"trade_order_id,omitempty"`
	TransactionType string `url:"transaction_type,omitempty"`
}

func (req *ListStoreTransactionsAPIReq) Verify() error {
	// Verify the api request params
	//if req.Limit == "" {
	//	return errors.New("limit can't be empty")
	//}
	return nil
}

func (req *ListStoreTransactionsAPIReq) Endpoint() string {
	return "payments/store/transactions.json"
}

type ListStoreTransactionsAPIResp struct {
	client.BaseAPIResponse
	Transactions []Transaction `json:"transactions,omitempty"`
}
