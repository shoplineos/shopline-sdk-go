package order

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetOrderRefundAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/query-order-specify-of-your-money-back-order?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/order/order-management/query-order-specify-of-your-money-back-order?version=v20251201
type GetOrderRefundAPIReq struct {
	OrderID  string `url:"order_id,omitempty"`
	RefundID string `url:"refund_id,omitempty"`
}

func (r GetOrderRefundAPIReq) Verify() error {
	if r.OrderID == "" {
		return errors.New("OrderID is required")
	}
	if r.RefundID == "" {
		return errors.New("RefundID is required")
	}
	return nil
}

func (r GetOrderRefundAPIReq) Endpoint() string {
	return fmt.Sprintf("orders/%s/refunds/%s.json", r.OrderID, r.RefundID)
}

type GetOrderRefundAPIResp struct {
	client.BaseAPIResponse
	Refund Refund `json:"refund,omitempty"`
}
