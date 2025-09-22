package oauth

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/manager"
	"log"
	"net/url"
)

// RefreshAccessToken
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#step6
func RefreshAccessToken(appkey, storeHandle string) (*client.TokenResponse, error) {
	if appkey == "" {
		return nil, fmt.Errorf("appKey is required")
	}
	if storeHandle == "" {
		return nil, fmt.Errorf("storeHandle is required")
	}

	accessToken, err := manager.GetApp(appkey).RefreshAccessToken(context.Background(), storeHandle)

	if err != nil {
		log.Printf("Refresh access token failed, appkey: %s, err: %v\n", appkey, err)
		return nil, err
	}

	return accessToken, nil
}

// CreateAccessToken
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#step-4-request-an-access-token
func CreateAccessToken(appkey, code string) (*client.TokenResponse, error) {
	if appkey == "" {
		return nil, fmt.Errorf("appKey is required")
	}
	if code == "" {
		return nil, fmt.Errorf("code is required")
	}

	accessToken, err := manager.GetApp(appkey).CreateAccessToken(context.Background(), code)

	if err != nil {
		log.Printf("Create access token failed, appkey: %s, err: %v\n", appkey, err)
		return nil, err
	}

	return accessToken, nil
}

// VerifySign verify callback url's params sign
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/generate-and-verify-signatures?version=v20260301
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/generate-and-verify-signatures?version=v20260301
func VerifySign(appkey string, query url.Values, sign string) bool {
	return manager.GetApp(appkey).VerifySign(query, sign)
}

// AuthorizeUrl
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E4%B8%89%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E6%8E%88%E6%9D%83%E7%A0%81
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#step2
func AuthorizeUrl(appkey, storeHandle, scope string) (string, error) {
	return manager.GetApp(appkey).AuthorizeUrl(storeHandle, scope)
}
