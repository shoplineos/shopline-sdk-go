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
	ID string
}

func (c DeleteStorefrontAccessTokenAPIReq) Verify() error {
	if c.ID == "" {
		return errors.New("storefront access token id is empty")
	}
	return nil
}

func (c DeleteStorefrontAccessTokenAPIReq) Endpoint() string {
	return fmt.Sprintf("storefront_access_tokens/%s.json", c.ID)
}

type DeleteStorefrontAccessTokenAPIResp struct {
	client.BaseAPIResponse
}
