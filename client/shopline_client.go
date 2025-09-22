package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/common"
	"github.com/shoplineos/shopline-sdk-go/config"
	"github.com/shoplineos/shopline-sdk-go/signature"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

type Client struct {
	StoreHandle string

	// Enable signature calculation, default is false
	EnableSign bool

	// app config
	App App

	// http client
	Client *http.Client

	// access Token
	Token string

	// Base URL for API requests.
	// This is set on a per-store basis which means that each store must have its own Client.
	// eg: https://shopName.myshopline.com
	baseURL *url.URL

	// URL Prefix, defaults to "admin/openapi"
	pathPrefix string

	// API version you're currently using of the api, defaults to "config.DefaultAPIVersion"
	apiVersion string
}

type App struct {
	AppKey      string
	AppSecret   string `json:"-"`
	RedirectUrl string // oauth redirect Url
	Scope       string // app scope

	Client *Client // API Client
}

type ShopLineRequestOption struct {

	// Enable signature calculation, default is false
	EnableSign bool

	// Timeout(option)
	Timeout time.Duration

	// API version(option)
	ApiVersion string
}

// ShopLineRequest request params, pagination see detail：
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
type ShopLineRequest struct {
	Headers map[string]string // http header
	Query   interface{}       // struct，http query params
	Body    interface{}       // struct，http body params

	Option *ShopLineRequestOption // option params
}

func (r *ShopLineRequest) isSignEnabled() bool {
	return r.Option != nil && r.Option.EnableSign
}

func (r *ShopLineRequest) isApiVersionPresent() bool {
	return r.Option != nil && r.Option.ApiVersion != ""
}

func (r *ShopLineRequest) isTimeoutPresent() bool {
	return r.Option != nil && r.Option.Timeout > 0
}

type CommonAPIRespData struct {
	// traceId
	TraceId string
}

// ShopLineResponse response
type ShopLineResponse struct {
	StatusCode int

	Errors string

	// traceId
	TraceId string

	//Headers    map[string]string

	// ResponseData
	Data interface{}

	// Pagination, see detail：
	// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
	// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
	Link string

	Pagination *Pagination
}

func (resp *ShopLineResponse) IsSuccess() bool {
	return resp.StatusCode == http.StatusOK
}

// ShopLineRequest Wrapper
type shopLineRequestWrapper struct {
	shopLineRequest      *ShopLineRequest
	requestBodyJsonBytes []byte // nil able
}

// Pagination of results
type Pagination struct {
	Next     *ListOptions
	Previous *ListOptions
}

// ListOptions
// General list options that can be used for most collections of entities.
type ListOptions struct {
	// PageInfo is used with new pagination search.
	PageInfo string `url:"page_info,omitempty"`

	SinceId *uint64 `url:"since_id,omitempty"`
	Limit   int     `url:"limit,omitempty"`
	Fields  string  `url:"fields,omitempty"`
}

// TokenResponse access Token struct
type TokenResponse struct {
	Code     int    `json:"code"`
	I18nCode string `json:"i18nCode"`
	Message  string `json:"message"`
	Data     struct {
		AccessToken string `json:"accessToken"`
		ExpiresTime string `json:"expireTime"`
		Scope       string `json:"scope"`
		//RefreshToken string `json:"refresh_token"`
	} `json:"data"`
}

// TimeoutInMillisecond default timeout time in millisecond
const (
	TimeoutInMillisecond = 10 * 1000 * time.Millisecond
	defaultApiPathPrefix = config.DefaultApiPathPrefix
	defaultApiVersion    = config.DefaultAPIVersion
)

// AuthorizeUrl
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E4%B8%89%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E6%8E%88%E6%9D%83%E7%A0%81
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/#step2
func (app App) AuthorizeUrl(storeHandle string, scope string) (string, error) {
	//shopUrl, err := url.Parse(common.GetStoreBaseUrl(storeHandle))
	//if err != nil {
	//	return "", err
	//}

	//scope := "read_products,write_products,read_orders,write_orders"

	//redirectUri := fmt.Sprintf("http://%s/auth/callback", shopUrl)

	//shopUrl.Path = "/admin/oauth-web/#/oauth/authorize"
	//query := shopUrl.Query()
	//query.Set("appKey", app.AppKey)
	//query.Set("redirectUri", app.RedirectUrl)
	//query.Set("scope", app.resolveScope(scope))
	//query.Set("responseType", "code")
	//shopUrl.RawQuery = query.Encode()

	redirectUri := url.QueryEscape(app.RedirectUrl)
	scope = app.resolveScope(scope)
	scope = url.QueryEscape(scope)
	baseUrl := fmt.Sprintf("https://%s.myshopline.com/admin/oauth-web/#/oauth/authorize?appKey=%s&responseType=code&scope=%s&redirectUri=%s", storeHandle, app.AppKey, scope, redirectUri)

	return baseUrl, nil
}

func (app App) VerifySign(params url.Values, receivedSign string) bool {
	isSignValid := signature.VerifySign(app.AppSecret, params, receivedSign)
	return isSignValid
}

// VerifyWebhookMessage Verify a message against a message HMAC
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
func (app App) VerifyWebhookMessage(message, messageMAC string) bool {
	return signature.VerifyWebhookMessage(app.AppSecret, message, messageMAC)
}

// VerifyWebhookRequest Verify a Webhook http request, sent by shopline.
// The body of the request is still readable after invoking the method.
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
func (app App) VerifyWebhookRequest(r *http.Request) bool {
	return signature.VerifyWebhookRequest(app.AppSecret, r)
}

func (app App) resolveScope(scope string) string {
	if scope != "" {
		return scope
	}
	return app.Scope
}

func (app App) GetStoreHandle() string {
	return app.Client.StoreHandle
}

// CreateAccessToken
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#%E7%AC%AC%E5%85%AD%E6%AD%A5app-%E8%AF%B7%E6%B1%82-access-token
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20251201#step-4-request-an-access-token
func (app App) CreateAccessToken(ctx context.Context, code string) (*TokenResponse, error) {

	// 1. build request
	requestBody := map[string]string{
		"code": code,
	}

	tokenResponse := &TokenResponse{}

	shopLineReq := &ShopLineRequest{
		Option: &ShopLineRequestOption{EnableSign: true},
		Body:   requestBody,
	}

	// 2. new http request
	httpReq, err := app.Client.NewHttpRequest(ctx, MethodPost, "admin/oauth/token/create", shopLineReq)
	if err != nil {
		return nil, err
	}

	// 3. execute
	_, err = app.Client.executeHttpRequest(shopLineReq, httpReq, tokenResponse)
	if err != nil {
		return nil, err
	}
	return tokenResponse, nil
}

// RefreshAccessToken
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20251201#step6
func (app App) RefreshAccessToken(ctx context.Context, storeHandle string) (*TokenResponse, error) {
	err := verifyForRefreshAccessToken(app.AppKey, app.AppSecret, storeHandle)
	if err != nil {
		return nil, err
	}

	// 1. build request
	tokenResponse := &TokenResponse{}

	shopLineReq := &ShopLineRequest{
		Option: &ShopLineRequestOption{EnableSign: true},
	}

	// 2. new http request
	httpReq, err := app.Client.NewHttpRequest(ctx, MethodPost, "admin/oauth/token/refresh", shopLineReq)
	if err != nil {
		return nil, err
	}

	// 3. execute
	_, err = app.Client.executeHttpRequest(shopLineReq, httpReq, tokenResponse)
	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}

func MustNewClient(app App, storeHandle, token string) *Client {
	c, err := NewClient(app, storeHandle, token)
	if err != nil {
		panic(err)
	}
	return c
}

func NewClient(app App, storeHandle, token string) (*Client, error) {
	baseURL, err := url.Parse(common.GetStoreBaseUrl(storeHandle))
	if err != nil {
		return nil, err
	}

	c := &Client{
		Client: &http.Client{
			Timeout: time.Millisecond * TimeoutInMillisecond,
		},
		App:         app,
		StoreHandle: storeHandle,
		baseURL:     baseURL,
		Token:       token,
		pathPrefix:  defaultApiPathPrefix,
		apiVersion:  defaultApiVersion,
	}

	return c, nil
}

// Put performs a PUT request for the given endpoint and saves the result in the given resource.
// accessToken:
//
//	中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201#%E7%AC%AC%E4%B8%83%E6%AD%A5app-%E6%94%B6%E5%88%B0-access-token
//	en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#step-5-receive-the-access-token
//
// endpoint: API request endpoint, eg: API "https://xx.myshopline.com/admin/openapi/v20251201/orders.json" 's endpoint is "orders.json"
// request: ShopLineRequest
// resource : API response, To specify the return type of the request
func (c *Client) Put(ctx context.Context, endpoint string, request *ShopLineRequest, resource interface{}) (*ShopLineResponse, error) {
	return c.Execute(ctx, MethodPut, endpoint, request, resource)
}

// Delete performs a DELETE request for the given endpoint and saves the result in the given resource.
// accessToken:
//
//	中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201#%E7%AC%AC%E4%B8%83%E6%AD%A5app-%E6%94%B6%E5%88%B0-access-token
//	en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#step-5-receive-the-access-token
//
// endpoint: API request endpoint, eg: API "https://xx.myshopline.com/admin/openapi/v20251201/orders.json" 's endpoint is "orders.json"
// request: ShopLineRequest
// resource : API response, To specify the return type of the request
func (c *Client) Delete(ctx context.Context, endpoint string, request *ShopLineRequest, resource interface{}) (*ShopLineResponse, error) {
	return c.Execute(ctx, MethodDelete, endpoint, request, resource)
}

// Post performs a POST request for the given endpoint and saves the result in the given resource.
// accessToken:
//
//	中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201#%E7%AC%AC%E4%B8%83%E6%AD%A5app-%E6%94%B6%E5%88%B0-access-token
//	en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#step-5-receive-the-access-token
//
// endpoint: API request endpoint, eg: API "https://xx.myshopline.com/admin/openapi/v20251201/orders.json" 's endpoint is "orders.json"
// request: ShopLineRequest
// resource : API response, To specify the return type of the request
func (c *Client) Post(ctx context.Context, endpoint string, request *ShopLineRequest, resource interface{}) (*ShopLineResponse, error) {
	return c.Execute(ctx, MethodPost, endpoint, request, resource)
}

// Get performs a GET request for the given endpoint and saves the result in the given resource.
// accessToken:
//
//	中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201#%E7%AC%AC%E4%B8%83%E6%AD%A5app-%E6%94%B6%E5%88%B0-access-token
//	en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#step-5-receive-the-access-token
//
// endpoint: API request endpoint, eg: API "https://xx.myshopline.com/admin/openapi/v20251201/orders.json" 's endpoint is "orders.json"
// request: ShopLineRequest
// resource : API response, To specify the return type of the request
func (c *Client) Get(ctx context.Context, endpoint string, request *ShopLineRequest, resource interface{}) (*ShopLineResponse, error) {
	return c.Execute(ctx, MethodGet, endpoint, request, resource)
}

// executeWithoutToken create access Token and refresh access Token
//func (c *Client) executeWithoutToken(ctx context.Context, method HTTPMethod, path string, request *ShopLineRequest) (*ShopLineResponse, error) {
//	return c.executeInternal(ctx, method, path, request)
//}

// Execute performs a http request for the given endpoint and saves the result in the given resource.
// accessToken:
//
//	中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201#%E7%AC%AC%E4%B8%83%E6%AD%A5app-%E6%94%B6%E5%88%B0-access-token
//	en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#step-5-receive-the-access-token
//
// method: HTTP Method
// endpoint: API request endpoint, eg: orders.json
// request: ShopLineRequest
// resource : API response, To specify the return type of the request
func (c *Client) Execute(ctx context.Context, method HTTPMethod, endpoint string, request *ShopLineRequest, resource interface{}) (*ShopLineResponse, error) {
	return c.executeInternal(ctx, method, endpoint, request, resource)
}

func (c *Client) executeInternal(ctx context.Context, method HTTPMethod, relPath string, request *ShopLineRequest, resource interface{}) (*ShopLineResponse, error) {
	_, _, err := c.verify(relPath, method, request)
	if err != nil {
		return nil, err
	}

	relPath = c.resolveUrlPath(relPath, request)

	httpReq, err := c.NewHttpRequest(ctx, method, relPath, request)
	if err != nil {
		return nil, err
	}

	return c.executeHttpRequest(request, httpReq, resource)
}

// return the http url's path '/admin/openapi/{version}/{relPath}'
// eg: /admin/openapi/20251201/orders.json
func (c *Client) resolveUrlPath(relPath string, request *ShopLineRequest) string {
	if strings.HasPrefix(relPath, "/") {
		// make sure it's a relative path
		relPath = strings.TrimLeft(relPath, "/")
	}

	// {version}/{relPath}
	relPath = path.Join(c.resolveApiVersion(request), relPath)

	// /admin/openapi/{version}/{relPath}
	relPath = path.Join(c.pathPrefix, relPath)
	return relPath
}

func (c *Client) executeHttpRequest(request *ShopLineRequest, httpReq *http.Request, resource interface{}) (*ShopLineResponse, error) {
	// invoke http request
	timeout := resolveTimeout(request)
	client := c.Client // &http.Client{Timeout: timeout}
	client.Timeout = timeout
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Printf("Failed to send request: %v\n", err)
		return nil, err
	}

	defer resp.Body.Close()

	log.Printf("Execute request finished！status: %d\n", resp.StatusCode)

	// build response
	shopLineResp, err := buildShopLineResponse(request, resp, resource)
	if err != nil {
		return shopLineResp, err
	}

	c.logDetailIfNecessary(httpReq.Method, httpReq.URL.String(), request, shopLineResp)

	return shopLineResp, nil
}

func (c *Client) NewHttpRequest(ctx context.Context, method HTTPMethod, path string, request *ShopLineRequest) (*http.Request, error) {
	// build request URL
	// eg：requestURL := fmt.Sprintf("https://%s.myshopline.com/admin/openapi/%s/products/%s.json", shopHandle, ApiVersion, productId)
	requestURL, err := c.buildFinalRequestUrl(path, request)
	if err != nil {
		log.Printf("Failed to build requestURL, path: %s, request: %v, err: %v\n", path, request, err)
		return nil, err
	}

	requestBodyJsonData, err := c.buildRequestBodyJsonData(request)
	if err != nil {
		log.Printf("Failed to serialize JSON, bodyParams:%v, err:%v\n", request.Body, err)
		return nil, err
	}

	// create HTTP request
	httpReq, err := http.NewRequest(string(method), requestURL, bytes.NewBuffer(requestBodyJsonData))
	if err != nil {
		log.Printf("Failed to create request: %v\n", err)
		return nil, err
	}

	shopLineRequestWrapper := &shopLineRequestWrapper{
		shopLineRequest:      request,
		requestBodyJsonBytes: requestBodyJsonData,
	}

	// set headers
	err = c.setHeaders(c.App.AppKey, c.App.AppSecret, httpReq, shopLineRequestWrapper)
	if err != nil {
		return nil, err
	}
	return httpReq, nil
}

func resolveTimeout(req *ShopLineRequest) time.Duration {
	if req.isTimeoutPresent() {
		return req.Option.Timeout
	}
	return TimeoutInMillisecond
}

// set Headers
func (c *Client) setHeaders(appKey string, appSecret string, httpReq *http.Request, requestWrapper *shopLineRequestWrapper) error {
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("appkey", appKey)
	httpReq.Header.Set("User-Agent", config.UserAgent)

	// create access Token & refresh access Token is not required
	if c.Token != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.Token)
	}

	timestamp := buildTimestamp()
	httpReq.Header.Set("timestamp", timestamp)

	request := requestWrapper.shopLineRequest

	// Signature calculation
	if request.isSignEnabled() {
		sign, err := buildSign(appKey, appSecret, timestamp, requestWrapper)
		if err != nil {
			return err
		}
		log.Printf("Build Sign enabled, sign: %s\n", sign)
		httpReq.Header.Set("sign", sign)
	}

	if request.Headers != nil {
		for key, value := range request.Headers {
			httpReq.Header.Set(key, value)
		}
	}

	return nil
}

// First get the App.AppKey in the request, if not exists then get the DefaultAppKey in app_constants.go
func resolveAppKey(appConfig App) string {
	if len(appConfig.AppKey) > 0 {
		return appConfig.AppKey
	}
	return config.DefaultAppKey
}

// First get the App.AppSecret in the request, if not exists then get the DefaultAppSecret in app_constants.go
func resolveAppSecret(appConfig App) string {
	if len(appConfig.AppSecret) > 0 {
		return appConfig.AppSecret
	}
	return config.DefaultAppSecret
}

// build sign string
func buildSign(appKey, appSecret, timestamp string, requestWrapper *shopLineRequestWrapper) (string, error) {
	bodyJsonString, err := buildBodyJsonString(requestWrapper.requestBodyJsonBytes)
	if err != nil {
		return "", err
	}

	sign := signature.GenerateSign(appKey, bodyJsonString, timestamp, appSecret)
	return sign, nil
}

// Build params convert to json string
func buildBodyJsonString(bodyParams []byte) (string, error) {
	if bodyParams == nil {
		return "", nil
	}

	//body, err := json.Marshal(bodyParams)
	//if err != nil {
	//	return "", err
	//}
	return string(bodyParams), nil
}

// Build timestamp
func buildTimestamp() string {
	timestamp := time.Now().Unix()
	timestampInt := strconv.FormatInt(timestamp, 13)
	return timestampInt
}

// Build shopline response
func buildShopLineResponse(shopLineRequest *ShopLineRequest, httpResp *http.Response, resource interface{}) (*ShopLineResponse, error) {
	shopLineResp := &ShopLineResponse{}
	shopLineResp.StatusCode = httpResp.StatusCode

	// Pagination, see detail：
	// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
	// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/paging-mechanism/?version=v20251201

	link := resolveLinkHeader(httpResp.Header)
	shopLineResp.Link = link

	shopLineResp.TraceId = httpResp.Header.Get("traceId")

	err := CheckHttpResponseError(httpResp)
	if err != nil {
		shopLineResp.Errors = err.Error()
		return shopLineResp, err
	}

	if err := json.NewDecoder(httpResp.Body).Decode(&resource); err != nil {
		//respData := &map[string]any{}
		//json.NewDecoder(httpResp.Body).Decode(respData)
		log.Printf("Failed to parse response body, statusCode: %d, body: %v, err: %v\n", httpResp.StatusCode, resource, err)

		return shopLineResp, err
	}
	shopLineResp.Data = resource

	pagination, err := parsePaginationIfNecessary(shopLineResp.Link)
	if err != nil {
		return shopLineResp, err
	}
	shopLineResp.Pagination = pagination

	setCommonAPIRespData(shopLineResp)
	return shopLineResp, nil
}

func setCommonAPIRespData(shopLineResp *ShopLineResponse) {
	apiResp := shopLineResp.Data
	if apiResp == nil {
		return
	}

	apiRespVal := reflect.ValueOf(apiResp)

	if apiRespVal.Kind() == reflect.Ptr {
		apiRespVal = apiRespVal.Elem()
	}

	if apiRespVal.Kind() != reflect.Struct {
		log.Printf("Invalid response data type(must struct ptr or struct): %T\n", apiResp)
		return
	}

	typ := apiRespVal.Type()

	setTraceId2APIResp(shopLineResp, typ, apiRespVal)

}

func setTraceId2APIResp(shopLineResp *ShopLineResponse, typ reflect.Type, apiRespVal reflect.Value) {
	_, ok := typ.FieldByName("TraceId")
	if ok {
		//log.Printf("TraceId name: %s\n", f.Name)
		traceFieldValue := apiRespVal.FieldByName("TraceId")
		traceFieldValue.SetString(shopLineResp.TraceId)
	}
}

// For Pagination
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
func resolveLinkHeader(header http.Header) string {
	if header.Get("link") != "" {
		return header.Get("link")
	}
	if header.Get("Link") != "" {
		return header.Get("Link")
	}
	return ""
}

func CheckHttpResponseError(resp *http.Response) error {
	// 200 <= StatusCode < 300
	if http.StatusOK <= resp.StatusCode && resp.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	shoplineError := struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Errors  interface{} `json:"errors"`
	}{}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// empty body, this probably means shopline returned an error with no response body
	if len(bodyBytes) > 0 {
		err := json.Unmarshal(bodyBytes, &shoplineError)
		if err != nil {
			return ResponseDecodingError{
				Body:    bodyBytes,
				Message: err.Error(),
				Status:  resp.StatusCode,
			}
		}
	}

	// Create the response error from the shopline error.
	responseError := ResponseError{
		Status:  resp.StatusCode,
		Code:    shoplineError.Code,
		Message: shoplineError.Message,
	}

	if shoplineError.Errors == nil {
		return wrapSpecificError(resp, responseError)
	}

	switch reflect.TypeOf(shoplineError.Errors).Kind() {
	case reflect.String:
		// Single string, use as message
		responseError.Message = shoplineError.Errors.(string)
	case reflect.Slice:
		// An array, parse each entry as a string and join them on the message
		// json always serializes JSON arrays into []interface{}
		for _, elem := range shoplineError.Errors.([]interface{}) {
			responseError.Errors = append(responseError.Errors, fmt.Sprint(elem))
		}
		responseError.Message = strings.Join(responseError.Errors, ", ")
	case reflect.Map:
		// A map, parse each error for each key in the map.
		// json always serializes into map[string]interface{} for objects
		for k, v := range shoplineError.Errors.(map[string]interface{}) {
			switch reflect.TypeOf(v).Kind() {
			// Check to make sure the interface is a slice
			// json always serializes JSON arrays into []interface{}
			case reflect.Slice:
				for _, elem := range v.([]interface{}) {
					// If the primary message of the response error is not set, use
					// any message.
					if responseError.Message == "" {
						responseError.Message = fmt.Sprintf("%v: %v", k, elem)
					}
					topicAndElem := fmt.Sprintf("%v: %v", k, elem)
					responseError.Errors = append(responseError.Errors, topicAndElem)
				}
			case reflect.String:
				elem := v.(string)
				if responseError.Message == "" {
					responseError.Message = fmt.Sprintf("%v: %v", k, elem)
				}
				topicAndElem := fmt.Sprintf("%v: %v", k, elem)
				responseError.Errors = append(responseError.Errors, topicAndElem)
			}
		}
	}

	return responseError
}

func wrapSpecificError(r *http.Response, err ResponseError) error {

	if err.Status == http.StatusNotAcceptable {
		err.Message = http.StatusText(err.Status)
	}

	return err
}

// ResponseError
// A general http response error that follows a similar layout to shopline's response errors
type ResponseError struct {
	Status  int
	Code    string
	Message string
	Errors  []string
}

// GetStatus returns http response status
func (e ResponseError) GetStatus() int {
	return e.Status
}

// GetMessage returns response error message
func (e ResponseError) GetMessage() string {
	return e.Message
}

// GetErrors returns response errors
func (e ResponseError) GetErrors() []string {
	return e.Errors
}

func (e ResponseError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	sort.Strings(e.Errors)
	s := strings.Join(e.Errors, ", ")

	if s != "" {
		return s
	}

	return "Unknown Error"
}

// ResponseDecodingError occurs when the http response body from shopline not be parsed.
type ResponseDecodingError struct {
	Body    []byte
	Message string
	Status  int
}

func (e ResponseDecodingError) Error() string {
	return e.Message
}

// linkRegex is used to parse the pagination link from shopline API search results.
var linkRegex = regexp.MustCompile(`^ *<([^>]+)>; rel="(previous|next)" *$`)

// Pagination, see detail：
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
func parsePaginationIfNecessary(linkHeader string) (*Pagination, error) {
	if linkHeader == "" {
		return nil, nil
	}
	pagination := new(Pagination)

	for _, link := range strings.Split(linkHeader, ",") {
		match := linkRegex.FindStringSubmatch(link)
		// Make sure the link is not empty or invalid
		if len(match) != 3 {
			// We expect 3 values:
			// match[0] = full match
			// match[1] is the URL and match[2] is either 'previous' or 'next'
			err := ResponseDecodingError{
				Message: "could not extract pagination link header",
			}
			return nil, err
		}

		rel, err := url.Parse(match[1])
		if err != nil {
			err = ResponseDecodingError{
				Message: "pagination does not contain a valid URL",
			}
			return nil, err
		}

		params, err := url.ParseQuery(rel.RawQuery)
		if err != nil {
			return nil, err
		}

		paginationListOptions := ListOptions{}

		paginationListOptions.PageInfo = params.Get("page_info")
		if paginationListOptions.PageInfo == "" {
			err = ResponseDecodingError{
				Message: "page_info is missing",
			}
			return nil, err
		}

		limit := params.Get("limit")
		if limit != "" {
			paginationListOptions.Limit, err = strconv.Atoi(params.Get("limit"))
			if err != nil {
				return nil, err
			}
		}

		fields := params.Get("fields")
		if fields != "" {
			paginationListOptions.Fields = params.Get("fields")
		}

		// 'rel' is either next or previous
		if match[2] == "next" {
			pagination.Next = &paginationListOptions
		} else {
			pagination.Previous = &paginationListOptions
		}
	}

	return pagination, nil
}

func (c *Client) logDetailIfNecessary(method string, apiURL string, req *ShopLineRequest, resp *ShopLineResponse) {
	if config.LogDetailEnabled {
		reqJsonData, _ := json.MarshalIndent(req, "", "  ")
		respJsonData, _ := json.MarshalIndent(resp, "", "  ")
		log.Printf("Current AccessToken: %s\n method: %s\n apiURL: %s\n Request: %s\n Response: %s\n", c.Token, method, apiURL, reqJsonData, respJsonData)
	}

}

// verify request params
func (c *Client) verify(url string, method HTTPMethod, request *ShopLineRequest) (string, string, error) {
	if request == nil {
		return "", "", fmt.Errorf("ShopLineRequest is required")
	}
	if method == "" {
		return "", "", fmt.Errorf("HTTP Method is required")
	}
	if url == "" {
		return "", "", fmt.Errorf("url is required")
	}

	appKey := resolveAppKey(c.App)
	if appKey == "" {
		return "", "", fmt.Errorf("appKey is required")
	}

	appSecret := resolveAppSecret(c.App)
	if appSecret == "" {
		return "", "", fmt.Errorf("appSecret is required")
	}
	return appKey, appSecret, nil
}

func verifyForRefreshAccessToken(appkey, appSecret, shopHandle string) error {
	if appkey == "" {
		return fmt.Errorf("appKey is required")
	}
	if appSecret == "" {
		return fmt.Errorf("appSecret is required")
	}

	if shopHandle == "" {
		return fmt.Errorf("shopHandle is required")
	}
	return nil
}

// Add the request query parameters to the http query parameters
func (c *Client) buildFinalRequestUrl(relPath string, request *ShopLineRequest) (string, error) {

	rel, err := url.Parse(relPath)
	if err != nil {
		return "", err
	}

	parsedURL := c.baseURL.ResolveReference(rel)

	if request.Query != nil {
		optionsQuery, err := query.Values(request.Query)
		if err != nil {
			return "", err
		}

		for k, values := range parsedURL.Query() {
			for _, v := range values {
				optionsQuery.Add(k, v)
			}
		}
		parsedURL.RawQuery = optionsQuery.Encode()
	}

	//parsedURL, err := url.Parse(baseURL)
	//if err != nil {
	//	log.Printf("Failed to parse URL, baseURL: %s, err: %v\n", baseURL, err)
	//	return baseURL, err
	//}

	requestURL := parsedURL.String()
	log.Printf("Final requestURL: %s\n", requestURL)

	return requestURL, nil
}

func (c *Client) resolveApiVersion(req *ShopLineRequest) string {
	if req.isApiVersionPresent() {
		return req.Option.ApiVersion
	}
	return defaultApiVersion
}

// body params convert to json string
func (c *Client) buildRequestBodyJsonData(request *ShopLineRequest) ([]byte, error) {
	if request == nil || request.Body == nil {
		return nil, nil
	}

	body := request.Body

	return json.Marshal(body)
}

func (c *Client) GetAppKey() string {
	return c.App.AppKey
}
