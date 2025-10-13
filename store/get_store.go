package store

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetStoreAPIReq
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/query-store-information?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/store/query-store-information?version=v20251201
type GetStoreAPIReq struct {
}

func (req *GetStoreAPIReq) Verify() error {
	// Verify the api request params
	return nil
}

func (req *GetStoreAPIReq) Endpoint() string {
	return "merchants/shop.json"
}

type GetStoreAPIResp struct {
	client.BaseAPIResponse
	Store Store `json:"data,omitempty"`
}

type Store struct {
	ID                   uint64            `json:"id"`                     // Store ID
	Name                 string            `json:"name"`                   // Store Name
	Domain               string            `json:"domain"`                 // Store Domain
	Handle               string            `json:"handle"`                 // Store Handle
	Currency             string            `json:"currency"`               // Currency
	Timezone             string            `json:"timezone"`               // Timezone
	CreatedAt            string            `json:"created_at"`             // Created Time(ISO 8601)
	UpdatedAt            string            `json:"updated_at"`             // Update Time(ISO 8601)
	ContactEmail         string            `json:"contact_email"`          // Contact Email
	ContactPhone         string            `json:"contact_phone"`          // Contact Phone
	Address              StoreAddress      `json:"address"`                // Store Address
	Language             string            `json:"language"`               // Store Language
	Status               string            `json:"status"`                 // Status(active/inactive)
	Plan                 StorePlan         `json:"plan"`                   // Store Plan
	Logo                 *StoreImage       `json:"logo,omitempty"`         // Store Logo
	Favicon              *StoreImage       `json:"favicon,omitempty"`      // Store Favicon
	SocialLinks          []SocialLink      `json:"social_links,omitempty"` // Social Links
	PaymentGateways      []PaymentGateway  `json:"payment_gateways"`       // Payment Gateway List
	ShippingCarriers     []ShippingCarrier `json:"shipping_carriers"`      // Shipping Carriers
	DefaultCountryCode   string            `json:"default_country_code"`   // Default Country Code
	DefaultProvinceCode  string            `json:"default_province_code"`  // Default Province Code
	HasPhysicalStore     bool              `json:"has_physical_store"`     // Has Physical Store
	BusinessRegistration BusinessReg       `json:"business_registration"`  // Business Registration
}

type StoreAddress struct {
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	City         string `json:"city"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	Zip          string `json:"zip"`
	CountryCode  string `json:"country_code"`
	ProvinceCode string `json:"province_code"`
}

type StorePlan struct {
	ID           string `json:"id"`            // Store Plan ID
	Name         string `json:"name"`          // Store Plan Name
	Type         string `json:"type"`          // Store Plan Type
	Price        string `json:"price"`         // Store Plan Price
	BillingCycle string `json:"billing_cycle"` // Store Plan Billing Cycle(month/year)
	Status       string `json:"status"`        // Store Plan Status
}

type StoreImage struct {
	ID        string `json:"id"`
	Src       string `json:"src"` // Image URL
	Alt       string `json:"alt"` // Image Alt
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	CreatedAt string `json:"created_at"`
}

type SocialLink struct {
	Platform string `json:"platform"` // Platform(facebook/instagram)
	URL      string `json:"url"`
}

type PaymentGateway struct {
	ID     string `json:"id"`     // Payment Gateway ID
	Name   string `json:"name"`   // Payment Gateway Name
	Status string `json:"status"` // Payment Gateway Status(active/inactive)
	Type   string `json:"type"`   // Payment Gateway Type
}

type ShippingCarrier struct {
	ID     string `json:"id"`     // Shipping Carrier ID
	Name   string `json:"name"`   // Shipping Carrier Name
	Status string `json:"status"` // Status(active/inactive)
}

type BusinessReg struct {
	RegisteredName    string `json:"registered_name"`    // Registered Name
	RegistrationNo    string `json:"registration_no"`    // Registration No
	RegisteredAddress string `json:"registered_address"` // Registered Address
}

// GetStoreInfo
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/store/query-store-information?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/store/query-store-information?version=v20251201
// Deprecated
// see StoreService
func GetStoreInfo(c *client.Client, apiReq *GetStoreAPIReq) (*GetStoreAPIResp, error) {

	// 1. API request
	shopLineReq := &client.ShopLineRequest{}

	// 2. API endpoint
	endpoint := apiReq.Endpoint()

	// 3. API response
	apiResp := &GetStoreAPIResp{}

	// 4. Call API
	_, err := c.Get(context.Background(), endpoint, shopLineReq, apiResp)
	//if err != nil {
	//	log.Printf("Failed to send request: %v\n", err)
	//	return nil, err
	//}

	return apiResp, err
}
