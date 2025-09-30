package order

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
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
		OrderId: "123",
	}

	apiResponse, err := OrderRefund(cli, apiReq)
	if err != nil {
		t.Errorf("OrderRefund returned an error %v", err)
	}
	assert.Equal(t, "123", apiResponse.Refund.OrderId)

}
