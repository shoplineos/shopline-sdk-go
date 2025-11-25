package client

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var paymentClient *PaymentClient

func setupPaymentClient() {
	setup()
	paymentClient = NewPaymentClient(client, privateKeyStr, publicKeyStr)
}

type PaymentSuccessfulNoticeAPIReqStruct struct {
	BaseAPIRequest
	ChannelId                 string `url:"-" json:"-"` // Required
	PaymentMethod             string `url:"-" json:"-"` // Required
	Amount                    string `json:"amount,omitempty" url:"-"`
	ChannelOrderTransactionId string `json:"channel_order_transaction_id,omitempty" url:"-"`
	Currency                  string `json:"currency,omitempty" url:"-"`
	OrderTransactionId        string `json:"order_transaction_id,omitempty" url:"-"`
	Status                    string `json:"status,omitempty" url:"-"`
}

func (req *PaymentSuccessfulNoticeAPIReqStruct) GetEndpoint() string {
	return fmt.Sprintf("payment/notify/%s/%s/pay.json", req.ChannelId, req.PaymentMethod)
}

func (req *PaymentSuccessfulNoticeAPIReqStruct) GetMethod() string {
	return "POST"
}

func (req *PaymentSuccessfulNoticeAPIReqStruct) GetData() interface{} {
	return req
}

// Verify the api request parameters
func (req *PaymentSuccessfulNoticeAPIReqStruct) Verify() error {
	return nil
}

func (req *PaymentSuccessfulNoticeAPIReqStruct) GetRequestOptions() *RequestOptions {
	opts := &RequestOptions{
		NotDecodeBody: true,
		EnableSign:    true,
		PathPrefix:    "payments_apps/openapi",
	}
	return opts
}

type PaymentSuccessfulNoticeAPIRespStruct struct {
	BaseAPIResponse
}

// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/payments-app-api/payment-successful-notice?version=v20260301
// En: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/payments-app-api/payment-successful-notice?version=v20260301
func TestCallPaymentSuccessfulNotify(t *testing.T) {
	setupPaymentClient()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payment/notify/1/m/pay.json", client.StoreHandle, "payments_apps/openapi", client.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	req := &PaymentSuccessfulNoticeAPIReqStruct{
		ChannelId:                 "1",
		PaymentMethod:             "m",
		Amount:                    "200",
		ChannelOrderTransactionId: "2",
		Currency:                  "USD",
		OrderTransactionId:        "2",
		Status:                    "2",
	}

	resp := &PaymentSuccessfulNoticeAPIRespStruct{}
	err := paymentClient.Call(context.Background(), req, resp)
	if err != nil {
		t.Error(err)
	}

}

func TestCallPaymentSuccessfulNotifySignError(t *testing.T) {
	setupPaymentClient()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payment/notify/1/m/pay.json", client.StoreHandle, "payments_apps/openapi", client.ApiVersion),
		httpmock.NewStringResponder(400, `{"errors":"sign error"}`))

	req := &PaymentSuccessfulNoticeAPIReqStruct{
		ChannelId:                 "1",
		PaymentMethod:             "m",
		Amount:                    "200",
		ChannelOrderTransactionId: "2",
		Currency:                  "USD",
		OrderTransactionId:        "2",
		Status:                    "2",
	}

	resp := &PaymentSuccessfulNoticeAPIRespStruct{}
	err := paymentClient.Call(context.Background(), req, resp)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Printf("err: %s/n", err.Error())
	assert.NotNil(t, err)
	assert.Equal(t, "sign error", err.Error())
}

type RefundSuccessfulNotificationAPIReqStruct struct {
	BaseAPIRequest
	ChannelId                  string `url:"-" json:"-"` // Required
	PaymentMethod              string `url:"-" json:"-"` // Required
	Currency                   string `json:"currency,omitempty" url:"-"`
	RefundTransactionId        string `json:"refund_transaction_id,omitempty" url:"-"`
	Status                     string `json:"status,omitempty" url:"-"`
	Amount                     int64  `json:"amount,omitempty" url:"-"`
	ChannelOrderTransactionId  string `json:"channel_order_transaction_id,omitempty" url:"-"`
	ChannelRefundTransactionId string `json:"channel_refund_transaction_id,omitempty" url:"-"`
}

func (req *RefundSuccessfulNotificationAPIReqStruct) GetEndpoint() string {
	return fmt.Sprintf("payment/notify/%s/%s/refund.json", req.ChannelId, req.PaymentMethod)
}

func (req *RefundSuccessfulNotificationAPIReqStruct) GetMethod() string {
	return "POST"
}

func (req *RefundSuccessfulNotificationAPIReqStruct) GetData() interface{} {
	return req
}

// Verify the api request parameters
func (req *RefundSuccessfulNotificationAPIReqStruct) Verify() error {
	return nil
}

func (req *RefundSuccessfulNotificationAPIReqStruct) GetRequestOptions() *RequestOptions {
	opts := &RequestOptions{
		NotDecodeBody: true,
		EnableSign:    true,
		PathPrefix:    "payments_apps/openapi",
	}
	return opts
}

func (req *RefundSuccessfulNotificationAPIReqStruct) NewAPIResp() *RefundSuccessfulNotificationAPIRespStruct {
	return &RefundSuccessfulNotificationAPIRespStruct{}
}

type RefundSuccessfulNotificationAPIRespStruct struct {
	BaseAPIResponse
}

// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/payments-app-api/refund-successful-notification?version=v20260301
// En: https://developer.shopline.com/docs/admin-rest-api/payments-app-api/refund-successful-notification?version=v20260301
func TestCallRefundSuccessfulNotify(t *testing.T) {
	setupPaymentClient()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payment/notify/1/m/refund.json", client.StoreHandle, "payments_apps/openapi", client.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	req := &RefundSuccessfulNotificationAPIReqStruct{
		ChannelId:                  "1",
		PaymentMethod:              "m",
		Amount:                     200,
		ChannelOrderTransactionId:  "2",
		ChannelRefundTransactionId: "2",
		Currency:                   "USD",
		RefundTransactionId:        "2",
		Status:                     "2",
	}

	resp := req.NewAPIResp()
	err := paymentClient.Call(context.Background(), req, resp)
	if err != nil {
		t.Error(err)
	}
}
