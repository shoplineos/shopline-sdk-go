package payment

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// MerchantRefundSuccessAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/payments-app-api/refund-successful-notification?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/payments-app-api/refund-successful-notification?version=v20251201
type MerchantRefundSuccessAPIReq struct {
	client.BaseAPIRequest
	ChannelId     string // required
	PaymentMethod string // required

	Amount                     uint64 `json:"amount,omitempty"`                        // required
	ChannelOrderTransactionId  string `json:"channel_order_transaction_id,omitempty"`  // required
	ChannelRefundTransactionId string `json:"channel_refund_transaction_id,omitempty"` // required
	Currency                   string `json:"currency,omitempty"`                      // required
	RefundTransactionId        string `json:"refund_transaction_id,omitempty"`         // required
	Status                     string `json:"status,omitempty"`                        // required
}

func (req *MerchantRefundSuccessAPIReq) Method() string {
	return "POST"
}

func (req *MerchantRefundSuccessAPIReq) GetData() interface{} {
	return req
}

func (req *MerchantRefundSuccessAPIReq) Verify() error {
	// Verify the api request params
	if req.ChannelId == "" {
		return errors.New("ChannelId is empty")
	}
	if req.PaymentMethod == "" {
		return errors.New("PaymentMethod is empty")
	}

	if req.Amount == 0 {
		return errors.New("amount is empty")
	}
	if req.ChannelOrderTransactionId == "" {
		return errors.New("channelOrderTransactionId is empty")
	}
	if req.ChannelRefundTransactionId == "" {
		return errors.New("channelRefundTransactionId is empty")
	}
	if req.Currency == "" {
		return errors.New("currency is empty")
	}
	if req.RefundTransactionId == "" {
		return errors.New("RefundTransactionId is empty")
	}
	if req.Status == "" {
		return errors.New("status is empty")
	}
	return nil
}

func (req *MerchantRefundSuccessAPIReq) Endpoint() string {
	return fmt.Sprintf("payment/notify/%s/%s/refund.json", req.ChannelId, req.PaymentMethod)
}

func (req *MerchantRefundSuccessAPIReq) GetRequestOptions() *client.RequestOptions {
	return &client.RequestOptions{
		NotDecodeBody: true,
	}
}

type MerchantRefundSuccessAPIResp struct {
	client.BaseAPIResponse
}
