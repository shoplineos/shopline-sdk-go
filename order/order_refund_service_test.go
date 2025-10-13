package order

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderRefund(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(client.MethodPost,
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/order/refund.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, `{"data":{"id":"5144142", "order_id":"123"}}`))

	apiReq := &RefundAPIReq{
		OrderID: "123",
	}

	apiResponse, err := GetOrderRefundService().Refund(context.Background(), apiReq)
	if err != nil {
		t.Errorf("OrderRefund refund returned an error %v", err)
	}
	assert.Equal(t, "123", apiResponse.Refund.OrderId)
}

func TestOrderRefundList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/1/refunds.json?order_id=1", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("order/refunds.json")))

	apiReq := &ListOrderRefundsAPIReq{
		OrderID: "1",
	}

	apiResponse, err := GetOrderRefundService().List(context.Background(), apiReq)
	if err != nil {
		t.Errorf("OrderRefund list returned an error %v", err)
	}

	assert.NotNil(t, apiResponse)
	assert.NotNil(t, apiResponse.Refunds)
	assert.NotEmpty(t, apiResponse.Refunds)

	refund := apiResponse.Refunds[0]
	assert.NotEmpty(t, refund)

	assert.Equal(t, "ref_1234567890abcdef01234567", refund.ID)

}

func TestOrderRefundGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/1/refunds/ref_1234567890abcdef01234567.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("order/refund.json")))

	apiReq := &GetOrderRefundAPIReq{
		OrderID:  "1",
		RefundID: "ref_1234567890abcdef01234567",
	}

	apiResponse, err := GetOrderRefundService().Get(context.Background(), apiReq)
	if err != nil {
		t.Errorf("OrderRefund Get returned an error %v", err)
	}

	assert.NotNil(t, apiResponse)
	assert.NotNil(t, apiResponse.Refund)

	refund := apiResponse.Refund
	assert.NotEmpty(t, refund)

	assert.Equal(t, "ref_1234567890abcdef01234567", refund.ID)

}

func TestOrderRefundCalc(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/1/refunds/calculate.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("order/calculate.json")))

	apiReq := &CalculateOrderRefundAPIReq{
		OrderID:  "1",
		Currency: "CNY",
	}

	apiResponse, err := GetOrderRefundService().Calculate(context.Background(), apiReq)
	if err != nil {
		t.Errorf("OrderRefund.Calculate returned an error %v", err)
	}

	assert.NotNil(t, apiResponse)
	assert.NotNil(t, apiResponse.CalculateRefundResult)

	refund := apiResponse.CalculateRefundResult
	assert.NotNil(t, refund)
	assert.NotNil(t, refund.TotalDutiesSet)
	assert.NotNil(t, refund.TotalDutiesSet.PresentmentMoney)

	assert.Equal(t, "USD", refund.Currency)
	assert.Equal(t, "USD", refund.TotalDutiesSet.PresentmentMoney.CurrencyCode)

}
