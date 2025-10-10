package support

import (
	"github.com/shoplineos/shopline-sdk-go/access"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/metafield"
	"github.com/shoplineos/shopline-sdk-go/order"
	"github.com/shoplineos/shopline-sdk-go/product"
)

// GetClientAwares Return client awares that you had register
func GetClientAwares() []client.Aware {
	var awares = []client.Aware{
		product.GetProductService(),
		order.GetOrderService(),
		metafield.GetMetafieldDefinitionService(),
		metafield.GetMetafieldService(),
		access.GetStorefrontAccessTokenService(),
		// you can add service here
	}
	return awares
}
