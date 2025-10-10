package order

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/test"
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
				VariantID: "1",
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

	expected := Order{ID: "1"}
	if o.Order.ID != expected.ID {
		t.Errorf("Order.Create returned id %s, expected %s", o.Order.ID, expected.ID)
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
