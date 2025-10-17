package order

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderServiceCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"order":{"id": "1"}}`))

	order := Order{
		LineItems: []LineItem{
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
	apiResp, err := GetOrderService().List(context.Background(), apiReq)
	if err != nil {
		t.Errorf("ListOrders error: %v", err)
	}

	if len(apiResp.Orders) != 1 {
		t.Errorf("Order.List got %v orders, expected: 1", len(apiResp.Orders))
	}

	order := apiResp.Orders[0]
	orderTests(t, order)
}

func TestListAttributionInfos(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/order_attribution_info.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("order/order_attribution_info.json")))

	apiReq := &ListOrderAttributionInfosAPIReq{
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
