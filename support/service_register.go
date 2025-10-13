package support

import (
	"github.com/shoplineos/shopline-sdk-go/access"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/metafield"
	"github.com/shoplineos/shopline-sdk-go/order"
	"github.com/shoplineos/shopline-sdk-go/payment"
	"github.com/shoplineos/shopline-sdk-go/product"
	"github.com/shoplineos/shopline-sdk-go/store"
	"github.com/shoplineos/shopline-sdk-go/webhook"
)

// GetClientAwares Return client awares that you had register
func GetClientAwares() []client.Aware {
	var awares = []client.Aware{
		product.GetProductService(),
		order.GetOrderService(),
		order.GetOrderRiskService(),
		order.GetOrderRefundService(),
		metafield.GetMetafieldDefinitionService(),
		metafield.GetMetafieldService(),
		access.GetStorefrontAccessTokenService(),
		webhook.GetWebhookService(),
		store.GetStoreService(),
		payment.GetMerchantAppService(),
		payment.GetPaymentStoreService(),
		// you can add service here
	}
	return awares
}
