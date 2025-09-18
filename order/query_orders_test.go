package order

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/get-orders?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/get-orders?version=v20251201
func TestQueryOrders(t *testing.T) {

	apiReq := &QueryOrdersAPIReq{
		Limit: "2", // 10 for every page
		//SortCondition:   "created_at:desc",
		//CreatedAtMin:    "2024-01-01T00:00:00+08:00",
		//FinancialStatus: "paid",
		//IDs: "21070984663426767352955294",
	}

	apiResp, err := QueryOrders(apiReq)

	if err != nil {
		fmt.Println("Query orders failed, err:", err)
		return
	}

	fmt.Printf("Find total %d orders\n", len(apiResp.Orders))
	for i, order := range apiResp.Orders {
		fmt.Printf("\nOrder %d:\n", i+1)
		fmt.Printf("Order ID: %s\n", order.ID)
		fmt.Printf("Order Name: %s\n", order.Name)
		fmt.Printf("Order Status: %s\n", order.Status)
		fmt.Printf("Order FinancialStatus: %s\n", order.FinancialStatus)
		fmt.Printf("Order CreatedAt: %s\n", order.CreatedAt)
		fmt.Printf("Order TotalPrice: %s %s\n", order.TotalPrice, order.Currency)
		fmt.Printf("Order Email: %s\n", order.Customer.Email)
		fmt.Printf("Order Items Count: %d\n", len(order.LineItems))
	}
	assert.Nil(t, err)
	assert.NotNil(t, apiResp)
}
