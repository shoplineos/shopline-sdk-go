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

type Transaction struct {
	Exchange Exchange `json:"exchange,omitempty"`

	FeeType             string              `json:"fee_type,omitempty"`
	PaymentMethod       string              `json:"payment_method,omitempty"`
	Status              string              `json:"status,omitempty"`
	TradeOrderId        string              `json:"trade_order_id,omitempty"`
	CreateTime          string              `json:"create_time,omitempty"`
	DisputeType         string              `json:"dispute_type,omitempty"`
	PaymentMethodOption PaymentMethodOption `json:"payment_method_option,omitempty"`
	PaymentMsg          PaymentMsg          `json:"payment_msg,omitempty"`

	SellerOrderId    string `json:"seller_order_id,omitempty"`
	SubPaymentMethod string `json:"sub_payment_method,omitempty"`
	UpdateTime       string `json:"update_time,omitempty"`
	Amount           string `json:"amount,omitempty"`
	ChannelDealId    string `json:"channel_deal_id,omitempty"`

	CreditCard      CreditCard     `json:"credit_card,omitempty"`
	Reason          string         `json:"reason,omitempty"`
	TransactionType string         `json:"transaction_type,omitempty"`
	AdditionalData  AdditionalData `json:"additional_data,omitempty"`
	Customer        Customer       `json:"customer,omitempty"`

	Fee        string `json:"fee,omitempty"`
	MerchantId string `json:"merchant_id,omitempty"`
	PaidAmount string `json:"paid_amount,omitempty"`
	SubStatus  string `json:"sub_status,omitempty"`
	Currency   string `json:"currency,omitempty"`
}

type Exchange struct {
	Amount string `json:"amount,omitempty"`

	// Example: USD/USD @1.0000000000
	Rate     string `json:"rate,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type PaymentMethodOption struct {
	Installment Installment `json:"installment,omitempty"`
}

type Installment struct {
	Count string `json:"count,omitempty"`
}

type PaymentMsg struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type CreditCard struct {
	Brand         string `json:"brand,omitempty"`
	IssuerCountry string `json:"issuer_country,omitempty"`
	Last4         string `json:"last4,omitempty"`
	AuthCode      string `json:"auth_code,omitempty"`
	Bin           string `json:"bin,omitempty"`
}

type AdditionalData struct {
	DisputeEvidenceUpdateDeadline string `json:"dispute_evidence_update_deadline,omitempty"`
	IsSettled                     bool   `json:"is_settled,omitempty"`
	ReserveHeld                   string `json:"reserve_held,omitempty"`
	ReserveReleaseRime            string `json:"reserve_release_time,omitempty"`
	SettleTime                    string `json:"settle_time,omitempty"`
	StatementTime                 string `json:"statement_time,omitempty"`
}

type Customer struct {
	PersonalInfo PersonalInfo `json:"personal_info,omitempty"`
}

type PersonalInfo struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
