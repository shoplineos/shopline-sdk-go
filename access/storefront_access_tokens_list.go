package access

import (
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListStorefrontAccessTokensAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/access/storefront-api/create-an-access-token?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/access/storefront-api/create-an-access-token?version=v20251201
type ListStorefrontAccessTokensAPIReq struct {
	client.BaseAPIRequest
}

func (r *ListStorefrontAccessTokensAPIReq) Method() string {
	return "GET"
}

func (r *ListStorefrontAccessTokensAPIReq) Verify() error {
	return nil
}

func (r *ListStorefrontAccessTokensAPIReq) Endpoint() string {
	return fmt.Sprintf("storefront_access_tokens.json")
}

type ListStorefrontAccessTokensAPIResp struct {
	client.BaseAPIResponse
	StorefrontAccessTokens []StorefrontAccessToken `json:"storefront_access_tokens,omitempty"`
}
