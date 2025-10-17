package order

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// RefundAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/order-refund?version=v20251201
// en：https://developer.shopline.com/docs/admin-rest-api/order/order-management/order-refund?version=v20251201
type RefundAPIReq struct {
	client.BaseAPIRequest

	OrderId         string          `json:"order_id,omitempty"`
	Notify          string          `json:"notify,omitempty"`
	ProcessedAt     string          `json:"processed_at,omitempty"` // 2023-04-12T22:59:52+08:00
	RefundLineItems []RefundLine    `json:"refund_line_items,omitempty"`
	RefundShipping  *RefundShipping `json:"shipping,omitempty"`
	Transactions    []Transaction   `json:"transactions,omitempty"`

	Currency string `json:"currency,omitempty"`
	Note     string `json:"note,omitempty"`
}

func (req *RefundAPIReq) GetMethod() string {
	return "POST"
}

func (req *RefundAPIReq) GetData() interface{} {
	return req
}

type RefundLine struct {
	LineItemId   string `json:"line_item_id"`
	Quantity     int    `json:"quantity,omitempty"`
	RefundAmount string `json:"refund_amount,omitempty"`
}

type RefundShipping struct {
	Amount     string `json:"amount,omitempty"` // 10.00
	FullRefund bool   `json:"full_refund,omitempty"`
}

type RefundAPIResp struct {
	client.BaseAPIResponse
	Refund Refund `json:"data,omitempty"`
}

type Refund struct {
	Id              string        `json:"id,omitempty"` // refund id
	OrderId         string        `json:"order_id,omitempty"`
	CreatedAt       string        `json:"created_at,omitempty"` // ISO 8601
	Note            string        `json:"note,omitempty"`
	RefundLineItems []RefundLine  `json:"refund_line_items,omitempty"`
	Transactions    []Transaction `json:"transactions,omitempty"`

	TotalDutiesSet   TotalDutiesSet    `json:"total_duties_set,omitempty"`
	OrderAdjustments []OrderAdjustment `json:"order_adjustments,omitempty"`
}

type OrderAdjustment struct {
	Id           string       `json:"id,omitempty"`
	AmountSet    AmountSet    `json:"amount_set,omitempty"`
	Kind         string       `json:"kind,omitempty"`
	RefundId     string       `json:"refund_id,omitempty"`
	Amount       string       `json:"amount,omitempty"`
	OrderId      string       `json:"order_id,omitempty"`
	TaxAmount    string       `json:"tax_amount,omitempty"`
	TaxAmountSet TaxAmountSet `json:"tax_amount_set,omitempty"`
	Reason       string       `json:"reason,omitempty"`
}

type TaxAmountSet struct {
	PresentmentMoney Money `json:"presentment_money,omitempty"`
	ShopMoney        Money `json:"shop_money,omitempty"`
}

type AmountSet struct {
	PresentmentMoney Money `json:"presentment_money,omitempty"`
	ShopMoney        Money `json:"shop_money,omitempty"`
}

type TotalDutiesSet struct {
	PresentmentMoney Money `json:"presentment_money,omitempty"`
	ShopMoney        Money `json:"shop_money,omitempty"`
}

type Money struct {
	Amount       string `json:"amount,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
}

func (req *RefundAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *RefundAPIReq) GetEndpoint() string {
	return "order/refund.json"
}

// OrderRefund
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/order-refund?version=v20251201
// en：https://developer.shopline.com/docs/admin-rest-api/order/order-management/order-refund?version=v20251201
// Deprecated
// see OrderService
func OrderRefund(c *client.Client, apiReq *RefundAPIReq) (*RefundAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := apiReq.GetEndpoint()

	// 3. API response data
	apiResp := &RefundAPIResp{}

	// 4. Call API
	_, err := c.Post(context.Background(), endpoint, shopLineReq, apiResp)
	//if err != nil {
	//	fmt.Printf("Execute Request failed，endpoint: %s, shopLineReq: %v, err: %v\n", endpoint, shopLineReq, err)
	//	return nil, err
	//}

	return apiResp, err
}
