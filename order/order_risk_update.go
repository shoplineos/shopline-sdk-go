package order

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// UpdateOrderRiskAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/update-order-risk-fraud?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/order/order-management/update-order-risk-fraud?version=v20251201
type UpdateOrderRiskAPIReq struct {
	OrderID         string          `url:"order_id,omitempty"`
	RiskID          string          `url:"id,omitempty"`
	UpdateOrderRisk UpdateOrderRisk `json:"risk,omitempty"`
}

type UpdateOrderRisk struct {
	Display        bool            `json:"display,omitempty"`
	CauseCancel    bool            `json:"cause_cancel,omitempty"`
	Recommendation string          `json:"recommendation,omitempty"`
	RiskDetailMsg  []RiskDetailMsg `json:"risk_detail_msg_list,omitempty"`
	Source         string          `json:"source,omitempty"`
	Score          string          `json:"score,omitempty"`
	CheckoutId     string          `json:"checkout_id,omitempty"`
}

func (r UpdateOrderRiskAPIReq) Verify() error {
	if r.OrderID == "" {
		return errors.New("order_id is required")
	}
	return nil
}

func (r UpdateOrderRiskAPIReq) Endpoint() string {
	return fmt.Sprintf("orders/v2/%s/risks/%s.json", r.OrderID, r.RiskID)
}

type UpdateOrderRiskAPIResp struct {
	client.BaseAPIResponse
	OrderRisk OrderRisk `json:"risk,omitempty"`
}
