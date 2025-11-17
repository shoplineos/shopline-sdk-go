package order

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"testing"
)

//
//func TestCreateOrder(t *testing.T) {
//
//	apiReq := &CreateOrderAPIReq{
//		Order: Order{
//			Name:        "D101-" + time.Now().Format("20060102"),
//			OrderNote:   "API Create test Order",
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
//				Address1:     "Xiaoguwei Street",
//				Address2:     "Apartment 5",
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
//				Address1:       "Xiaoguwei Street",
//				Address2:       "Apartment 5",
//				Zip:            "510036",
//			},
//		},
//	}
//
//	c := manager.GetDefaultClient()
//
//	shopLineResp, err := CreateOrder(c, apiReq)
//	if err != nil {
//		fmt.Println("Create order failed, err:", err)
//	} else {
//		fmt.Printf("Create order successful！orderID: %s\n", shopLineResp.Order.Id)
//	}
//
//	a := assert.New(t)
//	a.NotEmpty(shopLineResp)
//}

func TestOrderCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"order":{"id": "1"}}`))

	order := Order{
		LineItems: []CreateAnOrderAPIRespLineItem{
			{
				VariantId: "1",
				Quantity:  1,
			},
		},
	}

	apiReq := &CreateOrderAPIReq{
		Order: order,
	}
	o, err := GetOrderService().Create(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Order.Create returned error: %v", err)
	}

	expected := Order{Id: "1"}
	if o.Order.Id != expected.Id {
		t.Errorf("Order.Create returned id %s, expected %s", o.Order.Id, expected.Id)
	}
}

func TestCreateOrder(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"order":{"id": "1"}}`))

	order := CreateAnOrderAPIReqOrder{
		LineItems: []LineItem{
			{
				VariantId: "1",
				Quantity:  1,
			},
		},
	}

	apiReq := &CreateAnOrderAPIReq{
		Order: order,
	}

	apiResp := &CreateAnOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("Order.Create returned error: %v", err)
	}

	expected := Order{Id: "1"}
	if apiResp.Order.Id != expected.Id {
		t.Errorf("Order.Create returned id %s, expected %s", apiResp.Order.Id, expected.Id)
	}
}
