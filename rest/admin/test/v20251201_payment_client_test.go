package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/paymentsapp"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	// For test
	privateKeyStr = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQChLWJXdQp+jHuVwz/55pcXjeXIR3l8oYsBZCdNnLV8UlXN9c/cEdHafxmjK4tneY94yIv0y315VPMj41EDP/iYz/ciE02sujMS69uAsAeP6SioFV42Poyl933A/eow95UL2qDI2AoHtLzAnJk4bfl/iKfSP1bvkrWaR2zpV0jljCexwXa9yLv8MFXPsM2knfgQLqY67wrB6JMTbpDTMazz8BFmrdUP4+z2XWyTE6knLBi1hpuW5/NKP1I7D5v++8QX+BN5E0cooUdYuoCVqh+rOd8u96XVSK3HC8+pNYqYURZKVogt62QXdnKcSd5q34PqSCJzmsqf5cNbOqeijwi1AgMBAAECggEABqxb0nR5D5xhuVjSozy2u996Tu4vDtsCAWsXM0zK+Qxp2yoNJf0148N+uTI75oTSgCUVIfizGsz93fOsnpAT19qtqCNFxGJKnyl7sERGV+EtN5v4FEWUNxxVzSHZan2073JcO4/LhusNV2mhPWq7DKdiXWxI1u8x2WMRYXrliYPdq5BwI/0O297vEPJ+xZYe1vfZ0nc7wg/SbbYERbPDofeIgW4PspovVi1uSSgBg3n+90NSf5RjvbH5VeUxzeOniB75si8Vdvnx6+mP0ALyvr+EenqeQnYpUaqY2s7gi7KnmQyDYZ/Cw7kc0TR32jKowXln50scY1PWGP0V6Z0LiQKBgQDQZRk9HOnZmcW7fYaMyhKtBnboTb6RlmouDdz+j8eYI732+qTUjAbM6GtPvE5mJlO9ObMNK/s8ugb1Y++ZzZUKiOc/jYNgqT8tevLks2DlS0iNiPFtE7OjmA2sEk/zaoZqAscRIrmTravekLnc5OJpkOvsHGcjlO5ju7P4tbsdnQKBgQDF/wF+ZghZX8e47tQLP6nkMs0UMtbCMXjsbB82e58gy/28J/fus5XEvxJjA4wiUbZAnj/NQrYYT44HkwnKV6SkMSsZ6x7e5Had1Jn+3RYtarc0aVGE9xLfwQqvVSBEmLLRerfFX+5fbp6FKlM1nn0rxXpxRLt/JnTCmkluB7S3+QKBgEfq6+fcR5PR7pxCuKFzxzgxJ+4Jjn+90gzsudycD/ygMRm/7Axx+pLSjt4olUHJblK6S+F60Sxm4qnjADgq64mEL5IOK027etMeQB7PDNx0u6gkn3TOPMtzWRyOAUt28sY5CSwPuM2PPOYFOi9SShS2b8S/FJUB+7ctevGU/es9AoGBAKrom2Z7NrvHNMSKy9jF5KXJwEKuO7k3MTWLg0npXgvWajkPmzGeLSq+8GUtu7ooJJUUxOgurLbBfU1GfE4AZ2sf0h+2WFh4h3dn/GIGf81Gb8w7GRYYnF8u6EU+yvLLiJfQQW+Lhl00RHuYdGk1XMD63t2FQf/YtzMAMWBcIIApAoGATHeTyP12xFfL0+GIgpMMNKWxVkuwZdwqozVtVfp5vuNw/kkbM0dXohClg/CuMyRMp7Hh2h0fZPsGcwUPSEY3c2NF9vjGNR6/Rk4OyjeVSTZEsEZIfH7n1lkcbhbATQ65boWn/I3iiAASJX2CgkTmitO+fCwrbWkZ7H7CCrH9Aas=" // for test
	publicKeyStr  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoS1iV3UKfox7lcM/+eaXF43lyEd5fKGLAWQnTZy1fFJVzfXP3BHR2n8ZoyuLZ3mPeMiL9Mt9eVTzI+NRAz/4mM/3IhNNrLozEuvbgLAHj+koqBVeNj6Mpfd9wP3qMPeVC9qgyNgKB7S8wJyZOG35f4in0j9W75K1mkds6VdI5YwnscF2vci7/DBVz7DNpJ34EC6mOu8KweiTE26Q0zGs8/ARZq3VD+Ps9l1skxOpJywYtYablufzSj9SOw+b/vvEF/gTeRNHKKFHWLqAlaofqznfLvel1UitxwvPqTWKmFEWSlaILetkF3ZynEneat+D6kgic5rKn+XDWzqnoo8ItQIDAQAB"                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 // for test
)

var paymentClient *client.PaymentClient

func setupPaymentClient() {
	setup()
	paymentClient = client.NewPaymentClient(cli, privateKeyStr, publicKeyStr)
}

// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/payments-app-api/payment-successful-notice?version=v20260301
// En: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/payments-app-api/payment-successful-notice?version=v20260301
func TestCallPaymentSuccessfulNotify(t *testing.T) {
	setupPaymentClient()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payment/notify/1/m/pay.json", cli.StoreHandle, "payments_apps/openapi", cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	req := &paymentsapp.PaymentSuccessfulNoticeAPIReq{
		ChannelId:                 "1",
		PaymentMethod:             "m",
		Amount:                    "200",
		ChannelOrderTransactionId: "2",
		Currency:                  "USD",
		OrderTransactionId:        "2",
		Status:                    "2",
	}

	resp := req.NewAPIResp()
	err := paymentClient.Call(context.Background(), req, resp)
	if err != nil {
		t.Error(err)
	}

}

func TestCallPaymentSuccessfulNotifySignError(t *testing.T) {
	setupPaymentClient()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payment/notify/1/m/pay.json", cli.StoreHandle, "payments_apps/openapi", cli.ApiVersion),
		httpmock.NewStringResponder(400, `{"errors":"sign error"}`))

	req := &paymentsapp.PaymentSuccessfulNoticeAPIReq{
		ChannelId:                 "1",
		PaymentMethod:             "m",
		Amount:                    "200",
		ChannelOrderTransactionId: "2",
		Currency:                  "USD",
		OrderTransactionId:        "2",
		Status:                    "2",
	}

	resp := req.NewAPIResp()
	err := paymentClient.Call(context.Background(), req, resp)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Printf("err: %s/n", err.Error())
	assert.NotNil(t, err)
	assert.Equal(t, "sign error", err.Error())

}

// 中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/payments-app-api/refund-successful-notification?version=v20260301
// En: https://developer.shopline.com/docs/admin-rest-api/payments-app-api/refund-successful-notification?version=v20260301
func TestCallRefundSuccessfulNotify(t *testing.T) {
	setupPaymentClient()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payment/notify/1/m/refund.json", cli.StoreHandle, "payments_apps/openapi", cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	req := &paymentsapp.RefundSuccessfulNotificationAPIReq{
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
