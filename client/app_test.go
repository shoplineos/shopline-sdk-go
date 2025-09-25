package client

import (
	"context"
	"fmt"
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

	log.Printf("New AccessToken: %v\n", token.Data)
}

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

func TestVerifySign(t *testing.T) {

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
