package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

type IClient interface {
	Call(ctx context.Context, req APIRequest, resource interface{}) error
	NewHttpRequest(ctx context.Context, method HTTPMethod, path string, request *ShopLineRequest) (*http.Request, error)
}

type Client struct {
	StoreHandle string

	// App config
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

	// API version you're currently using of the api
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

// RequestOptions Request options
type RequestOptions struct {

	// Enable signature calculation, default is false
	EnableSign bool

	// Timeout(Optional)
	Timeout time.Duration

	// API version(Optional)
	ApiVersion string

	// When call an API successful, some APIs will return empty body
	// eg:https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/payments-app-api/merchant-activation-successful-notification?version=v20251201
	NotDecodeBody bool
}

// ShopLineRequest request parameters
// pagination for more details, see：
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
type ShopLineRequest struct {
	Headers map[string]string // Http header
	Query   interface{}       // Your own struct or an APIRequest, for http url query params
	Data    interface{}       // Your own struct or an APIRequest, for http body params
	Options *RequestOptions   // Option parameters
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

func (r *ShopLineRequest) NotDecodeBody() bool {
	return r.Options != nil && r.Options.NotDecodeBody
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

	// API response resource, the return type of the request when call APIs specify
	Data interface{}

	// Pagination, for more details, see：
	// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
	// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/paging-mechanism?version=v20251201
	Link string

	Pagination *Pagination

	ResponseBody string
}

func (resp *ShopLineResponse) HasNext() bool {
	return resp.Pagination != nil && resp.Pagination.Next != nil
}

func (resp *ShopLineResponse) IsSuccess() bool {
	// 200 <= StatusCode < 300
	// return http.StatusOK <= resp.StatusCode && resp.StatusCode < http.StatusMultipleChoices
	return resp.StatusCode == http.StatusOK
}

// ShopLineRequest Wrapper
type shopLineRequestWrapper struct {
	shopLineRequest      *ShopLineRequest
	requestBodyJsonBytes []byte // nil able
}

const (
	// TimeoutInMillisecond default timeout time in millisecond
	TimeoutInMillisecond = 10 * 1000 * time.Millisecond
	defaultApiPathPrefix = "admin/openapi"
)

func MustNewClient(app App, storeHandle, token string, opts ...Option) *Client {
	c, err := NewClient(app, storeHandle, token, opts...)
	if err != nil {
		panic(err)
	}
	return c
}

func NewClient(app App, storeHandle, token string, opts ...Option) (*Client, error) {
	return NewClientWithAwares(app, storeHandle, token, nil, opts...)
}

func MustNewClientWithAwares(app App, storeHandle, token string, awares []Aware, opts ...Option) *Client {
	if len(awares) == 0 {
		panic("The client awares is nil or empty, please see 'awares_loader.go'")
	}

	c, err := NewClientWithAwares(app, storeHandle, token, awares, opts...)
	if err != nil {
		panic(err)
	}
	return c
}

func NewClientWithAwares(app App, storeHandle, token string, awares []Aware, opts ...Option) (*Client, error) {
	baseURL, err := url.Parse(GetStoreBaseUrl(storeHandle))
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
		ApiVersion:  DefaultApiVersion,
		PathPrefix:  defaultApiPathPrefix,
	}

	for _, opt := range opts {
		opt(c)
	}

	for _, aware := range awares {
		aware.SetClient(c)
	}
	return c, nil
}

// Call an API
// resource : An API response resource, to specify the return type of the request, an APIResponse or your own resource
func (c *Client) Call(ctx context.Context, req APIRequest, resource interface{}) error {
	if req == nil {
		return errors.New("request is required")
	}
	if resource == nil {
		return errors.New("resource is required")
	}

	// 1. New a SHOPLINE API request
	shopLineReq := newShopLineRequest(req)

	// 2. Call an API
	_, err := c.Execute(ctx, HTTPMethod(req.GetMethod()), req.GetEndpoint(), shopLineReq, resource)
	return err
}

func newShopLineRequest(req APIRequest) *ShopLineRequest {
	shopLineReq := &ShopLineRequest{}
	if req == nil {
		return shopLineReq
	}

	shopLineReq.Query = req.GetQuery()
	shopLineReq.Data = req.GetData()

	shopLineReq.Headers = req.GetHeaders()
	shopLineReq.Options = req.GetRequestOptions()
	return shopLineReq
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
// See the method Call
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
// See the method Call
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
// See the method Call
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
// See the method Call
func (c *Client) Get(ctx context.Context, endpoint string, request *ShopLineRequest, resource interface{}) (*ShopLineResponse, error) {
	return c.Execute(ctx, MethodGet, endpoint, request, resource)
}

// Execute performs an http request for the given endpoint and saves the result in the given resource.
// accessToken:
//
//	中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201#%E7%AC%AC%E4%B8%83%E6%AD%A5app-%E6%94%B6%E5%88%B0-access-token
//	en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization/?version=v20251201#step-5-receive-the-access-token
//
// method: HTTP GetMethod
// endpoint: API request endpoint, eg: orders.json
// request: ShopLineRequest
// resource : API response, To specify the return type of the request
// See the method Call
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

// Return the http url's path '/admin/openapi/{version}/{relPath}'
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

	// Build response
	shopLineResp, err := buildShopLineResponse(httpReq.Method, request, resp, resource)
	if err != nil {
		return shopLineResp, err
	}

	c.logDetailIfNecessary(httpReq.Method, httpReq.URL.String(), request, shopLineResp)

	return shopLineResp, nil
}

func (c *Client) NewHttpRequest(ctx context.Context, method HTTPMethod, path string, request *ShopLineRequest) (*http.Request, error) {
	// Build request URL
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

	// Create HTTP request
	httpReq, err := http.NewRequest(string(method), requestURL, bytes.NewBuffer(requestBodyJsonData))
	if err != nil {
		log.Printf("Failed to create request: %v\n", err)
		return nil, err
	}

	httpReq = httpReq.WithContext(ctx)

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

// Set Headers
func (c *Client) setHeaders(appKey string, appSecret string, httpReq *http.Request, wrapper *shopLineRequestWrapper) error {
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("appkey", appKey)
	httpReq.Header.Set("User-Agent", DefaultUserAgent)

	// Create access Token & refresh access Token is not required
	if c.Token != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.Token)
	}

	timestamp := BuildTimestamp()
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

func resolveAppKey(app App) string {
	return app.AppKey
}

func resolveAppSecret(app App) string {
	return app.AppSecret
}

// Generate sign string
func generateSign(appKey, appSecret, timestamp string, requestBodyJsonBytes []byte) (string, error) {
	bodyJsonString, err := buildBodyJsonString(requestBodyJsonBytes)
	if err != nil {
		return "", err
	}

	sign := GenerateSign(appKey, bodyJsonString, timestamp, appSecret)
	return sign, nil
}

// Build body params to json string
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

// Build SHOPLINE response
func buildShopLineResponse(method string, request *ShopLineRequest, httpResp *http.Response, resource interface{}) (*ShopLineResponse, error) {
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

	if !isDeleteSuccess(method, httpResp.StatusCode) && !request.NotDecodeBody() { // delete method will return empty body
		// decode http body on success
		if err := json.NewDecoder(httpResp.Body).Decode(&resource); err != nil {
			//respData := &map[string]any{}
			//json.NewDecoder(httpResp.Body).Decode(respData)
			//log.Printf("Failed to parse json response body, statusCode: %d, body: %v, err: %v\n", httpResp.StatusCode, resource, err)
			bodyBytes, _ := io.ReadAll(httpResp.Body)
			body := string(bodyBytes)
			log.Printf("Use json decoder to decode response body failed, statusCode: %d, body: %v, err: %v\n", httpResp.StatusCode, body, err)
			shopLineResp.ResponseBody = body
			return shopLineResp, err
		}
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

func isDeleteSuccess(method string, statusCode int) bool {
	return MethodDelete == method && statusCode == http.StatusOK
}

// set common data to api response data
func setCommonAPIRespData(shopLineResp *ShopLineResponse) {
	apiRespData := shopLineResp.Data
	if apiRespData == nil {
		return
	}

	if _, ok := (apiRespData).(APIResponse); ok {
		apiResp := (apiRespData).(APIResponse)
		apiResp.SetTraceId(shopLineResp.TraceId)
		apiResp.SetPagination(shopLineResp.Pagination)
	}

	//apiRespVal := reflect.ValueOf(apiResp)
	//
	//if apiRespVal.Kind() == reflect.Ptr {
	//	apiRespVal = apiRespVal.Elem()
	//}
	//
	//if apiRespVal.Kind() != reflect.Struct {
	//	log.Printf("Invalid response data type(must struct ptr or struct): %T\n", apiResp)
	//	return
	//}
	//
	//typ := apiRespVal.Type()
	//
	//setTraceId2APIResp(shopLineResp, typ, apiRespVal)

}

// Set traceId to api response
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
		return "", "", fmt.Errorf("HTTP GetMethod is required")
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
		if _, ok := (request.Data).(APIRequest); ok {
			err := (request.Data).(APIRequest).Verify()
			if err != nil {
				return "", "", err
			}
		}
	}
	if request.Query != nil {
		if _, ok := (request.Query).(APIRequest); ok {
			err := (request.Query).(APIRequest).Verify()
			if err != nil {
				return "", "", err
			}
		}
	}
	return appKey, appSecret, nil
}

// Add the request query parameters to the http query parameters
func (c *Client) buildRequestUrl(method HTTPMethod, relPath string, request *ShopLineRequest) (string, error) {

	parsedURL, err := c.ResolveURL(relPath)
	if err != nil {
		return "", err
	}

	if request.Query != nil {
		optionsQuery, err := query.Values(request.Query)
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

// Serialize body parameters to json bytes
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
//eg:https://storeHandle.myshopline.com/admin/openapi/v20251201/orders/count.json
func (c *Client) ResolveURL(relPath string) (*url.URL, error) {
	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}
	parsedURL := c.baseURL.ResolveReference(rel)
	return parsedURL, nil
}
