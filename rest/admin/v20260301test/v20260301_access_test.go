package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/access"
	"github.com/stretchr/testify/assert"
)

func accessURL(cli *client.Client, path string) string {
	return fmt.Sprintf("https://%s.myshopline.com/%s/%s/%s",
		cli.StoreHandle, cli.PathPrefix, cli.ApiVersion, path)
}

func TestGetAListOfAccessTokens(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"storefront_access_tokens":[{"id":1001,"title":"For SHOPLINE Themes","access_token":"abc123","access_scope":"unauthenticated_read_message","created_at":"2023-01-01T00:00:00+08:00"}]}`
	httpmock.RegisterResponder("GET", accessURL(cli, "storefront_access_tokens.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &access.GetAListOfAccessTokensAPIReq{}
	apiResp := &access.GetAListOfAccessTokensAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.StorefrontAccessTokens, 1)
	assert.Equal(t, int64(1001), apiResp.StorefrontAccessTokens[0].Id)
	assert.Equal(t, "For SHOPLINE Themes", apiResp.StorefrontAccessTokens[0].Title)
	assert.Equal(t, "abc123", apiResp.StorefrontAccessTokens[0].AccessToken)
}

func TestCreateAnAccessToken(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"storefront_access_token":{"id":1002,"title":"For SHOPLINE Themes","access_token":"xyz789","access_scope":"unauthenticated_write_message, unauthenticated_read_message","created_at":"2023-01-01T00:00:00+08:00"}}`
	httpmock.RegisterResponder("POST", accessURL(cli, "storefront_access_tokens.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &access.CreateAnAccessTokenAPIReq{
		StorefrontAccessToken: access.CreateAnAccessTokenAPIReqStorefrontAccessToken{
			Title: "For SHOPLINE Themes",
		},
	}
	apiResp := &access.CreateAnAccessTokenAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, int64(1002), apiResp.StorefrontAccessToken.Id)
	assert.Equal(t, "For SHOPLINE Themes", apiResp.StorefrontAccessToken.Title)
	assert.Equal(t, "xyz789", apiResp.StorefrontAccessToken.AccessToken)
}

func TestDeleteAnAccessToken(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	tokenId := "1001"
	httpmock.RegisterResponder("DELETE", accessURL(cli, fmt.Sprintf("storefront_access_tokens/%s.json", tokenId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &access.DeleteAnAccessTokenAPIReq{TokenId: tokenId}
	apiResp := &access.DeleteAnAccessTokenAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteAnAccessToken_MissingTokenId(t *testing.T) {
	err := (&access.DeleteAnAccessTokenAPIReq{}).Verify()
	assert.EqualError(t, err, "TokenId is required")
}
