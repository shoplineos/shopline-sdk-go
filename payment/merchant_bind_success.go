package payment

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

type MerchantBindSuccessAPIReq struct {
	Handle string `json:"handle,omitempty"`  // required
	AppKey string `json:"app_key,omitempty"` // required
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
