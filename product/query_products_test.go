package product

import (
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"net/http"
	"reflect"
	"runtime"
	"testing"
)

//// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-products?version=v20251201
//// en: https://developer.shopline.com/docs/admin-rest-api/product/product/get-products?version=v20251201
//func TestQueryProducts(t *testing.T) {
//
//	apiReq := &QueryProductsAPIReq{
//		ProductCategory: "Clothes",
//		Status:          "active",
//		OrderBy:         "created_at_desc",
//		Limit:           3,
//		// Fields:         "id,title,status,vendor",
//	}
//
//	c := manager.GetDefaultClient()
//	apiResp, err := QueryProducts(c, apiReq)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Printf("apiResp:%v\n", apiResp)
//}
//
//// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-products?version=v20251201
//// en: https://developer.shopline.com/docs/admin-rest-api/product/product/get-products?version=v20251201
//func TestQueryProductsTitle(t *testing.T) {
//
//	apiReq := &QueryProductsAPIReq{
//		//ProductCategory: "Clothes",
//		//Status:          "active",
//		//OrderBy:         "created_at_desc",
//		//Limit:           3,
//		Fields: "id,title",
//	}
//
//	c := manager.GetDefaultClient()
//	apiResp, err := QueryProducts(c, apiReq)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Printf("apiResp:%v\n", apiResp)
//
//}
//
//// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-products?version=v20251201
//// en: https://developer.shopline.com/docs/admin-rest-api/product/product/get-products?version=v20251201
//func TestQueryProductsPagination(t *testing.T) {
//
//	requestParams := &QueryProductsAPIReq{
//		//ProductCategory: "Clothes",
//		//Status:          "active",
//		//OrderBy:         "created_at_desc",
//		Limit:    4,
//		Fields:   "title",
//		PageInfo: "eyJzaW5jZUlkIjoiMTYwNzA4MjgzODk5MzAwNDQzOTA0NTMzODAiLCJkaXJlY3Rpb24iOiJuZXh0IiwibGltaXQiOjR9",
//	}
//
//	c := manager.GetDefaultClient()
//	products, err := QueryProducts(c, requestParams)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Printf("products:%v\n", products)
//}

func TestProductList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"products": [{"Id":"1"},{"Id":"2"}]}`))

	requestParams := &QueryProductsAPIReq{}
	products, err := QueryProducts(cli, requestParams)
	if err != nil {
		t.Errorf("Product.List returned error: %v", err)
	}

	expected := []ProductRespData{{Id: "1"}, {Id: "2"}}
	if !reflect.DeepEqual(products.Products, expected) {
		t.Errorf("Product.List returned %+v, expected %+v", products.Products, expected)
	}
}

func TestProductListFilterByIds(t *testing.T) {
	setup()
	defer teardown()

	params := map[string]string{"ids": "1,2,3"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		params,
		httpmock.NewStringResponder(200, `{"products": [{"id":"1"},{"id":"2"},{"id":"3"}]}`))

	requestParams := &QueryProductsAPIReq{
		IDs: "1,2,3",
	}
	productsAPIResp, err := QueryProducts(cli, requestParams)

	if err != nil {
		t.Errorf("Product.List returned error: %v", err)
	}

	expected := []ProductRespData{{Id: "1"}, {Id: "2"}, {Id: "3"}}
	if !reflect.DeepEqual(productsAPIResp.Products, expected) {
		t.Errorf("Product.List returned %+v, expected %+v", productsAPIResp.Products, expected)
	}
}

func TestProductListError(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(500, ""))

	expectedErrMessage := "Unknown Error"

	requestParams := &QueryProductsAPIReq{}
	_, err := QueryProducts(cli, requestParams)

	if err == nil {
		t.Errorf("Product.List returned products, expected nil: %v", err)
	}

	if err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Product.List err returned %+v, expected %+v", err, expectedErrMessage)
	}
}

func TestProductListAll(t *testing.T) {
	setup()
	defer teardown()

	listURL := fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion)

	cases := []struct {
		name                string // case name
		expectedProducts    []ProductRespData
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
				`{"products": [{"id":"1"},{"id":"2"}]}`,
				`{"products": [{"id":"3"},{"id":"4"}]}`,
			},
			expectedProducts: []ProductRespData{{Id: "1"}, {Id: "2"}, {Id: "3"}, {Id: "4"}},
			expectedErr:      nil,
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
				`{"products": [{"id":"1"}]}`,
			},
			expectedProducts: []ProductRespData{{Id: "1"}},
			expectedErr:      nil,
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
				`{"products": []}`,
			},
			expectedProducts: []ProductRespData{},
			expectedErr:      errors.New("The page_info is missing"),
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

			requestParams := &QueryProductsAPIReq{}

			Products, err := QueryProductsAll(cli, requestParams)

			if !reflect.DeepEqual(Products, c.expectedProducts) {
				t.Errorf("test %d Product.ListAll orders returned %+v, expected %+v", i, Products, c.expectedProducts)
			}

			if (c.expectedErr != nil || err != nil) && err.Error() != c.expectedErr.Error() {
				t.Errorf(
					"test %d Product.ListAll err returned %+v, expected %+v",
					i,
					err,
					c.expectedErr,
				)
			}
		})
	}
}

func TestProductListWithPagination(t *testing.T) {
	setup()
	defer teardown()

	listURL := fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion)

	// The strconv.Atoi error changed in go 1.8, 1.7 is still being tested/supported.
	limitConversionErrorMessage := `strconv.Atoi: parsing "invalid": invalid syntax`
	if runtime.Version()[2:5] == "1.7" {
		limitConversionErrorMessage = `strconv.ParseInt: parsing "invalid": invalid syntax`
	}

	cases := []struct {
		body               string
		linkHeader         string
		expectedProducts   []ProductRespData
		expectedPagination *client.Pagination
		expectedErr        error
	}{
		// Expect empty pagination when there is no link header
		{
			`{"products": [{"id":"1"},{"id":"2"}]}`,
			"",
			[]ProductRespData{{Id: "1"}, {Id: "2"}},
			new(client.Pagination),
			nil,
		},
		// Invalid link header responses
		{
			"{}",
			"invalid link",
			[]ProductRespData(nil),
			nil,
			client.ResponseDecodingError{Message: "could not extract pagination link header"},
		},
		{
			"{}",
			`<:invalid.url>; rel="next"`,
			[]ProductRespData(nil),
			nil,
			client.ResponseDecodingError{Message: "pagination does not contain a valid URL"},
		},
		{
			"{}",
			`<http://valid.url?%invalid_query>; rel="next"`,
			[]ProductRespData(nil),
			nil,
			errors.New(`invalid URL escape "%in"`),
		},
		{
			"{}",
			`<http://valid.url>; rel="next"`,
			[]ProductRespData(nil),
			nil,
			client.ResponseDecodingError{Message: "The page_info is missing"},
		},
		{
			"{}",
			`<http://valid.url?page_info=foo&limit=invalid>; rel="next"`,
			[]ProductRespData(nil),
			nil,
			errors.New(limitConversionErrorMessage),
		},
		// Valid link header responses
		{
			`{"products": [{"id":"1"}]}`,
			`<http://valid.url?page_info=foo&limit=2>; rel="next"`,
			[]ProductRespData{{Id: "1"}},
			&client.Pagination{
				Next: &client.ListOptions{PageInfo: "foo", Limit: 2},
			},
			nil,
		},
		{
			`{"products": [{"id":"2"}]}`,
			`<http://valid.url?page_info=foo>; rel="next", <http://valid.url?page_info=bar>; rel="previous"`,
			[]ProductRespData{{Id: "2"}},
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

		requestParams := &QueryProductsAPIReq{}
		apiResp, err := QueryProducts(cli, requestParams)

		if c.expectedErr != nil || err != nil {
			if err.Error() != c.expectedErr.Error() {
				t.Errorf(
					"test %d Product.ListWithPagination err returned %+v, expected %+v",
					i,
					err,
					c.expectedErr,
				)
			}

		} else {

			products := apiResp.Products
			pagination := apiResp.Pagination

			if !reflect.DeepEqual(products, c.expectedProducts) {
				t.Errorf("test %d Product.ListWithPagination products returned %+v, expected %+v", i, products, c.expectedProducts)
			}

			if pagination != nil && !reflect.DeepEqual(pagination, c.expectedPagination) {
				t.Errorf(
					"test %d Product.ListWithPagination pagination returned %+v, expected %+v",
					i,
					pagination,
					c.expectedPagination,
				)
			}

		}

	}
}
