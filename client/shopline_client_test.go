package client

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/config"
	"log"
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

	//httpmock.ActivateNonDefault(client.Client)

	return client

}

func teardown() {
	httpmock.DeactivateAndReset()
}

// zh: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/delete-product?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/product/product/delete-product?version=v20251201
//func TestDeleteProduct(t *testing.T) {
//	setup()
//
//	productId := "16070822412199259745763380"
//	responseData := &map[string]any{}
//	shopLineReq := &ShopLineRequest{}
//
//	endpoint := fmt.Sprintf("products/%s.json", productId)
//
//	resp, err := client.Delete(context.Background(), endpoint, shopLineReq, responseData)
//	if err != nil {
//		t.Fatal(err)
//	}
//	a := assert.New(t)
//	a.Equal(200, resp.StatusCode)
//
//}

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
			Data: c.message,
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

func TestBuildShopLineResponse(t *testing.T) {

	// case 1
	httpResp := httpmock.NewStringResponse(200, `{"foo": "bar"}`)

	responseData := &map[string]any{}
	response, err := buildShopLineResponse(httpResp, responseData)
	a := assert.New(t)
	a.Nil(err)
	a.NotNil(response)
	a.NotNil(response.Data)
	a2 := (*responseData)["foo"].(string)
	a.Equal(a2, "bar")

	// case 2
	httpResp = httpmock.NewStringResponse(500, `{"errors": "system error"}`)

	response, err = buildShopLineResponse(httpResp, responseData)
	a.NotNil(err)
	a.Equal(err.Error(), "system error")
	a.NotNil(response)

}

type testRequestData struct {
	Name string `url:"name"`
}

func TestBuildFinalRequestUrl(t *testing.T) {
	requestData := &testRequestData{}
	requestData.Name = "lisi"
	shoplineReq := &ShopLineRequest{
		Data: requestData,
	}

	app := App{
		AppKey:    AppKeyForTest,
		AppSecret: AppSecretForTest,
	}

	c := MustNewClient(app, StoreHandelForTest, "access token")

	url, err := c.buildFinalRequestUrl(MethodGet, "test/foo", shoplineReq)
	a := assert.New(t)
	a.Nil(err)
	a.Equal("https://zwapptest.myshopline.com/test/foo?name=lisi", url)

}
