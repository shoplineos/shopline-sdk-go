package order

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListOrderRefundsAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/query-all-refund-orders-for-the-id-association?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/order/order-management/query-all-refund-orders-for-the-id-association?version=v20251201
type ListOrderRefundsAPIReq struct {
	OrderID string `url:"order_id,omitempty"`
	Limit   uint64 `url:"limit,omitempty"`
}

func (r ListOrderRefundsAPIReq) Verify() error {
	if r.OrderID == "" {
		return errors.New("order_id is required")
	}
	return nil
}

func (r ListOrderRefundsAPIReq) Endpoint() string {
	return fmt.Sprintf("orders/%s/refunds.json", r.OrderID)
}

type ListOrderRefundsAPIResp struct {
	client.BaseAPIResponse
	Refunds []Refund `json:"refunds,omitempty"`
}
