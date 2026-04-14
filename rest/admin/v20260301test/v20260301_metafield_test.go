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

func TestCreateMetafield(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"metafield":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field", "value":"single_line_text_field_value"}}`))

	req := &metafield2.CreateAMetafieldForAResourceAPIReq{
		OwnerId:  "123",
		Resource: "product",
		Metafield: metafield2.CreateAMetafieldForAResourceAPIReqMetafield{
			Description: "test desc",
			Key:         "key_test",
			Namespace:   "namespace_test",
			//OwnerId:       "123",
			//OwnerResource: "product",
			Type:  "single_line_text_field",
			Value: "single_line_text_field_value",
		},
	}

	apiResp := &metafield2.CreateAMetafieldForAResourceAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("Metafield.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, int64(1), apiResp.Metafield.Id)
	assert.Equal(t, "single_line_text_field_value", apiResp.Metafield.Value)
}

func TestUpdateMetafield(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("PUT",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"metafield":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field", "value":"single_line_text_field_value"}}`))

	req := &metafield2.UpdateAMetafieldForAResourceAPIReq{
		//Key:           "key_test",
		//Namespace:     "namespace_test",
		Id:       "1",
		OwnerId:  "123",
		Resource: "product",
		Metafield: metafield2.UpdateAMetafieldForAResourceAPIReqMetafield{
			Id:          1,
			Description: "test desc",
			//Key:           "key_test",
			//Namespace:     "namespace_test",
			//OwnerId:       "123",
			//OwnerResource: "product",
			Type:  "single_line_text_field",
			Value: "single_line_text_field_value",
		},
	}

	apiResp := &metafield2.UpdateAMetafieldForAResourceAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("Metafield.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, int64(1), apiResp.Metafield.Id)
	assert.Equal(t, "single_line_text_field_value", apiResp.Metafield.Value)
}

func TestDeleteMetafield(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("DELETE",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	req := &metafield2.DeleteAMetafieldAPIReq{
		Id:       "1",
		OwnerId:  "123",
		Resource: "product",
	}

	apiResp := &metafield2.DeleteAMetafieldAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("Metafield.Delete returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
}

func TestCountMetafield(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields/count.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"count": 1}`))

	req := &metafield2.GetTheMetafieldCountForAResourceAPIReq{
		OwnerId:  "123",
		Resource: "product",
	}

	apiResp := &metafield2.GetTheMetafieldCountForAResourceAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("Metafield.Count returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, apiResp.Count)
}

func TestListMetafields(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "./metafield/metafields.json")))

	req := &metafield2.GetMetafieldsForAResourceAPIReq{
		Resource: "product",
		OwnerId:  "123",
	}
	apiResp := &metafield2.GetMetafieldsForAResourceAPIResp{}
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

func TestListAllMetafields(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "./metafield/metafields.json")))

	req := &metafield2.GetMetafieldsForAResourceAPIReq{
		Resource: "product",
		OwnerId:  "123",
	}
	apiResp := &metafield2.GetMetafieldsForAResourceAPIResp{}
	metafields, err := client.ListAll(cli, context.Background(), req, apiResp, func(resp interface{}) []metafield2.Metafield {
		r := resp.(*metafield2.GetMetafieldsForAResourceAPIResp)
		return r.Metafields
	})

	if err != nil {
		t.Errorf("Metafield.List returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

	metafield := metafields[0]

	assert.NotNil(t, metafield.Id)
	assert.Equal(t, int64(1), metafield.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", metafield.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", metafield.UpdatedAt)
	assert.Equal(t, "single_line_text_field", metafield.Type)

}

func TestGetMetafield(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"metafield":{"id":1, "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field", "value":"single_line_text_field_value"}}`))

	req := &metafield2.GetAMetafieldForAResourceAPIReq{
		Id:       "1",
		OwnerId:  "123",
		Resource: "product",
	}

	//apiResp, err := GetMetafieldService().Get(context.Background(), req)
	apiResp := &metafield2.GetAMetafieldForAResourceAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("Metafield.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, int64(1), apiResp.Metafield.Id)
	assert.Equal(t, "single_line_text_field_value", apiResp.Metafield.Value)
}
