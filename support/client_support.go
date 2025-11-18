package support

import (
	"github.com/shoplineos/shopline-sdk-go/client"
)

// MustNewClient Will apply awares.GetClientAwares()
func MustNewClient(app client.App, storeHandle, token string, opts ...client.Option) client.IClient {
	c := client.MustNewClientWithAwares(app, storeHandle, token, GetClientAwares(), opts...)
	return c
}
