package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/test"
	metafield2 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/metafield"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateStoreMetafield(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()
	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafields.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"metafield":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field", "value":"single_line_text_field_value"}}`))

	req := &metafield2.CreateAMetafieldForTheStoreAPIReq{
		Metafield: metafield2.CreateAMetafieldForTheStoreAPIReqMetafield{
			Description: "test desc",
			Key:         "key_test",
			Namespace:   "namespace_test",
			//OwnerId:       "123",
			//OwnerResource: "product",
			Type:  "single_line_text_field",
			Value: "single_line_text_field_value",
		},
	}

	apiResp := &metafield2.CreateAMetafieldForTheStoreAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)
	if err != nil {
		t.Errorf("Metafield.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, int64(1), apiResp.Metafield.Id)
	assert.Equal(t, "single_line_text_field_value", apiResp.Metafield.Value)
}

func TestUpdateStoreMetafield(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("PUT",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafields/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"metafield":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field", "value":"single_line_text_field_value"}}`))

	req := &metafield2.UpdateAStoreMetafieldAPIReq{
		Id: "1",
		Metafield: metafield2.UpdateAStoreMetafieldAPIReqMetafield{
			Description: "test desc",
			//Key:           "key_test",
			//Namespace:     "namespace_test",
			//OwnerId:       "123",
			//OwnerResource: "product",
			Type:  "single_line_text_field",
			Value: "single_line_text_field_value",
		},
	}

	apiResp := &metafield2.UpdateAStoreMetafieldAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)
	if err != nil {
		t.Errorf("Metafield.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, int64(1), apiResp.Metafield.Id)
	assert.Equal(t, "single_line_text_field_value", apiResp.Metafield.Value)
}

func TestDeleteStoreMetafield(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("DELETE",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafields/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	req := &metafield2.DeleteAStoreMetafieldAPIReq{
		Id: "1",
	}

	apiResp := &metafield2.DeleteAStoreMetafieldAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)
	if err != nil {
		t.Errorf("Metafield.Delete returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
}

func TestCountStoreMetafield(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafields/count.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"count": 1}`))

	req := &metafield2.GetTheStoreMetafieldCountAPIReq{}

	apiResp := &metafield2.GetTheStoreMetafieldCountAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)
	if err != nil {
		t.Errorf("Metafield.Count returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, apiResp.Count)
}

func TestListStoreMetafields(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafields.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "./metafield/metafields.json")))

	req := &metafield2.GetStoreMetafieldsAPIReq{}

	apiResp := &metafield2.GetStoreMetafieldsAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("Metafield.List returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

	metafield := apiResp.Metafields[0]

	assert.NotNil(t, metafield.Id)
	assert.Equal(t, int64(1), metafield.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", metafield.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", metafield.UpdatedAt)
	assert.Equal(t, "single_line_text_field", metafield.Type)

}

func getResources(resp interface{}) []metafield2.Metafield {
	apiResp := resp.(*metafield2.GetStoreMetafieldsAPIResp)
	return apiResp.Metafields
}

func TestListAllStoreMetafields(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/metafields.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "./metafield/metafields.json")))

	req := &metafield2.GetStoreMetafieldsAPIReq{}

	apiResp := &metafield2.GetStoreMetafieldsAPIResp{}
	metafields, err := client.ListAll(cli, context.Background(), req, apiResp, getResources)
	if err != nil {
		t.Errorf("Metafield.List returned error: %v", err)
	}

	metafield := metafields[0]

	assert.NotNil(t, metafield.Id)
	assert.Equal(t, int64(1), metafield.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", metafield.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", metafield.UpdatedAt)
	assert.Equal(t, "single_line_text_field", metafield.Type)

}
