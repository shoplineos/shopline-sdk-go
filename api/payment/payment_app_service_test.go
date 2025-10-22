package payment

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantBindSuccess(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/app/notify/bind.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	apiReq := &MerchantBindSuccessAPIReq{
		Handle: "product",
		AppKey: "appkey",
	}

	apiResp, err := GetMerchantAppService().MerchantBindSuccess(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Payment.MerchantBindSuccess returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

}

func TestMerchantPaySuccess(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payment/notify/1/2/pay.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	apiReq := &MerchantPaySuccessAPIReq{
		ChannelId:                 "1",
		PaymentMethod:             "2",
		Amount:                    "11",
		ChannelOrderTransactionId: "11",
		Currency:                  "CNY",
		OrderTransactionId:        "1111",
		Status:                    "ok",
	}

	apiResp, err := GetMerchantAppService().MerchantPaySuccess(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Payment.MerchantPaySuccess returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

}

func TestMerchantRefundSuccess(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payment/notify/1/2/refund.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	apiReq := &MerchantRefundSuccessAPIReq{
		ChannelId:                  "1",
		PaymentMethod:              "2",
		Amount:                     100,
		ChannelOrderTransactionId:  "11",
		ChannelRefundTransactionId: "11",
		RefundTransactionId:        "1111",
		Currency:                   "CNY",
		Status:                     "ok",
	}

	apiResp, err := GetMerchantAppService().MerchantRefundSuccess(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Payment.MerchantRefundSuccess returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
}

func TestMerchantDeviceBindSuccess(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/app/notify/device_bind.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	apiReq := &MerchantDeviceBindSuccessAPIReq{
		ResourceId:       "1",
		BindingReference: "2",
		ConfigData:       "2222",
		Passthrough:      "11",
		Status:           "ok",
	}

	apiResp, err := GetMerchantAppService().MerchantDeviceBindSuccess(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Payment.MerchantDeviceBindSuccess returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

}
