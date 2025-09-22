package oauth

import (
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/config"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization?version=v20251201
func TestCreateAccessToken(t *testing.T) {

	code := "" // code

	accessToken, err := CreateAccessToken(config.DefaultAppKey, code)
	if err != nil {
		log.Fatalf("Failed to get access token: %v\n", err)
	} else {
		fmt.Printf("Access Token: %v\n", accessToken)
	}

	//log.Printf("Status Code: %d\n", shopLineResp.StatusCode)
	log.Printf("Access Token: %s\n", accessToken.Data.AccessToken)
	log.Printf("Expires In: %s seconds\n", accessToken.Data.ExpiresTime)
	//fmt.Printf("Refresh Token: %s\n", tokenResp.Data.RefreshToken)

	a := assert.New(t)
	a.Equal(200, accessToken.Code)
}

// https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/app-authorization/?lang=zh-hans-cn#%E7%AC%AC%E5%85%AB%E6%AD%A5app-%E8%AF%B7%E6%B1%82%E5%88%B7%E6%96%B0-access-token
func TestRefreshAccessToken(t *testing.T) {

	// replace app data
	storeHandle := config.DefaultStoreHandle
	appKey := config.DefaultAppKey

	accessToken, err := RefreshAccessToken(appKey, storeHandle)
	if err != nil {
		log.Fatalf("Failed to refresh access token: %v\n", err)
	} else {
		fmt.Printf("Refresh Access Token Result: %v\n", accessToken)
	}

	//log.Printf("Status Code: %d\n", accessToken)
	log.Printf("New Access Token: %s\n", accessToken.Data.AccessToken)
	log.Printf("Expires In: %s seconds\n", accessToken.Data.ExpiresTime)
	//fmt.Printf("Refresh Token: %s\n", tokenResp.Data.RefreshToken)

	a := assert.New(t)
	a.Equal(200, accessToken.Code)

}
