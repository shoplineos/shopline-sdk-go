package access

import (
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CreateStorefrontAccessTokenAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/access/storefront-api/create-an-access-token?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/access/storefront-api/create-an-access-token?version=v20251201
type CreateStorefrontAccessTokenAPIReq struct {
	StorefrontAccessToken CreateStorefrontAccessToken `json:"storefront_access_token,omitempty"`
}

func (c CreateStorefrontAccessTokenAPIReq) Verify() error {
	return nil
}

func (c CreateStorefrontAccessTokenAPIReq) Endpoint() string {
	return fmt.Sprintf("storefront_access_tokens.json")
}

type CreateStorefrontAccessTokenAPIResp struct {
	StorefrontAccessToken StorefrontAccessToken `json:"storefront_access_token,omitempty"`
	client.BaseAPIResponse
}

type CreateStorefrontAccessToken struct {
	Title string `json:"title,omitempty"`
}
