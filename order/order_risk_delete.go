package order

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// DeleteOrderRiskAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/remove-the-specify-of-the-order-risk-fraud?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/order/order-management/remove-the-specify-of-the-order-risk-fraud?version=v20251201
type DeleteOrderRiskAPIReq struct {
	OrderId string `url:"order_id,omitempty"`
	RiskId  string `url:"id,omitempty"`
}

func (r DeleteOrderRiskAPIReq) Verify() error {
	if r.OrderId == "" {
		return errors.New("order_id is required")
	}
	if r.RiskId == "" {
		return errors.New("risk_id is required")
	}
	return nil
}

func (r DeleteOrderRiskAPIReq) Endpoint() string {
	return fmt.Sprintf("orders/v2/%s/risks/%s.json", r.OrderId, r.RiskId)
}

type DeleteOrderRiskAPIResp struct {
	client.BaseAPIResponse
}
