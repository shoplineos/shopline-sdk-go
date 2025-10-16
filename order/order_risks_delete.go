package order

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// DeleteOrderRisksAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/delete-all-risk-fraud-related-to-the-order?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/order/order-management/delete-all-risk-fraud-related-to-the-order?version=v20251201
type DeleteOrderRisksAPIReq struct {
	client.BaseAPIRequest
	OrderId string `url:"order_id,omitempty"`
}

func (r *DeleteOrderRisksAPIReq) Method() string {
	return "DELETE"
}

func (r *DeleteOrderRisksAPIReq) GetData() interface{} {
	return r
}

func (r *DeleteOrderRisksAPIReq) Verify() error {
	if r.OrderId == "" {
		return errors.New("order_id is required")
	}
	return nil
}

func (r *DeleteOrderRisksAPIReq) Endpoint() string {
	return fmt.Sprintf("orders/v2/%s/risks.json", r.OrderId)
}

type DeleteOrderRisksAPIResp struct {
	client.BaseAPIResponse
}
