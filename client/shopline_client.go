package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/common"
	"github.com/shoplineos/shopline-sdk-go/config"
	"github.com/shoplineos/shopline-sdk-go/model"
	"github.com/shoplineos/shopline-sdk-go/signature"
	"log"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

type Client struct {
	StoreHandle string

	// app config
	App App

	// Http client
	Client *http.Client

	// Access Token
	Token string

	// Base URL for API requests.
	// This is set on a per-store basis which means that each store must have its own Client.
	// eg: https://storeHandle.myshopline.com
	baseURL *url.URL

	// URL Prefix, defaults to "admin/openapi"
	PathPrefix string

	// API version you're currently using of the api, defaults to "config.DefaultAPIVersion"
	ApiVersion string

	// Enable signature calculation, default is false
	EnableSign bool

	Options *Options // Options
}

type App struct {
	AppKey      string
	AppSecret   string `json:"-"`
	Scope       string // app scope
	RedirectUrl string // oauth redirect Url

	Client *Client // API Client
}

// ShopLineRequestOptions Request options
type ShopLineRequestOptions struct {

	// Enable signature calculation, default is false
	EnableSign bool

	// Timeout(Optional)
	Timeout time.Duration

	// API version(Optional)
	ApiVersion string
}

// ShopLineRequest request params, pagination see detail：
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
type ShopLineRequest struct {
	Headers map[string]string       // http header
	Data    interface{}             // your own struct or an APIRequest, for http url query params or body params
	Options *ShopLineRequestOptions // option params
}

func (r *ShopLineRequest) isSignEnabled() bool {
	return r.Options != nil && r.Options.EnableSign
}

func (r *ShopLineRequest) isSignDisabled() bool {
	return r.Options != nil && !r.Options.EnableSign
}

func (r *ShopLineRequest) isApiVersionPresent() bool {
	return r.Options != nil && r.Options.ApiVersion != ""
}

func (r *ShopLineRequest) isTimeoutPresent() bool {
	return r.Options != nil && r.Options.Timeout > 0
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

	// API Response Data, the return type of the request when call APIs specify
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

const (
	// TimeoutInMillisecond default timeout time in millisecond
	TimeoutInMillisecond = 10 * 1000 * time.Millisecond
	defaultApiPathPrefix = config.DefaultApiPathPrefix
	defaultApiVersion    = config.DefaultAPIVersion
)

// AuthorizeUrl
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E4%B8%89%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E6%8E%88%E6%9D%83%E7%A0%81
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/#step2
func (app App) AuthorizeUrl(storeHandle string, scope string) (string, error) {
	redirectUri := url.QueryEscape(app.RedirectUrl)
	scope = app.resolveScope(scope)
	scope = url.QueryEscape(scope)
	authorizeUrl := fmt.Sprintf("https://%s.myshopline.com/admin/oauth-web/#/oauth/authorize?appKey=%s&responseType=code&scope=%s&redirectUri=%s", storeHandle, app.AppKey, scope, redirectUri)

	return authorizeUrl, nil
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

	shopLineReq := &ShopLineRequest{
		Options: &ShopLineRequestOptions{EnableSign: true},
		Data:    requestBody,
	}

	// 2. new http request
	httpReq, err := app.Client.NewHttpRequest(ctx, MethodPost, "admin/oauth/token/create", shopLineReq)
	if err != nil {
		return nil, err
	}

	// 3. resource
	tokenResponse := &TokenResponse{}

	// 4. execute
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
	shopLineReq := &ShopLineRequest{
		Options: &ShopLineRequestOptions{EnableSign: true},
	}

	// 2. new http request
	httpReq, err := app.Client.NewHttpRequest(ctx, MethodPost, "admin/oauth/token/refresh", shopLineReq)
	if err != nil {
		return nil, err
	}

	// 3. resource
	tokenResponse := &TokenResponse{}

	// 4. execute
	_, err = app.Client.executeHttpRequest(shopLineReq, httpReq, tokenResponse)
	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}

func MustNewClient(app App, storeHandle, token string, opts ...Option) *Client {
	c, err := NewClient(app, storeHandle, token, opts...)
	if err != nil {
		panic(err)
	}
	return c
}

func NewClient(app App, storeHandle, token string, opts ...Option) (*Client, error) {
	baseURL, err := url.Parse(common.GetStoreBaseUrl(storeHandle))
	if err != nil {
		return nil, err
	}

	c := &Client{
		Client: &http.Client{
			Timeout: TimeoutInMillisecond,
		},
		App:         app,
		StoreHandle: storeHandle,
		baseURL:     baseURL,
		Token:       token,
		ApiVersion:  defaultApiVersion,
		PathPrefix:  defaultApiPathPrefix,
	}

	for _, opt := range opts {
		opt(c)
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

func (c *Client) executeInternal(ctx context.Context, method HTTPMethod, endpoint string, request *ShopLineRequest, resource interface{}) (*ShopLineResponse, error) {
	_, _, err := c.verify(endpoint, method, request)
	if err != nil {
		return nil, err
	}

	relPath := c.resolveUrlPath(endpoint, request)

	httpReq, err := c.NewHttpRequest(ctx, method, relPath, request)
	if err != nil {
		return nil, err
	}

	return c.executeHttpRequest(request, httpReq, resource)
}

// return the http url's path '/admin/openapi/{version}/{relPath}'
// eg: /admin/openapi/v20251201/orders.json
func (c *Client) resolveUrlPath(relPath string, request *ShopLineRequest) string {
	if strings.HasPrefix(relPath, "/") {
		// make sure it's a relative path
		relPath = strings.TrimLeft(relPath, "/")
	}

	// {version}/{relPath}
	relPath = path.Join(c.resolveApiVersion(request), relPath)

	// /admin/openapi/{version}/{relPath}
	relPath = path.Join(c.PathPrefix, relPath)
	return relPath
}

func (c *Client) executeHttpRequest(request *ShopLineRequest, httpReq *http.Request, resource interface{}) (*ShopLineResponse, error) {
	// Call http request
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
	shopLineResp, err := buildShopLineResponse(resp, resource)
	if err != nil {
		return shopLineResp, err
	}

	c.logDetailIfNecessary(httpReq.Method, httpReq.URL.String(), request, shopLineResp)

	return shopLineResp, nil
}

func (c *Client) NewHttpRequest(ctx context.Context, method HTTPMethod, path string, request *ShopLineRequest) (*http.Request, error) {
	// build request URL
	// eg：requestURL := fmt.Sprintf("https://%s.myshopline.com/admin/openapi/%s/products/%s.json", shopHandle, ApiVersion, productId)
	requestURL, err := c.buildRequestUrl(method, path, request)
	if err != nil {
		log.Printf("Failed to build requestURL, path: %s, request: %v, err: %v\n", path, request, err)
		return nil, err
	}

	requestBodyJsonData, err := c.serializeBodyDataIfNecessary(method, request)
	if err != nil {
		log.Printf("Failed to serialize JSON, bodyParams:%v, err:%v\n", request.Data, err)
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
		return req.Options.Timeout
	}
	return TimeoutInMillisecond
}

// set Headers
func (c *Client) setHeaders(appKey string, appSecret string, httpReq *http.Request, wrapper *shopLineRequestWrapper) error {
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("appkey", appKey)
	httpReq.Header.Set("User-Agent", config.UserAgent)

	// create access Token & refresh access Token is not required
	if c.Token != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.Token)
	}

	timestamp := common.BuildTimestamp()
	httpReq.Header.Set("timestamp", timestamp)

	request := wrapper.shopLineRequest

	// Signature calculation
	if c.isSignEnabled(request) {
		sign, err := generateSign(appKey, appSecret, timestamp, wrapper.requestBodyJsonBytes)
		if err != nil {
			return err
		}
		log.Printf("Generate Sign enabled, sign: %s\n", sign)
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
func resolveAppKey(app App) string {
	if len(app.AppKey) > 0 {
		return app.AppKey
	}
	return config.DefaultAppKey
}

// First get the App.AppSecret in the request, if not exists then get the DefaultAppSecret in app_constants.go
func resolveAppSecret(app App) string {
	if len(app.AppSecret) > 0 {
		return app.AppSecret
	}
	return config.DefaultAppSecret
}

// generate sign string
func generateSign(appKey, appSecret, timestamp string, requestBodyJsonBytes []byte) (string, error) {
	bodyJsonString, err := buildBodyJsonString(requestBodyJsonBytes)
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

// Build shopline response
func buildShopLineResponse(httpResp *http.Response, resource interface{}) (*ShopLineResponse, error) {
	shopLineResp := &ShopLineResponse{}
	shopLineResp.StatusCode = httpResp.StatusCode

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

	pagination, err := parsePagination(shopLineResp.Link)
	if err != nil {
		return shopLineResp, err
	}
	shopLineResp.Pagination = pagination

	setCommonAPIRespData(shopLineResp)
	return shopLineResp, nil
}

// set common data to api response data
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

// set traceId to api response
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

// ResponseDecodingError occurs when the http response body from shopline not be parsed.
type ResponseDecodingError struct {
	Body    []byte
	Message string
	Status  int
}

func (e ResponseDecodingError) Error() string {
	return e.Message
}

func (c *Client) logDetailIfNecessary(method string, apiURL string, req *ShopLineRequest, resp *ShopLineResponse) {
	if c.IsLogDetailEnabled() {
		reqJsonData, _ := json.MarshalIndent(req, "", "  ")
		respJsonData, _ := json.MarshalIndent(resp, "", "  ")
		log.Printf("Request detail, Current AccessToken: %s\n method: %s\n apiURL: %s\n Request: %s\n Response: %s\n", c.Token, method, apiURL, reqJsonData, respJsonData)
	}

}

// verify request params
func (c *Client) verify(endpoint string, method HTTPMethod, request *ShopLineRequest) (string, string, error) {
	if request == nil {
		return "", "", fmt.Errorf("ShopLineRequest is required")
	}
	if method == "" {
		return "", "", fmt.Errorf("HTTP Method is required")
	}
	if endpoint == "" {
		return "", "", fmt.Errorf("API endpoint is required")
	}

	appKey := resolveAppKey(c.App)
	if appKey == "" {
		return "", "", fmt.Errorf("appKey is required")
	}

	appSecret := resolveAppSecret(c.App)
	if appSecret == "" {
		return "", "", fmt.Errorf("appSecret is required")
	}

	if request.Data != nil {
		if _, ok := (request.Data).(model.APIRequest); ok {
			err := (request.Data).(model.APIRequest).Verify()
			if err != nil {
				return "", "", err
			}
		}
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
func (c *Client) buildRequestUrl(method HTTPMethod, relPath string, request *ShopLineRequest) (string, error) {

	parsedURL, err := c.ResolveURL(relPath)
	if err != nil {
		return "", err
	}

	// Get only
	if method == MethodGet && request.Data != nil {
		optionsQuery, err := query.Values(request.Data)
		if err != nil {
			return "", err
		}

		for k, values := range parsedURL.Query() {
			for _, v := range values {
				if v != "" { // filter the empty value
					optionsQuery.Add(k, v)
				}
			}
		}
		parsedURL.RawQuery = optionsQuery.Encode()
	}

	requestURL := parsedURL.String()
	log.Printf("Final to build requestURL: %s\n", requestURL)

	return requestURL, nil
}

func (c *Client) resolveApiVersion(req *ShopLineRequest) string {
	if req.isApiVersionPresent() {
		return req.Options.ApiVersion
	}
	return c.ApiVersion
}

// body params serialize to json bytes
func (c *Client) serializeBodyDataIfNecessary(method HTTPMethod, request *ShopLineRequest) ([]byte, error) {
	if method == MethodGet || request == nil || request.Data == nil {
		return nil, nil
	}

	body := request.Data

	return json.Marshal(body)
}

func (c *Client) GetAppKey() string {
	return c.App.AppKey
}

func (c *Client) IsLogDetailEnabled() bool {
	return c.Options != nil && c.Options.EnableLogDetail
}

func (c *Client) isSignEnabled(request *ShopLineRequest) bool {
	if request.isSignDisabled() { // in blacklist
		return false
	}

	return request.isSignEnabled() || c.EnableSign
}

// ResolveURL resolve the call api http url
// https://{storeHandle}.myshopline.com/{pathPrefix}/{version}/{endpoint}
//
//eg:https://storeHandle.myshopline.com/admin/openapi/v20251201/orders/count.j
//eg:https://storeHandle.myshopline.com/admin/openapi/v20251201/orders/count.json
func (c *Client) ResolveURL(relPath string) (*url.URL, error) {
	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}
	parsedURL := c.baseURL.ResolveReference(rel)
	return parsedURL, nil
}
