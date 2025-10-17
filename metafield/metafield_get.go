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
	client.BaseAPIRequest
	// enum：product、order、customer、collection、
	// shop、variant、draft_order
	// eg："product"
	OwnerResource string
	OwnerId       string
	Id            string `json:"id,omitempty"`
}

func (c *GetMetafieldAPIReq) GetMethod() string {
	return "GET"
}

func (c *GetMetafieldAPIReq) GetQuery() interface{} {
	return c
}

func (c *GetMetafieldAPIReq) Verify() error {
	if c.OwnerId == "" {
		return errors.New("OwnerId is required")
	}

	if c.OwnerResource == "" {
		return errors.New("OwnerResource is required")
	}
	if c.Id == "" {
		return errors.New("Id is required")
	}

	return nil
}

func (c *GetMetafieldAPIReq) GetEndpoint() string {
	return fmt.Sprintf("%s/%s/metafields/%s.json", c.OwnerResource, c.OwnerId, c.Id)
}

type GetMetafieldAPIResp struct {
	client.BaseAPIResponse
	Metafield Metafield `json:"metafield,omitempty"`
}
