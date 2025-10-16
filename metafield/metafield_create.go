package metafield

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// CreateMetafieldAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/create-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/create-a-metafield-definition?version=v20251201
type CreateMetafieldAPIReq struct {
	client.BaseAPIRequest
	Metafield CreateMetafield `json:"metafield"`
}

func (c *CreateMetafieldAPIReq) Method() string {
	return "POST"
}

func (c *CreateMetafieldAPIReq) Verify() error {
	if c.Metafield.OwnerId == "" {
		return errors.New("OwnerId is required")
	}

	if c.Metafield.OwnerResource == "" {
		return errors.New("OwnerResource is required")
	}

	return nil
}

func (c *CreateMetafieldAPIReq) Endpoint() string {
	return fmt.Sprintf("%s/%s/metafields.json", c.Metafield.OwnerResource, c.Metafield.OwnerId)
}

type CreateMetafieldAPIResp struct {
	Metafield Metafield `json:"metafield,omitempty"`
	client.BaseAPIResponse
}

type CreateMetafield struct {
	// enum：product、order、customer、collection、
	// shop、variant、draft_order
	// eg："product"
	OwnerResource string
	OwnerId       string

	Description string `json:"description,omitempty"`

	// eg："product_warranty_period"、"customer_vip_level"
	Key       string `json:"key,omitempty"`
	Namespace string `json:"namespace,omitempty"`

	// eg："single_line_text_field"
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
