package access

type StorefrontAccessToken struct {

	// Example: unauthenticated_write_message, unauthenticated_read_message
	AccessScope string `json:"access_scope,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"` // ISO 8601
	ID          int64  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
}
