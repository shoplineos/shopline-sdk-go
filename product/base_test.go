package product

import (
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/config"
)

var (
	c   *client.Client
	app client.App
)

func setup() {
	app = client.App{
		AppKey:    config.AppKeyForUnitTest,
		AppSecret: config.AppSecretForUnitTest,
	}

	c = client.MustNewClient(app, config.StoreHandelForUnitTest, config.AccessTokenForUnitTest)
	if c == nil {
		panic("client is nil")
	}

	app.Client = c

	httpmock.ActivateNonDefault(c.Client)
}

func teardown() {
	httpmock.DeactivateAndReset()
}
