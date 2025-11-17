package metafield

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/create-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/create-a-metafield-definition?version=v20251201
func TestCreateMetafieldDefinition(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafield_definition.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"definition":{"id":"1", "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00","access": {"admin":"MERCHANT_READ_WRITE"}, "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field"}}`))

	req := &CreateMetafieldDefinitionAPIReq{
		MetafieldDefinition: CreateMetafieldDefinition{
			Access:        Access{Admin: "MERCHANT_READ_WRITE"},
			Description:   "test desc",
			Key:           "key_test",
			Name:          "name_test",
			Namespace:     "namespace_test",
			OwnerResource: "product",
			Type:          "single_line_text_field",
		},
	}

	apiResp, err := GetMetafieldDefinitionService().Create(context.Background(), req)
	if err != nil {
		t.Errorf("MetafieldDefinition.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.NotNil(t, apiResp.MetafieldDefinition.Id)
	assert.Equal(t, "1", apiResp.MetafieldDefinition.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", apiResp.MetafieldDefinition.CreatedAt)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", apiResp.MetafieldDefinition.UpdatedAt)

}

// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/update-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/update-a-metafield-definition?version=v20251201
func TestUpdateMetafieldDefinition(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafield_definition.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"definition":{"id":"1", "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-10-09T14:48:44-04:00","access": {"admin":"MERCHANT_READ_WRITE"}, "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field"}}`))

	req := &UpdateMetafieldDefinitionAPIReq{
		MetafieldDefinition: UpdateMetafieldDefinition{
			Access:        Access{Admin: "MERCHANT_READ_WRITE"},
			Description:   "test desc",
			Key:           "key_test",
			Name:          "name_test",
			Namespace:     "namespace_test",
			OwnerResource: "product",
		},
	}

	apiResp, err := GetMetafieldDefinitionService().Update(context.Background(), req)
	if err != nil {
		t.Errorf("MetafieldDefinition.Update returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.NotNil(t, apiResp.MetafieldDefinition.Id)
	assert.Equal(t, "1", apiResp.MetafieldDefinition.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", apiResp.MetafieldDefinition.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", apiResp.MetafieldDefinition.UpdatedAt)
	assert.Equal(t, "single_line_text_field", apiResp.MetafieldDefinition.Type)

}

// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/get-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/get-a-metafield-definition?version=v20251201
func TestGetMetafieldDefinition(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafield_definition.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"definition":{"id":"1", "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-10-09T14:48:44-04:00","access": {"admin":"MERCHANT_READ_WRITE"}, "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field"}}`))

	req := &GetMetafieldDefinitionAPIReq{
		Id: "1",
	}

	apiResp, err := GetMetafieldDefinitionService().Get(context.Background(), req)
	if err != nil {
		t.Errorf("MetafieldDefinition.Get returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.NotNil(t, apiResp.MetafieldDefinition.Id)
	assert.Equal(t, "1", apiResp.MetafieldDefinition.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", apiResp.MetafieldDefinition.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", apiResp.MetafieldDefinition.UpdatedAt)
	assert.Equal(t, "single_line_text_field", apiResp.MetafieldDefinition.Type)

}

// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/get-metafield-definitions?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/get-metafield-definitions?version=v20251201
func TestListMetafieldDefinitions(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafield_definitions.json?owner_resource=product", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("metafield/metafield_definitions.json")))

	req := &ListMetafieldDefinitionAPIReq{
		OwnerResource: "product",
	}

	apiResp, err := GetMetafieldDefinitionService().List(context.Background(), req)
	if err != nil {
		t.Errorf("MetafieldDefinition.List returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.NotNil(t, apiResp.Data)

	metafieldDefinition := apiResp.Data.MetafieldDefinitions[0]

	assert.NotNil(t, metafieldDefinition.Id)
	assert.Equal(t, "1", metafieldDefinition.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", metafieldDefinition.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", metafieldDefinition.UpdatedAt)
	assert.Equal(t, "single_line_text_field", metafieldDefinition.Type)

}

// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/get-metafield-definitions?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/get-metafield-definitions?version=v20251201
func TestListAllMetafieldDefinitions(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafield_definitions.json?owner_resource=product", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("metafield/metafield_definitions.json")))

	req := &ListMetafieldDefinitionAPIReq{
		OwnerResource: "product",
	}

	apiResp, err := GetMetafieldDefinitionService().ListAll(context.Background(), req)
	if err != nil {
		t.Errorf("MetafieldDefinition.List returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

	metafieldDefinition := apiResp[0]

	assert.NotNil(t, metafieldDefinition.Id)
	assert.Equal(t, "1", metafieldDefinition.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", metafieldDefinition.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", metafieldDefinition.UpdatedAt)
	assert.Equal(t, "single_line_text_field", metafieldDefinition.Type)

}

// DeleteMetafieldDefinitionAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/delete-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/delete-a-metafield-definition?version=v20251201
func TestDeleteMetafieldDefinition(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafield_definition.json?id=1", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	req := &DeleteMetafieldDefinitionAPIReq{
		Id: "1",
	}

	apiResp, err := GetMetafieldDefinitionService().Delete(context.Background(), req)
	if err != nil {
		t.Errorf("MetafieldDefinition.Delete returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

}
