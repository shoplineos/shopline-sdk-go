package support

import (
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/v20251201/access"
	"github.com/shoplineos/shopline-sdk-go/rest/v20251201/metafield"
	"github.com/shoplineos/shopline-sdk-go/rest/v20251201/order"
	"github.com/shoplineos/shopline-sdk-go/rest/v20251201/product"
	"github.com/shoplineos/shopline-sdk-go/rest/v20251201/store"
	"github.com/shoplineos/shopline-sdk-go/rest/v20251201/webhook"
)

// GetClientAwares Return client awares that you had register
// Deprecated
func GetClientAwares() []client.Aware {
	var awares = []client.Aware{
		product.GetProductService(),
		order.GetOrderService(),
		order.GetOrderRiskService(),
		metafield.GetMetafieldDefinitionService(),
		metafield.GetMetafieldService(),
		access.GetStorefrontAccessTokenService(),
		webhook.GetWebhookService(),
		store.GetStoreService(),
		// you can add service here
	}
	return awares
}
