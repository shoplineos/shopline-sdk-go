# shopline sdk go

## Getting Started

[中文](#zh)、[English](#en)

### <span id="en">English</span>

**Note**: This is an unstable SDK for developers using https://developer.shopline.com, we are still improving. The library does not have implementations of all shopline resources. PRs for new resources and endpoints are welcome, or you can simply implement some yourself as-you-go. See the section "Using your own models" for more info.



#### Install
```
$ go get github.com/shoplineos/shopline-sdk-go
```

#### Use
```
import "github.com/shoplineos/shopline-sdk-go/client"
```


#### Init App and Client
```
  // 1. create app
  appInstance := client.App{
      AppKey:      "",              // replace your data
      AppSecret:   "",              // replace your data
      Scope:       "read_products,write_products,read_orders,write_orders", // replace your data
      RedirectUrl: "http://appdemo.myshopline.com/auth/callback",           // replace your data
  }
  
  handle := "zwapptest" // replace your data
  accessToken := "" // replace your data
  
  // 2. create client
  // use support to create client, it will regitster awares
  c := support.MustNewClient(appInstance, handle, accessToken)
  appInstance.Client = c
  
    
  // 3. use client to call API
  // 3.1 API request
  getProductCountAPIReq := &GetProductCountAPIReq{}
  shoplineReq := &client.ShopLineRequest{
      Query: getProductCountAPIReq,
  }

  // 3.2 API endpoint
  endpoint := getProductCountAPIReq.Endpoint()

  // 3.3 API response
  apiResp := &GetProductCountAPIResp{}

  // 3.4 Call API
  shoplineResp, err := c.Get(context.Background(), endpoint, shoplineReq, apiResp)
  fmt.Printf("count:%d", apiResp.Count)
```

#### OAuth

If you don't have an access token yet, you can obtain one with the oauth flow. Something like this will work:

```
// see server/main.go
// Create an oauth-authorize url for the App and redirect to it.
// In some request handler, you probably want something like this:
func InstallHandler(w http.ResponseWriter, r *http.Request) {
    appkey := r.URL.Query().Get("appkey")
    handle := r.URL.Query().Get("handle")
    lang := r.URL.Query().Get("lang")
    timestampStr := r.URL.Query().Get("timestamp")
    sign := r.URL.Query().Get("sign")
    
    app := manager.GetApp(appkey)
    // Verify the Sign
    isSignValid := app.VerifySign(r.URL.Query(), sign)
    if !isSignValid {
        log.Printf("sign verification failed, appkey: %s, sign: %s\n", appkey, sign)
        http.Error(w, "Invalid signature", http.StatusUnauthorized)
        return
    }
    
  
    url, err := app.AuthorizeUrl(handle, "")
    if err != nil {
        log.Printf("Authorize url error, appkey: %s, handle: %s, err: %v\n", appkey, handle, err)
        http.Error(w, "server error", http.StatusInternalServerError)
        return
    }
    
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Fetch a access token in the callback
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
    appkey := r.URL.Query().Get("appkey")
    code := r.URL.Query().Get("code")
    handle := r.URL.Query().Get("handle")
    customField := r.URL.Query().Get("customField")
    timestampStr := r.URL.Query().Get("timestamp")
    sign := r.URL.Query().Get("sign")
    
    
    app := manager.GetApp(appkey)
    // Check that the callback signature is valid
    isSignValid := app.VerifySign(r.URL.Query(), sign)
    if isSignValid {
        log.Println("sign verified successfully")
    } else {
        log.Printf("sign verification failed, sign=%s\n", sign)
        return
    }
    
    // create token
    token, err := app.CreateAccessToken(context.Background(), code)
    
    // Do something with the token, like store it in a DB or Cache.
}
```

#### API calls with an access token

With an access token, you can make API calls like this:

Get Product Count:
``` Get Product Count
  apiReq := &GetProductCountAPIReq{}
  apiResp, err := product.GetProductService().Count(context.Background(), apiReq)
  fmt.Printf("count:%d", apiResp.Count)
```

Query Products:
```Query Products
  requestParams := &ListProductsAPIReq{
      // IDs: "1,2,3",
  }
  productsAPIResp, err := product.GetProductService().List(context.Background(), requestParams)
```


Pagination:
```Query Products
  requestParams := &ListProductsAPIReq{
      // IDs: "1,2,3",
  }
  productsAPIResp, err := product.GetProductService().ListWithPagination(context.Background(), requestParams)
```

Query all products:
``` query all products
  requestParams := &ListProductsAPIReq{}
  Products, err := product.GetProductService().ListAll(context.Background(), requestParams)
  
```

Create product:
``` create product
// see create_product_test.go
// create product
// https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product/?version=v20251201
apiReq := &product.CreateProductAPIReq{
    Product: Product{
        Title:          "Test product - " + time.Now().Format("20060102150405"),
        BodyHTML:       "<p>This is a test product created via the API</p>",
        Subtitle:       "Limited time offer",
        Vendor:         "Test provider",
        Status:         "active",
        PublishedScope: "web",
        Tags:           []string{"Test", "New", "API Create"},

        // Options（Color、Size）
        Options: []Option{
            {Name: "color", Values: []string{"red", "blue"}},
            {Name: "size", Values: []string{"S", "M", "L"}},
        },

        Images: []Image{
            {Src: "https://example.com/product-main.jpg", Alt: "Main picture"},
            {Src: "https://example.com/product-detail.jpg", Alt: "Detail picture"},
        },

        // Variants（red S、red M、blue S）
        Variants: []Variant{
            {
                SKU:            "RED-S-001",
                Price:          "99.99",
                CompareAtPrice: "129.99",
                Option1:        "red",
                Option2:        "S",
                Weight:         "0.5",
                WeightUnit:     "kg",
                Taxable:        true,
            },
            {
                SKU:            "RED-M-002",
                Price:          "109.99",
                CompareAtPrice: "139.99",
                Option1:        "red",
                Option2:        "M",
                Weight:         "0.6",
                WeightUnit:     "kg",
                Taxable:        true,
            },
            {
                SKU:            "BLUE-S-003",
                Price:          "99.99",
                CompareAtPrice: "129.99",
                Option1:        "blue",
                Option2:        "S",
                Weight:         "0.5",
                WeightUnit:     "kg",
                Taxable:        true,
            },
        },
    },
}

apiResp, err := product.GetProductService().Create(context.Background(), apiReq)

```

#### Using your own models

Not all endpoints are implemented right now. In those case, feel free to implement them and make a Pull Request, or you
can create your own struct for the data and use the client to call APIs. This is how the existing endpoints are
implemented.

For example, let's say you want to fetch product count. There's a helper function Get specifically for fetching stuff so
this will work:

```
// see get_product_count.go
type GetProductCountAPIReq struct {
	Status       string `url:"status,omitempty"`
	CollectionId string `url:"collection_id,omitempty"`
	CreatedAtMin string `url:"created_at_min,omitempty"` // Minimum order creation time（ISO 8601）
	CreatedAtMax string `url:"created_at_max,omitempty"` // Max order creation time（ISO 8601）
	UpdatedAtMin string `url:"updated_at_min,omitempty"` // Minimum order update time（ISO 8601）
	UpdatedAtMax string `url:"updated_at_max,omitempty"` // Max order update time（ISO 8601）
}

func (req *GetProductCountAPIReq) Verify() error {
	// Verify the API request params
	return nil
}

func (req *GetProductCountAPIReq) Endpoint() string {
	endpoint := "products/count.json"
	return endpoint
}

type GetProductCountAPIResp struct {
	Count int `json:"count"`
	client.BaseAPIResponse
}

func GetProductsCount(c *client.Client, apiReq *GetProductCountAPIReq) (*GetProductCountAPIResp, error) {

    // 1. API request
    shoplineReq := &client.ShopLineRequest{
        Query: apiReq, // http url query params
        // Data: apiReq, // http body params
    }
    
    // 2. API endpoint
    endpoint := apiReq.Endpoint()
    
    // 3. API response
    apiResp := &GetProductCountAPIResp{}
    
    // 4. Call API
    shoplineResp, err := c.Get(context.Background(), endpoint, shoplineReq, apiResp)
    
    // option
    // apiResp.TraceId = shoplineResp.TraceId
    
    return apiResp, err
}

```


#### Use your own Service Interface
* step1:Define a Service Interface
  * See product_service.go or order_service.go
  ```
    type IOrderService interface {
        List(context.Context, *ListOrdersAPIReq) (*ListOrdersAPIResp, error)
        ListAll(context.Context, *ListOrdersAPIReq) ([]Order, error)
        ...
    }
  ```
* step2:Define a Service struct and implements the Service Interface
  ```
    type OrderService struct {
      client.BaseService
    }
  
    func (o *OrderService) List(ctx context.Context, apiReq *ListOrdersAPIReq) (*ListOrdersAPIResp, error) {
      ...
    }
  
    func (o *OrderService) ListAll(ctx context.Context, apiReq *ListOrdersAPIReq) (*ListOrdersAPIResp, error) {
      ...
    }

  ```
* step3:Create a Service Instance
  ```
    var serviceInst = &OrderService{}
    func GetOrderService() *OrderService {
      return serviceInst
    }
  ```
  
* step4:Register the Service
  * way1: Use client.WithClientAware
    ```
    cli = support.MustNewClient(app, cfg.DefaultStoreHandle, cfg.DefaultAccessToken, client.WithClientAware(order.GetOrderService()))
    ```
  * way2: modify the source code 'service_register.go'
    
  ``` see: service_register.go
    func GetClientAwares() []client.Aware {
      var awares = []client.Aware{
          order.GetOrderService(),
          // you can add service here
      }
      return awares
    }

  ```

#### Webhooks verification

In order to be sure that a webhook is sent from Shopline API you could easily verify it with the VerifyWebhookRequest
method.

For example:

```
// see server/main.go
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
    app := manager.GetDefaultApp()
    app.VerifyWebhookRequest(r)
    // do something
}

```

#### <span id="Test_for_Server">Test for Server</span>

The server backend example is only for local testing. Please complete the code if you want to use it in production.<br>

* If you haven't registered yet, please go to the registration platform
  first：https://developer.shopline.com/docs/apps/get-started/onboarding-guidelines-for-shopline-developer-s?version=v20251201 <br>
* If you haven't created an application yet, create one
  first：https://developer.shopline.com/docs/apps/application-management/creating-an-app?version=v20251201

##### step1、Replace App Data

Find "app_config.go" and replace the variable data inside.

```
// Init app_config.go
const (
    DefaultRedirectUrl       = "http://appdemo.myshopline.com/auth/callback" // DefaultRedirectUrl, replace real DefaultRedirectUrl for OAuth

    DefaultAppKey            = ""  // DefaultAppKey, replace real AppKey
    DefaultAppSecret         = ""  // DefaultAppSecret, replace real AppSecret
    DefaultAppScope          = ""  // DefaultAppScope, replace real AppScope
    DefaultStoreHandle       = ""  // replace real store handle
    DefaultAccessToken       = ""  // DefaultAccessToken for test
)

```

##### step2、Start the program(Only for local test)

Execute server/main.go. If successful, port 80 will be started locally. <br>
The console prints "Server started on :80," indicating successful startup.

##### step3、Local hosts bind 1 test domain name

127.0.0.1 appdemo.myshopline.com <br>
Mac: /etc/hosts <br>
Windows: C:\Windows\System32\drivers\etc\hosts

##### step4、Go to "App Settings" and set the App URL and App callback URL

path：[App list](https://developer.myshopline.com/app/list) -> App Detail -> App settings <br>
App loading mode: You can select the "Redirected" mode first, because the "embedded" mode must be the https
protocol.<br>
App URL：http://appdemo.myshopline.com/install <br>
App callback URL：http://appdemo.myshopline.com/auth/callback

##### step5、Test App, Receive the authorization code

* 5.1 path：[App list](https://developer.myshopline.com/app/list) -> App Detail -> Test App
* 5.2 Into App Detail，select「Test App」, click the「Install App」. At this point, the platform will first request
  our [App URL] -> [Platform Auth Page] -> [App Callback URL] in step 4.
  If successful, it will print: Auth callback received ... code: xxx, where code is what we will use to exchange for
  access token later.
  en：https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step3
  code path: server/main.go function 'CallbackHandler'

##### step6、App Create Access Token

* 6.1 Find the TestCreateAccessToken function in oauth_test.go and replace the code variable in the function with the
  code obtained in step 5.
* 6.2 Executing the TestCreateAccessToken function will request the platform to create an access token, and the result
  will be printed in the console if successful.
  en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step-4-request-an-access-token

```
// Create Access Token
code := "code"
appKey := ""
app := manager.GetApp(appKey)
token, err := app.CreateAccessToken(context.Background(), code)
// Do something with the token, like store it in a DB or Cache.

```

##### step7、Use Access Token to call an API

Use Access Token to call
the [Create Product](https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201)
API

* 7.1 Find app_config.go and replace the AccessToken variable value with the access token obtained in step 6.
* 7.2 Find create_product_test.go and execute the TestCreateProduct function. If successful, it will print "New product
  ID: xxx", eg: "New product ID: 16071495214036953630973380".
  For more examples, see the xxx_test.go files in each package.  <br>
  en: https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201

##### step8、App Refresh Access Token

Access tokens expire periodically, so we need to refresh them regularly. Executing the TestRefreshAccessToken function
in oauth_test.go will request the platform to refresh the access token. Successful refreshes will be printed to the
console.<br>
en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step6

```
// Refresh Access Token
storeHandle := ""
appKey := ""
app := manager.GetApp(appKey)
token, err := app.RefreshAccessToken(context.Background(), storeHandle)

// Do something with the token, like store it in a DB or Cache.

```

### <span id="zh">中文</span>

**注意**：对于使用 https://developer.shopline.com 的开发者来说，目前这是一个不稳定的 SDK，我们还在不断完善中。该库并未包含所有 Shopline 资源的实现。欢迎提交新资源和端点的 Pull Request，或者您也可以自行实现一些。更多信息，请参阅“使用您自己的数据模型”部分。


#### Install
```
$ go get github.com/shoplineos/shopline-sdk-go
```

#### Use
```
import "github.com/shoplineos/shopline-sdk-go/client"
```


#### 初始化 App 和 Client
```
  // 1. create app
  appInstance := client.App{
      AppKey:      "",              // replace your data
      AppSecret:   "",              // replace your data
      Scope:       "read_products,write_products,read_orders,write_orders", // replace your data
      RedirectUrl: "http://appdemo.myshopline.com/auth/callback",           // for OAuth replace your data
  }
  
  handle := "zwapptest" // replace your data
  accessToken := ""  // replace your data
  
  // 2. create client
  // c := client.MustNewClient(appInstance, handle, accessToken)
  // appInstance.Client = c

  // use support to create client, it will regitster awares
  c := support.MustNewClient(appInstance, handle, accessToken)
  appInstance.Client = c

  // 3. use client to call API
  // 3.1 API request
  getProductCountAPIReq := &GetProductCountAPIReq{}
  shoplineReq := &client.ShopLineRequest{
      Query: getProductCountAPIReq,
  }

  // 3.2 API endpoint
  endpoint := getProductCountAPIReq.Endpoint()

  // 3.3 API response
  apiResp := &GetProductCountAPIResp{}

  // 3.4 Call API
  shoplineResp, err := c.Get(context.Background(), endpoint, shoplineReq, apiResp)
  fmt.Printf("count:%d", apiResp.Count)
    
```

#### OAuth 认证

如果还没有 access token，可以通过 OAuth 流程来获取 access token，如下：

```
// 详细见：server/main.go
// Create an oauth-authorize url for the App and redirect to it.
// In some request handler, you probably want something like this:
func InstallHandler(w http.ResponseWriter, r *http.Request) {
    // 1. verify http method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. parse request params
	appkey := r.URL.Query().Get("appkey")
	handle := r.URL.Query().Get("handle")
	lang := r.URL.Query().Get("lang")
	timestampStr := r.URL.Query().Get("timestamp")
	sign := r.URL.Query().Get("sign")

	// 3. verify params
	if appkey == "" || timestampStr == "" || sign == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	// 4. TODO Verify the timestamp
	//timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	//if err != nil || time.Now().Unix()-timestamp > 6000 {
	//	http.Error(w, "Invalid timestamp", http.StatusBadRequest)
	//	return
	//}
	
	
	app := manager.GetApp(appkey)

	// 5. Verify the Sign
	isSignValid := app.VerifySign(r.URL.Query(), sign)
	if !isSignValid {
		log.Printf("sign verification failed, appkey: %s, sign: %s\n", appkey, sign)
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	// 6. TODO process biz logic
	log.Printf("install received - appkey: %s, handle: %s, lang: %s", appkey, handle, lang)

	w.Header().Set("Content-Type", "application/json")

	// App Request an authorization code
	// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#%E7%AC%AC%E4%B8%89%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E6%8E%88%E6%9D%83%E7%A0%81
	// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step2
	// url := fmt.Sprintf("https://%s.myshopline.com/admin/oauth-web/#/oauth/authorize?appKey=%s&responseType=code&scope=%s&redirectUri=%s", storeHandle, appKey, scope, redirectUri)

	url, err := app.AuthorizeUrl(handle, "")
	if err != nil {
		log.Printf("Authorize url error, appkey: %s, handle: %s, err: %v\n", appkey, handle, err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Fetch a access token in the callback
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
    appkey := r.URL.Query().Get("appkey")
    code := r.URL.Query().Get("code")
    handle := r.URL.Query().Get("handle")
    customField := r.URL.Query().Get("customField")
    timestampStr := r.URL.Query().Get("timestamp")
    sign := r.URL.Query().Get("sign")
    
    // Check that the callback signature is valid
    app := manager.GetApp(appkey)
	isSignValid := app.VerifySign(r.URL.Query(), sign)
    
    if isSignValid {
        log.Println("sign verified successfully")
    } else {
        log.Printf("sign verification failed, sign=%s\n", sign)
        return
    }
    
    // create token
    token, err := app.CreateAccessToken(context.Background(), code)
    
    // Do something with the token, like store it in a DB or Cache.
}
```

#### 使用 access token 调用 APIs

查询商品数量：
``` Get Product Count
  apiReq := &GetProductCountAPIReq{}
  apiResp, err := product.GetProductService().Count(context.Background(), apiReq)
  fmt.Printf("count:%d", apiResp.Count)
```

查询商品数据:
```Query Products
  requestParams := &ListProductsAPIReq{
      // IDs: "1,2,3",
  }
  productsAPIResp, err := product.GetProductService().List(context.Background(), requestParams)
```

分页：
```Query Products
  requestParams := &ListProductsAPIReq{
      //IDs: "1,2,3",
  }
  productsAPIResp, err := product.GetProductService().ListWithPagination(context.Background(), requestParams)

```


查询所有商品:
``` query all products
  requestParams := &ListProductsAPIReq{}
  Products, err := product.GetProductService().ListAll(context.Background(), requestParams)
  
```


创建商品:
```
// see create_product_test.go
// create product
// https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product/?version=v20251201
apiReq := &product.CreateProductAPIReq{
    Product: Product{
        Title:          "Test product - " + time.Now().Format("20060102150405"),
        BodyHTML:       "<p>This is a test product created via the API</p>",
        Subtitle:       "Limited time offer",
        Vendor:         "Test provider",
        Status:         "active",
        PublishedScope: "web",
        Tags:           []string{"Test", "New", "API Create"},

        // Options（Color、Size）
        Options: []Option{
            {Name: "color", Values: []string{"red", "blue"}},
            {Name: "size", Values: []string{"S", "M", "L"}},
        },

        // 商品图片
        Images: []Image{
            {Src: "https://example.com/product-main.jpg", Alt: "Main picture"},
            {Src: "https://example.com/product-detail.jpg", Alt: "Detail picture"},
        },

        // Variants（red S、red M、blue S）
        Variants: []Variant{
            {
                SKU:            "RED-S-001",
                Price:          "99.99",
                CompareAtPrice: "129.99",
                Option1:        "red",
                Option2:        "S",
                Weight:         "0.5",
                WeightUnit:     "kg",
                Taxable:        true,
            },
            {
                SKU:            "RED-M-002",
                Price:          "109.99",
                CompareAtPrice: "139.99",
                Option1:        "red",
                Option2:        "M",
                Weight:         "0.6",
                WeightUnit:     "kg",
                Taxable:        true,
            },
            {
                SKU:            "BLUE-S-003",
                Price:          "99.99",
                CompareAtPrice: "129.99",
                Option1:        "blue",
                Option2:        "S",
                Weight:         "0.5",
                WeightUnit:     "kg",
                Taxable:        true,
            },
        },
    },
}


apiResp, err := product.GetProductService().Create(context.Background(), apiReq)

```

#### 使用您自己的数据模型

目前为止不是所有的 API 都已经实现，您可以发起1个 Pull Request，或者自己实现数据模型对象。
下面这个例子是获取商品数量:

```
// see get_product_count.go
type GetProductCountAPIReq struct {
	Status       string `url:"status,omitempty"`
	CollectionId string `url:"collection_id,omitempty"`
	CreatedAtMin string `url:"created_at_min,omitempty"` // Minimum order creation time（ISO 8601）
	CreatedAtMax string `url:"created_at_max,omitempty"` // Max order creation time（ISO 8601）
	UpdatedAtMin string `url:"updated_at_min,omitempty"` // Minimum order update time（ISO 8601）
	UpdatedAtMax string `url:"updated_at_max,omitempty"` // Max order update time（ISO 8601）
}

func (req *GetProductCountAPIReq) Verify() error {
	// Verify the API request params
	return nil
}

func (req *GetProductCountAPIReq) Endpoint() string {
	endpoint := "products/count.json"
	return endpoint
}

type GetProductCountAPIResp struct {
	Count int `json:"count"`
	client.BaseAPIResponse
}

func GetProductsCount(c *client.Client, apiReq *GetProductCountAPIReq) (*GetProductCountAPIResp, error) {

    // 1. API request
    shoplineReq := &client.ShopLineRequest{
        Query: apiReq, // http url query params
        // Data: apiReq, // http body params
    }
    
    // 2. API endpoint
    endpoint := apiReq.Endpoint()
    
    // 3. API response
    apiResp := &GetProductCountAPIResp{}
    
    // 4. Call API
    shoplineResp, err := c.Get(context.Background(), endpoint, shoplineReq, apiResp)
    
    // option
    // apiResp.TraceId = shoplineResp.TraceId
    
    return apiResp, err
}

```

#### 实现你自己的 Service Interface
* step1:Define a Service Interface
  * See product_service.go or order_service.go
  ```
    type IOrderService interface {
        List(context.Context, *ListOrdersAPIReq) (*ListOrdersAPIResp, error)
        ListAll(context.Context, *ListOrdersAPIReq) ([]Order, error)
        ...
    }
  ```
* step2:Define a Service struct and implements the Service Interface
  ```
    type OrderService struct {
      client.BaseService
    }
  
    func (o *OrderService) List(ctx context.Context, apiReq *ListOrdersAPIReq) (*ListOrdersAPIResp, error) {
      ...
    }
  
    func (o *OrderService) ListAll(ctx context.Context, apiReq *ListOrdersAPIReq) (*ListOrdersAPIResp, error) {
      ...
    }

  ```
* step3:Create a Service Instance
  ```
    var serviceInst = &OrderService{}
    func GetOrderService() *OrderService {
      return serviceInst
    }
  ```

* step4:Register the Service
  * 方式1: 使用 client.WithClientAware
    ```
    cli = support.MustNewClient(app, cfg.DefaultStoreHandle, cfg.DefaultAccessToken, client.WithClientAware(order.GetOrderService()))
    ```
  * 方式2: modify the source code 'service_register.go'

  ``` see: service_register.go
    func GetClientAwares() []client.Aware {
      var awares = []client.Aware{
          order.GetOrderService(),
          // you can add service here
      }
      return awares
    }

  ```



#### 验证Webhooks

例子:

```
// see server/main.go
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
    app := manager.GetDefaultApp()
    app.VerifyWebhookRequest(r)
    // do something
}
```

#### <span id="服务端测试">服务端测试</span>

server端例子仅仅作为本地测试 demo 使用。如要在生产使用，请完善里面的代码。<br>

*

如果尚未入驻，请先去入驻：https://developer.shopline.com/zh-hans-cn/docs/apps/get-started/onboarding-guidelines-for-shopline-developer-s?version=v20251201 <br>

*

如果尚未创建应用，先去创建应用：https://developer.shopline.com/zh-hans-cn/docs/apps/application-management/creating-an-app?version=v20251201

##### step1、替换 App 数据

找到 app_config.go 替换里面的变量数据。

```
// Init app_config.go
const (
    DefaultRedirectUrl       = "http://appdemo.myshopline.com/auth/callback" // DefaultRedirectUrl, replace real DefaultRedirectUrl for OAuth

    DefaultAppKey            = ""  // DefaultAppKey, replace real AppKey
    DefaultAppSecret         = ""  // DefaultAppSecret, replace real AppSecret
    DefaultAppScope          = ""  // DefaultAppScope, replace real AppScope
    DefaultStoreHandle       = ""  // replace real store handle
    DefaultAccessToken       = ""  // DefaultAccessToken for test
)

```

##### step2、启动程序

执行 server/main.go，成功后，会在本地启动80端口。<br>
控制台打印“Server started on :80” 表示启动成功了。

##### step3、本地 hosts 绑定1个测试域名

127.0.0.1 appdemo.myshopline.com <br>
Mac径路: /etc/hosts <br>
Windows路径: C:\Windows\System32\drivers\etc\hosts

##### step4、去「应用设置」，设置 应用地址 和 应用回调地址

路径：[应用列表](https://developer.myshopline.com/app/list) -> 应用详情 -> 应用设置 <br>
应用打开方式：可以先选择“外跳”模式，因为“内嵌”模式，必须是https协议。 <br>
应用地址: http://appdemo.myshopline.com/install <br>
应用回调地址: http://appdemo.myshopline.com/auth/callback

##### step5、应用测试，接收授权码

* 5.1 路径：[应用列表](https://developer.myshopline.com/app/list) -> 应用详情 -> 应用测试
* 5.2 进入应用详情，选择「应用测试」，点击「安装应用」。此时，平台会先请求我们第4步的【应用地址】 -> 【平台认证页面】 -> 【应用回调地址】
  成功后会打印: Auth callback received ... code: xxx，其中 code 就是我们后续用来交换 access token。 <br>
  中文：https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E4%BA%94%E6%AD%A5app-%E6%94%B6%E5%88%B0%E6%8E%88%E6%9D%83%E7%A0%81  <br>
  代码：server/main.go 函数 CallbackHandler

##### step6、App 获取 Access Token

* 6.1 找到 oauth_test.go 里的 TestCreateAccessToken 函数，用第5步获取到的 code，替换函数里的 code 变量。
* 6.2 执行 TestCreateAccessToken 函数，会请求平台获取 access token，成功后会打印在控制台。
  中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E5%85%AD%E6%AD%A5app-%E8%AF%B7%E6%B1%82-access-token  <br>

```
// Create Access Token
appkey := "appkey"
code := "code"
app := manager.GetApp(appKey)
token, err := app.CreateAccessToken(context.Background(), code)

// Do something with the token, like store it in a DB or Cache.

```

##### step7、使用 Access Token 调用1个接口

使用 Access Token
调用 [创建商品](https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/create-a-product?version=v20251201)
接口。

* 7.1 找到 app_config.go，用第6步获取到的 access token，替换 AccessToken 变量值。
* 7.2 找到 create_product_test.go，执行 TestCreateProduct 函数，成功后会打印 “新商品ID: xxx”，如：新商品ID:
  16071495214036953630973380。
  更多例子见：各个包下面的 xxx_test.go 文件。  <br>
  中文: https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/create-a-product/?version=v20251201  <br>

##### step8、App 刷新 Access Token

Access Token 每隔一段时间会过期，因此我们需要定期刷新 Access Token。执行 oauth_test.go 里的 TestRefreshAccessToken
函数，会请求平台刷新 access token，成功后会打印在控制台。 <br>
中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E5%85%AB%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E5%88%B7%E6%96%B0-access-token  <br>

```
// Refresh Access Token
storeHandle := ""
appKey := ""
app := manager.GetApp(appKey)
token, err := app.RefreshAccessToken(context.Background(), storeHandle)

// Do something with the token, like store it in a DB or Cache.

```
