package payment

type Payout struct {
	Amount              string `json:"amount,omitempty"`
	Currency            string `json:"currency,omitempty"`
	PayoutTransactionNo string `json:"payout_transaction_no,omitempty"`
	Status              string `json:"status,omitempty"`
	Time                string `json:"time,omitempty"`
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

type BalanceTransaction struct {
	Id                       string `json:"id"`
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
