package order

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"testing"
)

//
//func TestQueryOrdersCount(t *testing.T) {
//
//	apiReq := &GetOrdersCountAPIRequest{
//		Status: "any",
//	}
//
//	c := manager.GetDefaultClient()
//	apiResp, err := QueryOrdersCount(c, apiReq)
//
//	fmt.Printf("Count: %v\n", apiResp)
//	if err != nil {
//		log.Printf("Request failed, error: %v", err)
//	}
//
//	a := assert.New(t)
//	a.Greater(apiResp.Count, 0)
//}

func TestOrderCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/count.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"count": 7}`))

	params := map[string]string{"created_at_min": "2025-09-30T10:14:36-00:00"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/count.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		params,
		httpmock.NewStringResponder(200, `{"count": 2}`))

	apiReq := &GetOrdersCountAPIRequest{}
	apiResp, err := QueryOrdersCount(cli, apiReq)

	if err != nil {
		t.Errorf("Order.Count returned error: %v", err)
	}

	expected := 7
	if apiResp.Count != expected {
		t.Errorf("Order.Count returned %d, expected %d", apiResp.Count, expected)
	}

	//date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	apiReq = &GetOrdersCountAPIRequest{
		CreatedAtMin: "2025-09-30T10:14:36-00:00",
	}
	apiResp, err = QueryOrdersCount(cli, apiReq)

	if err != nil {
		t.Errorf("Order.Count returned error: %v", err)
	}

	expected = 2
	if apiResp.Count != expected {
		t.Errorf("Order.Count returned %d, expected %d", apiResp.Count, expected)
	}
}
