package access

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// DeleteStorefrontAccessTokenAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/access/storefront-api/create-an-access-token?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/access/storefront-api/create-an-access-token?version=v20251201
type DeleteStorefrontAccessTokenAPIReq struct {
	client.BaseAPIRequest
	Id string
}

func (r *DeleteStorefrontAccessTokenAPIReq) GetMethod() string {
	return "DELETE"
}

func (r *DeleteStorefrontAccessTokenAPIReq) GetData() interface{} {
	return r
}

func (r *DeleteStorefrontAccessTokenAPIReq) Verify() error {
	if r.Id == "" {
		return errors.New("storefront access token id is empty")
	}
	return nil
}

func (r *DeleteStorefrontAccessTokenAPIReq) GetEndpoint() string {
	return fmt.Sprintf("storefront_access_tokens/%s.json", r.Id)
}

type DeleteStorefrontAccessTokenAPIResp struct {
	client.BaseAPIResponse
}
