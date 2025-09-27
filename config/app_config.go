package config

const (

	// DefaultRedirectUrl TODO replace real data
	DefaultRedirectUrl = "http://appdemo.myshopline.com/auth/callback" // DefaultRedirectUrl, replace real DefaultRedirectUrl

)

const (
	DefaultAppKey    = ""                                                      // DefaultAppKey, replace real AppKey
	DefaultAppSecret = ""                                                      // DefaultAppSecret, replace real AppSecret
	DefaultAppScope  = "read_products,write_products,read_orders,write_orders" // DefaultAppScope, replace real AppScope

	// DefaultAPIVersion see
	// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/api-versioning-guide?version=v20260301
	// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/api-versioning-guide?version=v20260301
	DefaultAPIVersion = "v20251201" // replace real API Version

	DefaultStoreHandle = "zwapptest" // replace real store handle

	// DefaultAccessToken for test
	DefaultAccessToken = ""

	DefaultApiPathPrefix = "admin/openapi"
)
const (
	UserAgent = "shopline-sdk-go/0.0.10"
)
