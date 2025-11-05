package order

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListOrderRisksAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/get-fraud-risks-for-an-order?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/order/order-management/get-fraud-risks-for-an-order?version=v20251201
// Deprecated
// See GetFraudRisksForAnOrderAPIReq
type ListOrderRisksAPIReq struct {
	client.BaseAPIRequest

	OrderId string `url:"order_id"`
}

func (r *ListOrderRisksAPIReq) GetMethod() string {
	return "GET"
}

func (r *ListOrderRisksAPIReq) Verify() error {
	if r.OrderId == "" {
		return errors.New("order_id is required")
	}
	return nil
}

func (r *ListOrderRisksAPIReq) GetEndpoint() string {
	return fmt.Sprintf("orders/v2/%s/risks.json", r.OrderId)
}

type ListOrderRisksAPIResp struct {
	client.BaseAPIResponse
	OrderRisks []OrderRisk `json:"risks,omitempty"`
}
