package order

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"testing"
)

//
//func TestUpdateOrder(t *testing.T) {
//
//	apiReq := &UpdateOrderAPIReq{
//		Order: Order{
//			Id:          "21071580496925210798359834",
//			Name:        "D101-" + time.Now().Format("20060102"),
//			OrderNote:   "API Create test Order updatezw",
//			Currency:    "USD",
//			SendReceipt: false,
//			ProcessedAt: time.Now().Format("2006-01-02T15:04:05+08:00"),
//
//			PriceInfo: PriceInfo{
//				TotalShippingPrice:         "8.00",
//				CurrentExtraTotalDiscounts: "8.00",
//				TaxesIncluded:              false,
//			},
//
//			// line item
//			LineItems: []LineItem{
//				{
//					ProductId: "16071506529459141648923380", // Product Id
//					VariantId: "18071506529466020307563380", // Variant Id
//					Title:     "beautiful skirt",            // Product title
//					Price:     "3.25",                       // Price
//					Quantity:  1,                            // Quantity
//					Taxable:   false,
//					TaxLine: TaxLine{
//						Price: "3.25",
//						Rate:  "0.020",
//						Title: "Tax name",
//					},
//					DiscountPrice: Discount{
//						Title:  "Discount name",
//						Amount: "1.00",
//					},
//				},
//			},
//
//			// Transaction
//			TransactionList: []Transaction{
//				{
//					Amount:      "3.25",
//					Gateways:    "PayPal",
//					ProcessedAt: time.Now().Format("2006-01-02T15:04:05+08:00"),
//					Status:      "unpaid",
//				},
//			},
//
//			ShippingAddress: ShippingAddress{
//				FirstName:    "Tom",
//				LastName:     "Washington",
//				Phone:        "13903004000",
//				Email:        "test001@Gmail.com",
//				Country:      "China",
//				CountryCode:  "CN",
//				Province:     "Guangdong Province",
//				ProvinceCode: "4220006",
//				City:         "Guangzhou City",
//				CityCode:     "510000",
//				Area:         "Panyu District",
//				AreaCode:     "510006",
//				Address1:     "Xiaoguwei Street 11",
//				Address2:     "Apartment 5 11",
//				Zip:          "510036",
//				Latitude:     "43",
//				Longitude:    "34",
//			},
//
//			// Customer
//			Customer: Customer{
//				Id:        "4201057495",
//				FirstName: "Tom",
//				LastName:  "Washington",
//				Phone:     "13903004000",
//				AreaCode:  "+86",
//				Email:     "test001@Gmail.com",
//			},
//
//			// ShippingLine
//			ShippingLine: ShippingLine{
//				Code:  "SF",
//				Title: "Shipping name",
//				Price: "3.25",
//				TaxLine: TaxLine{
//					Price: "100",
//					Rate:  "0.020",
//					Title: "Tax name",
//				},
//			},
//
//			BillingAddress: BillingAddress{
//				SameAsReceiver: false,
//				FirstName:      "Tom",
//				LastName:       "Washington",
//				Phone:          "13903004000",
//				Email:          "test001@Gmail.com",
//				Country:        "China",
//				CountryCode:    "US",
//				Province:       "Guangdong Province",
//				ProvinceCode:   "4220006",
//				City:           "Guangzhou City",
//				CityCode:       "510000",
//				Area:           "Panyu District",
//				AreaCode:       "510007",
//				Address1:       "Xiaoguwei Street 11",
//				Address2:       "Apartment 5 11",
//				Zip:            "510036",
//			},
//		},
//	}
//
//	c := manager.GetDefaultClient()
//
//	shopLineResp, err := UpdateOrder(c, apiReq)
//	if err != nil {
//		fmt.Println("Create order failed, err:", err)
//		return
//	}
//
//	fmt.Printf("Create order successful！orderID: %s\n", shopLineResp.Order.Id)
//	a := assert.New(t)
//	a.NotEmpty(shopLineResp)
//
//}

func TestOrderUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"order":{"id": "1"}}`))

	order := Order{
		Id:                "1",
		FinancialStatus:   "paid",
		FulfillmentStatus: "fulfilled",
	}

	apiReq := &UpdateOrderAPIReq{Order: order}
	o, err := GetOrderService().Update(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Order.Update returned error: %v", err)
	}

	expected := Order{Id: "1"}
	if o.Order.Id != expected.Id {
		t.Errorf("Order.Update returned id %s, expected %s", o.Order.Id, expected.Id)
	}
}
