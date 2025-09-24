package signature

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/url"
	"testing"
)

//
//var (
//	client *Client
//	app    App
//)
//
//func setup() *Client {
//	app = App{
//		AppKey:    config.DefaultAppKey,
//		AppSecret: config.DefaultAppSecret,
//	}
//
//	client = MustNewClient(app, config.DefaultStoreHandle, config.DefaultAccessToken)
//	if client == nil {
//		panic("client is nil")
//	}
//
//	app.Client = client
//	return client
//}

const (
	AppKeyForTest      = "LDLLJJLflja2039203flaflaLLFLE"
	AppSecretForTest   = "LDLLJJLflja2039203flaflaLLFLE"
	StoreHandelForTest = "zwapptest"
	LangForTest        = "zh"
)

// no body
func TestGenerateSign1(t *testing.T) {

	appKey := AppKeyForTest
	jsonBody := ""

	//timestamp := time.Now().Unix()
	timestampInt := "21c335b3a"

	//fmt.Printf("timestampInt: %s\n", timestampInt)

	sign := GenerateSign(appKey, string(jsonBody), timestampInt, AppSecretForTest)

	a := assert.New(t)

	fmt.Printf("sign: %s\n", sign)

	a.Equal("cf5edab6b92d740fe54b3ce9e2723788128384a1eb5c2a5d64c9210a31185b86", sign)

}

// has body
func TestGenerateSign2(t *testing.T) {

	appKey := AppKeyForTest

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

	sign := GenerateSign(appKey, string(jsonBody), timestampInt, AppSecretForTest)

	a := assert.New(t)
	//fmt.Printf("sign: %s\n", sign)
	a.Equal(sign, "4c4cf0ac52439d1308b58897f857dcba8545efb6c0e8091fba64b148218d2074")

}

// has body & no timestamp
func TestGenerateSignForWebhook(t *testing.T) {

	appKey := AppKeyForTest

	requestBody := "my secret message"

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}

	timestamp := ""
	// timestampInt := "21c335b3a"

	//fmt.Printf("timestampInt: %s\n", timestampInt)

	sign := GenerateSign(appKey, string(jsonBody), timestamp, AppSecretForTest)

	//fmt.Printf("sign: %s\n", sign)

	a := assert.New(t)
	a.Equal(sign, "7d8b4cbec432e0c7e901881cc0f2490a910d56eaa73fc94bcc4276d14a61fe88")

}

func TestGenerateSignForGet(t *testing.T) {
	appSecret := AppSecretForTest
	params := url.Values{}
	params.Add("appkey", AppKeyForTest)
	params.Add("handle", StoreHandelForTest)
	params.Add("lang", LangForTest)
	params.Add("timestamp", "21c335b3a")

	sign := GenerateSignForGet(appSecret, params)
	//fmt.Printf("sign: %s\n", sign)

	a := assert.New(t)
	a.Equal(sign, "654a42549daaad959b6488500f499fc0381c6cf24d8afceeb566729fc82fee6a")

}

func TestVerifySign(t *testing.T) {

	appSecret := AppSecretForTest
	params := url.Values{}
	params.Add("appkey", AppKeyForTest)
	params.Add("handle", StoreHandelForTest)
	params.Add("lang", LangForTest)
	params.Add("timestamp", "21c335b3a")

	receivedSign := "654a42549daaad959b6488500f499fc0381c6cf24d8afceeb566729fc82fee6a"
	success := VerifySign(appSecret, params, receivedSign)
	a := assert.New(t)
	a.Equal(success, true)

	receivedSign = "wrong"
	success = VerifySign(appSecret, params, receivedSign)
	a.False(success)

}
