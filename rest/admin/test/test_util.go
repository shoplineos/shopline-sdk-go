package test

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/config"
	"os"
)

var (
	cli *client.Client
	app client.App
)

func GetClient() *client.Client {
	return cli
}

func GetApp() client.App {
	return app
}

func LoadTestDataFromCurrentDir(filename string) []byte {
	return LoadTestDataV2("", filename)
}

func LoadTestData(filename string) []byte {
	return LoadTestDataV2("../../test/", filename)
}

func LoadTestDataV2(dir string, filename string) []byte {
	f, err := os.ReadFile(dir + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load test data %v", filename))
	}
	return f
}

func Setup() {
	setup()
}

func Teardown() {
	teardown()
}

func setup() {
	app = client.App{
		AppKey:    config.AppKeyForUnitTest,
		AppSecret: config.AppSecretForUnitTest,
	}

	cli = client.MustNewClient(app, config.StoreHandelForUnitTest, config.AccessTokenForUnitTest)
	if cli == nil {
		panic("client is nil")
	}

	app.Client = cli

	httpmock.ActivateNonDefault(cli.Client)
}

func teardown() {
	httpmock.DeactivateAndReset()
}
