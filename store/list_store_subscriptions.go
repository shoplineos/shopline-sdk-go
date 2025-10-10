package store

import (
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListStoreSubscriptionsAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/subscribe/get-active-store-plans?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/store/subscribe/get-active-store-plans?version=v20251201
type ListStoreSubscriptionsAPIReq struct {
	IncludeTrial bool `url:"include_trial,omitempty"`
}

func (req *ListStoreSubscriptionsAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *ListStoreSubscriptionsAPIReq) Endpoint() string {
	return "store/subscription"
}

type ListStoreSubscriptionsAPIResp struct {
	client.BaseAPIResponse
	Subscriptions []Subscription `json:"data,omitempty"`

	// Deprecated
	Message string `json:"message,omitempty"`
}

type Subscription struct {
	AutoRecurring bool   `json:"auto_recurring,omitempty"`
	BillingCycle  string `json:"billing_cycle,omitempty"`

	// Example: {"MCC_packageVersion":"pe","SLP":"Premium"}
	BusinessParameters map[string]string `json:"business_parameters,omitempty"`

	// Deprecated
	CancelledAt uint64 `json:"cancelled_at,omitempty"`

	CreatedAt uint64 `json:"created_at,omitempty"`

	// Deprecated
	Enable           bool   `json:"enable,omitempty"`
	EndAt            uint64 `json:"end_at,omitempty"`
	ExtendPeroid     uint32 `json:"extend_peroid,omitempty"`
	GracePeriod      uint32 `json:"grace_period,omitempty"`
	GracePeriodEndAt uint64 `json:"grace_period_end_at,omitempty"`
	MerchantEmail    string `json:"merchant_email,omitempty"`
	MerchantId       uint64 `json:"merchant_id,omitempty"`

	// Deprecated
	NextRecurringAt uint64 `json:"next_recurring_at,omitempty"`

	// Deprecated
	PaymentMethod string `json:"payment_method,omitempty"`

	ProductLine string `json:"product_line,omitempty"`

	// {"en":"Starter","jp":"スターター","malay":"Versi pertama","zh-hans-cn":"入门版","zh-hant-tw":"入門版"}
	ProductName map[string]string `json:"product_name,omitempty"`

	// Deprecated
	Remarks     string `json:"remarks,omitempty"`
	StartAt     uint64 `json:"start_at,omitempty"`
	Status      string `json:"status,omitempty"`
	StoreHandle string `json:"store_handle,omitempty"`
	StoreId     uint64 `json:"store_id,omitempty"`
	SubId       string `json:"sub_id,omitempty"`
	Type        string `json:"type,omitempty"`
}
