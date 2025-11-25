package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	paymentsapp2 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/paymentsapp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantBindSuccess(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/app/notify/bind.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	apiReq := &paymentsapp2.MerchantActivationSuccessfulNotificationAPIReq{
		Handle: "product",
		AppKey: "appkey",
	}

	apiResp := &paymentsapp2.MerchantActivationSuccessfulNotificationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("Payment.MerchantBindSuccess returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

}

//func TestMerchantPaySuccess(t *testing.T) {
//	setup()
//	defer teardown()
//
//	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payment/notify/1/2/pay.json", cli.StoreHandle, "payments_apps/openapi", cli.ApiVersion),
//		httpmock.NewStringResponder(200, ""))
//
//	apiReq := &paymentsapp2.PaymentSuccessfulNoticeAPIReq{
//		ChannelId:                 "1",
//		PaymentMethod:             "2",
//		Amount:                    "11",
//		ChannelOrderTransactionId: "11",
//		Currency:                  "CNY",
//		OrderTransactionId:        "1111",
//		Status:                    "ok",
//	}
//
//	apiResp := &paymentsapp2.PaymentSuccessfulNoticeAPIResp{}
//	err := cli.Call(context.Background(), apiReq, apiResp)
//
//	if err != nil {
//		t.Errorf("Payment.MerchantPaySuccess returned error: %v", err)
//	}
//
//	assert.NotNil(t, apiResp)
//
//}

//func TestMerchantRefundSuccess(t *testing.T) {
//	setup()
//	defer teardown()
//
//	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payment/notify/1/2/refund.json", cli.StoreHandle, "payments_apps/openapi", cli.ApiVersion),
//		httpmock.NewStringResponder(200, ""))
//
//	apiReq := &paymentsapp2.RefundSuccessfulNotificationAPIReq{
//		ChannelId:                  "1",
//		PaymentMethod:              "2",
//		Amount:                     100,
//		ChannelOrderTransactionId:  "11",
//		ChannelRefundTransactionId: "11",
//		RefundTransactionId:        "1111",
//		Currency:                   "CNY",
//		Status:                     "ok",
//	}
//
//	apiResp := &paymentsapp2.RefundSuccessfulNotificationAPIResp{}
//	err := cli.Call(context.Background(), apiReq, apiResp)
//
//	if err != nil {
//		t.Errorf("Payment.MerchantRefundSuccess returned error: %v", err)
//	}
//
//	assert.NotNil(t, apiResp)
//}

func TestMerchantDeviceBindSuccess(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/app/notify/device_bind.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	apiReq := &paymentsapp2.DeviceBindingSuccessNotificationAPIReq{
		ResourceId:       "1",
		BindingReference: "2",
		ConfigData:       "2222",
		Passthrough:      "11",
		Status:           "ok",
	}

	apiResp := &paymentsapp2.DeviceBindingSuccessNotificationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("Payment.MerchantDeviceBindSuccess returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

}
