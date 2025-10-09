package metafield

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CountMetafieldAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafields/resource-metafields/get-the-metafield-count-for-a-resource?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafields/resource-metafields/get-the-metafield-count-for-a-resource?version=v20251201
type CountMetafieldAPIReq struct {
	// enum：product、order、customer、collection、
	// shop、variant、draft_order
	// eg："product"
	OwnerResource string
	OwnerId       string
}

func (c CountMetafieldAPIReq) Verify() error {
	if c.OwnerId == "" {
		return errors.New("OwnerId is required")
	}

	if c.OwnerResource == "" {
		return errors.New("OwnerResource is required")
	}

	return nil
}

func (c CountMetafieldAPIReq) Endpoint() string {
	return fmt.Sprintf("%s/%s/metafields/count.json", c.OwnerResource, c.OwnerId)
}

type CountMetafieldAPIResp struct {
	Count int32 `json:"count,omitempty"`
	client.BaseAPIResponse
}
