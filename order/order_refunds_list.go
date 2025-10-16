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
	client.BaseAPIRequest

	OrderId string `url:"order_id,omitempty"`
	Limit   uint64 `url:"limit,omitempty"`
}

func (r *ListOrderRefundsAPIReq) Method() string {
	return "GET"
}

func (r *ListOrderRefundsAPIReq) GetQuery() interface{} {
	return r
}

func (r *ListOrderRefundsAPIReq) Verify() error {
	if r.OrderId == "" {
		return errors.New("order_id is required")
	}
	return nil
}

func (r *ListOrderRefundsAPIReq) Endpoint() string {
	return fmt.Sprintf("orders/%s/refunds.json", r.OrderId)
}

type ListOrderRefundsAPIResp struct {
	client.BaseAPIResponse
	Refunds []Refund `json:"refunds,omitempty"`
}
