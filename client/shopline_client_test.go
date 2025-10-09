package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/config"
	"github.com/shoplineos/shopline-sdk-go/test"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var (
	client *Client
	app    App
)

func setup() {
	app = App{
		AppKey:    config.DefaultAppKey,
		AppSecret: config.DefaultAppSecret,
	}

	client = MustNewClient(app, config.DefaultStoreHandle, config.DefaultAccessToken)
	if client == nil {
		panic("client is nil")
	}

	app.Client = client

	httpmock.ActivateNonDefault(client.Client)
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

func TestNewClient(t *testing.T) {
	app := App{
		AppKey:    config.AppKeyForUnitTest,
		AppSecret: config.AppSecretForUnitTest,
	}

	cli, err := NewClient(app, config.StoreHandelForUnitTest, config.AccessTokenForUnitTest)
	if err != nil {
		t.Errorf("NewClient() err = %v, expected nil", err)
	}

	assert.Equal(t, app.AppKey, cli.App.AppKey)
}

func TestNewClientOptions(t *testing.T) {
	app := App{
		AppKey:    config.AppKeyForUnitTest,
		AppSecret: config.AppSecretForUnitTest,
	}

	cli, err := NewClient(app, config.StoreHandelForUnitTest, config.AccessTokenForUnitTest, WithVersion("v20250601"))
	if err != nil {
		t.Errorf("NewClient() err = %v, expected nil", err)
	}

	assert.Equal(t, "v20250601", cli.ApiVersion)
}

//func TestNewClientAware(t *testing.T) {
//	app := App{
//		AppKey:    config.AppKeyForUnitTest,
//		AppSecret: config.AppSecretForUnitTest,
//	}
//
//	cli, err := NewClient(app, config.StoreHandelForUnitTest, config.AccessTokenForUnitTest, WithVersion("v20250601"), WithClientAware())
//	if err != nil {
//		t.Errorf("NewClient() err = %v, expected nil", err)
//	}
//
//	assert.Equal(t, "v20250601", cli.ApiVersion)
//}

func TestResolveUrlPath(t *testing.T) {
	setup()
	defer teardown()

	req := &ShopLineRequest{}

	path := client.resolveUrlPath("orders.json", req)
	assert.Equal(t, "admin/openapi/v20251201/orders.json", path)
}

func TestSerializeBodyDataIfNecessary(t *testing.T) {
	setup()
	defer teardown()

	requestBody := map[string]string{
		"code": "hello world",
	}

	actual, err := client.serializeBodyDataIfNecessary(MethodPost, &ShopLineRequest{
		Data: requestBody,
	})
	if err != nil {
		t.Fatalf("SerializeBodyDataIfNecessary() err = %v, expected nil", err)
	}

	expected, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("SerializeBodyDataIfNecessary() err = %v, expected nil", err)
	}

	assert.Equal(t, expected, actual)
}

type ProductStruct struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	BodyHTML string `json:"bodyHTML"`
	Vendor   string `json:"vendor"`
}

type CreateProductAPIReqStruct struct {
	Product ProductStruct `json:"product"`
}

type CreateProductAPIRespStruct struct {
	Product ProductStruct `json:"product"`
}

func TestExecuteInternal(t *testing.T) {
	setup()
	defer teardown()

	// test create product
	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/products.json", client.StoreHandle, client.PathPrefix, client.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("product/product.json")))

	// 1. build request
	product := ProductStruct{
		Title:    "Hello shopline Freestyle 111",
		BodyHTML: "<strong>Hello shopline!<\\/strong>",
		Vendor:   "shopline",
	}

	apiReq := &CreateProductAPIReqStruct{
		Product: product,
	}

	apiResp := &CreateProductAPIRespStruct{}

	shopLineReq := &ShopLineRequest{
		Data: apiReq,
	}

	// 2. Call API
	client.executeInternal(context.Background(), MethodPost, "products/products.json", shopLineReq, apiResp)

	//fmt.Printf("apiResp: %+v\n", apiResp.Product)
	assert.Equal(t, "111", apiResp.Product.Id)

}

func TestBuildBodyJsonString(t *testing.T) {
	requestBody := map[string]string{
		"code": "r0r42304fklajfjlafa",
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}

	jsonString, err := buildBodyJsonString(jsonBody)
	if err != nil {
		log.Fatalf("Failed to buildBodyJsonString: %v", err)
	}

	a := assert.New(t)
	a.Equal(jsonString, `{"code":"r0r42304fklajfjlafa"}`)
}

func TestGenerateSign(t *testing.T) {
	timestamp := "21c335b3a"
	requestBody := map[string]string{
		"code": "code",
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}

	sign, err := generateSign(config.AppKeyForUnitTest, config.AppSecretForUnitTest, timestamp, jsonBody)
	if err != nil {
		t.Fatalf("GenerateSign() err = %v, expected nil", err)
	}

	a := assert.New(t)
	//fmt.Printf("sign: %s\n", sign)
	a.Equal(sign, "4c4cf0ac52439d1308b58897f857dcba8545efb6c0e8091fba64b148218d2074")
}

func TestNewHttpRequest(t *testing.T) {
	setup()
	defer teardown()

	inURL, outURL := "foo?page=1", fmt.Sprintf("https://%s.myshopline.com/foo?page=1", client.StoreHandle)

	inBody := struct {
		Hello string `json:"hello"`
	}{Hello: "World"}

	outBody := `{"hello":"World"}`

	slReq := &ShopLineRequest{
		Data: inBody,
	}

	req, err := client.NewHttpRequest(context.Background(), MethodPost, inURL, slReq)
	if err != nil {
		t.Fatalf("NewHttpRequest(%v) err = %v, expected nil", inURL, err)
	}

	// Test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewHttpRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// Test body was JSON encoded
	body, _ := io.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewHttpRequest(%v) Body = %v, expected %v", inBody, string(body), outBody)
	}

	verifyUserAgent(t, req)
	verifyToken(t, req)
	verifyAccept(t, req)
	verifyContentType(t, req)
	verifyTimestamp(t, req)

}

func verifyTimestamp(t *testing.T, req *http.Request) {
	timestamp := req.Header.Get("timestamp")
	if timestamp == "" || len(timestamp) != 13 {
		t.Errorf("NewHttpRequest() timestamp = %v", timestamp)
	}
}

func verifyUserAgent(t *testing.T, req *http.Request) {
	// Test user-agent is attached to the request
	userAgent := req.Header.Get("User-Agent")
	if userAgent != config.UserAgent {
		t.Errorf("NewHttpRequest() User-Agent = %v, expected %v", userAgent, config.UserAgent)
	}
}

func verifyToken(t *testing.T, req *http.Request) {
	// Test token is attached to the request
	token := req.Header.Get("Authorization")
	expected := "Bearer " + client.Token
	if token != expected {
		t.Errorf("NewHttpRequest() Authorization Token = %v, expected %v", token, expected)
	}
}

func verifyAccept(t *testing.T, req *http.Request) {
	// Accept
	accept := req.Header.Get("Accept")
	acceptExpected := "application/json"
	if accept != acceptExpected {
		t.Errorf("NewHttpRequest() Accept = %v, expected %v", accept, acceptExpected)
	}
}

func verifyContentType(t *testing.T, req *http.Request) {
	// contentType
	contentType := req.Header.Get("Content-Type")
	contentTypeExpected := "application/json"
	if contentType != contentTypeExpected {
		t.Errorf("NewHttpRequest() ContentType = %v, expected %v", contentType, contentTypeExpected)
	}
}

func TestVerifyWebhookRequest(t *testing.T) {
	setup()
	defer teardown()

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
		req, err := client.NewHttpRequest(context.Background(), MethodPost, "", shoplineReq)
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
	response, err := buildShopLineResponse("Get", httpResp, responseData)
	a := assert.New(t)
	a.Nil(err)
	a.NotNil(response)
	a.NotNil(response.Data)
	a2 := (*responseData)["foo"].(string)
	a.Equal(a2, "bar")

	// case 2
	httpResp = httpmock.NewStringResponse(500, `{"errors": "system error"}`)

	response, err = buildShopLineResponse("Get", httpResp, responseData)
	a.NotNil(err)
	a.Equal(err.Error(), "system error")
	a.NotNil(response)

}

type testRequestData struct {
	Name string `url:"name"`
}

func TestBuildRequestUrl(t *testing.T) {
	requestData := &testRequestData{}
	requestData.Name = "lisi"
	shoplineReq := &ShopLineRequest{
		Query: requestData,
	}

	app := App{
		AppKey:    AppKeyForTest,
		AppSecret: AppSecretForTest,
	}

	c := MustNewClient(app, StoreHandelForTest, "access token")

	url, err := c.buildRequestUrl(MethodGet, "test/foo", shoplineReq)
	a := assert.New(t)
	a.Nil(err)
	a.Equal("https://zwapptest.myshopline.com/test/foo?name=lisi", url)

}
