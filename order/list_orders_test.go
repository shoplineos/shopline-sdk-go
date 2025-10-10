package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/test"
	"net/http"
	"reflect"
	"runtime"
	"testing"
)

func TestOrderList(t *testing.T) {
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

func TestOrderListOptions(t *testing.T) {
	setup()
	defer teardown()
	params := map[string]string{
		"fields": "id,name",
		"limit":  "250",
		//"page":   "10",
		"status": "any",
	}

	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		params,
		httpmock.NewBytesResponder(200, test.LoadTestData("order/orders.json")))

	apiReq := &ListOrdersAPIReq{
		Limit:  "250",
		Fields: "id,name",
		Status: "any",
	}
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

func TestOrderListAll(t *testing.T) {
	setup()
	defer teardown()

	listURL := fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion)

	cases := []struct {
		name                string
		expectedOrders      []Order
		expectedRequestURLs []string
		expectedLinkHeaders []string
		expectedBodies      []string
		expectedErr         error
	}{
		{
			name: "Pulls the next page",
			expectedRequestURLs: []string{
				listURL,
				fmt.Sprintf("%s?page_info=pg2", listURL),
			},
			expectedLinkHeaders: []string{
				`<http://valid.url?page_info=pg2>; rel="next"`,
				`<http://valid.url?page_info=pg1>; rel="previous"`,
			},
			expectedBodies: []string{
				`{"orders": [{"id":"1"},{"id":"2"}]}`,
				`{"orders": [{"id":"3"},{"id":"4"}]}`,
			},
			expectedOrders: []Order{{ID: "1"}, {ID: "2"}, {ID: "3"}, {ID: "4"}},
			expectedErr:    nil,
		},
		{
			name: "Stops when there is not a next page",
			expectedRequestURLs: []string{
				listURL,
			},
			expectedLinkHeaders: []string{
				`<http://valid.url?page_info=pg2>; rel="previous"`,
			},
			expectedBodies: []string{
				`{"orders": [{"id":"1"}]}`,
			},
			expectedOrders: []Order{{ID: "1"}},
			expectedErr:    nil,
		},
		{
			name: "Returns errors when required",
			expectedRequestURLs: []string{
				listURL,
			},
			expectedLinkHeaders: []string{
				`<http://valid.url?paage_info=pg2>; rel="previous"`,
			},
			expectedBodies: []string{
				`{"orders": []}`,
			},
			expectedOrders: []Order{},
			expectedErr:    errors.New("The page_info is missing"),
		},
	}

	for i, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if len(c.expectedRequestURLs) != len(c.expectedLinkHeaders) {
				t.Errorf(
					"test case must have the same number of expected request urls (%d) as expected link headers (%d)",
					len(c.expectedRequestURLs),
					len(c.expectedLinkHeaders),
				)

				return
			}

			if len(c.expectedRequestURLs) != len(c.expectedBodies) {
				t.Errorf(
					"test case must have the same number of expected request urls (%d) as expected bodies (%d)",
					len(c.expectedRequestURLs),
					len(c.expectedBodies),
				)

				return
			}

			for i := range c.expectedRequestURLs {
				response := &http.Response{
					StatusCode: 200,
					Body:       httpmock.NewRespBodyFromString(c.expectedBodies[i]),
					Header: http.Header{
						"Link": {c.expectedLinkHeaders[i]},
					},
				}

				httpmock.RegisterResponder("GET", c.expectedRequestURLs[i], httpmock.ResponderFromResponse(response))
			}

			apiReq := &ListOrdersAPIReq{}
			orders, err := GetOrderService().ListAll(context.Background(), apiReq)
			//if err != nil {
			//	t.Errorf("ListOrders error, index: %d, err: %v", i, err)
			//}

			if !reflect.DeepEqual(orders, c.expectedOrders) {
				t.Errorf("test %d Order.ListAll orders returned %+v, expected %+v", i, orders, c.expectedOrders)
			}

			if (c.expectedErr != nil || err != nil) && err.Error() != c.expectedErr.Error() {
				t.Errorf(
					"test %d Order.ListAll err returned %+v, expected %+v",
					i,
					err,
					c.expectedErr,
				)
			}
		})
	}
}

func TestOrderListWithPagination(t *testing.T) {
	setup()
	defer teardown()

	listURL := fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion)

	// The strconv.Atoi error changed in go 1.8, 1.7 is still being tested/supported.
	limitConversionErrorMessage := `strconv.Atoi: parsing "invalid": invalid syntax`
	if runtime.Version()[2:5] == "1.7" {
		limitConversionErrorMessage = `strconv.ParseInt: parsing "invalid": invalid syntax`
	}

	cases := []struct {
		body               string
		linkHeader         string
		expectedOrders     []Order
		expectedPagination *client.Pagination
		expectedErr        error
	}{
		// Expect empty pagination when there is no link header
		{
			`{"orders": [{"id":"1"},{"id":"2"}]}`,
			"",
			[]Order{{ID: "1"}, {ID: "2"}},
			new(client.Pagination),
			nil,
		},
		// Invalid link header responses
		{
			"{}",
			"invalid link",
			[]Order(nil),
			nil,
			client.ResponseDecodingError{Message: "could not extract pagination link header"},
		},
		{
			"{}",
			`<:invalid.url>; rel="next"`,
			[]Order(nil),
			nil,
			client.ResponseDecodingError{Message: "pagination does not contain a valid URL"},
		},
		{
			"{}",
			`<http://valid.url?%invalid_query>; rel="next"`,
			[]Order(nil),
			nil,
			errors.New(`invalid URL escape "%in"`),
		},
		{
			"{}",
			`<http://valid.url>; rel="next"`,
			[]Order(nil),
			nil,
			client.ResponseDecodingError{Message: "The page_info is missing"},
		},
		{
			"{}",
			`<http://valid.url?page_info=foo&limit=invalid>; rel="next"`,
			[]Order(nil),
			nil,
			errors.New(limitConversionErrorMessage),
		},
		// Valid link header responses
		{
			`{"orders": [{"id":"1"}]}`,
			`<http://valid.url?page_info=foo&limit=2>; rel="next"`,
			[]Order{{ID: "1"}},
			&client.Pagination{
				Next: &client.ListOptions{PageInfo: "foo", Limit: 2},
			},
			nil,
		},
		{
			`{"orders": [{"id":"2"}]}`,
			`<http://valid.url?page_info=foo>; rel="next", <http://valid.url?page_info=bar>; rel="previous"`,
			[]Order{{ID: "2"}},
			&client.Pagination{
				Next:     &client.ListOptions{PageInfo: "foo"},
				Previous: &client.ListOptions{PageInfo: "bar"},
			},
			nil,
		},
	}
	for i, c := range cases {
		response := &http.Response{
			StatusCode: 200,
			Body:       httpmock.NewRespBodyFromString(c.body),
			Header: http.Header{
				"Link": {c.linkHeader},
			},
		}

		httpmock.RegisterResponder("GET", listURL, httpmock.ResponderFromResponse(response))

		//orders, pagination, err := client.Order.ListWithPagination(context.Background(), nil)
		apiReq := &ListOrdersAPIReq{}
		apiResp, err := GetOrderService().List(context.Background(), apiReq)

		if apiResp != nil && apiResp.Orders != nil && !reflect.DeepEqual(apiResp.Orders, c.expectedOrders) {
			t.Errorf("test %d Order.ListWithPagination orders returned %+v, expected %+v", i, apiResp.Orders, c.expectedOrders)
		}

		if apiResp != nil && apiResp.Pagination != nil && !reflect.DeepEqual(apiResp.Pagination, c.expectedPagination) {
			t.Errorf(
				"test %d Order.ListWithPagination pagination returned %+v, expected %+v",
				i,
				apiResp.Pagination,
				c.expectedPagination,
			)
		}

		if (c.expectedErr != nil || err != nil) && err.Error() != c.expectedErr.Error() {
			t.Errorf(
				"test %d Order.ListWithPagination err returned %+v, expected %+v",
				i,
				err,
				c.expectedErr,
			)
		}
	}
}

func TestOrderListError(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(500, ""))

	expectedErrMessage := "Unknown Error"

	apiReq := &ListOrdersAPIReq{}
	apiResp, err := GetOrderService().List(context.Background(), apiReq)

	if apiResp != nil && apiResp.Orders != nil {
		t.Errorf("Order.List returned orders, expected nil: %v", err)
	}

	if err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Order.List err returned %+v, expected %+v", err, expectedErrMessage)
	}
}

//// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/get-orders?version=v20251201
//// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/get-orders?version=v20251201
//func TestQueryOrders(t *testing.T) {
//
//	apiReq := &ListOrdersAPIReq{
//		Limit: "2", // 10 for every page
//		//SortCondition:   "created_at:desc",
//		//CreatedAtMin:    "2024-01-01T00:00:00+08:00",
//		//FinancialStatus: "paid",
//		//IDs: "21070984663426767352955294",
//	}
//
//	c := manager.GetDefaultClient()
//
//	apiResp, err := ListOrders(c, apiReq)
//
//	if err != nil {
//		fmt.Println("Query orders failed, err:", err)
//		return
//	}
//
//	fmt.Printf("Find total %d orders\n", len(apiResp.Orders))
//	for i, order := range apiResp.Orders {
//		fmt.Printf("\nOrder %d:\n", i+1)
//		fmt.Printf("Order ID: %s\n", order.ID)
//		fmt.Printf("Order Name: %s\n", order.Name)
//		fmt.Printf("Order Status: %s\n", order.Status)
//		fmt.Printf("Order FinancialStatus: %s\n", order.FinancialStatus)
//		fmt.Printf("Order CreatedAt: %s\n", order.CreatedAt)
//		fmt.Printf("Order TotalPrice: %s %s\n", order.CurrentTotalPrice, order.Currency)
//		fmt.Printf("Order Email: %s\n", order.Customer.Email)
//		fmt.Printf("Order Items Count: %d\n", len(order.LineItems))
//	}
//	assert.Nil(t, err)
//	assert.NotNil(t, apiResp)
//}
