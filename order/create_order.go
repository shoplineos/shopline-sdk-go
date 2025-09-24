package order

import (
	"context"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type CreateOrderAPIReq struct {
	Order Order `json:"order"`
}

// CreateOrderAPIResp Define the request structure for creating an order (corresponding to the API request body)
type CreateOrderAPIResp struct {
	Order Order `json:"order"`
	client.CommonAPIRespData
}

// Order Order structure
type Order struct {
	ID                      string          `json:"id"`                                   // Order ID
	Status                  string          `json:"status"`                               // Order status
	Name                    string          `json:"name,omitempty"`                       // Order name
	OrderNote               string          `json:"order_note,omitempty"`                 // Order Note
	PriceInfo               PriceInfo       `json:"price_info,omitempty"`                 // Order price info
	Transactions            Transactions    `json:"transactions,omitempty"`               // Transactions
	Currency                string          `json:"currency,omitempty"`                   // Currency（eg：USD）
	LineItems               []LineItem      `json:"line_items"`                           // Line Items
	SendFulfillmentReceipt  bool            `json:"send_fulfillment_receipt,omitempty"`   // SendFulfillmentReceipt
	CustomerLocale          string          `json:"customer_locale,omitempty"`            // Customer Locale
	SendReceipt             bool            `json:"send_receipt,omitempty"`               // Whether to send order confirmation email
	TransactionList         []Transaction   `json:"transaction_list,omitempty"`           // Transaction List
	InventoryBehaviour      string          `json:"inventory_behaviour,omitempty"`        // Inventory Behaviour
	MarketRegionCountryCode string          `json:"market_region_country_code,omitempty"` // MarketRegion CountryCode
	ShippingAddress         ShippingAddress `json:"shipping_address,omitempty"`           // Shipping Address
	NoteAttributes          []NoteAttribute `json:"note_attributes,omitempty"`            // Note Attributes
	ProcessedAt             string          `json:"processed_at,omitempty"`               // Order Processed time（ISO 8601）
	ShippingLine            ShippingLine    `json:"shipping_line,omitempty"`              // Shipping Line
	ExchangeRate            string          `json:"exchange_rate,omitempty"`              // Exchange Rate
	FulfillmentStatus       string          `json:"fulfillment_status,omitempty"`         // Fulfillment Status
	BuyerNote               string          `json:"buyer_note,omitempty"`                 // Buyer Note
	CompanyLocationID       string          `json:"company_location_id,omitempty"`        // Company Location ID
	Customer                Customer        `json:"customer,omitempty"`                   // Customer info
	FinancialStatus         string          `json:"financial_status,omitempty"`           // Financial Status
	Note                    string          `json:"note,omitempty"`                       // Note
	BillingAddress          BillingAddress  `json:"billing_address,omitempty"`            // Billing Address

	CreatedAt  string `json:"created_at"`          // Create time
	UpdatedAt  string `json:"updated_at"`          // Updated time
	TotalPrice string `json:"current_total_price"` // Total Price
}

// PriceInfo PriceInfo structure
type PriceInfo struct {
	TotalShippingPrice         string `json:"total_shipping_price,omitempty"`          // Total Shipping Price
	CurrentExtraTotalDiscounts string `json:"current_extra_total_discounts,omitempty"` // Current ExtraTotal Discounts
	TaxesIncluded              bool   `json:"taxes_included,omitempty"`                // Taxes Included
}

// Transactions Transactions structure
type Transactions struct {
	ID string `json:"id,omitempty"` // 交易ID
}

// LineItem Line Item structure
type LineItem struct {
	ProductID        string   `json:"product_id,omitempty"`        // Product ID
	TaxLine          TaxLine  `json:"tax_line,omitempty"`          // Tax Line
	Taxable          bool     `json:"taxable,omitempty"`           // Taxable
	Title            string   `json:"title,omitempty"`             // Title
	DiscountPrice    Discount `json:"discount_price,omitempty"`    // Discount Price
	Quantity         int      `json:"quantity"`                    // Quantity
	RequiresShipping bool     `json:"requires_shipping,omitempty"` // Requires Shipping
	VariantID        string   `json:"variant_id,omitempty"`        // Variant ID
	LocationID       string   `json:"location_id,omitempty"`       // Location ID
	Price            string   `json:"price"`                       // Price
}

// TaxLine Tax Line structure
type TaxLine struct {
	Price string `json:"price,omitempty"` // Price
	Rate  string `json:"rate,omitempty"`  // Rate
	Title string `json:"title,omitempty"` // Title
}

// Discount structure
type Discount struct {
	Title  string `json:"title,omitempty"`  // Title
	Amount string `json:"amount,omitempty"` // Amount
}

// Transaction Transaction
type Transaction struct {
	Amount      string `json:"amount,omitempty"`       // Amount
	Gateways    string `json:"gateways,omitempty"`     // Pay Gateways
	ProcessedAt string `json:"processed_at,omitempty"` // Processed Time
	Status      string `json:"status,omitempty"`       // Status
}

type ShippingAddress struct {
	Phone        string `json:"phone,omitempty"`         // Phone
	Email        string `json:"email,omitempty"`         // Email
	Country      string `json:"country,omitempty"`       // Country
	Zip          string `json:"zip,omitempty"`           // Zip
	Area         string `json:"area,omitempty"`          // Area
	Longitude    string `json:"longitude,omitempty"`     // Longitude
	Company      string `json:"company,omitempty"`       // Company
	FirstName    string `json:"first_name,omitempty"`    // First Name
	Address1     string `json:"address1,omitempty"`      // Address1
	Latitude     string `json:"latitude,omitempty"`      // Latitude
	City         string `json:"city,omitempty"`          // City
	CountryCode  string `json:"country_code,omitempty"`  // Country Code
	Province     string `json:"province,omitempty"`      // Province
	Address2     string `json:"address2,omitempty"`      // Address2
	LastName     string `json:"last_name,omitempty"`     // Last Name
	ProvinceCode string `json:"province_code,omitempty"` // Province Code
	CityCode     string `json:"city_code,omitempty"`     // City Code
	AreaCode     string `json:"area_code,omitempty"`     // Area Code
}

type NoteAttribute struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}

// ShippingLine Shipping Line
type ShippingLine struct {
	Code    string  `json:"code,omitempty"`     // Code
	Price   string  `json:"price,omitempty"`    // Price
	TaxLine TaxLine `json:"tax_line,omitempty"` // Tax Line
	Title   string  `json:"title,omitempty"`    // Title
}

type Customer struct {
	ID        string `json:"id,omitempty"`         // ID
	LastName  string `json:"last_name,omitempty"`  // Last Name
	Phone     string `json:"phone,omitempty"`      // Phone
	AreaCode  string `json:"area_code,omitempty"`  // Area Code
	Email     string `json:"email,omitempty"`      // Email
	FirstName string `json:"first_name,omitempty"` // First Name
}

type BillingAddress struct {
	CityCode       string `json:"city_code,omitempty"`        // City Code
	SameAsReceiver bool   `json:"same_as_receiver,omitempty"` // Is Same As Receiver
	FirstName      string `json:"first_name,omitempty"`       // First Name
	Province       string `json:"province,omitempty"`         // Province
	Area           string `json:"area,omitempty"`             // Area
	Phone          string `json:"phone,omitempty"`            // Phone
	ProvinceCode   string `json:"province_code,omitempty"`    // Province Code
	Company        string `json:"company,omitempty"`          // Company
	CountryCode    string `json:"country_code,omitempty"`     // Country Code
	AreaCode       string `json:"area_code,omitempty"`        // Area Code
	Zip            string `json:"zip,omitempty"`              // Zip
	Address1       string `json:"address1,omitempty"`         // Address1
	LastName       string `json:"last_name,omitempty"`        // Last Name
	Address2       string `json:"address2,omitempty"`         // Address2
	City           string `json:"city,omitempty"`             // City
	Country        string `json:"country,omitempty"`          // Country
	Email          string `json:"email,omitempty"`            // Email
}

// CreateOrder
// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/order/order-management/create-an-order?version=v20251201
// en: https://developer.shopline.com/docs/admin-rest-api/order/order-management/create-an-order?version=v20251201
func CreateOrder(c *client.Client, apiReq *CreateOrderAPIReq) (*CreateOrderAPIResp, error) {
	// 1. API request
	shopLineReq := &client.ShopLineRequest{
		Data: apiReq, // API request data
	}

	// 2. API endpoint
	endpoint := "orders.json"

	// 3. API response data
	apiResp := &CreateOrderAPIResp{}

	// 4. Invoke API
	_, err := c.Post(context.Background(), endpoint, shopLineReq, apiResp)
	if err != nil {
		fmt.Printf("Execute Request failed，endpoint:%s, shopLineReq: %v, err: %v\n", endpoint, shopLineReq, err)
		return nil, err
	}

	return apiResp, nil
}
