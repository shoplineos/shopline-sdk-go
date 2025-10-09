package metafield

type MetafieldDefinition struct {
	ID        string `json:"id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"` // Create time
	UpdatedAt string `json:"updated_at,omitempty"` // Updated time

	Access      Access `json:"access,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	// eg："product_warranty_period"、"customer_vip_level"
	Key string `json:"key,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	// enum：product、order、customer、collection、
	// shop、variant、draft_order
	// eg："product"
	OwnerResource string `json:"owner_resource,omitempty"`

	// eg："single_line_text_field"
	Type string `json:"type,omitempty"`
}

type Access struct {
	Admin string `json:"admin,omitempty"`
}
