package main

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/manager"
	"log"
	"net/http"
)

// Process App install request，https://developer.myshopline.com/app/store-test?appKey=72a6746a3607e3cb26b336899b172403f0c1ba6c
func main() {

	// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#%E7%AC%AC%E4%BA%8C%E6%AD%A5app-%E9%AA%8C%E8%AF%81%E5%AE%89%E8%A3%85%E8%AF%B7%E6%B1%82
	// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step2
	http.HandleFunc("/install", InstallHandler)

	// 中文: 注册第五步 「App 收到授权码」 接口, https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#%E7%AC%AC%E4%BA%94%E6%AD%A5app-%E6%94%B6%E5%88%B0%E6%8E%88%E6%9D%83%E7%A0%81
	// en : https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step-5-receive-the-access-token
	http.HandleFunc("/auth/callback", CallbackHandler)

	http.HandleFunc("/webhook", WebhookHandler)

	// start the server
	log.Println("Server started on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

// WebhookHandler Verify a Webhook http request, sent by SHOPLine.
// The body of the request is still readable after invoking the method.
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	manager.GetDefaultApp().VerifyWebhookRequest(r)
	// do something
}

// InstallHandler
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#%E7%AC%AC%E4%BA%8C%E6%AD%A5app-%E9%AA%8C%E8%AF%81%E5%AE%89%E8%A3%85%E8%AF%B7%E6%B1%82
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step2
// After store install，will receive a request from SHOPLine platform：GET https://{appUrl}?appkey={appkey}&handle={handle}&lang={lang}&timestamp={timestamp}&sign={sign}
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
	//if err != nil || time.Now().Unix()-timestamp > 600 {
	//	http.Error(w, "Invalid timestamp", http.StatusBadRequest)
	//	return
	//}

	// 5. Verify the Sign
	app := manager.GetApp(appkey)
	isSignValid := app.VerifySign(r.URL.Query(), sign)
	if !isSignValid {
		log.Printf("sign verification failed, appkey: %s, sign: %s\n", appkey, sign)
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	// 6. TODO do biz logic
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

// CallbackHandler
// 中文: 注册第五步 「App 收到授权码」 接口, https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#%E7%AC%AC%E4%BA%94%E6%AD%A5app-%E6%94%B6%E5%88%B0%E6%8E%88%E6%9D%83%E7%A0%81
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/app-authorization?version=v20260301#step-5-receive-the-access-token
// http://{storeHandle}.myshopline.com/auth/callback?appkey=72a6746a3607e3cb26b336899b172403f0c1ba6c&code=use19b2f46d1e5ecb495dcfee9827b64b5202cbd61&handle=zwapptest&timestamp=1752563015566&sign=7dc202cc18c13f033e638eab348a071a3823626527311e5abc428b367a26c683
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	//  parse request params
	//appkey={appkey}&code={code}&customField={customField}&handle={handle}&timestamp={timestamp}&sign={sign}
	appkey := r.URL.Query().Get("appkey")
	code := r.URL.Query().Get("code")
	handle := r.URL.Query().Get("handle")
	customField := r.URL.Query().Get("customField")
	timestampStr := r.URL.Query().Get("timestamp")
	sign := r.URL.Query().Get("sign")

	app := manager.GetApp(appkey)
	isSignValid := app.VerifySign(r.URL.Query(), sign)
	if isSignValid {
		log.Println("sign verified successfully")
	} else {
		log.Printf("sign verification failed, sign=%s\n", sign)
		return
	}

	// todo do biz logic
	log.Printf("Auth callback received - appkey: %s, handle: %s, code: %s, customField: %s, timestampStr: %s, sign: %s", appkey, handle, code, customField, timestampStr, sign)

	// create a new token
	token, err := app.CreateAccessToken(context.Background(), code)
	if err != nil {
		log.Printf("Create access token error, appkey: %s, handle: %s, code: %s, err: %v\n", appkey, handle, code, err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Access Token created: %v\n", token)

	accessToken := token.Data.AccessToken
	cli := manager.GetClient(appkey, handle)
	if cli != nil {
		// reset token
		cli.Token = accessToken
	} else {
		// Create a new API client
		cli, err = client.NewClient(app, handle, accessToken)
		if err != nil {
			log.Printf("new client error, appkey: %s, err: %v\n", appkey, err)
		}
		// todo to register client
		manager.Register(cli)
	}

}
