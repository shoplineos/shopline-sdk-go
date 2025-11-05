package order

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/rest/test"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestOrderServiceCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"order":{"id": "1"}}`))

	order := Order{
		LineItems: []CreateAnOrderAPIRespLineItem{
			{
				VariantId: "1",
				Quantity:  1,
			},
		},
	}

	apiReq := &CreateOrderAPIReq{
		Order: order,
	}
	o, err := GetOrderService().Create(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Order.Create returned error: %v", err)
	}

	expected := Order{Id: "1"}
	if o.Order.Id != expected.Id {
		t.Errorf("Order.Create returned id %s, expected %s", o.Order.Id, expected.Id)
	}
}

func TestOrderServiceList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("order/orders.json")))

	apiReq := &ListOrdersAPIReq{}
	//apiReq := &GetOrdersAPIReq{}
	apiResp, err := GetOrderService().List(context.Background(), apiReq)
	if err != nil {
		t.Errorf("ListOrders error: %v", err)
	}

	if len(apiResp.Orders) != 1 {
		t.Errorf("Order.List got %v orders, expected: 1", len(apiResp.Orders))
	}

	order := apiResp.Orders[0]
	orderTestsForListOrdersAPI(t, order)
}

func orderTestsForListOrdersAPI(t *testing.T, order ListOrdersAPIRespOrder) {
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
		httpmock.NewBytesResponder(200, test.LoadTestData("order/order_attribution_info.json")))

	apiReq := &ListOrderAttributionInfosAPIRequest{
		OrderIds: []string{"1"},
	}
	apiResp, err := GetOrderService().ListAttributionInfos(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Order.ListAttributionInfos error: %v", err)
	}

	if len(apiResp.OrderAttributionInfos) != 3 {
		t.Errorf("Order.List got %v orders, expected: 3", len(apiResp.OrderAttributionInfos))
	}

	attributionInfo := apiResp.OrderAttributionInfos[0]
	assert.NotNil(t, attributionInfo)
	assert.Equal(t, "21071058873779430237494658", attributionInfo.OrderSeq)

}
