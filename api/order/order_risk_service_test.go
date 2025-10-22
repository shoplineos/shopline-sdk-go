package order

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/v2/21056921623197625897868288/risks.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("order/risks.json")))

	apiReq := &ListOrderRisksAPIReq{
		OrderId: "21056921623197625897868288",
	}

	apiResponse, err := GetOrderRiskService().List(context.Background(), apiReq)
	if err != nil {
		t.Errorf("OrderRisk list returned an error %v", err)
	}
	assert.NotNil(t, apiResponse)
	assert.NotEmpty(t, apiResponse.OrderRisks)
	risk := apiResponse.OrderRisks[0]
	assert.Equal(t, "21056921623197625897868288", risk.OrderId)
}

func TestGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/v2/21056921623197625897868288/risks/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("order/risk.json")))

	apiReq := &GetOrderRiskAPIReq{
		OrderId: "21056921623197625897868288",
		RiskId:  "1",
	}

	apiResponse, err := GetOrderRiskService().Get(context.Background(), apiReq)
	if err != nil {
		t.Errorf("OrderRisk get returned an error %v", err)
	}
	assert.NotNil(t, apiResponse)
	assert.NotEmpty(t, apiResponse.OrderRisk)
	assert.Equal(t, "21056921623197625897868288", apiResponse.OrderRisk.OrderId)
	assert.Equal(t, "1", apiResponse.OrderRisk.Id)
}

func TestCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/v2/21056921623197625897868288/risks.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("order/risk.json")))

	apiReq := &CreateOrderRiskAPIReq{
		OrderId: "21056921623197625897868288",
		CreateOrderRisk: CreateOrderRisk{
			Display: true,
		},
	}

	apiResponse, err := GetOrderRiskService().Create(context.Background(), apiReq)
	if err != nil {
		t.Errorf("OrderRisk create returned an error %v", err)
	}
	assert.NotNil(t, apiResponse)
	assert.NotEmpty(t, apiResponse.OrderRisk)
	assert.Equal(t, "21056921623197625897868288", apiResponse.OrderRisk.OrderId)
	assert.Equal(t, "1", apiResponse.OrderRisk.Id)
}

func TestUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/v2/21056921623197625897868288/risks/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewBytesResponder(200, test.LoadTestData("order/risk.json")))

	apiReq := &UpdateOrderRiskAPIReq{
		OrderId: "21056921623197625897868288",
		RiskId:  "1",
		UpdateOrderRisk: UpdateOrderRisk{
			Display: true,
		},
	}

	apiResponse, err := GetOrderRiskService().Update(context.Background(), apiReq)
	if err != nil {
		t.Errorf("OrderRisk Update returned an error %v", err)
	}
	assert.NotNil(t, apiResponse)
	assert.NotEmpty(t, apiResponse.OrderRisk)
	assert.Equal(t, "21056921623197625897868288", apiResponse.OrderRisk.OrderId)
	assert.Equal(t, "1", apiResponse.OrderRisk.Id)
}

func TestDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/v2/21056921623197625897868288/risks/1.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	apiReq := &DeleteOrderRiskAPIReq{
		OrderId: "21056921623197625897868288",
		RiskId:  "1",
	}

	apiResponse, err := GetOrderRiskService().Delete(context.Background(), apiReq)
	if err != nil {
		t.Errorf("OrderRisk Delete returned an error %v", err)
	}
	assert.NotNil(t, apiResponse)
}

func TestDeleteAll(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE",
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/v2/21056921623197625897868288/risks.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	apiReq := &DeleteOrderRisksAPIReq{
		OrderId: "21056921623197625897868288",
	}

	apiResponse, err := GetOrderRiskService().DeleteAll(context.Background(), apiReq)
	if err != nil {
		t.Errorf("OrderRisk DeleteAll returned an error %v", err)
	}
	assert.NotNil(t, apiResponse)
}
