package support

import (
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/access"
	metafield2 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/metafield"
	order2 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/order"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/product"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/store"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/webhook"
)

// GetClientAwares Return client awares that you had register
// Deprecated
func GetClientAwares() []client.Aware {
	var awares = []client.Aware{
		product.GetProductService(),
		order2.GetOrderService(),
		order2.GetOrderRiskService(),
		metafield2.GetMetafieldDefinitionService(),
		metafield2.GetMetafieldService(),
		access.GetStorefrontAccessTokenService(),
		webhook.GetWebhookService(),
		store.GetStoreService(),
		// you can add service here
	}
	return awares
}
