package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/test"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/localization"
	"github.com/stretchr/testify/assert"
)

func localizationURL(cli *client.Client, path string) string {
	return fmt.Sprintf("https://%s.myshopline.com/%s/%s/%s",
		cli.StoreHandle, cli.PathPrefix, cli.ApiVersion, path)
}

func TestGetStoreLanguages(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"default_language":"en","supported_languages":["en","vi"]}}`
	httpmock.RegisterResponder("GET", localizationURL(cli, "store/languages.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &localization.GetStoreLanguagesAPIReq{}
	apiResp := &localization.GetStoreLanguagesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "en", apiResp.Data.DefaultLanguage)
	assert.Equal(t, []string{"en", "vi"}, apiResp.Data.SupportedLanguages)
}

func TestAddStoreLanguages(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"default_language":"en","supported_languages":["en","vi"]}}`
	httpmock.RegisterResponder("POST", localizationURL(cli, "store/languages.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &localization.AddStoreLanguagesAPIReq{Languages: []string{"vi"}}
	apiResp := &localization.AddStoreLanguagesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "en", apiResp.Data.DefaultLanguage)
	assert.Contains(t, apiResp.Data.SupportedLanguages, "vi")
}

func TestDeleteStoreLanguages(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("DELETE", localizationURL(cli, "store/languages.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &localization.DeleteStoreLanguagesAPIReq{Languages: []string{"ja"}}
	apiResp := &localization.DeleteStoreLanguagesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryStoreSAvailableLanguages(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"available_locales":[{"iso_code":"en","name":"English"},{"iso_code":"vi","name":"Vietnamese"}]}}`
	httpmock.RegisterResponder("GET", localizationURL(cli, "store/available_locales.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &localization.QueryStoreSAvailableLanguagesAPIReq{}
	apiResp := &localization.QueryStoreSAvailableLanguagesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Data.AvailableLocales, 2)
	assert.Equal(t, "en", apiResp.Data.AvailableLocales[0].IsoCode)
	assert.Equal(t, "English", apiResp.Data.AvailableLocales[0].Name)
}

func TestGetStoreTranslationData(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"resource_id":"1636453548928","resource_type":"PRODUCT","default_content_list":[{"resource_content":"Title content","resource_content_type":"STRING","resource_field":"title"}],"content_list":[{"locale":"en","market":"2805652120189036673434","resource_content":"Translated title","resource_field":"title","updated_at":"2023-09-07T15:50:00Z"}]}}`
	httpmock.RegisterResponder("GET", localizationURL(cli, "ugc/resource.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &localization.GetStoreTranslationDataAPIReq{
		ResourceId:   "1636453548928",
		ResourceType: "PRODUCT",
	}
	apiResp := &localization.GetStoreTranslationDataAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "PRODUCT", apiResp.Data.ResourceType)
	assert.Equal(t, "1636453548928", apiResp.Data.ResourceId)
	assert.Len(t, apiResp.Data.DefaultContentList, 1)
	assert.Equal(t, "Title content", apiResp.Data.DefaultContentList[0].ResourceContent)
	assert.Len(t, apiResp.Data.ContentList, 1)
	assert.Equal(t, "Translated title", apiResp.Data.ContentList[0].ResourceContent)
}

func TestGetStoreTranslationData_MissingResourceId(t *testing.T) {
	err := (&localization.GetStoreTranslationDataAPIReq{}).Verify()
	assert.EqualError(t, err, "ResourceId is required")
}

func TestGetStoreTranslationData_MissingResourceType(t *testing.T) {
	err := (&localization.GetStoreTranslationDataAPIReq{ResourceId: "1636453548928"}).Verify()
	assert.EqualError(t, err, "ResourceType is required")
}

func TestUpdateStoreTranslationData(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"resource_list":[{"locale":"en","market":"2805652120189036673434","outdated":false,"resource_content":"Translated title","resource_field":"title","resource_id":"1636453548928","resource_type":"PRODUCT","updated_at":"2023-09-07T15:50:00Z"}],"fail_list":[]}}`
	httpmock.RegisterResponder("PUT", localizationURL(cli, "ugc/resource.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &localization.UpdateStoreTranslationDataAPIReq{
		ResourceList: []localization.ResourceList{
			{
				ResourceId:   "1636453548928",
				ResourceType: "PRODUCT",
				Locale:       "en",
				Market:       "2805652120189036673434",
				ContentList: []localization.UpdateStoreTranslationDataAPIReqContentList{
					{ResourceField: "title", ResourceContent: "Translated title"},
				},
			},
		},
	}
	apiResp := &localization.UpdateStoreTranslationDataAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Data.ResourceList, 1)
	assert.Equal(t, "PRODUCT", apiResp.Data.ResourceList[0].ResourceType)
	assert.Equal(t, "Translated title", apiResp.Data.ResourceList[0].ResourceContent)
}

func TestQueryStoreTranslationDataInBulk(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"resource_list":[{"resource_id":"1636453548928","resource_type":"PRODUCT","default_content_list":[{"resource_content":"Title content","resource_content_type":"STRING","resource_field":"title"}],"content_list":[{"locale":"en","market":"2805652120189036673434","outdated":false,"resource_content":"Translated title","resource_field":"title","updated_at":"2023-09-07T15:50:00Z"}]}]}}`
	httpmock.RegisterResponder("POST", localizationURL(cli, "ugc/resources.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &localization.QueryStoreTranslationDataInBulkAPIReq{
		Locale: "en",
		Market: "2805652120189036673434",
		ResourceList: []localization.QueryStoreTranslationDataInBulkAPIReqResourceList{
			{ResourceType: "PRODUCT", ResourceId: "1636453548928"},
		},
	}
	apiResp := &localization.QueryStoreTranslationDataInBulkAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Data.ResourceList, 1)
	assert.Equal(t, "PRODUCT", apiResp.Data.ResourceList[0].ResourceType)
	assert.Equal(t, "1636453548928", apiResp.Data.ResourceList[0].ResourceId)
	assert.Len(t, apiResp.Data.ResourceList[0].ContentList, 1)
}

func TestDeleteStoreTranslations(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("DELETE", localizationURL(cli, "ugc/resource.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &localization.DeleteStoreTranslationsAPIReq{
		Locale: "en",
		Market: "2805652120189036673434",
		ResourceList: []localization.DeleteStoreTranslationsAPIReqResourceList{
			{
				ResourceId:   "1636453548928",
				ResourceType: "PRODUCT",
				ContentList: []localization.DeleteStoreTranslationsAPIReqContentList{
					{ResourceField: "title"},
				},
			},
		},
	}
	apiResp := &localization.DeleteStoreTranslationsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}
