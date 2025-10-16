package payment

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// MerchantPaySuccessAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/payments-app-api/payment-successful-notice?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/payments-app-api/payment-successful-notice?version=v20251201
type MerchantPaySuccessAPIReq struct {
	client.BaseAPIRequest
	ChannelId     string // required
	PaymentMethod string // required

	Amount                    string `json:"amount,omitempty"`                       // required
	ChannelOrderTransactionId string `json:"channel_order_transaction_id,omitempty"` // required
	Currency                  string `json:"currency,omitempty"`                     // required
	OrderTransactionId        string `json:"order_transaction_id,omitempty"`         // required
	Status                    string `json:"status,omitempty"`                       // required
}

func (req *MerchantPaySuccessAPIReq) Method() string {
	return "POST"
}

func (req *MerchantPaySuccessAPIReq) Verify() error {
	// Verify the api request params
	if req.ChannelId == "" {
		return errors.New("ChannelId is empty")
	}
	if req.PaymentMethod == "" {
		return errors.New("PaymentMethod is empty")
	}
	if req.Amount == "" {
		return errors.New("amount is empty")
	}
	if req.ChannelOrderTransactionId == "" {
		return errors.New("channelOrderTransactionId is empty")
	}
	if req.Currency == "" {
		return errors.New("currency is empty")
	}
	if req.OrderTransactionId == "" {
		return errors.New("OrderTransactionId is empty")
	}
	if req.Status == "" {
		return errors.New("status is empty")
	}
	return nil
}

func (req *MerchantPaySuccessAPIReq) Endpoint() string {
	return fmt.Sprintf("payment/notify/%s/%s/pay.json", req.ChannelId, req.PaymentMethod)
}

type MerchantPaySuccessAPIResp struct {
	client.BaseAPIResponse
}
