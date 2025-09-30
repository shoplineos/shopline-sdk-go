package manager

import (
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/order"
	"github.com/shoplineos/shopline-sdk-go/product"
)

// GetClientAwares Return client awares that you had register
func GetClientAwares() []client.Aware {
	var awares = []client.Aware{
		product.GetProductService(),
		order.GetOrderService(),
		// you can add service here
	}
	return awares
}
