package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/test"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/bulkoperation"
	"github.com/stretchr/testify/assert"
)

func bulkURL(cli *client.Client, path string) string {
	return fmt.Sprintf("https://%s.myshopline.com/%s/%s/%s",
		cli.StoreHandle, cli.PathPrefix, cli.ApiVersion, path)
}

func TestCreateABulkQueryTask(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"id":"task-001","status":"CREATED","type":"QUERY","storeId":"store-001","createdAt":"2023-04-06T10:48:29+08:00"}`
	httpmock.RegisterResponder("POST", bulkURL(cli, "bulk_operation_run_query.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &bulkoperation.CreateABulkQueryTaskAPIReq{QueryPath: "products/products.json"}
	apiResp := &bulkoperation.CreateABulkQueryTaskAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "task-001", apiResp.Id)
	assert.Equal(t, "CREATED", apiResp.Status)
	assert.Equal(t, "QUERY", apiResp.Type)
}

func TestCreateABulkQueryTask_MissingQueryPath(t *testing.T) {
	err := (&bulkoperation.CreateABulkQueryTaskAPIReq{}).Verify()
	assert.EqualError(t, err, "QueryPath is required")
}

func TestCreateABulkMutationTask(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"id":"task-002","status":"CREATED","type":"MUTATION_GENERAL","storeId":"store-001","createdAt":"2023-04-07T10:48:29+08:00"}`
	httpmock.RegisterResponder("POST", bulkURL(cli, "bulk_operation_run_mutation_general.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &bulkoperation.CreateABulkMutationTaskAPIReq{}
	apiResp := &bulkoperation.CreateABulkMutationTaskAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "task-002", apiResp.Id)
	assert.Equal(t, "CREATED", apiResp.Status)
	assert.Equal(t, "MUTATION_GENERAL", apiResp.Type)
}

func TestCreateAChangeTypeBulkTaskToBeDeprecated(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"id":"task-003","status":"CREATED","type":"MUTATION","storeId":"store-001","createdAt":"2023-04-06T10:48:29+08:00"}`
	httpmock.RegisterResponder("POST", bulkURL(cli, "bulk_operation_run_mutation.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &bulkoperation.CreateAChangeTypeBulkTaskToBeDeprecatedAPIReq{}
	apiResp := &bulkoperation.CreateAChangeTypeBulkTaskToBeDeprecatedAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "task-003", apiResp.Id)
	assert.Equal(t, "CREATED", apiResp.Status)
	assert.Equal(t, "MUTATION", apiResp.Type)
}

func TestGetAValidBulkTask(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"id":"task-001","status":"RUNNING","type":"QUERY","storeId":"store-001","createdAt":"2023-04-06T10:48:29+08:00","completedCount":50}`
	httpmock.RegisterResponder("GET", bulkURL(cli, "current_bulk_operation.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &bulkoperation.GetAValidBulkTaskAPIReq{}
	apiResp := &bulkoperation.GetAValidBulkTaskAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "task-001", apiResp.Id)
	assert.Equal(t, "RUNNING", apiResp.Status)
	assert.Equal(t, "QUERY", apiResp.Type)
	assert.Equal(t, 50, apiResp.CompletedCount)
}

func TestCancelABulkTask(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"id":"task-001","status":"CANCELING","type":"QUERY","storeId":"store-001","createdAt":"2023-04-06T10:48:29+08:00"}`
	httpmock.RegisterResponder("POST", bulkURL(cli, "bulk_operation_run_cancel.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &bulkoperation.CancelABulkTaskAPIReq{}
	apiResp := &bulkoperation.CancelABulkTaskAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "task-001", apiResp.Id)
	assert.Equal(t, "CANCELING", apiResp.Status)
}
