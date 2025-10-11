package payment

import (
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListStoreBalanceTransactionsAPIReq
// 中文:https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-payments/balance-transactions?version=v20251201
// En:https://developer.shopline.com/docs/admin-rest-api/shopline-payments/balance-transactions?version=v20251201
type ListStoreBalanceTransactionsAPIReq struct {
	PageInfo string `url:"page_info,omitempty"`
	SinceId  string `url:"since_id,omitempty"`

	Limit    string `url:"limit,omitempty"`
	PayoutId string `url:"payout_id,omitempty"`

	PayoutStatus string `url:"payout_status,omitempty"`

	Status string `url:"status,omitempty"`
}

func (req *ListStoreBalanceTransactionsAPIReq) Verify() error {
	// Verify the api request params
	//if req.Limit == "" {
	//	return errors.New("limit can't be empty")
	//}
	return nil
}

func (req *ListStoreBalanceTransactionsAPIReq) Endpoint() string {
	return "payments/store/balance_transactions.json"
}

type ListStoreBalanceTransactionsAPIResp struct {
	client.BaseAPIResponse
	BalanceTransactions []BalanceTransaction `json:"transactions,omitempty"`
}

type BalanceTransaction struct {
	ID                       string `json:"id"`
	Amount                   string `json:"amount,omitempty"`
	Currency                 string `json:"currency,omitempty"`
	SourceOrderTransactionId string `json:"source_order_transaction_id,omitempty"`
	Fee                      string `json:"fee,omitempty"`
	Net                      string `json:"net,omitempty"`
	SettlementAmount         string `json:"settlement_amount,omitempty"`
	ProcessedAt              string `json:"processed_at,omitempty"`
	Type                     string `json:"type,omitempty"`
	SourcePaymentId          string `json:"source_payment_id,omitempty"`
	ExchangeRate             string `json:"exchange_rate,omitempty"`
	PayoutId                 string `json:"payout_id,omitempty"`
	SettlementCurrency       string `json:"settlement_currency,omitempty"`
	SourceType               string `json:"source_type,omitempty"`
	TransactionTime          string `json:"transaction_time,omitempty"`
	PayoutStatus             string `json:"payout_status,omitempty"`
	SourceOrderId            string `json:"source_order_id,omitempty"`
}
