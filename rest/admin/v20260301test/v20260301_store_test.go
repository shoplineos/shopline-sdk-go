package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/test"
	store2 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStore(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/merchants/shop.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "../test/store/store.json")))

	apiReq := &store2.QueryStoreInformationAPIReq{}
	apiResp := &store2.QueryStoreInformationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp.Data)
	assert.Equal(t, int64(1), apiResp.Data.Id)
}

func TestListCurrencies(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/currency/currencies.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"currencies":[{"rate_update_at":"2025-09-22T14:48:44-04:00","currency":"CNY", "enabled":true}]}`))

	apiReq := &store2.QueryStoreSettlementCurrencyAPIReq{}

	apiResp := &store2.QueryStoreSettlementCurrencyAPIResp{}
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
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/store/staff/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "../test/store/store_staff.json")))

	apiReq := &store2.GetAStaffMemberAPIReq{
		Uid: "1",
	}
	apiResp := &store2.GetAStaffMemberAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)
	assert.Equal(t, "1", apiResp.Data.Uid)
}

func TestListStaffs(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/store/list/staff.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "../test/store/staffs.json")))

	apiReq := &store2.GetAllStoreStaffMembersAPIReq{
		Limit: "1",
	}

	apiResp := &store2.GetAllStoreStaffMembersAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, len(apiResp.Data))

	staff := apiResp.Data[0]
	assert.NotNil(t, staff)
	assert.Equal(t, "1", staff.Uid)
}

func TestGetOperationLog(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/store/operation_logs/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "../test/store/log.json")))

	apiReq := &store2.GetAStoreOperationLogAPIReq{
		Id: "1",
	}
	apiResp := &store2.GetAStoreOperationLogAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)

	assert.Equal(t, "1", apiResp.Data.Id)
}

func TestListOperationLogs(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/store/operation_logs.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "../test/store/logs.json")))

	apiReq := &store2.GetStoreOperationLogsAPIReq{
		CreatedAtMin: "2025-09-22T14:48:44-04:00",
		CreatedAtMax: "2025-10-10T14:48:44-04:00",
	}
	apiResp := &store2.GetStoreOperationLogsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, len(apiResp.Data))

	log := apiResp.Data[0]
	assert.Equal(t, "1", log.Id)
}

func TestCountOperationLogs(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/store/operation_logs/count.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"data":{"count":1}}`))

	apiReq := &store2.GetStoreOperationLogCountAPIReq{
		CreatedAtMin: "2025-09-22T14:48:44-04:00",
		CreatedAtMax: "2025-10-10T14:48:44-04:00",
	}

	apiResp := &store2.GetStoreOperationLogCountAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)

	assert.Equal(t, int64(1), apiResp.Data.Count)
}

func TestListSubscriptions(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/store/subscription", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestDataV2("", "../test/store/subscriptions.json")))

	apiReq := &store2.GetActiveStorePlansAPIReq{}
	apiResp := &store2.GetActiveStorePlansAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, apiResp)
	assert.Equal(t, 1, len(apiResp.Data))

	subscription := apiResp.Data[0]
	assert.Equal(t, "SUB6025977562148654208", subscription.SubId)
}
