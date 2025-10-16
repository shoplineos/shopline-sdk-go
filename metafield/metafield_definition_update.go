package metafield

import "github.com/shoplineos/shopline-sdk-go/client"

// UpdateMetafieldDefinitionAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/update-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/update-a-metafield-definition?version=v20251201
type UpdateMetafieldDefinitionAPIReq struct {
	client.BaseAPIRequest
	MetafieldDefinition UpdateMetafieldDefinition `json:"definition"`
}

func (u *UpdateMetafieldDefinitionAPIReq) Method() string {
	return "PUT"
}

func (u *UpdateMetafieldDefinitionAPIReq) Verify() error {
	return nil
}

func (u *UpdateMetafieldDefinitionAPIReq) Endpoint() string {
	return "metafield_definition.json"
}

type UpdateMetafieldDefinitionAPIResp struct {
	client.BaseAPIResponse
	MetafieldDefinition MetafieldDefinition `json:"definition,omitempty"`
}

type UpdateMetafieldDefinition struct {
	Access Access `json:"access,omitempty"`

	Description string `json:"description,omitempty"`

	// eg："product_warranty_period"、"customer_vip_level"
	Key       string `json:"key,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`

	// enum：product、order、customer、collection、
	// shop、variant、draft_order
	// eg："product"
	OwnerResource string `json:"owner_resource,omitempty"`
}
