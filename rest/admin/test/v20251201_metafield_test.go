package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	metafield2 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/metafield"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"metafield":{"id":"1", "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field", "value":"single_line_text_field_value"}}`))

	req := &metafield2.CreateMetafieldAPIReq{
		Metafield: metafield2.CreateMetafield{
			Description:   "test desc",
			Key:           "key_test",
			Namespace:     "namespace_test",
			OwnerId:       "123",
			OwnerResource: "product",
			Type:          "single_line_text_field",
			Value:         "single_line_text_field_value",
		},
	}

	apiResp := &metafield2.CreateMetafieldAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("Metafield.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, "1", apiResp.Metafield.Id)
	assert.Equal(t, "single_line_text_field_value", apiResp.Metafield.Value)
}

func TestUpdateMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"metafield":{"id":"1", "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field", "value":"single_line_text_field_value"}}`))

	req := &metafield2.UpdateMetafieldAPIReq{
		Metafield: metafield2.UpdateMetafield{
			Id:            "1",
			Description:   "test desc",
			Key:           "key_test",
			Namespace:     "namespace_test",
			OwnerId:       "123",
			OwnerResource: "product",
			Type:          "single_line_text_field",
			Value:         "single_line_text_field_value",
		},
	}

	apiResp := &metafield2.UpdateMetafieldAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("Metafield.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, "1", apiResp.Metafield.Id)
	assert.Equal(t, "single_line_text_field_value", apiResp.Metafield.Value)
}

func TestDeleteMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	req := &metafield2.DeleteMetafieldAPIReq{
		Id:            "1",
		OwnerId:       "123",
		OwnerResource: "product",
	}

	//apiResp, err := GetMetafieldService().Delete(context.Background(), req)
	apiResp := &metafield2.DeleteMetafieldAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("Metafield.Delete returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
}

func TestCountMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields/count.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"count": 1}`))

	req := &metafield2.CountMetafieldAPIReq{
		OwnerId:       "123",
		OwnerResource: "product",
	}

	//apiResp, err := GetMetafieldService().Count(context.Background(), req)
	apiResp := &metafield2.CountMetafieldAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("Metafield.Count returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, apiResp.Count)
}

func TestListMetafields(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "metafield/metafields.json")))

	req := &metafield2.ListMetafieldAPIReq{
		OwnerResource: "product",
		OwnerId:       "123",
	}
	apiResp := &metafield2.ListMetafieldAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)
	//apiResp, err := GetMetafieldService().List(context.Background(), req)

	if err != nil {
		t.Errorf("Metafield.List returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

	metafield := apiResp.Metafields[0]

	assert.NotNil(t, metafield.Id)
	assert.Equal(t, "1", metafield.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", metafield.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", metafield.UpdatedAt)
	assert.Equal(t, "single_line_text_field", metafield.Type)

}

func TestListAllMetafields(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "metafield/metafields.json")))

	req := &metafield2.ListMetafieldAPIReq{
		OwnerResource: "product",
		OwnerId:       "123",
	}
	apiResp := &metafield2.ListMetafieldAPIResp{}
	metafields, err := client.ListAll(cli, context.Background(), req, apiResp, func(resp interface{}) []metafield2.Metafield {
		r := resp.(*metafield2.ListMetafieldAPIResp)
		return r.Metafields
	})

	if err != nil {
		t.Errorf("Metafield.List returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

	metafield := metafields[0]

	assert.NotNil(t, metafield.Id)
	assert.Equal(t, "1", metafield.Id)
	assert.Equal(t, "2025-09-22T14:48:44-04:00", metafield.CreatedAt)
	assert.Equal(t, "2025-10-09T14:48:44-04:00", metafield.UpdatedAt)
	assert.Equal(t, "single_line_text_field", metafield.Type)

}

func TestGetMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/product/123/metafields/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"metafield":{"id":"1", "created_at":"2025-09-22T14:48:44-04:00", "updated_at":"2025-09-22T14:48:44-04:00", "description":"test desc", "key":"key_test", "name":"name_test", "namespace":"namespace_test", "owner_resource":"product", "type":"single_line_text_field", "value":"single_line_text_field_value"}}`))

	req := &metafield2.GetMetafieldAPIReq{
		Id:            "1",
		OwnerId:       "123",
		OwnerResource: "product",
	}

	//apiResp, err := GetMetafieldService().Get(context.Background(), req)
	apiResp := &metafield2.GetMetafieldAPIResp{}
	err := cli.Call(context.Background(), req, apiResp)

	if err != nil {
		t.Errorf("Metafield.Create returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.Equal(t, "1", apiResp.Metafield.Id)
	assert.Equal(t, "single_line_text_field_value", apiResp.Metafield.Value)
}
