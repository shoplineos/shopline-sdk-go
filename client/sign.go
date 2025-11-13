package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// GenerateSignForCreateAccessToken Generate a signature for create access token
func GenerateSignForCreateAccessToken(appKey, code, timestamp, appSecret string) string {
	requestBody := map[string]string{
		"code": code,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}
	return GenerateSign(appKey, string(jsonBody), timestamp, appSecret)
}

// GenerateSignForRefreshAccessToken Generate a signature for refresh access token
func GenerateSignForRefreshAccessToken(appKey, timestamp, appSecret string) string {
	return GenerateSign(appKey, "", timestamp, appSecret)
}

// GenerateSign
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/generate-and-verify-signatures?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/generate-and-verify-signatures?version=v20251201
func GenerateSign(appKey, body, timestamp, appSecret string) string {
	source := body + timestamp

	// Create HMAC object
	hash := hmac.New(sha256.New, []byte(appSecret))

	// Write data
	hash.Write([]byte(source))

	// Compute hash value and to hex
	return hex.EncodeToString(hash.Sum(nil))
}

// VerifySign verify callback url's params sign
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/generate-and-verify-signatures?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/generate-and-verify-signatures?version=v20251201
func VerifySign(appSecret string, params url.Values, receivedSign string) bool {
	expectedSign := GenerateSignForGet(appSecret, params)

	// Compare the sign
	return hmac.Equal([]byte(expectedSign), []byte(receivedSign))
}

// GenerateSignForGet
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/generate-and-verify-signatures?version=v20251201
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/generate-and-verify-signatures?version=v20251201
func GenerateSignForGet(appSecret string, params url.Values) string {
	// 1. Copy the parameters and remove the sign field
	paramCopy := make(map[string]string)
	for k, v := range params {
		if k != "sign" {
			paramCopy[k] = v[0]
		}
	}

	// 2. sort
	keys := make([]string, 0, len(paramCopy))
	for k := range paramCopy {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 3. Build params
	var paramStrBuilder strings.Builder
	for _, k := range keys {
		if paramStrBuilder.Len() > 0 {
			paramStrBuilder.WriteByte('&')
		}
		paramStrBuilder.WriteString(k)
		paramStrBuilder.WriteByte('=')
		paramStrBuilder.WriteString(paramCopy[k])
	}
	paramStr := paramStrBuilder.String()

	// 4. Compute HMAC-SHA256 sign
	hash := hmac.New(sha256.New, []byte(appSecret))
	hash.Write([]byte(paramStr))
	expectedSign := hex.EncodeToString(hash.Sum(nil))
	return expectedSign
}

// VerifyWebhookRequest Verify a Webhook http request, sent by SHOPLINE.
// The body of the request is still readable after invoking the method.
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
func VerifyWebhookRequest(appSecret string, r *http.Request) bool {
	shoplineSha256 := r.Header.Get("X-Shopline-Hmac-Sha256")
	actualMac := []byte(shoplineSha256)

	mac := hmac.New(sha256.New, []byte(appSecret))
	requestBody, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	mac.Write(requestBody)
	macSum := mac.Sum(nil)

	s := hex.EncodeToString(macSum)
	expectedMac := []byte(s)

	return hmac.Equal(actualMac, expectedMac)
}

// VerifyWebhookMessage Verify a Webhook message against a message HMAC
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
// en: https://developer.shopline.com/docs/apps/api-instructions-for-use/webhooks/overview?version=v20260301#%E8%AF%B7%E6%B1%82header
func VerifyWebhookMessage(appSecret, message, messageMAC string) bool {
	mac := hmac.New(sha256.New, []byte(appSecret))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)

	// SHOPLINE HMAC is in hex so it needs to be decoded
	actualMac, _ := hex.DecodeString(messageMAC)

	return hmac.Equal(actualMac, expectedMAC)
}
