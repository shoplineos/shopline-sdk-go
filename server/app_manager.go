// Package manager for server-side
package main

import (
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/config"
	"github.com/shoplineos/shopline-sdk-go/support"
	"log"
	"sync"
)

var (
	defaultAppInstance           client.App
	defaultClientInstance        *client.Client
	defaultWebhookClientInstance *client.WebhookClient

	Apps           sync.Map // map[string]client2.App
	Clients        sync.Map // map[string]*client2.Client
	WebhookClients sync.Map // map[string]*client2.Client
)

// init default App & Client
func init() {
	ok := verifyAppDefaultConfig()
	if !ok {
		panic("Verify App default data failed, please check your config 'app_config.go'")
	}

	defaultAppInstance = client.App{
		AppKey:      config.DefaultAppKey,
		AppSecret:   config.DefaultAppSecret,
		RedirectUrl: config.DefaultRedirectUrl,
		Scope:       config.DefaultAppScope,
	}

	defaultClientInstance = support.MustNewClient(defaultAppInstance, config.DefaultStoreHandle, config.DefaultAccessToken)
	if defaultClientInstance == nil {
		panic("client is nil")
	}

	defaultAppInstance.Client = defaultClientInstance

	//Clients = make(map[string]*client2.Client)
	//Clients[defaultAppInstance.AppKey] = defaultClientInstance
	key := resolveCacheKey(defaultAppInstance.AppKey, defaultClientInstance.StoreHandle)
	Clients.Store(key, defaultClientInstance)

	//Apps = make(map[string]client2.App)
	//Apps[defaultAppInstance.AppKey] = defaultAppInstance
	Apps.Store(defaultAppInstance.AppKey, defaultAppInstance)

	defaultWebhookClientInstance = client.NewWebhookClient(defaultAppInstance)
	WebhookClients.Store(defaultAppInstance.AppKey, defaultWebhookClientInstance)

	fmt.Printf("Init default client success! appkey: %s, storeHandle: %s\n", defaultAppInstance.AppKey, defaultClientInstance.StoreHandle)
}

func verifyAppDefaultConfig() bool {
	if config.DefaultAppKey == "" || config.DefaultAppSecret == "" || config.DefaultStoreHandle == "" || config.DefaultAccessToken == "" {
		return false
	}
	if config.DefaultRedirectUrl == "" || config.DefaultAPIVersion == "" {
		return false
	}
	return true
}

func resolveCacheKey(appKey string, storeHandle string) string {
	return appKey + ":" + storeHandle
}

func GetClient(appKey, storeHandle string) *client.Client {
	key := resolveCacheKey(appKey, storeHandle)
	var c, ok = Clients.Load(key)
	if !ok {
		log.Printf("get client failed! cachekey: %s\n", key)
		return nil
	}
	return c.(*client.Client)
}

func GetDefaultClient() *client.Client {
	return defaultClientInstance
}

func GetDefaultWebhookClient() *client.WebhookClient {
	return defaultWebhookClientInstance
}

func GetWebhookClient(appKey string) *client.WebhookClient {
	cli, ok := WebhookClients.Load(appKey)
	if !ok {
		log.Printf("Get WebhookClient failed! appKey: %s\n", appKey)
		return nil
	}
	return cli.(*client.WebhookClient)
}

func GetApp(appKey string) client.App {
	app, ok := Apps.Load(appKey)
	if !ok {
		log.Printf("Get App failed! appKey: %s\n", appKey)
		return client.App{}
	}
	return app.(client.App)
}

func GetDefaultApp() client.App {
	return defaultAppInstance
}

func Register(client *client.Client) {
	//Clients[client.App.AppKey+":"+client.StoreHandle] = client
	Clients.Store(resolveCacheKey(client.App.AppKey, client.StoreHandle), client)
	fmt.Printf("Register client success! appkey: %s, storeHandle: %s\n", client.App.AppKey, client.StoreHandle)
}
