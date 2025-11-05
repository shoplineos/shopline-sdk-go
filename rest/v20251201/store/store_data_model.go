package store

type Store struct {
	Id                   uint64            `json:"id"`                     // Store Id
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
	Id           string `json:"id"`            // Store Plan Id
	Name         string `json:"name"`          // Store Plan Name
	Type         string `json:"type"`          // Store Plan Type
	Price        string `json:"price"`         // Store Plan Price
	BillingCycle string `json:"billing_cycle"` // Store Plan Billing Cycle(month/year)
	Status       string `json:"status"`        // Store Plan Status
}

type StoreImage struct {
	Id        string `json:"id"`
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
	Id     string `json:"id"`     // Payment Gateway Id
	Name   string `json:"name"`   // Payment Gateway Name
	Status string `json:"status"` // Payment Gateway Status(active/inactive)
	Type   string `json:"type"`   // Payment Gateway Type
}

type ShippingCarrier struct {
	Id     string `json:"id"`     // Shipping Carrier Id
	Name   string `json:"name"`   // Shipping Carrier Name
	Status string `json:"status"` // Status(active/inactive)
}

type BusinessReg struct {
	RegisteredName    string `json:"registered_name"`    // Registered Name
	RegistrationNo    string `json:"registration_no"`    // Registration No
	RegisteredAddress string `json:"registered_address"` // Registered Address
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

type StoreStaff struct {
	AccountOwner bool   `json:"account_owner,omitempty"`
	Email        string `json:"email,omitempty"`
	Name         string `json:"name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	UID          string `json:"uid,omitempty"`
	UserType     string `json:"user_type,omitempty"`
}

type OperationLog struct {
	Author      string `json:"author,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	Detail      string `json:"detail,omitempty"`
	Id          string `json:"id,omitempty"`
	SubjectId   string `json:"subject_id,omitempty"`
	SubjectType string `json:"subject_type,omitempty"`
	Verb        string `json:"verb,omitempty"`
}
