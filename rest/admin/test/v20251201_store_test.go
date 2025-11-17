package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	store2 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStore(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/merchants/shop.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "store/store.json")))

	apiReq := &store2.GetStoreAPIReq{}
	apiResp := &store2.GetStoreAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp.Store)
	assert.Equal(t, uint64(1), apiResp.Store.Id)
}

func TestListCurrencies(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/currency/currencies.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"currencies":[{"rate_update_at":"2025-09-22T14:48:44-04:00","currency":"CNY", "enabled":true}]}`))

	apiReq := &store2.ListStoreCurrenciesAPIReq{}

	apiResp := &store2.ListStoreCurrenciesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, len(apiResp.Currencies))

	currency := apiResp.Currencies[0]

	assert.Equal(t, "CNY", currency.Currency)
}

func TestGetStaff(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/store/staff/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "store/store_staff.json")))

	apiReq := &store2.GetStoreStaffAPIReq{
		UID: "1",
	}
	apiResp := &store2.GetStoreStaffAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)
	assert.Equal(t, "1", apiResp.StoreStaff.UID)
}

func TestListStaffs(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/store/list/staff.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "store/staffs.json")))

	apiReq := &store2.ListStoreStaffsAPIReq{
		Limit: "1",
	}

	apiResp := &store2.ListStoreStaffsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, len(apiResp.StoreStaffs))

	staff := apiResp.StoreStaffs[0]
	assert.NotNil(t, staff)
	assert.Equal(t, "1", staff.UID)
}

func TestGetOperationLog(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/store/operation_logs/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "store/log.json")))

	apiReq := &store2.GetStoreOperationLogAPIReq{
		ID: "1",
	}
	apiResp := &store2.GetStoreOperationLogAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)

	assert.Equal(t, "1", apiResp.OperationLog.Id)
}

func TestListOperationLogs(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/store/operation_logs.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "store/logs.json")))

	apiReq := &store2.ListStoreOperationLogsAPIReq{
		CreatedAtMin: "2025-09-22T14:48:44-04:00",
		CreatedAtMax: "2025-10-10T14:48:44-04:00",
	}
	apiResp := &store2.ListStoreOperationLogsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, len(apiResp.OperationLogs))

	log := apiResp.OperationLogs[0]
	assert.Equal(t, "1", log.Id)
}

func TestCountOperationLogs(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/store/operation_logs/count.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"data":{"count":1}}`))

	apiReq := &store2.CountStoreOperationLogsAPIReq{
		CreatedAtMin: "2025-09-22T14:48:44-04:00",
		CreatedAtMax: "2025-10-10T14:48:44-04:00",
	}

	apiResp := &store2.CountStoreOperationLogsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)

	assert.Equal(t, 1, apiResp.CountOperationLogData.Count)
}

func TestListSubscriptions(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/store/subscription", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "store/subscriptions.json")))

	apiReq := &store2.ListStoreSubscriptionsAPIReq{}
	apiResp := &store2.ListStoreSubscriptionsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, len(apiResp.Subscriptions))

	subscription := apiResp.Subscriptions[0]
	assert.Equal(t, "SUB6025977562148654208", subscription.SubId)
}
