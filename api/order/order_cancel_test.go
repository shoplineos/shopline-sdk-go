package order

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/test"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

//// 21071580496925210798359834
//func TestCancelOrder(t *testing.T) {
//
//	apiReq := &CancelOrderAPIReq{
//		OrderId: "123",
//	}
//
//	c := manager.GetDefaultClient()
//
//	apiResp, err := CancelOrder(c, apiReq)
//	if err != nil {
//		fmt.Println("Cancel order failed, err:", err)
//	} else {
//		fmt.Printf("Cancel order successful！orderID: %s\n", apiResp.Order.Id)
//	}
//
//	a := assert.New(t)
//	a.NotNil(err)
//
//}
//
//// 21071580496925210798359834
//func TestCancelOrderCase2(t *testing.T) {
//
//	apiReq := &CancelOrderAPIReq{
//		OrderId: "21071580496925210798359811",
//	}
//
//	c := manager.GetDefaultClient()
//
//	apiResp, err := CancelOrder(c, apiReq)
//	if err != nil {
//		fmt.Println("Cancel order failed, err:", err)
//	} else {
//		fmt.Printf("Cancel order successful！orderID: %s\n", apiResp.Order.Id)
//	}
//
//	a := assert.New(t)
//	a.NotNil(err)
//
//}
//
//// 21071580496925210798359834
//func TestCancelOrderCase3(t *testing.T) {
//
//	apiReq := &CancelOrderAPIReq{
//		OrderId: "21071580496925210798359834",
//	}
//
//	c := manager.GetDefaultClient()
//
//	apiResp, err := CancelOrder(c, apiReq)
//	if err != nil {
//		fmt.Println("Cancel order failed, err:", err)
//	} else {
//		fmt.Printf("Cancel order successful！orderID: %s\n", apiResp.Order.Id)
//	}
//
//	a := assert.New(t)
//	a.NotNil(err)
//
//}

func orderTests(t *testing.T, order Order) {
	// Check that dates are parsed
	//d := time.Date(2016, time.May, 17, 4, 14, 36, 0, time.UTC)
	if !assert.Equal(t, "2025-09-30T10:14:36-00:00", order.CreatedAt) {
		t.Errorf("Order.CreatedAt returned %+v, expected %+v", order.CreatedAt, "2025-09-30T10:14:36-00:00")
	}

	// Check null dates
	if order.ProcessedAt != "" {
		t.Errorf("Order.ProcessedAt returned %+v, expected %+v", order.ProcessedAt, nil)
	}

	// Check prices
	//p := decimal.NewFromFloat(10)
	if !strings.EqualFold("9.50", order.CurrentTotalPrice) {
		t.Errorf("Order.CurrentTotalPrice returned %+v, expected %+v", order.CurrentTotalPrice, "9.50")
	}

	if !strings.EqualFold("1.00", order.CurrentTotalTax) {
		t.Errorf("Order.CurrentTotalTax returned %+v, expected %+v", order.CurrentTotalTax, "1.00")
	}

	//Check customer
	//if order.Customer == nil {
	//	t.Error("Expected Customer to not be nil")
	//}

	if order.Customer.Email != "john@test.com" {
		t.Errorf("Customer.Email, expected %v, actual %v", "john@test.com", order.Customer.Email)
	}

	//ptp := decimal.NewFromFloat(9)
	lineItem := order.LineItems[0]
	if !assert.Equal(t, "1.00", lineItem.Price) {
		t.Errorf("Order.LineItems[0].Price, expected %v, actual %v", "1.00", lineItem.Price)
	}
}

func TestOrderCancel(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/123/cancel.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("order/order.json")))

	apiReq := &CancelOrderAPIReq{
		OrderId: "123",
	}

	apiResp, err := GetOrderService().Cancel(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Order.CancelledAt returned error: %v", err)
	}

	// Check that dates are parsed
	//timezone, _ := time.LoadLocation("America/New_York")
	//d := time.Date(2016, time.May, 17, 4, 14, 36, 0, timezone)
	d := "2025-09-30T10:14:36-00:00"
	if !assert.Equal(t, d, apiResp.Order.CancelledAt) {
		t.Errorf("Order.CancelledAt returned %+v, expected %+v", apiResp.Order.CancelledAt, d)
	}

	orderTests(t, apiResp.Order)
}
