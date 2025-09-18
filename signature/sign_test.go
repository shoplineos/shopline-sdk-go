package signature

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"shoplineapp/config"
	"testing"
)

//var (
//	client *client2.Client
//	app    client2.App
//)
//
//func setup() *client2.Client {
//	app = client2.App{
//		AppKey:    config.DefaultAppKey,
//		AppSecret: config.DefaultAppSecret,
//	}
//
//	client = client2.MustNewClient(app, config.DefaultStoreHandle, config.DefaultAccessToken)
//	if client == nil {
//		panic("client is nil")
//	}
//
//	app.Client = client
//	return client
//}

// no body
func TestGenerateSign1(t *testing.T) {

	appKey := ""
	jsonBody := ""

	//timestamp := time.Now().Unix()
	timestampInt := "21c335b3a"

	//fmt.Printf("timestampInt: %s\n", timestampInt)

	sign := GenerateSign(appKey, string(jsonBody), timestampInt, config.DefaultAppSecret)

	a := assert.New(t)

	fmt.Printf("sign: %s\n", sign)

	a.Equal("5c8d81bee7b236a6e57055264ae6e93195a1efad72b5175470f8dffe18ae015d", sign)

}

// has body
func TestGenerateSign2(t *testing.T) {

	appKey := ""

	requestBody := map[string]string{
		"code": "code",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}

	//timestamp := time.Now().Unix()
	timestampInt := "21c335b3a"

	//fmt.Printf("timestampInt: %s\n", timestampInt)

	sign := GenerateSign(appKey, string(jsonBody), timestampInt, config.DefaultAppSecret)

	a := assert.New(t)

	fmt.Printf("sign: %s\n", sign)

	a.Equal(sign, "55cb3704385db767b865c23f4f2201f4abdb817ccc5e3cce5cbbb86e555fecc0")

}

// has body & no timestamp
func TestGenerateSignForWebhook(t *testing.T) {

	appKey := config.DefaultAppKey

	requestBody := "my secret message"

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}

	timestamp := ""
	// timestampInt := "21c335b3a"

	//fmt.Printf("timestampInt: %s\n", timestampInt)

	sign := GenerateSign(appKey, string(jsonBody), timestamp, config.DefaultAppSecret)

	fmt.Printf("sign: %s\n", sign)

	a := assert.New(t)
	a.Equal(sign, "1487c504fdb834b0ec315fc038e6f16ed0b51d5e7eb3f1d3aca862139673788b")

}
