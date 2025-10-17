# SHOPLINE API SDK for Go


[中文](#zh)、[English](#en)

## <span id="en">English</span>

For developers using [SHOPLINE](https://developer.shopline.com), this SDK provides support for Go [SHOPLINE apps](https://apps.shopline.com/) to access the [SHOPLINE Admin API](https://developer.shopline.com/docs/apps/api-instructions-for-use/rest-admin-api/overview?version=v20260301) by simplifying common tasks and processes.

Current supported <span id="en-actions">actions</span>:
1. Creating access tokens for the Admin API via [OAuth](https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301)
2. Making requests to the [REST Admin API](https://developer.shopline.com/docs/apps/api-instructions-for-use/rest-admin-api/overview?version=v20260301)
3. Registering and processing [Webhooks](https://developer.shopline.com/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301)

**Note**: This SDK is in its Beta phase, meaning that it is actively being developed and improved. This SDK doesn't yet implement all [SHOPLINE](https://developer.shopline.com) resources. We welcome you to submit Pull Requests to add new resources or endpoints, or you can implement your own by following the instructions in the <a href="#en-use-your-own-model">Using your own data model</a> section. We look forward to your contributions!

### Install the SDK
```
$ go get github.com/shoplineos/shopline-sdk-go
```

### Use the SDK
```
import "github.com/shoplineos/shopline-sdk-go/client"
```


### Init App and Client
```
    // 1. Create an app instance
  appInstance := client.App{
      AppKey:      "",              // Replace with your data
      AppSecret:   "",              // Replace with your data
      Scope:       "read_products,write_products,read_orders,write_orders", // Replace with your data
      RedirectUrl: "http://appdemo.myshopline.com/auth/callback",           // Replace with your data
  }
  
  handle := "zwapptest" // Replace with your data
  accessToken := "" // Replace with your data
  
  // 2. Create the client instance
  // Use the support tool to create the client; it will register relevant handlers
  c := support.MustNewClient(appInstance, handle, accessToken)
  appInstance.Client = c
  
  
  // 3. Use the client to call the API
  // 3.1 API request
  getProductCountAPIReq := &GetProductCountAPIReq{}
  
  // 3.2 API response data
  apiResp := &GetProductCountAPIResp{}

  // 3.3 Call the API
  err := client.Call(context.Background(), apiReq, apiResp)

  fmt.Printf("count:%d", apiResp.Count)
```

### OAuth

If you don't have an access token yet, you can obtain one with the OAuth flow. The following shows an example:

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

### API calls with an access token

With an access token, you can make API calls like this:
```
  // Use the client to call the API
  // 1 API request
  getProductCountAPIReq := &GetProductCountAPIReq{}
  
  // 2 API response data
  apiResp := &GetProductCountAPIResp{}

  // 3 Call the API
  err := client.Call(context.Background(), apiReq, apiResp)

  fmt.Printf("count:%d", apiResp.Count)
```

The following are examples of product APIs:

[Get Product Count](https://developer.shopline.com/docs/admin-rest-api/product/product/get-product-count?version=v20251201):

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
// See create_product_test.go
// Create a product
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

// Use Service to call the API or to use client.Call
apiResp, err := product.GetProductService().Create(context.Background(), apiReq)

```

### <span id="en-use-your-own-model">Using your own data model</span>

Not all API endpoints are implemented, so feel free to add them yourself and submit a Pull Request. Alternatively, you can create your own struct for the data and use the client to call APIs.

For example, if you want to retrieve a product count, there is a helper function called Get that is specifically designed for such retrievals:

```
// See get_product_count.go
type GetProductCountAPIReq struct {
    client.BaseAPIRequest
	Status       string `url:"status,omitempty"`
	CollectionId string `url:"collection_id,omitempty"`
	CreatedAtMin string `url:"created_at_min,omitempty"` // Minimum product creation time（ISO 8601）
	CreatedAtMax string `url:"created_at_max,omitempty"` // Max product creation time（ISO 8601）
	UpdatedAtMin string `url:"updated_at_min,omitempty"` // Minimum product update time（ISO 8601）
	UpdatedAtMax string `url:"updated_at_max,omitempty"` // Max product update time（ISO 8601）
}

func (req *GetProductCountAPIReq) Method() string {
	return "GET"
}

func (req *GetProductCountAPIReq) GetQuery() string {
	return req
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

func TestGetProductsCount(c *client.Client) (*GetProductCountAPIResp, error) {

    // 1 API request
    getProductCountAPIReq := &GetProductCountAPIReq{}
    
    // 2 API response data
    apiResp := &GetProductCountAPIResp{}
  
    // 3 Call the API
    err := c.Call(context.Background(), apiReq, apiResp)

    return apiResp, err
}

```


### Implementing your own service interface(Optional)

Step1: Define a service interface

Refer to product_service.go or order_service.go for examples.
  ```
    type IOrderService interface {
        List(context.Context, *ListOrdersAPIReq) (*ListOrdersAPIResp, error)
        ListAll(context.Context, *ListOrdersAPIReq) ([]Order, error)
        ...
    }
  ```
Step2: Define a service struct and implement the service interface
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

Step3: Create a service instance
  ```
    var serviceInst = &OrderService{}
    func GetOrderService() *OrderService {
      return serviceInst
    }
  ```
  
Step4: Register the service
  * Method 1: Use client.WithClientAware to register.
    ```
    cli = support.MustNewClient(app, cfg.DefaultStoreHandle, cfg.DefaultAccessToken, client.WithClientAware(order.GetOrderService()))
    ```
  * Method 2: Modify the source code in service_register.go.
    
  ``` see: service_register.go
    func GetClientAwares() []client.Aware {
      var awares = []client.Aware{
          order.GetOrderService(),
          // you can add service here
      }
      return awares
    }

  ```

### Webhooks verification

In order to be sure that a webhook is sent from SHOPLINE you could easily verify it with the VerifyWebhookRequest
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

### <span id="Test_for_Server">Server-side testing</span>

Server-side testing is primarily used by developers to verify app functionality in a local environment. The following sample code serves as a basic demonstration and should not be used directly in a production environment. Complete the code before using it in production.<br>

* If you have not yet registered on the SHOPLINE Partner Portal, refer to [Onboarding guidelines for SHOPLINE developers](https://developer.shopline.com/docs/apps/get-started/onboarding-guidelines-for-shopline-developer-s?version=v20251201)  to complete your registration.<br>
* If you have not yet created an app, refer to the guide on [Create an App](https://developer.shopline.com/docs/apps/application-management/creating-an-app?version=v20251201) to complete this process.

#### Step 1: Replace App Data

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

#### Step 2: Start the program

Execute server/main.go. If successful, port 80 will be started locally. <br>
The console prints "Server started on :80," indicating successful startup.

#### Step 3: Local hosts bind 1 test domain name

127.0.0.1 appdemo.myshopline.com <br>
Mac: /etc/hosts <br>
Windows: C:\Windows\System32\drivers\etc\hosts

#### Step 4: Set the app URL and app callback URL

1. Locate [Apps](https://developer.myshopline.com/app/list) in the [Partner Portal](https://developer.myshopline.com/) .
2. Click on the name of the app you want to test.
3. Click on **App settings** under **Basic settings**.
4. Set the **App URL** and **App callback** URL to:
      ○ App URL: http://appdemo.myshopline.com/install
      ○ App callback URL: http://appdemo.myshopline.com/auth/callback
5. Choose **Redirect** as the app loading mode because the **Embedded** mode must use the HTTPS protocol.


#### Step 5: App Receive the authorization code

1. Locate [Apps](https://developer.myshopline.com/app/list)  in the [Partner Portal](https://developer.myshopline.com/) 
2. Click on the name of the app you want to test.
3. Click on **Test App**.
4. Click on **Install App**. At this point, the system will request the app URL set in step four, the browser will automatically redirect to the platform authentication page, and finally, it will callback to the app callback URL you set.
5. Upon successful operation, **Auth callback received ... code: xxx**  will be printed. The code will be used later to exchange for an access token.

For more details, refer to [Receive the authorization code](https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step3) </br>
For a code example, see: server/main.go function 'CallbackHandler'

#### Step 6: App requests an access token

1. Find the TestCreateAccessToken function in oauth_test.go and replace the code variable in the function with the
  code obtained in step 5.
2. Executing the TestCreateAccessToken function will request the system to create an access token, and the result
  will be printed in the console if successful.
  For more details, see: [Request an access token](https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step-4-request-an-access-token)

```
// Create an access token
code := "code"
appKey := ""
app := manager.GetApp(appKey)
token, err := app.CreateAccessToken(context.Background(), code)
// Store the token in a database or cache for future use. 

```

#### Step 7: Use the access token to call an API

Use access token to call the [Create Product](https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201) API

1. Find app_config.go and replace the AccessToken variable value with the access token obtained in step 6.
2. Find create_product_test.go and execute the TestCreateProduct function. If successful, it will print "New product
  ID: xxx", eg: "New product ID: 16071495214036953630973380".
  
For more examples, see the xxx_test.go files in each package.  <br>
For more details, see: [Create a product](https://developer.shopline.com/docs/admin-rest-api/product/product/create-a-product?version=v20251201)

#### Step 8: App refresh access token

Access tokens expire periodically, so it is necessary to refresh them regularly to maintain access. Executing the TestRefreshAccessToken function in oauth_test.go will request the platform to refresh the access token. Upon success, the result will be printed to the console.<br>
For more details, see: [App refresh access token](https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step6)

```
// Refresh the access token
storeHandle := ""
appKey := ""
app := manager.GetApp(appKey)
token, err := app.RefreshAccessToken(context.Background(), storeHandle)

// Store the token in a database or cache for future use. 
```

## <span id="zh">中文</span>

对于使用 [SHOPLINE](https://developer.shopline.com) 的开发者来说，当前 SDK 封装了通用对接逻辑，旨在帮助开发者高效地构建应用，让你能更专注于业务功能的实现。

当前支持的能力：
1. 通过 [OAuth](https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20260301) 创建 Admin API 用到的 access tokens
2. 调用 [REST Admin API](https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/rest-admin-api/overview?version=v20260301)
3. 注册和处理 [Webhooks](https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301)

**注意**：当前 SDK 尚处于 Beta 阶段，功能持续完善中，已支持 OAuth 授权、商品、订单等核心 API，并提供了清晰的客户端初始化、请求构建与响应处理模式。我们欢迎社区贡献代码，你可通过提交 Pull Request 补充新的资源或端点。如果遇到尚未实现的 API，可以参照 <a href="#zh-use-your-own-model">使用你自己的数据模型</a> 自行封装。


### 安装 SDK
```
$ go get github.com/shoplineos/shopline-sdk-go
```

### 使用 SDK
```
import "github.com/shoplineos/shopline-sdk-go/client"
```


### 初始化 App 和 Client
```
  // 1. 创建应用实例
  appInstance := client.App{
      AppKey:      "",              
      AppSecret:   "",             
      Scope:       "read_products,write_products,read_orders,write_orders",
      RedirectUrl: "http://appdemo.myshopline.com/auth/callback",           
  }
  
  handle := "zwapptest"
  accessToken := ""
  
  // 2. 创建客户端实例
  // c := client.MustNewClient(appInstance, handle, accessToken)
  // appInstance.Client = c

  // use support to create client, it will regitster awares
  c := support.MustNewClient(appInstance, handle, accessToken)
  appInstance.Client = c

  // 3. 调用 API
  // 3.1 API 请求
  getProductCountAPIReq := &GetProductCountAPIReq{}
  
  // 3.2 指定 API response data
  apiResp := &GetProductCountAPIResp{}

  // 3.3 调用 API
  err := c.Call(context.Background(), apiReq, apiResp)
  fmt.Printf("count:%d", apiResp.Count)
    
```

### OAuth 授权

如果还没有 access token，可以通过 OAuth 流程来获取 access token，流程如下：

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

### 使用 access token 调用 APIs

```
  // 使用客户端调用 API
  // 1 指定 API request
  getProductCountAPIReq := &GetProductCountAPIReq{}
  
  // 2 指定 API 响应数据
  apiResp := &GetProductCountAPIResp{}

  // 3 调用 API
  err := client.Call(context.Background(), apiReq, apiResp)
  fmt.Printf("count:%d", apiResp.Count)
```

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

// 使用 service 调用 API 或者 直接使用 client.Call
apiResp, err := product.GetProductService().Create(context.Background(), apiReq)

```


### <span id="zh-use-your-own-model">使用你自己的数据模型</span>

目前为止不是所有的 API 都已经实现，你可以发起1个 Pull Request，或者自己实现数据模型对象。
下面这个例子是获取商品数量:

```
// 详细见：get_product_count.go
type GetProductCountAPIReq struct {
    client.BaseAPIRequest // 基础结构体
	Status       string `url:"status,omitempty"`
	CollectionId string `url:"collection_id,omitempty"`
	CreatedAtMin string `url:"created_at_min,omitempty"`
	CreatedAtMax string `url:"created_at_max,omitempty"`
	UpdatedAtMin string `url:"updated_at_min,omitempty"`
	UpdatedAtMax string `url:"updated_at_max,omitempty"`
}

func (req *GetProductCountAPIReq) Method() string {
	return "GET"
}

func (req *GetProductCountAPIReq) GetQuery() string {
	return req
}

func (req *GetProductCountAPIReq) Verify() error {
	// 验证请求参数
	return nil
}

func (req *GetProductCountAPIReq) Endpoint() string {
	endpoint := "products/count.json"
	return endpoint
}

type GetProductCountAPIResp struct {
	client.BaseAPIResponse
	Count int `json:"count"`
}

func GetProductsCount(c *client.Client, apiReq *GetProductCountAPIReq) (*GetProductCountAPIResp, error) {
    // 1. 指定 API 响应数据
	apiResp := &GetProductCountAPIResp{}

	// 2. 调用 API
	err := c.Call(context.Background(), apiReq, apiResp)

    return apiResp, err
}

```

### 实现你自己的服务接口 Service Interface（可选）

步骤1: 定义一个服务接口
  * See product_service.go or order_service.go
  ```
    type IOrderService interface {
        List(context.Context, *ListOrdersAPIReq) (*ListOrdersAPIResp, error)
        ListAll(context.Context, *ListOrdersAPIReq) ([]Order, error)
        // 你可以在这里添加更多方法
    }
  ```

步骤2: 定义一个 Service 结构体并实现服务接口
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

步骤3: 创建服务实例
  ```
    var serviceInst = &OrderService{}
    func GetOrderService() *OrderService {
      return serviceInst
    }
  ```


步骤4: 注册服务
  * 方法1: 使用 client.WithClientAware
    ```
    cli = support.MustNewClient(app, cfg.DefaultStoreHandle, cfg.DefaultAccessToken, client.WithClientAware(order.GetOrderService()))
    ```
  * 方法2: 修改 'service_register.go' 源代码

  ``` see: service_register.go
    func GetClientAwares() []client.Aware {
      var awares = []client.Aware{
          order.GetOrderService(),
          // 可以在这里添加更多服务
      }
      return awares
    }

  ```


### 验证 Webhooks
为了确保 Webhook 确实从 SHOPLINE 发送的，你可以使用 VerifyWebhookRequest 方法轻松验证。更多信息见：[Webhooks](https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview/?version=v20260301) </br>
例子:

```
// 详细见：server/main.go
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
    app := manager.GetDefaultApp()
    app.VerifyWebhookRequest(r)
    // 处理业务
}
```

### <span id="服务端测试">服务端测试</span>

服务端测试主要用于开发者在本地环境中验证应用功能，以下示例代码仅作为基础演示，不可直接用于生产环境。如需使用，请先完善代码。<br>

* 如尚未入驻 SHOPLINE 合作伙伴后台，参考 [SHOPLINE 开发者入驻指引](https://developer.shopline.com/zh-hans-cn/docs/apps/get-started/onboarding-guidelines-for-shopline-developer-s?version=v20251201) 完成入驻
* 如尚未创建应用，参考 [创建应用](https://developer.shopline.com/zh-hans-cn/docs/apps/application-management/creating-an-app?version=v20251201) 完成应用创建。

#### 步骤一：替换 App 数据

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

#### 步骤二：启动程序

执行 server/main.go，成功后，会在本地启动80端口。<br>
控制台打印“Server started on :80” 表示启动成功了。

#### 步骤三：本地 hosts 绑定1个测试域名

127.0.0.1 appdemo.myshopline.com <br>
Mac径路: /etc/hosts <br>
Windows路径: C:\Windows\System32\drivers\etc\hosts

#### 步骤四：在 合作伙伴后台 设置 应用地址 和 应用回调地址
1. 在 [合作伙伴后台](https://developer.shopline.com) 中找到 [应用](https://developer.myshopline.com/app/list)。
2. 点击需要测试的应用名称。
3. 点击 **基础设置** 下的 **应用设置**。
4. 设置 **应用地址** 和 **应用回调地址** 分别为 http://appdemo.myshopline.com/install 和  http://appdemo.myshopline.com/auth/callback。
5. **应用打开方式** 选择 外跳，因为 **内嵌** 模式必须是 HTTPS 协议。

#### 步骤五：应用接收授权码

1. 在 [合作伙伴后台](https://developer.shopline.com) -> [应用列表](https://developer.myshopline.com/app/list) 中找到 **应用**。
2. 点击需要测试的应用名称。
3. 点击 **应用测试**。
4. 点击 **安装应用**。此时，平台会请求步骤四设置的应用地址，浏览器自动跳转平台认证页面，最终回调至你设置的应用回调地址。
5. 操作成功后会打印 Auth callback received ... code: xxx，其中 code 后续会用来交换 access token。
   更多详情可参考 App 收到授权码。
   代码示例见 server/main.go，函数 CallbackHandler。


#### 步骤六：App 获取 access token

1. 找到 oauth_test.go 里的 TestCreateAccessToken 函数，用第5步获取到的 code，替换函数里的 code 变量。
2. 执行 TestCreateAccessToken 函数，会请求平台获取 access token，成功后会打印在控制台。
  可参考: [应用请求 access token](https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E5%85%AD%E6%AD%A5app-%E8%AF%B7%E6%B1%82-access-token)  <br>

```
// 创建 access token
appkey := "appkey"
code := "code"
app := manager.GetApp(appKey)
token, err := app.CreateAccessToken(context.Background(), code)

// 可以将 token 保存到 数据库 或者 缓存

```

#### 步骤七：使用 access token 调用接口

使用 access token
调用 [创建商品](https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/product/product/create-a-product?version=v20251201)
接口。

1. 找到 app_config.go，用第6步获取到的 access token，替换 AccessToken 变量值。
2. 找到 create_product_test.go，执行 TestCreateProduct 函数，成功后会打印 “新商品ID: xxx”，如：新商品ID:
  16071495214036953630973380。
  更多例子见：各个包下面的 xxx_test.go 文件。  <br>

#### 步骤八：App 刷新 access token

access token 每隔一段时间会过期，因此你需要定期刷新 access token。执行 oauth_test.go 里的 TestRefreshAccessToken
函数，会请求平台刷新 access token，成功后会打印在控制台。 <br>
详细见: [App 刷新 access token](https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E5%85%AB%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E5%88%B7%E6%96%B0-access-token)  <br>

```
// 刷新 access token
storeHandle := ""
appKey := ""
app := manager.GetApp(appKey)
token, err := app.RefreshAccessToken(context.Background(), storeHandle)

// 将 access token 存在数据库或者缓存里，方便后续使用
```
