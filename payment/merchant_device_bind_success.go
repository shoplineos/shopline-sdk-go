package payment

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// MerchantDeviceBindSuccessAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/payments-app-api/device-binding-success-notification?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/payments-app-api/device-binding-success-notification?version=v20251201
type MerchantDeviceBindSuccessAPIReq struct {
	client.BaseAPIRequest
	ResourceId       string `json:"resource_id,omitempty"`       // required
	BindingReference string `json:"binding_reference,omitempty"` // required
	Status           string `json:"status,omitempty"`            // required
	ConfigData       string `json:"config_data,omitempty"`
	Passthrough      string `json:"passthrough,omitempty"`
}

func (req *MerchantDeviceBindSuccessAPIReq) Method() string {
	return "POST"
}

func (req *MerchantDeviceBindSuccessAPIReq) Verify() error {
	// Verify the api request params
	if req.ResourceId == "" {
		return errors.New("ResourceId is empty")
	}
	if req.BindingReference == "" {
		return errors.New("BindingReference is empty")
	}
	if req.Status == "" {
		return errors.New("status is empty")
	}
	return nil
}

func (req *MerchantDeviceBindSuccessAPIReq) Endpoint() string {
	return "app/notify/device_bind.json"
}

type MerchantDeviceBindSuccessAPIResp struct {
	client.BaseAPIResponse
}
