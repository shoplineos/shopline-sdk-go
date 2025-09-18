package client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"shoplineapp/config"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var (
	client *Client
	app    App
)

func setup() *Client {
	app = App{
		AppKey:    config.DefaultAppKey,
		AppSecret: config.DefaultAppSecret,
	}

	client = MustNewClient(app, config.DefaultStoreHandle, config.DefaultAccessToken)
	if client == nil {
		panic("client is nil")
	}

	app.Client = client
	return client
}

func TestCreateAccessToken(t *testing.T) {
	setup()

	token, err := app.CreateAccessToken(context.Background(), "abc")
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("DefaultAccessToken: %v\n", token)
}

func TestRefreshAccessToken(t *testing.T) {
	setup()

	token, err := app.RefreshAccessToken(context.Background(), config.DefaultStoreHandle)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("DefaultAccessToken: %v\n", token)
}

// product detail
// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/query-single-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/query-single-product?version=v20251201
func TestGetProductDetail(t *testing.T) {

	setup()

	shopLineReq := &ShopLineRequest{}
	responseData := &map[string]any{}
	productId := "16070822412102455208483380"
	path := fmt.Sprintf("products/%s.json", productId)
	resp, err := client.Get(context.Background(), path, shopLineReq, responseData)

	if err != nil {
		t.Fatal(err)
	}

	//fmt.Println(resp)

	a := assert.New(t)
	a.Equal(200, resp.StatusCode)
}

// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/delete-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/delete-product?version=v20251201
func TestDeleteProduct(t *testing.T) {
	setup()

	productId := "16070822412199259745763380"
	responseData := &map[string]any{}
	shopLineReq := &ShopLineRequest{}

	path := fmt.Sprintf("products/%s.json", productId)

	resp, err := client.Delete(context.Background(), path, shopLineReq, responseData)
	if err != nil {
		t.Fatal(err)
	}
	a := assert.New(t)
	a.Equal(200, resp.StatusCode)

}

func TestVerifyWebhookRequest(t *testing.T) {
	setup()

	cases := []struct {
		hmac     string
		message  string
		expected bool
	}{
		{"1487c504fdb834b0ec315fc038e6f16ed0b51d5e7eb3f1d3aca862139673788b", "my secret message", true},
		{"a2344333", "my secret message", false},
		{"432r2238480", "", false},
		{"HFHFHKJFKFKLL1122=", "", false},
		{"", "", false},
		{"JLFJOO@(@(mm3302392838929fkfkak=", "my secret message", false},
	}

	for i, c := range cases {

		//testClient := MustNewClient(App{}, "", "")
		shoplineReq := &ShopLineRequest{
			Body: c.message,
		}
		req, err := client.NewHttpRequest(context.Background(), MethodGet, "", shoplineReq)
		if err != nil {
			t.Fatalf("Webhook.verify err = %v, expected true", err)
		}
		if c.hmac != "" {
			req.Header.Add("X-Shopline-Hmac-Sha256", c.hmac)
		}

		isValid := app.VerifyWebhookRequest(req)

		if isValid != c.expected {
			t.Errorf("Webhook.verify was expecting, idx: %d, expected:%t, but got %t", i, c.expected, isValid)
		} else {
			log.Printf("Webhook.verify was successful, idx: %d, %t", i, isValid)
		}
	}
}

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
