package order

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListOrderAttributionInfosAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/get-order-attribution-in-bulk?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/order/order-management/get-order-attribution-in-bulk?version=v20251201
type ListOrderAttributionInfosAPIReq struct {
	OrderIDs []string `json:"orders,omitempty"`
}

func (r ListOrderAttributionInfosAPIReq) Verify() error {
	if r.OrderIDs == nil || len(r.OrderIDs) == 0 {
		return errors.New("OrderIDs is required")
	}
	return nil
}

func (r ListOrderAttributionInfosAPIReq) Endpoint() string {
	return "orders/order_attribution_info.json"
}

type ListOrderAttributionInfosAPIResp struct {
	client.BaseAPIResponse
	OrderAttributionInfos []OrderAttributionInfo `json:"data,omitempty"`
}

type OrderAttributionInfo struct {
	DaysForConversion uint32           `json:"days_for_conversion,omitempty"`
	OrderSeq          string           `json:"order_seq,omitempty"`
	OrderTime         string           `json:"order_time,omitempty"`
	TotalTraffics     uint32           `json:"total_traffics,omitempty"`
	FirstInteraction  FirstInteraction `json:"first_interaction,omitempty"`
	LastInteraction   LastInteraction  `json:"last_interaction,omitempty"`
}

type FirstInteraction struct {
	FirstInteractTime      string `json:"first_interact_time,omitempty"`
	FirstInteractionSource string `json:"first_interaction_source,omitempty"`
	FirstLandingPageUrl    string `json:"first_landing_page_url,omitempty"`
	FirstReferrerName      string `json:"first_referrer_name,omitempty"`
	FirstReferrerUrl       string `json:"first_referrer_url,omitempty"`

	FirstUtmParameters FirstUtmParameters `json:"first_utm_parameters,omitempty"`
}

type FirstUtmParameters struct {
	FirstUtmCampaign string `json:"first_utm_campaign,omitempty"`
	FirstUtmContent  string `json:"first_utm_content,omitempty"`
	FirstUtmMedium   string `json:"first_utm_medium,omitempty"`
	FirstUtmName     string `json:"first_utm_name,omitempty"`
	FirstUtmSource   string `json:"first_utm_source,omitempty"`
	FirstUtmTerm     string `json:"first_utm_term,omitempty"`
}

type LastInteraction struct {
	LastInteractTime      string `json:"last_interact_time,omitempty"`
	LastInteractionSource string `json:"last_interaction_source,omitempty"`
	LastLandingPageUrl    string `json:"last_landing_page_url,omitempty"`
	LastReferrerName      string `json:"last_referrer_name,omitempty"`
	LastReferrerUrl       string `json:"last_referrer_url,omitempty"`

	LastUtmParameters LastUtmParameters `json:"last_utm_parameters,omitempty"`
}

type LastUtmParameters struct {
	LastUtmCampaign string `json:"last_utm_campaign,omitempty"`
	LastUtmContent  string `json:"last_utm_content,omitempty"`
	LastUtmMedium   string `json:"last_utm_medium,omitempty"`
	LastUtmName     string `json:"last_utm_name,omitempty"`
	LastUtmSource   string `json:"last_utm_source,omitempty"`
	LastUtmTerm     string `json:"last_utm_term,omitempty"`
}
