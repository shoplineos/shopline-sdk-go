package test

//
//var (
//	cli *client.Client
//	app client.App
//)
//
//func Setup() {
//	setup()
//}
//
//func Teardown() {
//	teardown()
//}
//
//func setup() {
//	app = client.App{
//		AppKey:    config.AppKeyForUnitTest,
//		AppSecret: config.AppSecretForUnitTest,
//	}
//
//	cli = client.MustNewClient(app, config.StoreHandelForUnitTest, config.AccessTokenForUnitTest)
//	if cli == nil {
//		panic("client is nil")
//	}
//
//	app.Client = cli
//
//	httpmock.ActivateNonDefault(cli.Client)
//}
//
//func teardown() {
//	httpmock.DeactivateAndReset()
//}
