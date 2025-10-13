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
	OrderID string `url:"order_id,omitempty"`
}

func (r DeleteOrderRisksAPIReq) Verify() error {
	if r.OrderID == "" {
		return errors.New("order_id is required")
	}
	return nil
}

func (r DeleteOrderRisksAPIReq) Endpoint() string {
	return fmt.Sprintf("orders/v2/%s/risks.json", r.OrderID)
}

type DeleteOrderRisksAPIResp struct {
	client.BaseAPIResponse
}
