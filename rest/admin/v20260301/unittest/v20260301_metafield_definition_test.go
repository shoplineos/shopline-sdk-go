package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	metafield2 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/metafield"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/create-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/create-a-metafield-definition?version=v20251201
func TestCreateMetafieldDefinition(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafield_definition.json", cli.StoreHandle, cli.PathPrefix, ApiVersion),
		httpmock.NewStringResponder(200, `{"definition":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00","access": {"admin":"MERCHANT_READ_WRITE"}, "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field"}}`))

	req := &metafield2.CreateAMetafieldDefinitionAPIReq{
		Definition: metafield2.CreateAMetafieldDefinitionAPIReqDefinition{
			Access:        metafield2.Access{Admin: "MERCHANT_READ_WRITE"},
			Description:   "test desc",
			Key:           "key_test",
			Name:          "name_test",
			Namespace:     "namespace_test",
			OwnerResource: "product",
			Type:          "single_line_text_field",
		},
	}

	//apiResp, err := GetMetafieldDefinitionService().Create(context.Background(), req)
	apiResp := &metafield2.CreateAMetafieldDefinitionAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("MetafieldDefinition.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.NotNil(t, apiResp.Definition.Id)
	assert.Equal(t, int64(1), apiResp.Definition.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", apiResp.Definition.CreatedAt)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", apiResp.Definition.UpdatedAt)

}

// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/update-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/update-a-metafield-definition?version=v20251201
func TestUpdateMetafieldDefinition(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("PUT",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafield_definition.json", cli.StoreHandle, cli.PathPrefix, ApiVersion),
		httpmock.NewStringResponder(200, `{"definition":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-10-09T14:48:44-04:00","access": {"admin":"MERCHANT_READ_WRITE"}, "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field"}}`))

	req := &metafield2.UpdateAMetafieldDefinitionAPIReq{
		Definition: metafield2.UpdateAMetafieldDefinitionAPIReqDefinition{
			Access:        metafield2.Access{Admin: "MERCHANT_READ_WRITE"},
			Description:   "test desc",
			Key:           "key_test",
			Name:          "name_test",
			Namespace:     "namespace_test",
			OwnerResource: "product",
		},
	}

	apiResp := &metafield2.UpdateAMetafieldDefinitionAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("MetafieldDefinition.Update returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.NotNil(t, apiResp.Definition.Id)
	assert.Equal(t, int64(1), apiResp.Definition.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", apiResp.Definition.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", apiResp.Definition.UpdatedAt)
	assert.Equal(t, "single_line_text_field", apiResp.Definition.Type)

}

// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/get-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/get-a-metafield-definition?version=v20251201
func TestGetMetafieldDefinition(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafield_definition.json", cli.StoreHandle, cli.PathPrefix, ApiVersion),
		httpmock.NewStringResponder(200, `{"definition":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-10-09T14:48:44-04:00","access": {"admin":"MERCHANT_READ_WRITE"}, "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field"}}`))

	req := &metafield2.GetAMetafieldDefinitionAPIReq{
		Id: "1",
	}

	apiResp := &metafield2.GetAMetafieldDefinitionAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("MetafieldDefinition.Get returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.NotNil(t, apiResp.Definition.Id)
	assert.Equal(t, int64(1), apiResp.Definition.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", apiResp.Definition.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", apiResp.Definition.UpdatedAt)
	assert.Equal(t, "single_line_text_field", apiResp.Definition.Type)

}

// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/get-metafield-definitions?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/get-metafield-definitions?version=v20251201
func TestListMetafieldDefinitions(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafield_definitions.json?owner_resource=product", cli.StoreHandle, cli.PathPrefix, ApiVersion),
		httpmock.NewBytesResponder(200, client.LoadTestDataV2("", "./metafield/metafield_definitions.json")))

	req := &metafield2.GetMetafieldDefinitionsAPIReq{
		OwnerResource: "product",
	}

	apiResp := &metafield2.GetMetafieldDefinitionsAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("MetafieldDefinition.List returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.NotNil(t, apiResp.Data)

	metafieldDefinition := apiResp.Data.MetafieldDefinitions[0]

	assert.NotNil(t, metafieldDefinition.Id)
	assert.Equal(t, int64(1), metafieldDefinition.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", metafieldDefinition.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", metafieldDefinition.UpdatedAt)
	assert.Equal(t, "single_line_text_field", metafieldDefinition.Type)

}

// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/get-metafield-definitions?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/get-metafield-definitions?version=v20251201
func TestListAllMetafieldDefinitions(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafield_definitions.json?owner_resource=product", cli.StoreHandle, cli.PathPrefix, ApiVersion),
		httpmock.NewBytesResponder(200, client.LoadTestDataV2("", "./metafield/metafield_definitions.json")))

	req := &metafield2.GetMetafieldDefinitionsAPIReq{
		OwnerResource: "product",
	}

	apiResp := &metafield2.GetMetafieldDefinitionsAPIResp{}
	definitions, err := client.ListAll(cli, context.Background(), req, apiResp, func(resp interface{}) []metafield2.MetafieldDefinition {
		r := (resp.(*metafield2.GetMetafieldDefinitionsAPIResp)).Data
		return r.MetafieldDefinitions
	})
	//err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("MetafieldDefinition.List returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

	metafieldDefinition := definitions[0]

	assert.NotNil(t, metafieldDefinition.Id)
	assert.Equal(t, int64(1), metafieldDefinition.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", metafieldDefinition.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", metafieldDefinition.UpdatedAt)
	assert.Equal(t, "single_line_text_field", metafieldDefinition.Type)

}

// DeleteMetafieldDefinitionAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/delete-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/delete-a-metafield-definition?version=v20251201
func TestDeleteMetafieldDefinition(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("DELETE",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafield_definition.json?id=1", cli.StoreHandle, cli.PathPrefix, ApiVersion),
		httpmock.NewStringResponder(200, ""))

	req := &metafield2.DeleteAMetafieldDefinitionAPIReq{
		Id: "1",
	}

	apiResp := &metafield2.DeleteAMetafieldDefinitionAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("MetafieldDefinition.Delete returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

}
