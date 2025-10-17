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
	client.BaseAPIRequest

	OrderId string `url:"order_id,omitempty"`
	RiskId  string `url:"id,omitempty"`
}

func (r *GetOrderRiskAPIReq) GetMethod() string {
	return "GET"
}

func (r *GetOrderRiskAPIReq) GetQuery() interface{} {
	return r
}

func (r *GetOrderRiskAPIReq) Verify() error {
	if r.OrderId == "" {
		return errors.New("order_id is required")
	}
	if r.RiskId == "" {
		return errors.New("risk id is required")
	}
	return nil
}

func (r *GetOrderRiskAPIReq) GetEndpoint() string {
	return fmt.Sprintf("orders/v2/%s/risks/%s.json", r.OrderId, r.RiskId)
}

type GetOrderRiskAPIResp struct {
	client.BaseAPIResponse
	OrderRisk OrderRisk `json:"risk,omitempty"`
}
