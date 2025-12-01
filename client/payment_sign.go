package client

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type PaymentSignatureAlgorithm struct {
	PrivateKey string
	PublicKey  string

	RsaPrivateKey *rsa.PrivateKey
	RsaPublicKey  *rsa.PublicKey
}

// NewPaymentSignatureAlgorithm
// 中文: https://developer.shopline.com/zh-hans-cn/docs/apps/payment-application/payment-application-signature-logic?version=v20260301
// En: https://developer.shopline.com/docs/apps/payment-application/payment-application-signature-logic?version=v20260301
func NewPaymentSignatureAlgorithm(privateKey string, publicKey string) *PaymentSignatureAlgorithm {
	alg := &PaymentSignatureAlgorithm{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}

	alg.InitKeys()
	return alg
}

func (a *PaymentSignatureAlgorithm) InitKeys() {
	a.InitPrivateKey()
	a.InitPublicKey()
}

func (a *PaymentSignatureAlgorithm) InitPrivateKey() {
	derBytes, err := base64.StdEncoding.DecodeString(a.PrivateKey)
	if err != nil {
		derBytes, err = base64.URLEncoding.DecodeString(a.PrivateKey)
		if err != nil {
			log.Printf("failed to decode private key from Base64: %v\n", err)
			return
		}
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(derBytes)
	if err != nil {
		log.Printf("failed to parse PKCS8 private key: %v\n", err)
		return
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		log.Println("not an RSA private key")
		return
	}
	a.RsaPrivateKey = rsaPrivateKey
}

func (a *PaymentSignatureAlgorithm) InitPublicKey() {
	derBytes, err := base64.StdEncoding.DecodeString(a.PublicKey)
	if err != nil {
		log.Printf("failed to decode public key from Base64: %v\n", err)
		return
	}

	publicKey, err := x509.ParsePKIXPublicKey(derBytes)
	if err != nil {
		log.Printf("failed to parse PKIX public key: %v\n", err)
		return
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		log.Println("not an RSA public key")
		return
	}
	a.RsaPublicKey = rsaPublicKey
}

// Signature Generate signature
func (a *PaymentSignatureAlgorithm) Signature(params map[string]interface{}) (string, error) {
	signatureSource := BuildSignatureSourceString(params)
	return a.SignWithPrivateKey(signatureSource)
}

func (a *PaymentSignatureAlgorithm) SignatureSource(sourceString string) (string, error) {
	return a.SignWithPrivateKey(sourceString)
}

// BuildSignatureSourceString Generate text to be signed.
func BuildSignatureSourceString(params map[string]interface{}) string {
	content := &strings.Builder{}
	appendToContent(content, params, false)
	return content.String()
}

// appendToContent
// isSubMap Used to identify whether the currently processed map is a nested sub-map.
func appendToContent(content *strings.Builder, sourceObj map[string]interface{}, isSubMap bool) {
	if sourceObj == nil || len(sourceObj) == 0 {
		return
	}

	// Sort keys
	keys := make([]string, 0, len(sourceObj))
	for k := range sourceObj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, key := range keys {
		value := sourceObj[key]

		if content.Len() > 0 && !isSubMap {
			content.WriteString("&")
		}

		switch v := value.(type) {
		case []interface{}:
			if len(v) > 0 {
				if isScalarValue(v[0]) {
					content.WriteString(key)
					content.WriteString("=")
					for j, item := range v {
						if j > 0 {
							content.WriteString(",")
						}
						content.WriteString(toString(item))
					}
				} else {
					for _, item := range v {
						if subMap, ok := item.(map[string]interface{}); ok {
							appendToContent(content, subMap, true)
						}
					}
				}
			}
		case map[string]interface{}:
			// Process nested map
			content.WriteString(key)
			content.WriteString("=")
			appendToContent(content, v, true)
		default:
			if isScalarValue(v) {
				content.WriteString(key)
				content.WriteString("=")
				content.WriteString(toString(v))
			}
			// Ignore the unknow type
		}

		if isSubMap && i < len(keys)-1 {
			content.WriteString("&")
		}
	}
}

// isScalarValue
func isScalarValue(value interface{}) bool {
	switch value.(type) {
	case string, int, int64, float64, bool:
		return true
	default:
		switch value.(type) {
		case int8, int16, int32, uint, uint8, uint16, uint32, uint64:
			return true
		case float32:
			return true
		}
		return false
	}
}

// toString
func toString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		if float64(int(v)) == v {
			return strconv.Itoa(int(v))
		}
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// SignWithPrivateKey use private key to sign
func (a *PaymentSignatureAlgorithm) SignWithPrivateKey(signSourceStr string) (string, error) {

	hasher := sha1.New()
	hasher.Write([]byte(signSourceStr))
	hash := hasher.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, a.RsaPrivateKey, crypto.SHA1, hash)
	if err != nil {
		return "", fmt.Errorf("failed to sign: %v", err)
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// CheckSign use public key to Verify
func (a *PaymentSignatureAlgorithm) CheckSign(params map[string]interface{}, signedStr string) (bool, error) {
	sourceString := BuildSignatureSourceString(params)
	valid, err := a.CheckSignWithPublicKey(sourceString, signedStr)
	if err != nil {
		return false, err
	}
	return valid, nil
}

// CheckSignWithPublicKey use public key to Verify
func (a *PaymentSignatureAlgorithm) CheckSignWithPublicKey(signSourceStr string, signedStr string) (bool, error) {
	// Decode Base64
	signature, err := base64.StdEncoding.DecodeString(signedStr)
	if err != nil {
		return false, fmt.Errorf("failed to decode base64 signature: %v", err)
	}

	hasher := sha1.New()
	hasher.Write([]byte(signSourceStr))
	hash := hasher.Sum(nil)

	// Use public key to Verify
	err = rsa.VerifyPKCS1v15(a.RsaPublicKey, crypto.SHA1, hash, signature)
	if err != nil {
		// Verify failed
		return false, nil
	}

	// Verify success
	return true, nil
}
