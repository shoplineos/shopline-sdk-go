package store

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStore(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/merchants/shop.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("store/store.json")))

	apiReq := &GetStoreAPIReq{}
	storeInfo, err := GetStoreService().Get(context.Background(), apiReq)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, storeInfo.Store)
	assert.Equal(t, uint64(1), storeInfo.Store.Id)
}

func TestListCurrencies(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/currency/currencies.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"currencies":[{"rate_update_at":"2025-09-22T14:48:44-04:00","currency":"CNY", "enabled":true}]}`))

	apiReq := &ListStoreCurrenciesAPIReq{}
	apiResp, err := GetStoreService().ListCurrencies(context.Background(), apiReq)
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
		httpmock.NewBytesResponder(200, test.LoadTestData("store/store_staff.json")))

	apiReq := &GetStoreStaffAPIReq{
		UID: "1",
	}
	apiResp, err := GetStoreService().GetStaff(context.Background(), apiReq)
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
		httpmock.NewBytesResponder(200, test.LoadTestData("store/staffs.json")))

	apiReq := &ListStoreStaffsAPIReq{
		Limit: "1",
	}
	apiResp, err := GetStoreService().ListStaffs(context.Background(), apiReq)
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
		httpmock.NewBytesResponder(200, test.LoadTestData("store/log.json")))

	apiReq := &GetStoreOperationLogAPIReq{
		ID: "1",
	}
	apiResp, err := GetStoreService().GetOperationLog(context.Background(), apiReq)
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
		httpmock.NewBytesResponder(200, test.LoadTestData("store/logs.json")))

	apiReq := &ListStoreOperationLogsAPIReq{
		CreatedAtMin: "2025-09-22T14:48:44-04:00",
		CreatedAtMax: "2025-10-10T14:48:44-04:00",
	}
	apiResp, err := GetStoreService().ListOperationLogs(context.Background(), apiReq)
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

	apiReq := &CountStoreOperationLogsAPIReq{
		CreatedAtMin: "2025-09-22T14:48:44-04:00",
		CreatedAtMax: "2025-10-10T14:48:44-04:00",
	}
	apiResp, err := GetStoreService().CountOperationLogs(context.Background(), apiReq)
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
		httpmock.NewBytesResponder(200, test.LoadTestData("store/subscriptions.json")))

	apiReq := &ListStoreSubscriptionsAPIReq{}
	apiResp, err := GetStoreService().ListSubscriptions(context.Background(), apiReq)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, len(apiResp.Subscriptions))

	subscription := apiResp.Subscriptions[0]
	assert.Equal(t, "SUB6025977562148654208", subscription.SubId)
}
