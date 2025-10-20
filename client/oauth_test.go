package client

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/config"
	"github.com/stretchr/testify/assert"
	"log"
	"net/url"
	"testing"
)

const (
	AppKeyForTest      = "LDLLJJLflja2039203flaflaLLFLE"
	AppSecretForTest   = "LDLLJJLflja2039203flaflaLLFLE"
	StoreHandelForTest = "zwapptest"
	LangForTest        = "zh"
)

func TestCreateAccessToken(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://zwapptest.myshopline.com/admin/oauth/token/create",
		httpmock.NewStringResponder(200, `{
	"code": 200,
	"i18nCode": "SUCCESS",
	"message": "ok",
	"data": {
		"accessToken": "test-token",
		"expireTime": "2025-10-22",
		"scope": "ddd"
	}
}`))

	apiResp, err := app.CreateAccessToken(context.Background(), "test_code")
	if err != nil {
		t.Fatalf("App.CreateAccessToken(): %v", err)
	}

	expected := "test-token"
	if apiResp.Data.AccessToken != expected {
		t.Errorf("Token = %v, expected %v", apiResp.Data.AccessToken, expected)
	}
}

// https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201
func TestCreateAccessTokenError(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://zwapptest.myshopline.com/admin/oauth/token/create",
		httpmock.NewStringResponder(500, `{"errors": "system error"}`))

	_, err := app.CreateAccessToken(context.Background(), "test_code")
	a := assert.New(t)
	a.NotNil(err)
	a.Equal("system error", err.Error())

}

// https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E5%85%AB%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E5%88%B7%E6%96%B0-access-token
func TestRefreshAccessToken(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://zwapptest.myshopline.com/admin/oauth/token/refresh",
		httpmock.NewStringResponder(200, `{
	"code": 200,
	"i18nCode": "SUCCESS",
	"message": "ok",
	"data": {
		"accessToken": "test-token",
		"expireTime": "2025-10-22",
		"scope": "ddd"
	}
}`))

	accessToken, err := app.RefreshAccessToken(context.Background(), config.StoreHandelForUnitTest)
	if err != nil {
		log.Fatalf("Failed to refresh access token: %v\n", err)
	} else {
		fmt.Printf("Refresh Access Token Result: %v\n", accessToken)
	}

	//log.Printf("Status Code: %d\n", accessToken)
	log.Printf("New Access Token: %s\n", accessToken.Data.AccessToken)
	log.Printf("Expires In: %s seconds\n", accessToken.Data.ExpiresTime)
	//fmt.Printf("Refresh Token: %s\n", tokenResp.Body.RefreshToken)

	a := assert.New(t)
	a.Equal(200, accessToken.Code)

}

// https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E5%85%AB%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E5%88%B7%E6%96%B0-access-token
func TestRefreshAccessTokenError(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://zwapptest.myshopline.com/admin/oauth/token/refresh",
		httpmock.NewStringResponder(500, `{
	"code": "500",
	"i18nCode": "error",
	"message": "error"
}`))

	_, err := app.RefreshAccessToken(context.Background(), config.StoreHandelForUnitTest)
	a := assert.New(t)
	a.NotNil(err)
	a.Equal("error", err.Error())
}

//func TestRefreshAccessToken(t *testing.T) {
//	setup()
//	defer teardown()
//
//	token, err := app.RefreshAccessToken(context.Background(), config.DefaultStoreHandle)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	log.Printf("New AccessToken: %v\n", token.Data)
//}

func TestAuthorizeUrl(t *testing.T) {
	setup()

	storeHandle := "testStore"
	scope := "read_orders"
	redirectUri := app.RedirectUrl

	baseUrl := fmt.Sprintf("https://%s.myshopline.com/admin/oauth-web/#/oauth/authorize?appKey=%s&responseType=code&scope=%s&redirectUri=%s", storeHandle, app.AppKey, scope, redirectUri)

	authorizeUrl, err := app.AuthorizeUrl(storeHandle, scope)
	a := assert.New(t)
	a.Nil(err)
	a.Equal(baseUrl, authorizeUrl)

}

func TestAppVerifySign(t *testing.T) {

	localApp := App{
		AppKey:    AppKeyForTest,
		AppSecret: AppSecretForTest,
	}

	params := url.Values{}
	params.Add("appkey", AppKeyForTest)
	params.Add("handle", StoreHandelForTest)
	params.Add("lang", LangForTest)
	params.Add("timestamp", "21c335b3a")

	receivedSign := "654a42549daaad959b6488500f499fc0381c6cf24d8afceeb566729fc82fee6a"
	success := localApp.VerifySign(params, receivedSign)
	a := assert.New(t)
	a.Equal(success, true)

	receivedSign = "wrong"

	success = localApp.VerifySign(params, receivedSign)
	a.False(success)

}
