# shopline sdk go

## Getting Started

[中文](#zh)、[English](#en)

### <span id="en">English</span>

**Note**: For developers using [SHOPLINE](https://developer.shopline.com), the SDK is currently under active development and is a preview release with ongoing improvements. This library doesn't yet implement all SHOPLINE resources. We welcome you to submit Pull Requests to add new resources or endpoints, or you can implement your own by following the instructions in the <a href="#en-use-your-own-model">Using Your Own Data Model</a> section. We look forward to your contributions!


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
    // 1. Create an app
  appInstance := client.App{
      AppKey:      "",              // replace your data
      AppSecret:   "",              // replace your data
      Scope:       "read_products,write_products,read_orders,write_orders", // replace your data
      RedirectUrl: "http://appdemo.myshopline.com/auth/callback",           // replace your data
  }
  
  handle := "zwapptest" // replace your data
  accessToken := "" // replace your data
  
  // 2. Create the client instance
  // Use the support library to create the client; it will register relevant handlers
  c := support.MustNewClient(appInstance, handle, accessToken)
  appInstance.Client = c
  
    
  // 3. Use the client to call the API
  // 3.1 API request
  getProductCountAPIReq := &GetProductCountAPIReq{}
  shoplineReq := &client.ShopLineRequest{
      Query: getProductCountAPIReq,
  }

  // 3.2 API endpoint
  endpoint := getProductCountAPIReq.Endpoint()

  // 3.3 API response
  apiResp := &GetProductCountAPIResp{}

  // 3.4 Call an API
  shoplineResp, err := c.Get(context.Background(), endpoint, shoplineReq, apiResp)
  fmt.Printf("count:%d", apiResp.Count)
```

#### OAuth

If you don't have an Access Token yet, you can obtain one with the oauth flow. Something like this will work:

```
// For more details, see: server/main.go
// Create an OAuth authorization URL for the app and redirect to it.
// The following shows an example:
func InstallHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Verify HTTP method
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
  
    // 2. Parse request parameters
    appkey := r.URL.Query().Get("appkey")
    handle := r.URL.Query().Get("handle")
    lang := r.URL.Query().Get("lang")
    timestampStr := r.URL.Query().Get("timestamp")
    sign := r.URL.Query().Get("sign")
  
    // 3. Verify parameters
    if appkey == "" || timestampStr == "" || sign == "" {
        http.Error(w, "Missing required parameters", http.StatusBadRequest)
        return
    }
  
    // 4. TODO: Verify the timestamp
    // timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
    // if err != nil || time.Now().Unix()-timestamp > 6000 {
    //     http.Error(w, "Invalid timestamp", http.StatusBadRequest)
    //     return
    // }
  
    app := manager.GetApp(appkey)
  
    // 5. Verify the signature
    isSignValid := app.VerifySign(r.URL.Query(), sign)
    if !isSignValid {
        log.Printf("sign verification failed, appkey: %s, sign: %s\n", appkey, sign)
        http.Error(w, "Invalid signature", http.StatusUnauthorized)
        return
    }
  
    // 6. TODO: Process business logic
    log.Printf("install received - appkey: %s, handle: %s, lang: %s", appkey, handle, lang)
  
    w.Header().Set("Content-Type", "application/json")
  
    // App requests an authorization code
    // 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#%E7%AC%AC%E4%B8%89%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E6%8E%88%E6%9D%83%E7%A0%81
    // en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step2
    // url := fmt.Sprintf("https://%s.myshopline.com/admin/oauth-web/#/oauth/authorize?appKey=%s&responseType=code&scope=%s&redirectUri=%s", storeHandle, appKey, scope, redirectUri)

    url, err := app.AuthorizeUrl(handle, "")
    if err != nil {
        log.Printf("Authorize URL error, appkey: %s, handle: %s, err: %v\n", appkey, handle, err)
        http.Error(w, "server error", http.StatusInternalServerError)
        return
    }
  
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Fetch an access token in the callback
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
        log.Println("Sign verified successfully")
    } else {
        log.Printf("Sign verification failed, sign=%s\n", sign)
        return
    }
    
    // Create an access token
    token, err := app.CreateAccessToken(context.Background(), code)
    
    // Do something with the token, like store it in a DB or Cache.
}
```

#### API calls with an Access Token

With an Access Token, you can make API calls like this:

Get Product Count:
https://developer.shopline.com/docs/admin-rest-api/product/product/get-product-count?version=v20251201
``` Get Product Count
  apiReq := &GetProductCountAPIReq{}
  apiResp, err := product.GetProductService().Count(context.Background(), apiReq)
  fmt.Printf("count:%d", apiResp.Count)
```

[Get Products](https://developer.shopline.com/docs/admin-rest-api/product/product/get-products?version=v20251201):
```Get Products
  requestParams := &ListProductsAPIReq{
      // IDs: "1,2,3",
  }
  productsAPIResp, err := product.GetProductService().List(context.Background(), requestParams)
```


[Get Products Pagination](https://developer.shopline.com/docs/admin-rest-api/product/product/get-products?version=v20251201):
```Get Products
  requestParams := &ListProductsAPIReq{
      // IDs: "1,2,3",
  }
  productsAPIResp, err := product.GetProductService().ListWithPagination(context.Background(), requestParams)
```

[Query all products](https://developer.shopline.com/docs/admin-rest-api/product/product/get-products?version=v20251201):
``` query all products
  requestParams := &ListProductsAPIReq{}
  Products, err := product.GetProductService().ListAll(context.Background(), requestParams)
  
```

[Create a product](https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201):
``` create a product
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

#### <span id="en-use-your-own-model">Using Your Own Data Model</span>

Not all endpoints are implemented right now. In those case, feel free to implement them and make a Pull Request, or you
can create your own struct for the data and use the client to call APIs. This is how the existing endpoints are
implemented.

For example, let's say you want to fetch product count. There's a helper function Get specifically for fetching stuff so
this will work:

```
// See get_product_count.go
type GetProductCountAPIReq struct {
	Status       string `url:"status,omitempty"`
	CollectionId string `url:"collection_id,omitempty"`
	CreatedAtMin string `url:"created_at_min,omitempty"` // Minimum product creation time（ISO 8601）
	CreatedAtMax string `url:"created_at_max,omitempty"` // Max product creation time（ISO 8601）
	UpdatedAtMin string `url:"updated_at_min,omitempty"` // Minimum product update time（ISO 8601）
	UpdatedAtMax string `url:"updated_at_max,omitempty"` // Max product update time（ISO 8601）
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
    
    // 4. Call an API
    shoplineResp, err := c.Get(context.Background(), endpoint, shoplineReq, apiResp)
    
    // option
    // apiResp.TraceId = shoplineResp.TraceId
    
    return apiResp, err
}

```


#### Implementing your own service interface
* step 1: Define a service interface
  * Refer to product_service.go or order_service.go for examples.
  ```
    type IOrderService interface {
        List(context.Context, *ListOrdersAPIReq) (*ListOrdersAPIResp, error)
        ListAll(context.Context, *ListOrdersAPIReq) ([]Order, error)
        ...
    }
  ```
* step 2: Define a service struct and implement the service interface
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
* step 3: Create a service instance
  ```
    var serviceInst = &OrderService{}
    func GetOrderService() *OrderService {
      return serviceInst
    }
  ```
  
* step 4: Register the service
  * method 1: Use client.WithClientAware to register.
    ```
    cli = support.MustNewClient(app, cfg.DefaultStoreHandle, cfg.DefaultAccessToken, client.WithClientAware(order.GetOrderService()))
    ```
  * method 2: Modify the source code in service_register.go.
    
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

In order to be sure that a webhook is sent from SHOPLINE API you could easily verify it with the VerifyWebhookRequest
method. For more details, see: [Webhooks](https://developer.shopline.com/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301)

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

Server-side testing is primarily used by developers to verify application functionality in a local environment. The following sample code is for demonstration purposes only and cannot be used directly in a production environment.
. Please complete the code if you want to use it in production.<br>

* If you haven't registered yet, please go to the registration platform
  first：[Onboarding guidelines for SHOPLINE developers](https://developer.shopline.com/docs/apps/get-started/onboarding-guidelines-for-shopline-developer-s?version=v20251201) <br>
* If you haven't created an application yet, create one
  first：[Create an App](https://developer.shopline.com/docs/apps/application-management/creating-an-app?version=v20251201)

##### step 1、Replace App Data

Find "app_config.go" and replace the variable data inside.

```
// Initialize app_config.go
const (
    DefaultRedirectUrl       = "http://appdemo.myshopline.com/auth/callback" // Placeholder for the OAuth Redirect URL, replace with the actual URL.

    DefaultAppKey            = ""  // Replace with the actual app key.
    DefaultAppSecret         = ""  // Replace with the actual app secret.
    DefaultAppScope          = ""  // Replace with the actual app scope.
    DefaultStoreHandle       = ""  // Replace with the actual store handle.
    DefaultAccessToken       = ""  // Replace with the actual token for testing.
)

```

##### step 2、Start the program(Only for local test)

Execute server/main.go. If successful, port 80 will be started locally. <br>
The console prints "Server started on :80," indicating successful startup.

##### step 3、Local hosts bind 1 test domain name

127.0.0.1 appdemo.myshopline.com <br>
Mac: /etc/hosts <br>
Windows: C:\Windows\System32\drivers\etc\hosts

##### step 4、Go to "App Settings" and set the App URL and App callback URL

1. Find [Apps]([App list](https://developer.myshopline.com/app/list) ) in the [Partner Portal](https://developer.myshopline.com/) .

2. Click the name of the app you want to test.

3. Click App Settings under Basic Settings.

4. Set the App URL and App Callback URL to http://appdemo.myshopline.com/install and http://appdemo.myshopline.com/auth/callback, respectively.

5. Select 'Redirected' for the App Opening Method, as the 'Embedded' mode requires HTTPS.


##### step 5、App Receive the authorization code

1. Path：[Partner Portal](https://developer.myshopline.com/) -> [App list](https://developer.myshopline.com/app/list) -> App Detail -> Test App
2. Into App Detail，select「Test App」, click the「Install App」. At this point, the platform will first request
  our [App URL] -> [Platform Auth Page] -> [App Callback URL] in step 4.
  If successful, it will print: Auth callback received ... code: xxx, where code is what we will use to exchange for
  Access Token later.
  For more details, see：[Receive the authorization code](https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step3)
  Example: server/main.go function 'CallbackHandler'

##### step 6、App Create Access Token

1. Find the TestCreateAccessToken function in oauth_test.go and replace the code variable in the function with the
  code obtained in step 5.
2. Executing the TestCreateAccessToken function will request the platform to create an Access Token, and the result
  will be printed in the console if successful.
  For more details, see: [App Create Access Token](https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step-4-request-an-access-token)

```
// Create an Access Token
code := "code"
appKey := ""
app := manager.GetApp(appKey)
token, err := app.CreateAccessToken(context.Background(), code)
// Store the token in a database or cache for future use. 

```

##### step 7、Use Access Token to call an API

Use Access Token to call the [Create Product](https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201) API

1. Find app_config.go and replace the AccessToken variable value with the Access Token obtained in step 6.
2. Find create_product_test.go and execute the TestCreateProduct function. If successful, it will print "New product
  ID: xxx", eg: "New product ID: 16071495214036953630973380".
  For more examples, see the xxx_test.go files in each package.  <br>
  For more details, see: [Create a product](https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201)

##### step 8、App Refresh Access Token

Access tokens expire periodically, so we need to refresh them regularly. Executing the TestRefreshAccessToken function
in oauth_test.go will request the platform to refresh the Access Token. Successful refreshes will be printed to the
console.<br>
For more details, see: [App Refresh Access Token](https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step6)

```
// Refresh the access token
storeHandle := ""
appKey := ""
app := manager.GetApp(appKey)
token, err := app.RefreshAccessToken(context.Background(), storeHandle)

// Store the token in a database or cache for future use. 
```

### <span id="zh">中文</span>

**注意**：对于使用 [SHOPLINE](https://developer.shopline.com) 的开发者来说，当前 SDK 封装了通用对接逻辑，旨在帮助开发者高效地构建应用，让你能更专注于业务功能的实现。当前 SDK 尚处于 Beta 阶段，功能持续完善中，已支持 OAuth 授权、商品、订单等核心 API，并提供了清晰的客户端初始化、请求构建与响应处理模式。我们欢迎社区贡献代码，您可通过提交 Pull Request 补充新的资源或端点。如果遇到尚未实现的 API，可以参照 <a href="#zh-use-your-own-model">使用你自己的数据模型</a> 自行封装。


#### 安装
```
$ go get github.com/shoplineos/shopline-sdk-go
```

#### 使用
```
import "github.com/shoplineos/shopline-sdk-go/client"
```


#### 初始化 App 和 Client
```
  // 1. 创建应用实例
  appInstance := client.App{
      AppKey:      "",              // replace your data
      AppSecret:   "",              // replace your data
      Scope:       "read_products,write_products,read_orders,write_orders", // replace your data
      RedirectUrl: "http://appdemo.myshopline.com/auth/callback",           // for OAuth replace your data
  }
  
  handle := "zwapptest" // replace your data
  accessToken := ""  // replace your data
  
  // 2. 创建客户端实例
  // c := client.MustNewClient(appInstance, handle, accessToken)
  // appInstance.Client = c

  // use support to create client, it will regitster awares
  c := support.MustNewClient(appInstance, handle, accessToken)
  appInstance.Client = c

  // 3. 调用 API
  // 3.1 API 请求
  getProductCountAPIReq := &GetProductCountAPIReq{}
  shoplineReq := &client.ShopLineRequest{
      Query: getProductCountAPIReq,
  }

  // 3.2 API 端点
  endpoint := getProductCountAPIReq.Endpoint()

  // 3.3 API 响应
  apiResp := &GetProductCountAPIResp{}

  // 3.4 调用 API
  shoplineResp, err := c.Get(context.Background(), endpoint, shoplineReq, apiResp)
  fmt.Printf("count:%d", apiResp.Count)
    
```

#### OAuth 授权

如果还没有 Access Token，可以通过 OAuth 流程来获取 Access Token，流程如下：

```
// 详细见：server/main.go
// 为应用创建一个 OAuth 授权 URL 并重定向到该 URL。
// 如下展示了一个样例：
func InstallHandler(w http.ResponseWriter, r *http.Request) {
    // 1. 验证 HTTP 方法
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // 2. 解析请求参数
    appkey := r.URL.Query().Get("appkey")
    handle := r.URL.Query().Get("handle")
    lang := r.URL.Query().Get("lang")
    timestampStr := r.URL.Query().Get("timestamp")
    sign := r.URL.Query().Get("sign")

    // 3. 验证参数
    if appkey == "" || timestampStr == "" || sign == "" {
        http.Error(w, "Missing required parameters", http.StatusBadRequest)
        return
    }
  
    // 4. TODO: 验证时间戳
    // timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
    // if err != nil || time.Now().Unix()-timestamp > 6000 {
    //     http.Error(w, "Invalid timestamp", http.StatusBadRequest)
    //     return
    // }

    app := manager.GetApp(appkey)

    // 5. 验证签名
    isSignValid := app.VerifySign(r.URL.Query(), sign)
    if !isSignValid {
        log.Printf("sign verification failed, appkey: %s, sign: %s\n", appkey, sign)
        http.Error(w, "Invalid signature", http.StatusUnauthorized)
        return
    }

    // 6. TODO: 执行业务逻辑
    log.Printf("install received - appkey: %s, handle: %s, lang: %s", appkey, handle, lang)

    w.Header().Set("Content-Type", "application/json")

    // 应用请求授权码
    // 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#%E7%AC%AC%E4%B8%89%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E6%8E%88%E6%9D%83%E7%A0%81
    // en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step2
    // url := fmt.Sprintf("https://%s.myshopline.com/admin/oauth-web/#/oauth/authorize?appKey=%s&responseType=code&scope=%s&redirectUri=%s", storeHandle, appKey, scope, redirectUri)

    url, err := app.AuthorizeUrl(handle, "")
    if err != nil {
        log.Printf("Authorize URL error, appkey: %s, handle: %s, err: %v\n", appkey, handle, err)
        http.Error(w, "server error", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// 在回调中获取访问令牌
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
    appkey := r.URL.Query().Get("appkey")
    code := r.URL.Query().Get("code")
    handle := r.URL.Query().Get("handle")
    customField := r.URL.Query().Get("customField")
    timestampStr := r.URL.Query().Get("timestamp")
    sign := r.URL.Query().Get("sign")
    
    // 检查回调签名是否有效
    app := manager.GetApp(appkey)
    isSignValid := app.VerifySign(r.URL.Query(), sign)
    
    if isSignValid {
        log.Println("Sign verified successfully")
    } else {
        log.Printf("Sign verification failed, sign=%s\n", sign)
        return
    }
    
    // 创建访问令牌
    token, err := app.CreateAccessToken(context.Background(), code)
    
    // 使用令牌执行某些操作，例如将其存储在数据库或缓存中。
}
```

#### 使用 Access Token 调用 APIs
以下是以商品相关接口为例：</br>
[查询商品数量](https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-product-count?version=v20251201)：
``` Get Product Count
  apiReq := &GetProductCountAPIReq{}
  apiResp, err := product.GetProductService().Count(context.Background(), apiReq)
  fmt.Printf("count:%d", apiResp.Count)
```

[批量查询商品数据](https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-products?version=v20251201):
```Query Products
  requestParams := &ListProductsAPIReq{
      // IDs: "1,2,3",
  }
  productsAPIResp, err := product.GetProductService().List(context.Background(), requestParams)
```

[分页查询商品数据](https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-products?version=v20251201)：
```Query Products
  requestParams := &ListProductsAPIReq{
      //IDs: "1,2,3",
  }
  productsAPIResp, err := product.GetProductService().ListWithPagination(context.Background(), requestParams)

```

[查询所有商品](https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/get-products?version=v20251201):
``` Query all products
  requestParams := &ListProductsAPIReq{}
  Products, err := product.GetProductService().ListAll(context.Background(), requestParams)
  
```


[创建商品](https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/create-a-product?version=v20251201):
```
// 详细见：create_product_test.go
// 创建商品
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


#### <span id="zh-use-your-own-model">使用你自己的数据模型</span>

目前为止不是所有的 API 都已经实现，你可以发起1个 Pull Request，或者自己实现数据模型对象。
下面这个例子是获取商品数量:

```
// 详细见：get_product_count.go
type GetProductCountAPIReq struct {
	Status       string `url:"status,omitempty"`
	CollectionId string `url:"collection_id,omitempty"`
	CreatedAtMin string `url:"created_at_min,omitempty"` // Minimum product creation time（ISO 8601）
	CreatedAtMax string `url:"created_at_max,omitempty"` // Max product creation time（ISO 8601）
	UpdatedAtMin string `url:"updated_at_min,omitempty"` // Minimum product update time（ISO 8601）
	UpdatedAtMax string `url:"updated_at_max,omitempty"` // Max product update time（ISO 8601）
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
    shoplineReq := &client.ShopLineRequest{
        Query: apiReq, // http url query params
        // Data: apiReq, // http body params
    }
    endpoint := apiReq.Endpoint()
    apiResp := &GetProductCountAPIResp{}
    shoplineResp, err := c.Get(context.Background(), endpoint, shoplineReq, apiResp)
    return apiResp, err
}

```

#### 实现你自己的服务接口 Service Interface
* 步骤 1: 定义一个服务接口
  * See product_service.go or order_service.go
  ```
    type IOrderService interface {
        List(context.Context, *ListOrdersAPIReq) (*ListOrdersAPIResp, error)
        ListAll(context.Context, *ListOrdersAPIReq) ([]Order, error)
        // 你可以在这里添加更多方法
    }
  ```
* 步骤 2: 定义一个 Service 结构体并实现服务接口
  ```
    type OrderService struct {
      client.BaseService
    }
  
    func (o *OrderService) List(ctx context.Context, apiReq *ListOrdersAPIReq) (*ListOrdersAPIResp, error) {
        // 在这里定义函数实现逻辑
    }
  
    func (o *OrderService) ListAll(ctx context.Context, apiReq *ListOrdersAPIReq) (*ListOrdersAPIResp, error) {
        // 在这里定义函数实现逻辑
    }

  ```
* 步骤 3: 创建服务实例
  ```
    var serviceInst = &OrderService{}
    func GetOrderService() *OrderService {
      return serviceInst
    }
  ```

* 步骤 4: 注册服务
  * 方法 1: 使用 client.WithClientAware
    ```
    cli = support.MustNewClient(app, cfg.DefaultStoreHandle, cfg.DefaultAccessToken, client.WithClientAware(order.GetOrderService()))
    ```
  * 方法 2: 修改 'service_register.go' 源代码

  ``` see: service_register.go
    func GetClientAwares() []client.Aware {
      var awares = []client.Aware{
          order.GetOrderService(),
          // 可以在这里添加更多服务
      }
      return awares
    }

  ```


#### 验证 Webhooks
为了确保 Webhook 确实从 SHOPLINE 发送的，你可以使用 VerifyWebhookRequest 方法轻松验证。更多信息见：[Webhooks](https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview/?version=v20260301) </br>
例子:

```
// 详细见：server/main.go
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
    app := manager.GetDefaultApp()
    app.VerifyWebhookRequest(r)
    // do something
}
```

#### <span id="服务端测试">服务端测试</span>

服务端测试主要用于开发者在本地环境中验证应用功能，以下示例代码仅作为基础演示，不可直接用于生产环境。如需使用，请先完善代码。<br>

* 如尚未入驻 SHOPLINE 开发者平台，参考 [SHOPLINE 开发者入驻指引](https://developer.shopline.com/zh-hans-cn/docs/apps/get-started/onboarding-guidelines-for-shopline-developer-s?version=v20251201) 完成入驻
* 如尚未创建应用，参考 [创建应用](https://developer.shopline.com/zh-hans-cn/docs/apps/application-management/creating-an-app?version=v20251201) 完成应用创建。

##### 步骤 一：替换 App 数据

找到 app_config.go 替换里面的变量数据。

```
// 初始化 app_config.go
const (
    DefaultRedirectUrl       = "http://appdemo.myshopline.com/auth/callback" // OAuth 重定向 URL 的占位符，替换为实际的 URL

    DefaultAppKey            = ""  // 替换为实际的 app key
    DefaultAppSecret         = ""  // 替换为实际的 app secret
    DefaultAppScope          = ""  // 替换为实际的 app scope
    DefaultStoreHandle       = ""  // 替换为实际的 store handle
    DefaultAccessToken       = ""  // 替换为实际的令牌
)

```

##### 步骤 二：启动程序

执行 server/main.go，成功后，会在本地启动80端口。<br>
控制台打印“Server started on :80” 表示启动成功了。

##### 步骤 三：本地 hosts 绑定1个测试域名

127.0.0.1 appdemo.myshopline.com <br>
Mac径路: /etc/hosts <br>
Windows路径: C:\Windows\System32\drivers\etc\hosts

##### 步骤 四：在 合作伙伴后台 设置 应用地址 和 应用回调地址
1. 在 [合作伙伴后台](https://developer.shopline.com) 中找到 [应用](https://developer.myshopline.com/app/list)。
2. 点击需要测试的应用名称。
3. 点击 基础设置 下的 应用设置。
4. 设置 应用地址 和 应用回调地址 分别为 http://appdemo.myshopline.com/install 和  http://appdemo.myshopline.com/auth/callback。
5. 应用打开方式 选择 外跳，因为 内嵌 模式必须是 HTTPS 协议。

##### 步骤 五：应用接收授权码

1. 路径：[合作伙伴后台](https://developer.shopline.com) -> [应用列表](https://developer.myshopline.com/app/list) -> 应用详情 -> 应用测试
2. 进入应用详情，选择「应用测试」，点击「安装应用」。此时，平台会先请求我们第4步的【应用地址】 -> 【平台认证页面】 -> 【应用回调地址】
  成功后会打印: Auth callback received ... code: xxx，其中 code 就是我们后续用来交换 Access Token。 <br>
  详细文档见：[App 收到授权码](https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E4%BA%94%E6%AD%A5app-%E6%94%B6%E5%88%B0%E6%8E%88%E6%9D%83%E7%A0%81)  <br>
  代码示例见：server/main.go 函数 CallbackHandler

##### 步骤 六：App 获取 Access Token

1. 找到 oauth_test.go 里的 TestCreateAccessToken 函数，用第5步获取到的 code，替换函数里的 code 变量。
2. 执行 TestCreateAccessToken 函数，会请求平台获取 Access Token，成功后会打印在控制台。
  可参考: [应用请求 Access Token](https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E5%85%AD%E6%AD%A5app-%E8%AF%B7%E6%B1%82-access-token)  <br>

```
// 创建 Access Token
appkey := "appkey"
code := "code"
app := manager.GetApp(appKey)
token, err := app.CreateAccessToken(context.Background(), code)

// Do something with the token, like store it in a DB or Cache.

```

##### 步骤 七：使用 Access Token 调用1个接口

使用 Access Token
调用 [创建商品](https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/create-a-product?version=v20251201)
接口。

1. 找到 app_config.go，用第6步获取到的 Access Token，替换 AccessToken 变量值。
2. 找到 create_product_test.go，执行 TestCreateProduct 函数，成功后会打印 “新商品ID: xxx”，如：新商品ID:
  16071495214036953630973380。
  更多例子见：各个包下面的 xxx_test.go 文件。  <br>

##### 步骤 八：App 刷新 Access Token

Access Token 每隔一段时间会过期，因此我们需要定期刷新 Access Token。执行 oauth_test.go 里的 TestRefreshAccessToken
函数，会请求平台刷新 Access Token，成功后会打印在控制台。 <br>
详细见: [App 刷新 Access Token](https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E5%85%AB%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E5%88%B7%E6%96%B0-access-token)  <br>

```
// 刷新 Access Token
storeHandle := ""
appKey := ""
app := manager.GetApp(appKey)
token, err := app.RefreshAccessToken(context.Background(), storeHandle)

// 将 Access Token 存在数据库或者缓存里，方便后续使用
```
