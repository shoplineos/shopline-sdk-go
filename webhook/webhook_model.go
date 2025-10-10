package webhook

type Webhook struct {
	Id         uint64 `json:"id,omitempty"`
	Address    string `json:"address,omitempty"`
	Topic      string `json:"topic,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	ApiVersion string `json:"api_version,omitempty"`
}
