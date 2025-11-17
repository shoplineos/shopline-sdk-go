package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	payment2 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/payment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStoreBalance(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payments/store/balance.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"balance":[{"amount":"12.34", "currency":"CNY"}]}`))

	apiReq := &payment2.BalanceAPIReq{}
	apiResp := &payment2.BalanceAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("Payment.GetStoreBalance returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.NotNil(t, apiResp.Balance)
	assert.Equal(t, 1, len(apiResp.Balance))

	balance := apiResp.Balance[0]
	assert.Equal(t, "CNY", balance.Currency)
}

func TestListStorePayouts(t *testing.T) {
	setup()
	defer teardown()
	//end_time=2025-09-30T00:00:00+08:00&limit=1&start_time=2025-04-30T00:00:00+08:00
	//httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payments/store/payouts.json?limit=1", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
	httpmock.RegisterResponder("GET", "https://zwapptest.myshopline.com/admin/openapi/v20251201/payments/store/payouts.json?end_time=2025-09-30T00%3A00%3A00%2B08%3A00&limit=1&start_time=2025-04-30T00%3A00%3A00%2B08%3A00",
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "payment/payouts.json")))

	apiReq := &payment2.GetPayoutsAPIReq{
		Limit: "1",
		//Status:    "SUCCESS",
		StartTime: "2025-04-30T00:00:00+08:00",
		EndTime:   "2025-09-30T00:00:00+08:00",
	}

	apiResp := &payment2.GetPayoutsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)
	if err != nil {
		t.Errorf("Payment.ListStorePayouts returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.NotNil(t, apiResp.Payouts)
	assert.Equal(t, 1, len(apiResp.Payouts))

	payout := apiResp.Payouts[0]
	assert.Equal(t, "12.34", payout.Amount)
	assert.Equal(t, "SUCCESS", payout.Status)
}

func TestListStoreBalanceTransactions(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payments/store/balance_transactions.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "payment/balance_transactions.json")))

	apiReq := &payment2.BalanceTransactionsAPIReq{
		Limit:        "1",
		PayoutStatus: "ok",
	}

	apiResp := &payment2.BalanceTransactionsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("Payment.ListStoreBalanceTransactions returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

	assert.NotNil(t, apiResp.Transactions)
	assert.Equal(t, 1, len(apiResp.Transactions))

	trans := apiResp.Transactions[0]
	assert.Equal(t, "1.00", trans.Amount)
	assert.Equal(t, "202206161756376480011100000", trans.Id)
	assert.Equal(t, "in progress", trans.PayoutStatus)
}

func TestGetStorePayouts(t *testing.T) {
	setup()
	defer teardown()

	// https://zwapptest.myshopline.com/admin/openapi/v20251201/payments/store/payouts.json?end_time=2025-05-30&limit=2&payout_transaction_no=2222&start_time=2025-04-30
	//httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payments/store/payout.json?end_time=2025-05-30&limit=2&payout_transaction_no=2222&start_time=2025-04-30", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
	httpmock.RegisterResponder("GET", "https://zwapptest.myshopline.com/admin/openapi/v20251201/payments/store/payouts.json?end_time=2025-05-30&limit=2&payout_transaction_no=2222&start_time=2025-04-30",
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "payment/payouts.json")))

	apiReq := &payment2.GetPayoutsAPIReq{
		PayoutTransactionNo: "2222",
		StartTime:           "2025-04-30",
		EndTime:             "2025-05-30",
		Limit:               "2",
	}

	apiResp := &payment2.GetPayoutsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("Payment.GetStorePayouts returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.NotNil(t, apiResp.Payouts)
	assert.Equal(t, 1, len(apiResp.Payouts))

	payout := apiResp.Payouts[0]
	assert.Equal(t, "12.34", payout.Amount)
	assert.Equal(t, "SUCCESS", payout.Status)

}

func TestListStoreTransactions(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payments/store/transactions.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, LoadTestDataV2("", "payment/transactions.json")))

	apiReq := &payment2.QueryStoreTransactionRecordsAPIReq{
		Limit:   "1",
		DateMin: "2025-04-30T00:00:00+08:00",
		DateMax: "2025-09-30T00:00:00+08:00",
	}

	apiResp := &payment2.QueryStoreTransactionRecordsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	if err != nil {
		t.Errorf("Payment.ListStoreTransactions returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

	assert.NotNil(t, apiResp.Transactions)
	assert.Equal(t, 1, len(apiResp.Transactions))

	trans := apiResp.Transactions[0]
	assert.Equal(t, "10.10", trans.Amount)
	assert.Equal(t, "10010062529029006852253184000", trans.TradeOrderId)
	assert.Equal(t, "SUCCEEDED", trans.Status)

}
