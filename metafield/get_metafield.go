package metafield

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetMetafieldAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafields/resource-metafields/get-a-metafield-for-a-resource?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafields/resource-metafields/get-a-metafield-for-a-resource?version=v20251201
type GetMetafieldAPIReq struct {
	// enum：product、order、customer、collection、
	// shop、variant、draft_order
	// eg："product"
	OwnerResource string
	OwnerId       string
	ID            string `json:"id,omitempty"`
}

func (c GetMetafieldAPIReq) Verify() error {
	if c.OwnerId == "" {
		return errors.New("OwnerId is required")
	}

	if c.OwnerResource == "" {
		return errors.New("OwnerResource is required")
	}
	if c.ID == "" {
		return errors.New("ID is required")
	}

	return nil
}

func (c GetMetafieldAPIReq) Endpoint() string {
	return fmt.Sprintf("%s/%s/metafields/%s.json", c.OwnerResource, c.OwnerId, c.ID)
}

type GetMetafieldAPIResp struct {
	client.BaseAPIResponse
	Metafield Metafield `json:"metafield,omitempty"`
}
