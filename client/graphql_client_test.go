package client

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGraphQLQuery(t *testing.T) {
	setupStorefrontGraphQLClient()
	defer teardown()
	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("https://zwapptest.myshopline.com/%s/%s/graphql.json", "storefront/graph", client.ApiVersion),
		httpmock.NewStringResponder(200, `{"data":{"foo":"bar"}}`),
	)

	resp := struct {
		Foo string `json:"foo"`
	}{}
	err := graphQLClient.Query(context.Background(), "query {}", nil, &resp)
	if err != nil {
		t.Errorf("GraphQL.Query returned error: %v", err)
	}

	expectedFoo := "bar"
	if resp.Foo != expectedFoo {
		t.Errorf("resp.Foo returned %s expected %s", resp.Foo, expectedFoo)
	}
}

func TestGraphQLQueryWithError(t *testing.T) {
	setupStorefrontGraphQLClient()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("https://zwapptest.myshopline.com/%s/%s/graphql.json", "storefront/graph", client.ApiVersion),
		httpmock.NewStringResponder(200, `{"errors":[{"message":"oops"}]}`),
	)

	resp := struct {
		Foo string `json:"foo"`
	}{}
	err := graphQLClient.Query(context.Background(), "query {}", nil, &resp)

	if err == nil {
		t.Error("GraphQL.Query should return error!")
	}

	expectedError := "oops"
	if err.Error() != expectedError {
		t.Errorf("GraphQL.Query returned error message %s but expected %s", err.Error(), expectedError)
	}
}

func TestGraphQLQueryWithRetries(t *testing.T) {
	setupStorefrontGraphQLClient()
	defer teardown()

	type MyStruct struct {
		Foo string `json:"foo"`
	}

	//var retries int

	cases := []struct {
		description string
		responder   httpmock.Responder
		expected    interface{}
		retries     int
	}{
		{
			description: "no retries",
			responder: func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, `{"data":{"foo":"bar"}}`), nil
			},
			expected: MyStruct{Foo: "bar"},
			retries:  1,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {

			requestURL := fmt.Sprintf("https://zwapptest.myshopline.com/%s/%s/graphql.json", "storefront/graph", client.ApiVersion)

			httpmock.RegisterResponder(
				"POST",
				requestURL,
				c.responder,
			)

			resp := MyStruct{}
			err := graphQLClient.Query(context.Background(), "query {}", nil, &resp)

			callCountInfo := httpmock.GetCallCountInfo()

			attempts := callCountInfo[fmt.Sprintf("POST %s", requestURL)]

			if attempts != c.retries {
				t.Errorf("GraphQL.Query attempts equal %d but expected %d", attempts, c.retries)
			}

			if err != nil {
				if !reflect.DeepEqual(err, c.expected) {
					t.Errorf("GraphQL.Query got error %#v but expected %#v", err, c.expected)
				}
			} else if !reflect.DeepEqual(resp, c.expected) {
				t.Errorf("GraphQL.Query responsed %#v but expected %#v", resp, c.expected)
			}
		})
	}
}

func TestGraphQLQueryWithMultipleErrors(t *testing.T) {
	setupStorefrontGraphQLClient()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("https://zwapptest.myshopline.com/%s/%s/graphql.json", "storefront/graph", client.ApiVersion),
		httpmock.NewStringResponder(200, `{"errors":[{"message":"oops"},{"message":"I did it again"}]}`),
	)

	resp := struct {
		Foo string `json:"foo"`
	}{}
	err := graphQLClient.Query(context.Background(), "query {}", nil, &resp)

	if err == nil {
		t.Error("GraphQL.Query should return error!")
	}

	expectedError := "I did it again, oops"
	if err.Error() != expectedError {
		t.Errorf("GraphQL.Query returned error message %s but expected %s", err.Error(), expectedError)
	}
}

func TestGraphQLCostRetryAfterSeconds(t *testing.T) {
	cases := []struct {
		description string
		GraphQLCost GraphQLCost
		expected    float64
	}{
		{
			"last query passed, does not need to be throttled",
			GraphQLCost{
				RequestedQueryCost: 300,
				ActualQueryCost:    makeIntPointer(50),
				ThrottleStatus: GraphQLThrottleStatus{
					MaximumAvailable:   1000,
					CurrentlyAvailable: 400,
					RestoreRate:        50,
				},
			},
			0,
		},
		{
			"last query failed, needs to be throttled",
			GraphQLCost{
				RequestedQueryCost: 300,
				ActualQueryCost:    nil,
				ThrottleStatus: GraphQLThrottleStatus{
					MaximumAvailable:   1000,
					CurrentlyAvailable: 200,
					RestoreRate:        50,
				},
			},
			2,
		},
		{
			"last query passed, does not need to be throttled",
			GraphQLCost{
				RequestedQueryCost: 300,
				ActualQueryCost:    makeIntPointer(50),
				ThrottleStatus: GraphQLThrottleStatus{
					MaximumAvailable:   1000,
					CurrentlyAvailable: 200,
					RestoreRate:        50,
				},
			},
			0,
		},
		{
			"last query passed, needs to be throttled",
			GraphQLCost{
				RequestedQueryCost: 300,
				ActualQueryCost:    makeIntPointer(100),
				ThrottleStatus: GraphQLThrottleStatus{
					MaximumAvailable:   1000,
					CurrentlyAvailable: 50,
					RestoreRate:        50,
				},
			},
			1,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			s := c.GraphQLCost.RetryAfterSeconds()

			if s != c.expected {
				t.Errorf("GraphQLCost.RetryAfterSeconds returned %f expected %f (%s)", s, c.expected, c.description)
			}
		})
	}
}

func makeIntPointer(v int) *int {
	return &v
}
