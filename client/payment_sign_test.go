package client

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const (
	// For test
	privateKeyStr = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQChLWJXdQp+jHuVwz/55pcXjeXIR3l8oYsBZCdNnLV8UlXN9c/cEdHafxmjK4tneY94yIv0y315VPMj41EDP/iYz/ciE02sujMS69uAsAeP6SioFV42Poyl933A/eow95UL2qDI2AoHtLzAnJk4bfl/iKfSP1bvkrWaR2zpV0jljCexwXa9yLv8MFXPsM2knfgQLqY67wrB6JMTbpDTMazz8BFmrdUP4+z2XWyTE6knLBi1hpuW5/NKP1I7D5v++8QX+BN5E0cooUdYuoCVqh+rOd8u96XVSK3HC8+pNYqYURZKVogt62QXdnKcSd5q34PqSCJzmsqf5cNbOqeijwi1AgMBAAECggEABqxb0nR5D5xhuVjSozy2u996Tu4vDtsCAWsXM0zK+Qxp2yoNJf0148N+uTI75oTSgCUVIfizGsz93fOsnpAT19qtqCNFxGJKnyl7sERGV+EtN5v4FEWUNxxVzSHZan2073JcO4/LhusNV2mhPWq7DKdiXWxI1u8x2WMRYXrliYPdq5BwI/0O297vEPJ+xZYe1vfZ0nc7wg/SbbYERbPDofeIgW4PspovVi1uSSgBg3n+90NSf5RjvbH5VeUxzeOniB75si8Vdvnx6+mP0ALyvr+EenqeQnYpUaqY2s7gi7KnmQyDYZ/Cw7kc0TR32jKowXln50scY1PWGP0V6Z0LiQKBgQDQZRk9HOnZmcW7fYaMyhKtBnboTb6RlmouDdz+j8eYI732+qTUjAbM6GtPvE5mJlO9ObMNK/s8ugb1Y++ZzZUKiOc/jYNgqT8tevLks2DlS0iNiPFtE7OjmA2sEk/zaoZqAscRIrmTravekLnc5OJpkOvsHGcjlO5ju7P4tbsdnQKBgQDF/wF+ZghZX8e47tQLP6nkMs0UMtbCMXjsbB82e58gy/28J/fus5XEvxJjA4wiUbZAnj/NQrYYT44HkwnKV6SkMSsZ6x7e5Had1Jn+3RYtarc0aVGE9xLfwQqvVSBEmLLRerfFX+5fbp6FKlM1nn0rxXpxRLt/JnTCmkluB7S3+QKBgEfq6+fcR5PR7pxCuKFzxzgxJ+4Jjn+90gzsudycD/ygMRm/7Axx+pLSjt4olUHJblK6S+F60Sxm4qnjADgq64mEL5IOK027etMeQB7PDNx0u6gkn3TOPMtzWRyOAUt28sY5CSwPuM2PPOYFOi9SShS2b8S/FJUB+7ctevGU/es9AoGBAKrom2Z7NrvHNMSKy9jF5KXJwEKuO7k3MTWLg0npXgvWajkPmzGeLSq+8GUtu7ooJJUUxOgurLbBfU1GfE4AZ2sf0h+2WFh4h3dn/GIGf81Gb8w7GRYYnF8u6EU+yvLLiJfQQW+Lhl00RHuYdGk1XMD63t2FQf/YtzMAMWBcIIApAoGATHeTyP12xFfL0+GIgpMMNKWxVkuwZdwqozVtVfp5vuNw/kkbM0dXohClg/CuMyRMp7Hh2h0fZPsGcwUPSEY3c2NF9vjGNR6/Rk4OyjeVSTZEsEZIfH7n1lkcbhbATQ65boWn/I3iiAASJX2CgkTmitO+fCwrbWkZ7H7CCrH9Aas=" // for test
	publicKeyStr  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoS1iV3UKfox7lcM/+eaXF43lyEd5fKGLAWQnTZy1fFJVzfXP3BHR2n8ZoyuLZ3mPeMiL9Mt9eVTzI+NRAz/4mM/3IhNNrLozEuvbgLAHj+koqBVeNj6Mpfd9wP3qMPeVC9qgyNgKB7S8wJyZOG35f4in0j9W75K1mkds6VdI5YwnscF2vci7/DBVz7DNpJ34EC6mOu8KweiTE26Q0zGs8/ARZq3VD+Ps9l1skxOpJywYtYablufzSj9SOw+b/vvEF/gTeRNHKKFHWLqAlaofqznfLvel1UitxwvPqTWKmFEWSlaILetkF3ZynEneat+D6kgic5rKn+XDWzqnoo8ItQIDAQAB"                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 // for test
)

func TestSignature(t *testing.T) {

	params := map[string]interface{}{
		"key":   "value",
		"name":  "test",
		"age":   25,
		"score": 99.5,
		"isMan": true,
		"hobbies": []string{
			"reading",
			"sports",
		},
		"address": map[string]interface{}{
			"province": "Beijing",
			"city":     "Beijing",
		},
	}

	alg := NewPaymentSignatureAlgorithm(privateKeyStr, publicKeyStr)

	signatureSourceString := BuildSignatureSourceString(params)
	fmt.Printf("String to be signed: %s\n", signatureSourceString)

	signature, err := alg.SignatureSource(signatureSourceString)
	if err != nil {
		log.Fatalf("Failed to generate signature: %v", err)
	}

	fmt.Printf("Generated signature: %s\n", signature)

	isValid, err := alg.CheckSignWithPublicKey(signatureSourceString, signature)
	if err != nil {
		log.Fatalf("Error in the verification process: %v", err)
	}
	//fmt.Printf("Verify result: %v\n", isValid)
	assert.True(t, isValid)
}

func TestDecodePublicKey(t *testing.T) {
	//fmt.Println(len(publicKeyStr))
	_, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		log.Fatalf("failed to decode public key from Base64: %v", err)
	}
}

func TestBuildSignatureSourceString(t *testing.T) {

	params := map[string]interface{}{
		"key":   "value",
		"name":  "test",
		"age":   25,
		"score": 99.5,
		"isMan": true,
		"hobbies": []string{
			"reading",
			"sports",
		},
		"address": map[string]interface{}{
			"province": "Beijing",
			"city":     "Beijing",
		},
	}

	signatureSourceString := BuildSignatureSourceString(params)

	assert.Equal(t, "address=city=Beijing&province=Beijing&age=25&&isMan=true&key=value&name=test&score=99.5", signatureSourceString)

}

func TestCheckSign(t *testing.T) {

	params := map[string]interface{}{
		"key":   "value",
		"name":  "test",
		"age":   25,
		"score": 99.5,
		"isMan": true,
		"hobbies": []string{
			"reading",
			"sports",
		},
		"address": map[string]interface{}{
			"province": "Beijing",
			"city":     "Beijing",
		},
	}

	alg := NewPaymentSignatureAlgorithm(privateKeyStr, publicKeyStr)
	valid, err := alg.CheckSign(params, "IxeB79esGylecDSi+OWuIfDOCHybhJBpEt/TLQV3vomL9G0y9eSBubUS9NPvcwFSZehILBZyGT34uDQC/3bNeANScwXIcO3xx3NSQSbP5c8oYkLUfkCqLmE1W/DVfsvZ4leAqnJPA+lH4mPoHf56n/6HMeAXO0AmcKkVeK8l5Bf5DwvWKJeq2HkGRSFj/uIbk+wvdSrFE9scjgtaRLdF3QUJwJa/Jyxt0C3jpFwv+4QaP0Gq1zLLZLkKpsIBQZg3EN/i9599Yal196YDwQ3AqiHeMrhrJHcXwXGsfik/4gSpJEAWfi+kUGQpDNEJG/RmDSbfUF+0aylhM8A/AufCdg==")
	if err != nil {
		t.Fatalf("Failed to CheckSign: %v", err)
	}

	assert.True(t, valid)
}
