package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/test"
	access2 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/access"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateStorefrontAccessToken(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/storefront_access_tokens.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"storefront_access_token":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "access_scope":"test desc", "access_token":"access_token", "title":"name_test"}}`))

	req := &access2.CreateAnAccessTokenAPIReq{
		StorefrontAccessToken: access2.CreateAnAccessTokenAPIReqStorefrontAccessToken{
			Title: "test desc",
		},
	}

	apiResp := &access2.CreateAnAccessTokenAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("StorefrontAccessToken.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, int64(1), apiResp.StorefrontAccessToken.Id)
	assert.Equal(t, "access_token", apiResp.StorefrontAccessToken.AccessToken)
}

func TestDeleteStorefrontAccessToken(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("DELETE",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/storefront_access_tokens/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	req := &access2.DeleteAnAccessTokenAPIReq{
		TokenId: "1",
	}

	apiResp := &access2.DeleteAnAccessTokenAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)
	if err != nil {
		t.Errorf("StorefrontAccessToken.Delete returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
}

func TestListStorefrontAccessTokens(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/storefront_access_tokens.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "../test/access/storefront_access_tokens.json")))

	req := &access2.GetAListOfAccessTokensAPIReq{}

	apiResp := &access2.GetAListOfAccessTokensAPIResp{}
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
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/storefront_access_tokens.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "../test/access/storefront_access_tokens.json")))

	req := &access2.GetAListOfAccessTokensAPIReq{}

	apiResp := &access2.GetAListOfAccessTokensAPIResp{}
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
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/storefront_access_tokens.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "../test/access/storefront_access_tokens.json")))

	req := &access2.GetAListOfAccessTokensAPIReq{}

	apiResp := &access2.GetAListOfAccessTokensAPIResp{}
	tokens, err := client.ListAll(cli, context.Background(), req, apiResp, func(resp interface{}) []access2.StorefrontAccessToken {
		r := resp.(*access2.GetAListOfAccessTokensAPIResp)
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
