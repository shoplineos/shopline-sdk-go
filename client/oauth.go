package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// TokenResponse access Token struct
type TokenResponse struct {
	Code     int    `json:"code"`
	I18nCode string `json:"i18nCode"`
	Message  string `json:"message"`
	Data     struct {
		AccessToken string `json:"accessToken"`
		ExpiresTime string `json:"expireTime"`
		Scope       string `json:"scope"`
		//RefreshToken string `json:"refresh_token"`
	} `json:"data"`
}

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
	isSignValid := VerifySign(app.AppSecret, params, receivedSign)
	return isSignValid
}

// VerifyWebhookMessage Verify a message against a message HMAC
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
func (app App) VerifyWebhookMessage(message, messageMAC string) bool {
	return VerifyWebhookMessage(app.AppSecret, message, messageMAC)
}

// VerifyWebhookRequest Verify a Webhook http request, sent by SHOPLINE.
// The body of the request is still readable after invoking the method.
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
func (app App) VerifyWebhookRequest(r *http.Request) bool {
	return VerifyWebhookRequest(app.AppSecret, r)
}

func (app App) resolveScope(scope string) string {
	if scope != "" {
		return scope
	}
	return app.Scope
}

//func (app App) GetStoreHandle() string {
//	return app.Client.StoreHandle
//}

// CreateAccessToken
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#%E7%AC%AC%E5%85%AD%E6%AD%A5app-%E8%AF%B7%E6%B1%82-access-token
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20251201#step-4-request-an-access-token
func (app App) CreateAccessToken(ctx context.Context, code string) (*TokenResponse, error) {

	return app.Client.CreateAccessToken(ctx, code)
}

// RefreshAccessToken
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20251201#step6
func (app App) RefreshAccessToken(ctx context.Context, storeHandle string) (*TokenResponse, error) {
	err := verifyForRefreshAccessToken(app.AppKey, app.AppSecret, storeHandle)
	if err != nil {
		return nil, err
	}

	return app.Client.RefreshAccessToken(ctx, storeHandle)
}

func verifyForRefreshAccessToken(appkey, appSecret, shopHandle string) error {
	if appkey == "" {
		return fmt.Errorf("appKey is required")
	}
	if appSecret == "" {
		return fmt.Errorf("appSecret is required")
	}

	if shopHandle == "" {
		return fmt.Errorf("shopHandle is required")
	}
	return nil
}
