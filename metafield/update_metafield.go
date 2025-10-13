package metafield

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// UpdateMetafieldAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafields/resource-metafields/update-a-metafield-for-a-resource?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafields/resource-metafields/update-a-metafield-for-a-resource?version=v20251201
type UpdateMetafieldAPIReq struct {
	Metafield UpdateMetafield `json:"metafield"`
}

func (c UpdateMetafieldAPIReq) Verify() error {
	if c.Metafield.OwnerId == "" {
		return errors.New("OwnerId is required")
	}

	if c.Metafield.OwnerResource == "" {
		return errors.New("OwnerResource is required")
	}
	if c.Metafield.Id == "" {
		return errors.New("Id is required")
	}

	return nil
}

func (c UpdateMetafieldAPIReq) Endpoint() string {
	return fmt.Sprintf("%s/%s/metafields/%s.json", c.Metafield.OwnerResource, c.Metafield.OwnerId, c.Metafield.Id)
}

type UpdateMetafieldAPIResp struct {
	Metafield Metafield `json:"metafield,omitempty"`
	client.BaseAPIResponse
}

type UpdateMetafield struct {
	// enum：product、order、customer、collection、
	// shop、variant、draft_order
	// eg："product"
	OwnerResource string
	OwnerId       string
	Id            string `json:"id,omitempty"`

	Description string `json:"description,omitempty"`

	// eg："product_warranty_period"、"customer_vip_level"
	Key       string `json:"key,omitempty"`
	Namespace string `json:"namespace,omitempty"`

	// eg："single_line_text_field"
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
