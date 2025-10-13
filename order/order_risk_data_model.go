package order

type OrderRisk struct {
	Display        bool            `json:"display,omitempty"`
	CauseCancel    bool            `json:"cause_cancel,omitempty"`
	OrderID        string          `json:"order_id,omitempty"`
	Recommendation string          `json:"recommendation,omitempty"`
	RiskDetailMsg  []RiskDetailMsg `json:"risk_detail_msg_list,omitempty"`
	Source         string          `json:"source,omitempty"`
	ID             string          `json:"id,omitempty"`
	Score          string          `json:"score,omitempty"`
	CheckoutId     string          `json:"checkout_id,omitempty"`
}

type RiskDetailMsg struct {
	Message         string `json:"message,omitempty"`
	RiskDetailLevel string `json:"risk_detail_level,omitempty"`
}
