package payment

import (
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/config"
)

var (
	cli *client.Client
	app client.App
)

func setup() {
	app = client.App{
		AppKey:    config.AppKeyForUnitTest,
		AppSecret: config.AppSecretForUnitTest,
	}

	cli = client.MustNewClient(app, config.StoreHandelForUnitTest, config.AccessTokenForUnitTest, client.WithClientAwares([]client.Aware{
		GetMerchantAppService(),
		GetPaymentStoreService(),
	}))
	if cli == nil {
		panic("client is nil")
	}

	app.Client = cli

	httpmock.ActivateNonDefault(cli.Client)
}

func teardown() {
	httpmock.DeactivateAndReset()
}
