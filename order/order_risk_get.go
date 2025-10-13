package order

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetOrderRiskAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/query-risk-fraud-for-order-specify?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/order/order-management/query-risk-fraud-for-order-specify?version=v20251201
type GetOrderRiskAPIReq struct {
	OrderID string `url:"order_id,omitempty"`
	RiskID  string `url:"id,omitempty"`
}

func (r GetOrderRiskAPIReq) Verify() error {
	if r.OrderID == "" {
		return errors.New("order_id is required")
	}
	if r.RiskID == "" {
		return errors.New("risk id is required")
	}
	return nil
}

func (r GetOrderRiskAPIReq) Endpoint() string {
	return fmt.Sprintf("orders/v2/%s/risks/%s.json", r.OrderID, r.RiskID)
}

type GetOrderRiskAPIResp struct {
	client.BaseAPIResponse
	OrderRisk OrderRisk `json:"risk,omitempty"`
}
