package payment

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStoreBalance(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payments/store/balance.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"balance":[{"amount":"12.34", "currency":"CNY"}]}`))

	apiReq := &GetStoreBalanceAPIReq{}

	apiResp, err := GetPaymentStoreService().GetStoreBalance(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Payment.GetStoreBalance returned error: %v", err)
	}

	assert.NotNil(t, apiResp)
	assert.NotNil(t, apiResp.Balances)
	assert.Equal(t, 1, len(apiResp.Balances))

	balance := apiResp.Balances[0]
	assert.Equal(t, "CNY", balance.Currency)
}

func TestListStorePayouts(t *testing.T) {
	setup()
	defer teardown()
	//end_time=2025-09-30T00:00:00+08:00&limit=1&start_time=2025-04-30T00:00:00+08:00
	//httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payments/store/payouts.json?limit=1", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
	httpmock.RegisterResponder("GET", "https://zwapptest.myshopline.com/admin/openapi/v20251201/payments/store/payouts.json?end_time=2025-09-30T00%3A00%3A00%2B08%3A00&limit=1&start_time=2025-04-30T00%3A00%3A00%2B08%3A00",
		httpmock.NewBytesResponder(200, test.LoadTestData("payment/payouts.json")))

	apiReq := &ListStorePayoutsAPIReq{
		Limit: "1",
		//Status:    "SUCCESS",
		StartTime: "2025-04-30T00:00:00+08:00",
		EndTime:   "2025-09-30T00:00:00+08:00",
	}

	apiResp, err := GetPaymentStoreService().ListStorePayouts(context.Background(), apiReq)
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
		httpmock.NewBytesResponder(200, test.LoadTestData("payment/balance_transactions.json")))

	apiReq := &ListStoreBalanceTransactionsAPIReq{
		Limit:  "1",
		Status: "ok",
	}

	apiResp, err := GetPaymentStoreService().ListStoreBalanceTransactions(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Payment.ListStoreBalanceTransactions returned error: %v", err)
	}

	assert.NotNil(t, apiResp)

	assert.NotNil(t, apiResp.BalanceTransactions)
	assert.Equal(t, 1, len(apiResp.BalanceTransactions))

	trans := apiResp.BalanceTransactions[0]
	assert.Equal(t, "1.00", trans.Amount)
	assert.Equal(t, "202206161756376480011100000", trans.Id)
	assert.Equal(t, "in progress", trans.PayoutStatus)
}

func TestGetStorePayout(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.myshopline.com/%s/%s/payments/store/payout.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("payment/payouts.json")))

	apiReq := &GetStorePayoutAPIReq{
		PayoutId: "1",
	}

	apiResp, err := GetPaymentStoreService().GetStorePayout(context.Background(), apiReq)
	if err != nil {
		t.Errorf("Payment.GetStorePayout returned error: %v", err)
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
		httpmock.NewBytesResponder(200, test.LoadTestData("payment/transactions.json")))

	apiReq := &ListStoreTransactionsAPIReq{
		Limit:   "1",
		DateMin: "2025-04-30T00:00:00+08:00",
		DateMax: "2025-09-30T00:00:00+08:00",
	}

	apiResp, err := GetPaymentStoreService().ListStoreTransactions(context.Background(), apiReq)
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
