package order

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetOrderRefundAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/query-order-specify-of-your-money-back-order?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/order/order-management/query-order-specify-of-your-money-back-order?version=v20251201
// Deprecated
type GetOrderRefundAPIReq struct {
	client.BaseAPIRequest

	OrderId  string `url:"order_id,omitempty"`
	RefundId string `url:"refund_id,omitempty"`
}

func (r *GetOrderRefundAPIReq) GetMethod() string {
	return "GET"
}

func (r *GetOrderRefundAPIReq) GetQuery() interface{} {
	return r
}

func (r *GetOrderRefundAPIReq) Verify() error {
	if r.OrderId == "" {
		return errors.New("OrderId is required")
	}
	if r.RefundId == "" {
		return errors.New("RefundId is required")
	}
	return nil
}

func (r *GetOrderRefundAPIReq) GetEndpoint() string {
	return fmt.Sprintf("orders/%s/refunds/%s.json", r.OrderId, r.RefundId)
}

type GetOrderRefundAPIResp struct {
	client.BaseAPIResponse
	Refund Refund `json:"refund,omitempty"`
}
