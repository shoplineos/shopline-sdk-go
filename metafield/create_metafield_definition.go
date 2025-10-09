package metafield

// CreateMetafieldDefinitionAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/create-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/create-a-metafield-definition?version=v20251201
type CreateMetafieldDefinitionAPIReq struct {
	MetafieldDefinition CreateMetafieldDefinition `json:"definition"`
}

type CreateMetafieldDefinitionAPIResp struct {
	MetafieldDefinition MetafieldDefinition `json:"definition"`
}

type CreateMetafieldDefinition struct {
	//ID        string `json:"id,omitempty"`
	//CreatedAt string `json:"created_at,omitempty"` // Create time
	//UpdatedAt string `json:"updated_at,omitempty"` // Updated time

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

	// eg："single_line_text_field"
	Type string `json:"type,omitempty"`
}
