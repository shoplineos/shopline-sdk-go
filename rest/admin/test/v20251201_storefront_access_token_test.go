package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	access2 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/access"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateStorefrontAccessToken(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/storefront_access_tokens.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"storefront_access_token":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "access_scope":"test desc", "access_token":"access_token", "title":"name_test"}}`))

	req := &access2.CreateStorefrontAccessTokenAPIReq{
		StorefrontAccessToken: access2.CreateStorefrontAccessToken{
			Title: "test desc",
		},
	}

	apiResp := &access2.CreateStorefrontAccessTokenAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("StorefrontAccessToken.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, int64(1), apiResp.StorefrontAccessToken.Id)
	assert.Equal(t, "access_token", apiResp.StorefrontAccessToken.AccessToken)
}

func TestDeleteStorefrontAccessToken(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/storefront_access_tokens/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	req := &access2.DeleteStorefrontAccessTokenAPIReq{
		Id: "1",
	}

	apiResp := &access2.DeleteStorefrontAccessTokenAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)
	if err != nil {
		t.Errorf("StorefrontAccessToken.Delete returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
}

func TestListStorefrontAccessTokens(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/storefront_access_tokens.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "access/storefront_access_tokens.json")))

	req := &access2.ListStorefrontAccessTokensAPIReq{}

	apiResp := &access2.ListStorefrontAccessTokensAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("StorefrontAccessToken.List returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

	storefrontAccessToken := apiResp.StorefrontAccessTokens[0]

	assert.Equal(t, int64(1), storefrontAccessToken.Id)
	assert.Equal(t, "access_token", storefrontAccessToken.AccessToken)
}

func TestListWithPaginationStorefrontAccessTokens(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/storefront_access_tokens.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "access/storefront_access_tokens.json")))

	req := &access2.ListStorefrontAccessTokensAPIReq{}

	apiResp := &access2.ListStorefrontAccessTokensAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("StorefrontAccessToken.List returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

	storefrontAccessToken := apiResp.StorefrontAccessTokens[0]

	assert.Equal(t, int64(1), storefrontAccessToken.Id)
	assert.Equal(t, "access_token", storefrontAccessToken.AccessToken)
}

func TestListAllStorefrontAccessTokens(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/storefront_access_tokens.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "access/storefront_access_tokens.json")))

	req := &access2.ListStorefrontAccessTokensAPIReq{}

	apiResp := &access2.ListStorefrontAccessTokensAPIResp{}
	tokens, err := client.ListAll(cli, context.Background(), req, apiResp, func(resp interface{}) []access2.StorefrontAccessToken {
		r := resp.(*access2.ListStorefrontAccessTokensAPIResp)
		return r.StorefrontAccessTokens
	})

	if err != nil {
		t.Errorf("StorefrontAccessToken.List returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

	storefrontAccessToken := tokens[0]

	assert.Equal(t, int64(1), storefrontAccessToken.Id)
	assert.Equal(t, "access_token", storefrontAccessToken.AccessToken)
}
