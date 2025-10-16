package payment

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// MerchantBindSuccessAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/payments-app-api/merchant-activation-successful-notification?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/payments-app-api/merchant-activation-successful-notification?version=v20251201
type MerchantBindSuccessAPIReq struct {
	client.BaseAPIRequest

	Handle string `json:"handle,omitempty"`  // required
	AppKey string `json:"app_key,omitempty"` // required
}

func (req *MerchantBindSuccessAPIReq) Method() string {
	return "POST"
}

func (req *MerchantBindSuccessAPIReq) GetData() interface{} {
	return req
}

func (req *MerchantBindSuccessAPIReq) GetRequestOptions() *client.RequestOptions {
	opts := &client.RequestOptions{
		NotDecodeBody: true,
	}
	return opts
}

func (req *MerchantBindSuccessAPIReq) Verify() error {
	// Verify the api request params
	if req.Handle == "" {
		return errors.New("MerchantBindSuccessAPIReq.Handle is empty")
	}
	if req.AppKey == "" {
		return errors.New("MerchantBindSuccessAPIReq.AppKey is empty")
	}
	return nil
}

func (req *MerchantBindSuccessAPIReq) Endpoint() string {
	return "app/notify/bind.json"
}

type MerchantBindSuccessAPIResp struct {
	client.BaseAPIResponse
}
