package order

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CalculateOrderRefundAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/order-refund-trial?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/order/order-management/order-refund-trial?version=v20251201
type CalculateOrderRefundAPIReq struct {
	client.BaseAPIRequest

	OrderId         string           `url:"order_id,omitempty"`
	Currency        string           `json:"currency,omitempty"`
	RefundLineItems []RefundLineItem `json:"refund_line_items,omitempty"`

	RefundShipping RefundShipping `json:"shipping,omitempty"`
}

func (r *CalculateOrderRefundAPIReq) Method() string {
	return "POST"
}

func (r *CalculateOrderRefundAPIReq) Verify() error {
	if r.OrderId == "" {
		return errors.New("OrderId is required")
	}
	return nil
}

func (r *CalculateOrderRefundAPIReq) Endpoint() string {
	return fmt.Sprintf("orders/%s/refunds/calculate.json", r.OrderId)
}

type RefundLineItem struct {
	LineItemID string `json:"line_item_id"`
	Quantity   uint64 `json:"quantity,omitempty"`
	Type       string `json:"type,omitempty"`
}

type CalculateOrderRefundAPIResp struct {
	client.BaseAPIResponse
	CalculateRefundResult CalculateRefundResult `json:"refund,omitempty"`
}

type CalculateRefundResult struct {
	RefundTransaction []CalculateRefundTransaction `json:"transactons,omitempty"`

	Currency string `json:"currency,omitempty"`

	RefundLineItems []CalculateRefundLineItem `json:"refund_line_items,omitempty"`
	RefundShipping  CalculateRefundShipping   `json:"shipping,omitempty"`

	TotalDutiesSet CalculateTotalDutiesSet `json:"total_duties_set,omitempty"`
}

type CalculateRefundTransaction struct {
	Gateway           string `json:"gateway,omitempty"` // Pay Gateways
	Kind              string `json:"kind,omitempty"`
	MaximumRefundable string `json:"maximum_refundable,omitempty"`
	OrderId           string `json:"order_id,omitempty"`
	ParentId          string `json:"parent_id,omitempty"`
	Amount            string `json:"amount,omitempty"` // Amount
	Currency          string `json:"currency,omitempty"`
}

type CalculateRefundLineItem struct {
	Type                string `json:"type,omitempty"`
	DiscountTotalAmount string `json:"discount_total_amount,omitempty"`
	LineItemId          string `json:"line_item_id,omitempty"`
	Price               string `json:"price,omitempty"`
	Quantity            uint32 `json:"quantity,omitempty"`
	Subtotal            string `json:"subtotal,omitempty"`
	TotalTax            string `json:"total_tax,omitempty"`
}

type CalculateRefundShipping struct {
	Amount            string `json:"amount,omitempty"`
	MaximumRefundable string `json:"maximum_refundable,omitempty"`
	Tax               string `json:"tax,omitempty"`
}

type CalculateTotalDutiesSet struct {
	PresentmentMoney CalculatePresentmentMoney `json:"presentment_money,omitempty"`
	ShopMoney        CalculateShopMoney        `json:"shop_money,omitempty"`
}

type CalculatePresentmentMoney struct {
	Amount       string `json:"amount,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
}
type CalculateShopMoney struct {
	Amount       string `json:"amount,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
}
