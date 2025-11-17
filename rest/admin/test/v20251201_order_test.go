package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	order3 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/order"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestOrderCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"order":{"id": "1"}}`))

	order := order3.CreateAnOrderAPIReqOrder{
		LineItems: []order3.LineItem{
			{
				VariantId: "1",
				Quantity:  1,
			},
		},
	}

	apiReq := &order3.CreateAnOrderAPIReq{
		Order: order,
	}

	apiResp := &order3.CreateAnOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)
	if err != nil {
		t.Errorf("Order.Create returned error: %v", err)
	}

	expected := order3.Order{Id: "1"}
	if apiResp.Order.Id != expected.Id {
		t.Errorf("Order.Create returned id %s, expected %s", apiResp.Order.Id, expected.Id)
	}
}

func TestGetOrders(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "order/orders.json")))

	apiReq := &order3.GetOrdersAPIReq{}
	apiResp := &order3.GetOrdersAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("GetOrders error: %v", err)
	}

	if len(apiResp.Orders) != 1 {
		t.Errorf("Order.List got %v orders, expected: 1", len(apiResp.Orders))
	}

	order := apiResp.Orders[0]
	orderTestsForListOrdersAPI(t, order)
}

func orderTestsForListOrdersAPI(t *testing.T, order order3.GetOrdersAPIRespOrder) {
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

func TestListAttributionInfos(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/order_attribution_info.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "order/order_attribution_info.json")))

	apiReq := &order3.GetOrderAttributionInBulkAPIReq{
		Orders: []string{"1"},
	}
	apiResp := &order3.GetOrderAttributionInBulkAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)
	if err != nil {
		t.Errorf("Order.ListAttributionInfos error: %v", err)
	}

	if len(apiResp.Data) != 3 {
		t.Errorf("Order.ListAttributionInfos got %v orders, expected: 3", len(apiResp.Data))
	}

	attributionInfo := apiResp.Data[0]
	assert.NotNil(t, attributionInfo)
	assert.Equal(t, "21071058873779430237494658", attributionInfo.OrderSeq)

}
