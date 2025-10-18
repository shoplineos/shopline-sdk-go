package client

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/signature"
	"net/http"
	"net/url"
)

// AuthorizeUrl
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E4%B8%89%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E6%8E%88%E6%9D%83%E7%A0%81
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/#step2
func (app App) AuthorizeUrl(storeHandle string, scope string) (string, error) {
	redirectUri := url.QueryEscape(app.RedirectUrl)
	scope = app.resolveScope(scope)
	scope = url.QueryEscape(scope)
	authorizeUrl := fmt.Sprintf("https://%s.myshopline.com/admin/oauth-web/#/oauth/authorize?appKey=%s&responseType=code&scope=%s&redirectUri=%s", storeHandle, app.AppKey, scope, redirectUri)

	return authorizeUrl, nil
}

func (app App) VerifySign(params url.Values, receivedSign string) bool {
	isSignValid := signature.VerifySign(app.AppSecret, params, receivedSign)
	return isSignValid
}

// VerifyWebhookMessage Verify a message against a message HMAC
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
func (app App) VerifyWebhookMessage(message, messageMAC string) bool {
	return signature.VerifyWebhookMessage(app.AppSecret, message, messageMAC)
}

// VerifyWebhookRequest Verify a Webhook http request, sent by shopline.
// The body of the request is still readable after invoking the method.
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
func (app App) VerifyWebhookRequest(r *http.Request) bool {
	return signature.VerifyWebhookRequest(app.AppSecret, r)
}

func (app App) resolveScope(scope string) string {
	if scope != "" {
		return scope
	}
	return app.Scope
}

func (app App) GetStoreHandle() string {
	return app.Client.StoreHandle
}

// CreateAccessToken
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#%E7%AC%AC%E5%85%AD%E6%AD%A5app-%E8%AF%B7%E6%B1%82-access-token
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20251201#step-4-request-an-access-token
func (app App) CreateAccessToken(ctx context.Context, code string) (*TokenResponse, error) {

	// 1. Build request
	requestBody := map[string]string{
		"code": code,
	}

	shopLineReq := &ShopLineRequest{
		Options: &RequestOptions{EnableSign: true},
		Data:    requestBody,
	}

	// 2. New http request
	httpReq, err := app.Client.NewHttpRequest(ctx, MethodPost, "admin/oauth/token/create", shopLineReq)
	if err != nil {
		return nil, err
	}

	// 3. Specify resource
	tokenResponse := &TokenResponse{}

	// 4. Execute http
	_, err = app.Client.executeHttpRequest(shopLineReq, httpReq, tokenResponse)
	if err != nil {
		return nil, err
	}
	return tokenResponse, nil
}

// RefreshAccessToken
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20251201#step6
func (app App) RefreshAccessToken(ctx context.Context, storeHandle string) (*TokenResponse, error) {
	err := verifyForRefreshAccessToken(app.AppKey, app.AppSecret, storeHandle)
	if err != nil {
		return nil, err
	}

	// 1. Build request
	shopLineReq := &ShopLineRequest{
		Options: &RequestOptions{EnableSign: true},
	}

	// 2. New http request
	httpReq, err := app.Client.NewHttpRequest(ctx, MethodPost, "admin/oauth/token/refresh", shopLineReq)
	if err != nil {
		return nil, err
	}

	// 3. Specify resource
	tokenResponse := &TokenResponse{}

	// 4. Execute http
	_, err = app.Client.executeHttpRequest(shopLineReq, httpReq, tokenResponse)
	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}

//// RefreshAccessToken
//// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201
//// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#step6
//func RefreshAccessToken(app client.App, storeHandle string) (*client.TokenResponse, error) {
//	if storeHandle == "" {
//		return nil, fmt.Errorf("storeHandle is required")
//	}
//
//	accessToken, err := app.RefreshAccessToken(context.Background(), storeHandle)
//
//	if err != nil {
//		log.Printf("Refresh access token failed, appkey: %s, err: %v\n", app.AppKey, err)
//		return nil, err
//	}
//
//	return accessToken, nil
//}
//
//// CreateAccessToken
//// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201
//// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#step-4-request-an-access-token
//func CreateAccessToken(app client.App, code string) (*client.TokenResponse, error) {
//	if code == "" {
//		return nil, fmt.Errorf("code is required")
//	}
//
//	accessToken, err := app.CreateAccessToken(context.Background(), code)
//
//	if err != nil {
//		log.Printf("Create access token failed, appkey: %s, err: %v\n", app.AppKey, err)
//		return nil, err
//	}
//
//	return accessToken, nil
//}
//
//// VerifySign verify callback url's params sign
//// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/generate-and-verify-signatures?version=v20260301
//// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/generate-and-verify-signatures?version=v20260301
//func VerifySign(app client.App, query url.Values, sign string) bool {
//	return app.VerifySign(query, sign)
//}
//
//// AuthorizeUrl
//// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E4%B8%89%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E6%8E%88%E6%9D%83%E7%A0%81
//// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#step2
//func AuthorizeUrl(app client.App, storeHandle, scope string) (string, error) {
//	return app.AuthorizeUrl(storeHandle, scope)
//}
