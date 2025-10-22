package support

import (
	"github.com/shoplineos/shopline-sdk-go/api/access"
	metafield2 "github.com/shoplineos/shopline-sdk-go/api/metafield"
	order2 "github.com/shoplineos/shopline-sdk-go/api/order"
	payment2 "github.com/shoplineos/shopline-sdk-go/api/payment"
	"github.com/shoplineos/shopline-sdk-go/api/product"
	"github.com/shoplineos/shopline-sdk-go/api/store"
	"github.com/shoplineos/shopline-sdk-go/api/webhook"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetClientAwares Return client awares that you had register
func GetClientAwares() []client.Aware {
	var awares = []client.Aware{
		product.GetProductService(),
		order2.GetOrderService(),
		order2.GetOrderRiskService(),
		order2.GetOrderRefundService(),
		metafield2.GetMetafieldDefinitionService(),
		metafield2.GetMetafieldService(),
		access.GetStorefrontAccessTokenService(),
		webhook.GetWebhookService(),
		store.GetStoreService(),
		payment2.GetMerchantAppService(),
		payment2.GetPaymentStoreService(),
		// you can add service here
	}
	return awares
}
