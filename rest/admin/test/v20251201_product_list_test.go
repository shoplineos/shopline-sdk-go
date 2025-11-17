package test

import (
	"context"
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/product"
	"net/http"
	"reflect"
	"runtime"
	"testing"
)

func TestProductList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"products": [{"Id":"1"},{"Id":"2"}]}`))

	apiReq := &product.GetProductsAPIReq{}
	apiResp := &product.GetProductsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("Product.List returned error: %v", err)
	}

	expected := []product.Product{{Id: "1"}, {Id: "2"}}
	if !reflect.DeepEqual(apiResp.Products, expected) {
		t.Errorf("Product.List returned %+v, expected %+v", apiResp.Products, expected)
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

	apiReq := &product.GetProductsAPIReq{
		Ids: "1,2,3",
	}
	apiResp := &product.GetProductsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("Product.List returned error: %v", err)
	}

	expected := []product.Product{{Id: "1"}, {Id: "2"}, {Id: "3"}}
	if !reflect.DeepEqual(apiResp.Products, expected) {
		t.Errorf("Product.List returned %+v, expected %+v", apiResp.Products, expected)
	}
}

func TestProductListError(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(500, ""))

	expectedErrMessage := "Unknown Error"

	apiReq := &product.GetProductsAPIReq{}
	apiResp := &product.GetProductsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

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
		expectedProducts    []product.Product
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
			expectedProducts: []product.Product{{Id: "1"}, {Id: "2"}, {Id: "3"}, {Id: "4"}},
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
			expectedProducts: []product.Product{{Id: "1"}},
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
			expectedProducts: []product.Product{},
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

			requestParams := &product.GetProductsAPIReq{}
			apiResp := &product.GetProductsAPIResp{}
			Products, err := client.ListAll(cli, context.Background(), requestParams, apiResp, func(resp interface{}) []product.Product {
				r := resp.(*product.GetProductsAPIResp)
				return r.Products
			})

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
		expectedProducts   []product.Product
		expectedPagination *client.Pagination
		expectedErr        error
	}{
		// Expect empty pagination when there is no link header
		{
			`{"products": [{"id":"1"},{"id":"2"}]}`,
			"",
			[]product.Product{{Id: "1"}, {Id: "2"}},
			new(client.Pagination),
			nil,
		},
		// Invalid link header responses
		{
			"{}",
			"invalid link",
			[]product.Product(nil),
			nil,
			client.ResponseDecodingError{Message: "could not extract pagination link header"},
		},
		{
			"{}",
			`<:invalid.url>; rel="next"`,
			[]product.Product(nil),
			nil,
			client.ResponseDecodingError{Message: "pagination does not contain a valid URL"},
		},
		{
			"{}",
			`<http://valid.url?%invalid_query>; rel="next"`,
			[]product.Product(nil),
			nil,
			errors.New(`invalid URL escape "%in"`),
		},
		{
			"{}",
			`<http://valid.url>; rel="next"`,
			[]product.Product(nil),
			nil,
			client.ResponseDecodingError{Message: "The page_info is missing"},
		},
		{
			"{}",
			`<http://valid.url?page_info=foo&limit=invalid>; rel="next"`,
			[]product.Product(nil),
			nil,
			errors.New(limitConversionErrorMessage),
		},
		// Valid link header responses
		{
			`{"products": [{"id":"1"}]}`,
			`<http://valid.url?page_info=foo&limit=2>; rel="next"`,
			[]product.Product{{Id: "1"}},
			&client.Pagination{
				Next: &client.ListOptions{PageInfo: "foo", Limit: 2},
			},
			nil,
		},
		{
			`{"products": [{"id":"2"}]}`,
			`<http://valid.url?page_info=foo>; rel="next", <http://valid.url?page_info=bar>; rel="previous"`,
			[]product.Product{{Id: "2"}},
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

		requestParams := &product.GetProductsAPIReq{}
		apiResp := &product.GetProductsAPIResp{}
		err := cli.Call(context.Background(), requestParams, apiResp)
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
