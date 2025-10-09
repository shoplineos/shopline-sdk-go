package metafield

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// DeleteMetafieldAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafields/resource-metafields/delete-a-metafield?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafields/resource-metafields/delete-a-metafield?version=v20251201
type DeleteMetafieldAPIReq struct {
	// enum：product、order、customer、collection、
	// shop、variant、draft_order
	// eg："product"
	OwnerResource string
	OwnerId       string
	ID            string
}

func (c DeleteMetafieldAPIReq) Verify() error {
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

func (c DeleteMetafieldAPIReq) Endpoint() string {
	return fmt.Sprintf("%s/%s/metafields/%s.json", c.OwnerResource, c.OwnerId, c.ID)
}

type DeleteMetafieldAPIResp struct {
	client.BaseAPIResponse
}
