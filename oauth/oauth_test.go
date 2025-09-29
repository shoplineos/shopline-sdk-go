package oauth

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/config"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	cli *client.Client
	app client.App
)

func setup() {
	app = client.App{
		AppKey:    config.AppKeyForUnitTest,
		AppSecret: config.AppSecretForUnitTest,
	}

	cli = client.MustNewClient(app, config.StoreHandelForUnitTest, config.AccessTokenForUnitTest, client.WithEnableLogDetail(true))
	if cli == nil {
		panic("client is nil")
	}

	app.Client = cli

	httpmock.ActivateNonDefault(cli.Client)
}

func teardown() {
	httpmock.DeactivateAndReset()
}

func TestAppCreateAccessToken(t *testing.T) {
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

	apiResp, err := CreateAccessToken(app, "test_code")
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

	_, err := CreateAccessToken(app, "test_code")
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

	accessToken, err := RefreshAccessToken(app, config.StoreHandelForUnitTest)
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

	_, err := RefreshAccessToken(app, config.StoreHandelForUnitTest)
	a := assert.New(t)
	a.NotNil(err)
	a.Equal("error", err.Error())
}

//https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E5%85%AB%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E5%88%B7%E6%96%B0-access-token
//func TestRefreshAccessTokenFromShopLine(t *testing.T) {
//
//	// replace app data
//	storeHandle := config.DefaultStoreHandle
//	appKey := config.DefaultAppKey
//
//	app := manager.GetApp(appKey)
//
//	accessToken, err := RefreshAccessToken(app, storeHandle)
//	if err != nil {
//		log.Fatalf("Failed to refresh access token: %v\n", err)
//	} else {
//		fmt.Printf("Refresh Access Token Result: %v\n", accessToken)
//	}
//
//	//log.Printf("Status Code: %d\n", accessToken)
//	log.Printf("New Access Token: %s\n", accessToken.Data.AccessToken)
//	log.Printf("Expires In: %s seconds\n", accessToken.Data.ExpiresTime)
//	//fmt.Printf("Refresh Token: %s\n", tokenResp.Body.RefreshToken)
//
//	a := assert.New(t)
//	a.Equal(200, accessToken.Code)
//
//}
