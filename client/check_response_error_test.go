package client

import (
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
)

type errReader struct{}

var testErr = errors.New("test-error")

func (errReader) Read([]byte) (int, error) {
	return 0, testErr
}

func (errReader) Close() error {
	return nil
}

func TestCheckResponseError(t *testing.T) {
	cases := []struct {
		resp     *http.Response
		expected error
	}{
		{
			httpmock.NewStringResponse(200, `{"foo": "bar"}`),
			nil,
		},
		{
			httpmock.NewStringResponse(299, `{"foo": "bar"}`),
			nil,
		},
		{
			httpmock.NewStringResponse(400, `{"errors": "bad request"}`),
			ResponseError{Status: 400, Message: "bad request"},
		},
		{
			httpmock.NewStringResponse(400, `{"errors": "order is wrong"}`),
			ResponseError{Status: 400, Message: "order is wrong", Errors: []string{"order: order is wrong"}},
		},
		{
			httpmock.NewStringResponse(400, `{"errors": "collection_id: collection_id is wrong"}`),
			ResponseError{Status: 400, Message: "collection_id: collection_id is wrong", Errors: []string{"collection_id: collection_id is wrong"}},
		},
		{
			httpmock.NewStringResponse(400, `{errors:bad request}`),
			errors.New("invalid character 'e' looking for beginning of object key string"),
		},
		{
			&http.Response{StatusCode: 400, Body: errReader{}},
			testErr,
		},
		{
			httpmock.NewStringResponse(422, `{"errors": "Unprocessable Entity - ok"}`),
			ResponseError{Status: 422, Message: "Unprocessable Entity - ok"},
		},
		{
			httpmock.NewStringResponse(500, `{"errors": "terrible error"}`),
			ResponseError{Status: 500, Message: "terrible error"},
		},
		{
			httpmock.NewStringResponse(500, `{"errors": "This action requires read_customers scope"}`),
			ResponseError{Status: 500, Message: "This action requires read_customers scope"},
		},
	}

	for _, c := range cases {
		actual := CheckHttpResponseError(c.resp)
		if fmt.Sprint(actual) != fmt.Sprint(c.expected) {
			t.Errorf("CheckHttpResponseError(): expected [%v], actual [%v]", c.expected, actual)
		}
	}
}
