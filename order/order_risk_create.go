package order

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CreateOrderRiskAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/add-risk-fraud?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/order/order-management/add-risk-fraud?version=v20251201
type CreateOrderRiskAPIReq struct {
	client.BaseAPIRequest
	OrderId         string          `url:"order_id,omitempty"`
	CreateOrderRisk CreateOrderRisk `json:"risk,omitempty"`
}

type CreateOrderRisk struct {
	Display        bool            `json:"display,omitempty"`
	CauseCancel    bool            `json:"cause_cancel,omitempty"`
	Recommendation string          `json:"recommendation,omitempty"`
	RiskDetailMsg  []RiskDetailMsg `json:"risk_detail_msg_list,omitempty"`
	Source         string          `json:"source,omitempty"`
	Score          string          `json:"score,omitempty"`
	CheckoutId     string          `json:"checkout_id,omitempty"`
}

func (r *CreateOrderRiskAPIReq) Method() string {
	return "POST"
}

func (r *CreateOrderRiskAPIReq) Verify() error {
	if r.OrderId == "" {
		return errors.New("order_id is required")
	}
	return nil
}

func (r *CreateOrderRiskAPIReq) Endpoint() string {
	return fmt.Sprintf("orders/v2/%s/risks.json", r.OrderId)
}

type CreateOrderRiskAPIResp struct {
	client.BaseAPIResponse
	OrderRisk OrderRisk `json:"risk,omitempty"`
}
