package test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/test"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/order"
	"github.com/stretchr/testify/assert"
)

func orderTestsForListOrdersAPI(t *testing.T, o order.GetOrdersAPIRespOrder) {
	if !assert.Equal(t, "2025-09-30T10:14:36-00:00", o.CreatedAt) {
		t.Errorf("Order.CreatedAt returned %+v, expected %+v", o.CreatedAt, "2025-09-30T10:14:36-00:00")
	}
	if o.ProcessedAt != "" {
		t.Errorf("Order.ProcessedAt returned %+v, expected %+v", o.ProcessedAt, nil)
	}
	if !strings.EqualFold("9.50", o.CurrentTotalPrice) {
		t.Errorf("Order.CurrentTotalPrice returned %+v, expected %+v", o.CurrentTotalPrice, "9.50")
	}
	if !strings.EqualFold("1.00", o.CurrentTotalTax) {
		t.Errorf("Order.CurrentTotalTax returned %+v, expected %+v", o.CurrentTotalTax, "1.00")
	}
	if o.Customer.Email != "john@test.com" {
		t.Errorf("Customer.Email, expected %v, actual %v", "john@test.com", o.Customer.Email)
	}
	lineItem := o.LineItems[0]
	if !assert.Equal(t, "1.00", lineItem.Price) {
		t.Errorf("Order.LineItems[0].Price, expected %v, actual %v", "1.00", lineItem.Price)
	}
}

func orderURL(cli *client.Client, path string) string {
	return fmt.Sprintf("https://%s.myshopline.com/%s/%s/%s",
		cli.StoreHandle, cli.PathPrefix, cli.ApiVersion, path)
}

// ══════════════════════════════════════════════════════════════════════════════
// Orders
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateAnOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"order":{"id":"order-001","order_number":"1001"}}`
	httpmock.RegisterResponder("POST", orderURL(cli, "orders.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &order.CreateAnOrderAPIReq{}
	apiResp := &order.CreateAnOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "order-001", apiResp.Order.Id)
}

func TestGetOrders(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"orders":[{"id":"order-001","order_number":"1001"}]}`
	httpmock.RegisterResponder("GET", orderURL(cli, "orders.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &order.GetOrdersAPIReq{}
	apiResp := &order.GetOrdersAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Orders, 1)
	assert.Equal(t, "order-001", apiResp.Orders[0].Id)
}

func TestGetAnOrderCount(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "orders/count.json"),
		httpmock.NewStringResponder(200, `{"count":5}`))

	apiReq := &order.GetAnOrderCountAPIReq{}
	apiResp := &order.GetAnOrderCountAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, int64(5), apiResp.Count)
}

func TestUpdateAnOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("PUT", orderURL(cli, fmt.Sprintf("orders/%s.json", orderId)),
		httpmock.NewStringResponder(200, `{"order":{"id":"order-001"}}`))

	apiReq := &order.UpdateAnOrderAPIReq{Id: orderId}
	apiResp := &order.UpdateAnOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestUpdateAnOrder_MissingId(t *testing.T) {
	err := (&order.UpdateAnOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestDeleteAnOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("DELETE", orderURL(cli, fmt.Sprintf("orders/%s.json", orderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.DeleteAnOrderAPIReq{OrderId: orderId}
	apiResp := &order.DeleteAnOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteAnOrder_MissingOrderId(t *testing.T) {
	err := (&order.DeleteAnOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestCancelOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("orders/%s/cancel.json", orderId)),
		httpmock.NewStringResponder(200, `{"order":{"id":"order-001"}}`))

	apiReq := &order.CancelOrderAPIReq{Id: orderId}
	apiResp := &order.CancelOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCancelOrder_MissingId(t *testing.T) {
	err := (&order.CancelOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestOrderArchive(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("PUT", orderURL(cli, fmt.Sprintf("orders/%s/hide_mark.json", orderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.OrderArchiveAPIReq{Id: orderId}
	apiResp := &order.OrderArchiveAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestOrderArchive_MissingId(t *testing.T) {
	err := (&order.OrderArchiveAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestCancelOrderArchive(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("PUT", orderURL(cli, fmt.Sprintf("orders/%s/cancel_hide_mark.json", orderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.CancelOrderArchiveAPIReq{Id: orderId}
	apiResp := &order.CancelOrderArchiveAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCancelOrderArchive_MissingId(t *testing.T) {
	err := (&order.CancelOrderArchiveAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestGetOrderPayment(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderSeq := "SEQ-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("orders/%s/transactions.json", orderSeq)),
		httpmock.NewStringResponder(200, `{"transactions":[{"id":"txn-001"}]}`))

	apiReq := &order.GetOrderPaymentAPIReq{OrderSeq: orderSeq}
	apiResp := &order.GetOrderPaymentAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestGetOrderPayment_MissingOrderSeq(t *testing.T) {
	err := (&order.GetOrderPaymentAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderSeq is required")
}

func TestQueryAllRefundOrdersForTheIdAssociation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("orders/%s/refunds.json", orderId)),
		httpmock.NewStringResponder(200, `{"refunds":[{"id":"refund-001"}]}`))

	apiReq := &order.QueryAllRefundOrdersForTheIdAssociationAPIReq{OrderId: orderId}
	apiResp := &order.QueryAllRefundOrdersForTheIdAssociationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryAllRefundOrdersForTheIdAssociation_MissingOrderId(t *testing.T) {
	err := (&order.QueryAllRefundOrdersForTheIdAssociationAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestQueryOrderSpecifyOfYourMoneyBackOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	refundId := "refund-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("orders/%s/refunds/%s.json", orderId, refundId)),
		httpmock.NewStringResponder(200, `{"refund":{"id":"refund-001"}}`))

	apiReq := &order.QueryOrderSpecifyOfYourMoneyBackOrderAPIReq{OrderId: orderId, RefundId: refundId}
	apiResp := &order.QueryOrderSpecifyOfYourMoneyBackOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryOrderSpecifyOfYourMoneyBackOrder_MissingOrderId(t *testing.T) {
	err := (&order.QueryOrderSpecifyOfYourMoneyBackOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestQueryOrderSpecifyOfYourMoneyBackOrder_MissingRefundId(t *testing.T) {
	err := (&order.QueryOrderSpecifyOfYourMoneyBackOrderAPIReq{OrderId: "order-001"}).Verify()
	assert.EqualError(t, err, "RefundId is required")
}

func TestQueryNumberOfFulfillment(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("orders/%s/fulfillments/count.json", orderId)),
		httpmock.NewStringResponder(200, `{"count":3}`))

	apiReq := &order.QueryNumberOfFulfillmentAPIReq{OrderId: orderId}
	apiResp := &order.QueryNumberOfFulfillmentAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryNumberOfFulfillment_MissingOrderId(t *testing.T) {
	err := (&order.QueryNumberOfFulfillmentAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestQueryAllFulfillmentsForAnOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("orders/%s/fulfillments.json", orderId)),
		httpmock.NewStringResponder(200, `{"fulfillments":[{"id":"ful-001"}]}`))

	apiReq := &order.QueryAllFulfillmentsForAnOrderAPIReq{OrderId: orderId}
	apiResp := &order.QueryAllFulfillmentsForAnOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryAllFulfillmentsForAnOrder_MissingOrderId(t *testing.T) {
	err := (&order.QueryAllFulfillmentsForAnOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Order Refund
// ══════════════════════════════════════════════════════════════════════════════

func TestOrderRefund(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "order/refund.json"),
		httpmock.NewStringResponder(200, `{"refund":{"id":"refund-001"}}`))

	apiReq := &order.OrderRefundAPIReq{OrderId: "order-001"}
	apiResp := &order.OrderRefundAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestOrderRefund_MissingOrderId(t *testing.T) {
	err := (&order.OrderRefundAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestOrderRefundTrial(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("orders/%s/refunds/calculate.json", orderId)),
		httpmock.NewStringResponder(200, `{"refund":{"id":"refund-001"}}`))

	apiReq := &order.OrderRefundTrialAPIReq{OrderId: orderId}
	apiResp := &order.OrderRefundTrialAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestOrderRefundTrial_MissingOrderId(t *testing.T) {
	err := (&order.OrderRefundTrialAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Draft Orders
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateADraftOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "orders/draft_orders.json"),
		httpmock.NewStringResponder(200, `{"draft_order":{"id":"draft-001"}}`))

	apiReq := &order.CreateADraftOrderAPIReq{}
	apiResp := &order.CreateADraftOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryDraftOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "orders/draft_orders.json"),
		httpmock.NewStringResponder(200, `{"draft_orders":[{"id":"draft-001"}]}`))

	apiReq := &order.QueryDraftOrderAPIReq{}
	apiResp := &order.QueryDraftOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDraftOrderQuantityQuery(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "orders/draft_orders/count.json"),
		httpmock.NewStringResponder(200, `{"count":2}`))

	apiReq := &order.DraftOrderQuantityQueryAPIReq{}
	apiResp := &order.DraftOrderQuantityQueryAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestEditDraftOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("PUT", orderURL(cli, "orders/draft_orders.json"),
		httpmock.NewStringResponder(200, `{"draft_order":{"id":"draft-001"}}`))

	apiReq := &order.EditDraftOrderAPIReq{}
	apiResp := &order.EditDraftOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteDraftOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	draftOrderId := "draft-001"
	httpmock.RegisterResponder("DELETE", orderURL(cli, fmt.Sprintf("orders/draft_orders/%s.json", draftOrderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.DeleteDraftOrderAPIReq{DraftOrderId: draftOrderId}
	apiResp := &order.DeleteDraftOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteDraftOrder_MissingDraftOrderId(t *testing.T) {
	err := (&order.DeleteDraftOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "DraftOrderId is required")
}

func TestCompleteDraftOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	draftId := "draft-001"
	httpmock.RegisterResponder("PUT", orderURL(cli, fmt.Sprintf("orders/draft_orders/%s/complete.json", draftId)),
		httpmock.NewStringResponder(200, `{"draft_order":{"id":"draft-001"}}`))

	apiReq := &order.CompleteDraftOrderAPIReq{Id: draftId}
	apiResp := &order.CompleteDraftOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCompleteDraftOrder_MissingId(t *testing.T) {
	err := (&order.CompleteDraftOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestSendADraftOrderInvoiceEmail(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	draftId := "draft-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("orders/draft_orders/%s/send_invoice.json", draftId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.SendADraftOrderInvoiceEmailAPIReq{Id: draftId, To: "test@example.com"}
	apiResp := &order.SendADraftOrderInvoiceEmailAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestSendADraftOrderInvoiceEmail_MissingId(t *testing.T) {
	err := (&order.SendADraftOrderInvoiceEmailAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestSendADraftOrderInvoiceEmail_MissingTo(t *testing.T) {
	err := (&order.SendADraftOrderInvoiceEmailAPIReq{Id: "draft-001"}).Verify()
	assert.EqualError(t, err, "To is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Fulfillments
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateAFulfillmentBasedOnAnOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("orders/%s/fulfillments.json", orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment":{"id":"ful-001"}}`))

	apiReq := &order.CreateAFulfillmentBasedOnAnOrderAPIReq{OrderId: orderId}
	apiResp := &order.CreateAFulfillmentBasedOnAnOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCreateAFulfillmentBasedOnAnOrder_MissingOrderId(t *testing.T) {
	err := (&order.CreateAFulfillmentBasedOnAnOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestCreateAnFulfillmentForTheFulfillmentOrderSingleOrBatch(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "fulfillments/fulfillments.json"),
		httpmock.NewStringResponder(200, `{"fulfillment":{"id":"ful-001"}}`))

	apiReq := &order.CreateAnFulfillmentForTheFulfillmentOrderSingleOrBatchAPIReq{OrderId: "order-001"}
	apiResp := &order.CreateAnFulfillmentForTheFulfillmentOrderSingleOrBatchAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCreateAnFulfillmentForTheFulfillmentOrderSingleOrBatch_MissingOrderId(t *testing.T) {
	err := (&order.CreateAnFulfillmentForTheFulfillmentOrderSingleOrBatchAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestGetSpecifyFulfillmentUnderFulfillmentOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("orders/%s/%s/fulfillments.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment":{"id":"ful-001"}}`))

	apiReq := &order.GetSpecifyFulfillmentUnderFulfillmentOrderAPIReq{FulfillmentId: fulfillmentId, OrderId: orderId}
	apiResp := &order.GetSpecifyFulfillmentUnderFulfillmentOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestGetSpecifyFulfillmentUnderFulfillmentOrder_MissingFulfillmentId(t *testing.T) {
	err := (&order.GetSpecifyFulfillmentUnderFulfillmentOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestGetSpecifyFulfillmentUnderFulfillmentOrder_MissingOrderId(t *testing.T) {
	err := (&order.GetSpecifyFulfillmentUnderFulfillmentOrderAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestCancelFulfillment(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillments/%s/%s/cancel.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment":{"id":"ful-001"}}`))

	apiReq := &order.CancelFulfillmentAPIReq{FulfillmentId: fulfillmentId, OrderId: orderId}
	apiResp := &order.CancelFulfillmentAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCancelFulfillment_MissingFulfillmentId(t *testing.T) {
	err := (&order.CancelFulfillmentAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestCancelFulfillment_MissingOrderId(t *testing.T) {
	err := (&order.CancelFulfillmentAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestCancelShipping(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("orders/%s/fulfillments/%s/cancel.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment":{"id":"ful-001"}}`))

	apiReq := &order.CancelShippingAPIReq{FulfillmentId: fulfillmentId, OrderId: orderId}
	apiResp := &order.CancelShippingAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCancelShipping_MissingFulfillmentId(t *testing.T) {
	err := (&order.CancelShippingAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestCancelShipping_MissingOrderId(t *testing.T) {
	err := (&order.CancelShippingAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestUpdateTrackingNumber(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("orders/%s/fulfillments/%s/update_tracking.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment":{"id":"ful-001"}}`))

	apiReq := &order.UpdateTrackingNumberAPIReq{FulfillmentId: fulfillmentId, OrderId: orderId}
	apiResp := &order.UpdateTrackingNumberAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestUpdateTrackingNumber_MissingFulfillmentId(t *testing.T) {
	err := (&order.UpdateTrackingNumberAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestUpdateTrackingNumber_MissingOrderId(t *testing.T) {
	err := (&order.UpdateTrackingNumberAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestUpdateTrackingInformation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillments/%s/%s/update_tracking.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment":{"id":"ful-001"}}`))

	apiReq := &order.UpdateTrackingInformationAPIReq{FulfillmentId: fulfillmentId, OrderId: orderId}
	apiResp := &order.UpdateTrackingInformationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestUpdateTrackingInformation_MissingFulfillmentId(t *testing.T) {
	err := (&order.UpdateTrackingInformationAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestUpdateTrackingInformation_MissingOrderId(t *testing.T) {
	err := (&order.UpdateTrackingInformationAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestCreateFulfillmentShippingLogisticsEvent(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillments/%s/%s/events/create.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment_event":{"id":"event-001"}}`))

	apiReq := &order.CreateFulfillmentShippingLogisticsEventAPIReq{
		FulfillmentId: fulfillmentId,
		OrderId:       orderId,
		Status:        "in_transit",
	}
	apiResp := &order.CreateFulfillmentShippingLogisticsEventAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCreateFulfillmentShippingLogisticsEvent_MissingFulfillmentId(t *testing.T) {
	err := (&order.CreateFulfillmentShippingLogisticsEventAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestCreateFulfillmentShippingLogisticsEvent_MissingOrderId(t *testing.T) {
	err := (&order.CreateFulfillmentShippingLogisticsEventAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestCreateFulfillmentShippingLogisticsEvent_MissingStatus(t *testing.T) {
	err := (&order.CreateFulfillmentShippingLogisticsEventAPIReq{FulfillmentId: "ful-001", OrderId: "order-001"}).Verify()
	assert.EqualError(t, err, "Status is required")
}

func TestQueryFulfillmentAndDeliveryLogisticsEvents(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("fulfillments/%s/%s/events/search.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment_events":[{"id":"event-001"}]}`))

	apiReq := &order.QueryFulfillmentAndDeliveryLogisticsEventsAPIReq{FulfillmentId: fulfillmentId, OrderId: orderId}
	apiResp := &order.QueryFulfillmentAndDeliveryLogisticsEventsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryFulfillmentAndDeliveryLogisticsEvents_MissingFulfillmentId(t *testing.T) {
	err := (&order.QueryFulfillmentAndDeliveryLogisticsEventsAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestQueryFulfillmentAndDeliveryLogisticsEvents_MissingOrderId(t *testing.T) {
	err := (&order.QueryFulfillmentAndDeliveryLogisticsEventsAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestDeleteFulfillmentDeliveryLogisticsEvent(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	eventId := "event-001"
	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("DELETE", orderURL(cli, fmt.Sprintf("fulfillments/%s/%s/events/%s.json", eventId, fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.DeleteFulfillmentDeliveryLogisticsEventAPIReq{
		EventId:       eventId,
		FulfillmentId: fulfillmentId,
		OrderId:       orderId,
	}
	apiResp := &order.DeleteFulfillmentDeliveryLogisticsEventAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteFulfillmentDeliveryLogisticsEvent_MissingEventId(t *testing.T) {
	err := (&order.DeleteFulfillmentDeliveryLogisticsEventAPIReq{}).Verify()
	assert.EqualError(t, err, "EventId is required")
}

func TestDeleteFulfillmentDeliveryLogisticsEvent_MissingFulfillmentId(t *testing.T) {
	err := (&order.DeleteFulfillmentDeliveryLogisticsEventAPIReq{EventId: "event-001"}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestDeleteFulfillmentDeliveryLogisticsEvent_MissingOrderId(t *testing.T) {
	err := (&order.DeleteFulfillmentDeliveryLogisticsEventAPIReq{EventId: "event-001", FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Fulfillment App
// ══════════════════════════════════════════════════════════════════════════════

func TestFulfillmentAppAcceptsCancellationRequestForShipment(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillments/%s/%s/cancellation_request_accept.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.FulfillmentAppAcceptsCancellationRequestForShipmentAPIReq{FulfillmentId: fulfillmentId, OrderId: orderId}
	apiResp := &order.FulfillmentAppAcceptsCancellationRequestForShipmentAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestFulfillmentAppAcceptsCancellationRequestForShipment_MissingFulfillmentId(t *testing.T) {
	err := (&order.FulfillmentAppAcceptsCancellationRequestForShipmentAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestFulfillmentAppAcceptsCancellationRequestForShipment_MissingOrderId(t *testing.T) {
	err := (&order.FulfillmentAppAcceptsCancellationRequestForShipmentAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestFulfillmentAppAcceptsRequestToShip(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillments/%s/%s/request_accept.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.FulfillmentAppAcceptsRequestToShipAPIReq{FulfillmentId: fulfillmentId, OrderId: orderId}
	apiResp := &order.FulfillmentAppAcceptsRequestToShipAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestFulfillmentAppAcceptsRequestToShip_MissingFulfillmentId(t *testing.T) {
	err := (&order.FulfillmentAppAcceptsRequestToShipAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestFulfillmentAppAcceptsRequestToShip_MissingOrderId(t *testing.T) {
	err := (&order.FulfillmentAppAcceptsRequestToShipAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestFulfillmentAppRejectsShipmentCancellationRequest(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillments/%s/%s/cancellation_request_reject.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.FulfillmentAppRejectsShipmentCancellationRequestAPIReq{FulfillmentId: fulfillmentId, OrderId: orderId}
	apiResp := &order.FulfillmentAppRejectsShipmentCancellationRequestAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestFulfillmentAppRejectsShipmentCancellationRequest_MissingFulfillmentId(t *testing.T) {
	err := (&order.FulfillmentAppRejectsShipmentCancellationRequestAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestFulfillmentAppRejectsShipmentCancellationRequest_MissingOrderId(t *testing.T) {
	err := (&order.FulfillmentAppRejectsShipmentCancellationRequestAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestFulfillmentAppRejectsShipmentRequest(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillments/%s/%s/request_reject.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.FulfillmentAppRejectsShipmentRequestAPIReq{
		FulfillmentId: fulfillmentId,
		OrderId:       orderId,
		Reason:        "incorrect_address",
	}
	apiResp := &order.FulfillmentAppRejectsShipmentRequestAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestFulfillmentAppRejectsShipmentRequest_MissingFulfillmentId(t *testing.T) {
	err := (&order.FulfillmentAppRejectsShipmentRequestAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestFulfillmentAppRejectsShipmentRequest_MissingOrderId(t *testing.T) {
	err := (&order.FulfillmentAppRejectsShipmentRequestAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestFulfillmentAppRejectsShipmentRequest_MissingReason(t *testing.T) {
	err := (&order.FulfillmentAppRejectsShipmentRequestAPIReq{FulfillmentId: "ful-001", OrderId: "order-001"}).Verify()
	assert.EqualError(t, err, "Reason is required")
}

func TestFulfillmentAppRequestsShipmentCancellation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillments/%s/%s/request_cancellation.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.FulfillmentAppRequestsShipmentCancellationAPIReq{FulfillmentId: fulfillmentId, OrderId: orderId}
	apiResp := &order.FulfillmentAppRequestsShipmentCancellationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestFulfillmentAppRequestsShipmentCancellation_MissingFulfillmentId(t *testing.T) {
	err := (&order.FulfillmentAppRequestsShipmentCancellationAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestFulfillmentAppRequestsShipmentCancellation_MissingOrderId(t *testing.T) {
	err := (&order.FulfillmentAppRequestsShipmentCancellationAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestShipAccordingToFulfillmentOrderRequest(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentId := "ful-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillments/%s/%s/request.json", fulfillmentId, orderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.ShipAccordingToFulfillmentOrderRequestAPIReq{FulfillmentId: fulfillmentId, OrderId: orderId}
	apiResp := &order.ShipAccordingToFulfillmentOrderRequestAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestShipAccordingToFulfillmentOrderRequest_MissingFulfillmentId(t *testing.T) {
	err := (&order.ShipAccordingToFulfillmentOrderRequestAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentId is required")
}

func TestShipAccordingToFulfillmentOrderRequest_MissingOrderId(t *testing.T) {
	err := (&order.ShipAccordingToFulfillmentOrderRequestAPIReq{FulfillmentId: "ful-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Fulfillment Orders
// ══════════════════════════════════════════════════════════════════════════════

func TestFulfillmentOrderListQuery(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "fulfillment_orders/fulfillment_orders_search.json"),
		httpmock.NewStringResponder(200, `{"fulfillment_orders":[{"id":"fo-001"}]}`))

	apiReq := &order.FulfillmentOrderListQueryAPIReq{}
	apiResp := &order.FulfillmentOrderListQueryAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestGetAllFulfillmentOrders(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("fulfillment_orders/%s/fulfillment_orders.json", orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment_orders":[{"id":"fo-001"}]}`))

	apiReq := &order.GetAllFulfillmentOrdersAPIReq{OrderId: orderId}
	apiResp := &order.GetAllFulfillmentOrdersAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestGetAllFulfillmentOrders_MissingOrderId(t *testing.T) {
	err := (&order.GetAllFulfillmentOrdersAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestGetAllFulfillmentsForTheFulfillmentOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentOrderId := "fo-001"
	orderId := "order-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("fulfillment_orders/%s/%s/fulfillments.json", fulfillmentOrderId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillments":[{"id":"ful-001"}]}`))

	apiReq := &order.GetAllFulfillmentsForTheFulfillmentOrderAPIReq{FulfillmentOrderId: fulfillmentOrderId, OrderId: orderId}
	apiResp := &order.GetAllFulfillmentsForTheFulfillmentOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestGetAllFulfillmentsForTheFulfillmentOrder_MissingFulfillmentOrderId(t *testing.T) {
	err := (&order.GetAllFulfillmentsForTheFulfillmentOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentOrderId is required")
}

func TestGetAllFulfillmentsForTheFulfillmentOrder_MissingOrderId(t *testing.T) {
	err := (&order.GetAllFulfillmentsForTheFulfillmentOrderAPIReq{FulfillmentOrderId: "fo-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestGetSpecifyFulfillmentOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentOrderId := "fo-001"
	orderId := "order-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("fulfillment_orders/%s/%s.json", fulfillmentOrderId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment_order":{"id":"fo-001"}}`))

	apiReq := &order.GetSpecifyFulfillmentOrderAPIReq{FulfillmentOrderId: fulfillmentOrderId, OrderId: orderId}
	apiResp := &order.GetSpecifyFulfillmentOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestGetSpecifyFulfillmentOrder_MissingFulfillmentOrderId(t *testing.T) {
	err := (&order.GetSpecifyFulfillmentOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentOrderId is required")
}

func TestGetSpecifyFulfillmentOrder_MissingOrderId(t *testing.T) {
	err := (&order.GetSpecifyFulfillmentOrderAPIReq{FulfillmentOrderId: "fo-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestFulfillmentOrderReleaseOnholdStatus(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentOrderId := "fo-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillment_orders/%s/%s/release_hold.json", fulfillmentOrderId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment_order":{"id":"fo-001"}}`))

	apiReq := &order.FulfillmentOrderReleaseOnholdStatusAPIReq{FulfillmentOrderId: fulfillmentOrderId, OrderId: orderId}
	apiResp := &order.FulfillmentOrderReleaseOnholdStatusAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestFulfillmentOrderReleaseOnholdStatus_MissingFulfillmentOrderId(t *testing.T) {
	err := (&order.FulfillmentOrderReleaseOnholdStatusAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentOrderId is required")
}

func TestFulfillmentOrderReleaseOnholdStatus_MissingOrderId(t *testing.T) {
	err := (&order.FulfillmentOrderReleaseOnholdStatusAPIReq{FulfillmentOrderId: "fo-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestPerformanceAppointmentMarkerOnhold(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentOrderId := "fo-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillment_orders/%s/%s/hold.json", fulfillmentOrderId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment_hold":{"id":"hold-001"}}`))

	apiReq := &order.PerformanceAppointmentMarkerOnholdAPIReq{FulfillmentOrderId: fulfillmentOrderId, OrderId: orderId}
	apiResp := &order.PerformanceAppointmentMarkerOnholdAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestPerformanceAppointmentMarkerOnhold_MissingFulfillmentOrderId(t *testing.T) {
	err := (&order.PerformanceAppointmentMarkerOnholdAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentOrderId is required")
}

func TestPerformanceAppointmentMarkerOnhold_MissingOrderId(t *testing.T) {
	err := (&order.PerformanceAppointmentMarkerOnholdAPIReq{FulfillmentOrderId: "fo-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestReBookTheDateOfTheSale(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentOrderId := "fo-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillment_orders/%s/%s/reschedule.json", fulfillmentOrderId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment_order":{"id":"fo-001"}}`))

	apiReq := &order.ReBookTheDateOfTheSaleAPIReq{FulfillmentOrderId: fulfillmentOrderId, OrderId: orderId}
	apiResp := &order.ReBookTheDateOfTheSaleAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestReBookTheDateOfTheSale_MissingFulfillmentOrderId(t *testing.T) {
	err := (&order.ReBookTheDateOfTheSaleAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentOrderId is required")
}

func TestReBookTheDateOfTheSale_MissingOrderId(t *testing.T) {
	err := (&order.ReBookTheDateOfTheSaleAPIReq{FulfillmentOrderId: "fo-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestSetUpAFulfillmentOrderDeadline(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillment_orders/%s/set_fulfillment_orders_deadline.json", orderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.SetUpAFulfillmentOrderDeadlineAPIReq{
		OrderId:             orderId,
		FulfillmentDeadline: "2026-05-01T10:00:00+08:00",
	}
	apiResp := &order.SetUpAFulfillmentOrderDeadlineAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestSetUpAFulfillmentOrderDeadline_MissingOrderId(t *testing.T) {
	err := (&order.SetUpAFulfillmentOrderDeadlineAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestSetUpAFulfillmentOrderDeadline_MissingFulfillmentDeadline(t *testing.T) {
	err := (&order.SetUpAFulfillmentOrderDeadlineAPIReq{OrderId: "order-001"}).Verify()
	assert.EqualError(t, err, "FulfillmentDeadline is required")
}

func TestUpdateInventoryLocationForFulfillmentOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	fulfillmentOrderId := "fo-001"
	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("fulfillment_orders/%s/%s/move.json", fulfillmentOrderId, orderId)),
		httpmock.NewStringResponder(200, `{"fulfillment_order":{"id":"fo-001"}}`))

	apiReq := &order.UpdateInventoryLocationForFulfillmentOrderAPIReq{FulfillmentOrderId: fulfillmentOrderId, OrderId: orderId}
	apiResp := &order.UpdateInventoryLocationForFulfillmentOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestUpdateInventoryLocationForFulfillmentOrder_MissingFulfillmentOrderId(t *testing.T) {
	err := (&order.UpdateInventoryLocationForFulfillmentOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "FulfillmentOrderId is required")
}

func TestUpdateInventoryLocationForFulfillmentOrder_MissingOrderId(t *testing.T) {
	err := (&order.UpdateInventoryLocationForFulfillmentOrderAPIReq{FulfillmentOrderId: "fo-001"}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Fulfillment Services
// ══════════════════════════════════════════════════════════════════════════════

func TestFulfillmentServiceList(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "fulfillment_service.json"),
		httpmock.NewStringResponder(200, `{"fulfillment_services":[{"id":"fs-001","name":"My Service"}]}`))

	apiReq := &order.FulfillmentServiceListAPIReq{}
	apiResp := &order.FulfillmentServiceListAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestFulfillmentServiceDetails(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	serviceId := "fs-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("fulfillment_service/%s.json", serviceId)),
		httpmock.NewStringResponder(200, `{"fulfillment_service":{"id":"fs-001","name":"My Service"}}`))

	apiReq := &order.FulfillmentServiceDetailsAPIReq{Id: serviceId}
	apiResp := &order.FulfillmentServiceDetailsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestFulfillmentServiceDetails_MissingId(t *testing.T) {
	err := (&order.FulfillmentServiceDetailsAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestCreateFulfillmentService(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "fulfillment_service.json"),
		httpmock.NewStringResponder(200, `{"fulfillment_service":{"id":"fs-001","name":"My Service"}}`))

	apiReq := &order.CreateFulfillmentServiceAPIReq{Name: "My Service"}
	apiResp := &order.CreateFulfillmentServiceAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCreateFulfillmentService_MissingName(t *testing.T) {
	err := (&order.CreateFulfillmentServiceAPIReq{}).Verify()
	assert.EqualError(t, err, "Name is required")
}

func TestModifyFulfillmentService(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("PUT", orderURL(cli, "fulfillment_service.json"),
		httpmock.NewStringResponder(200, `{"fulfillment_service":{"id":"fs-001"}}`))

	apiReq := &order.ModifyFulfillmentServiceAPIReq{Id: "fs-001"}
	apiResp := &order.ModifyFulfillmentServiceAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestModifyFulfillmentService_MissingId(t *testing.T) {
	err := (&order.ModifyFulfillmentServiceAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestDeleteFulfillmentService(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	serviceId := "fs-001"
	httpmock.RegisterResponder("DELETE", orderURL(cli, fmt.Sprintf("fulfillment_service/%s.json", serviceId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.DeleteFulfillmentServiceAPIReq{Id: serviceId}
	apiResp := &order.DeleteFulfillmentServiceAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteFulfillmentService_MissingId(t *testing.T) {
	err := (&order.DeleteFulfillmentServiceAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestCreateFulfillmentLocation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "fulfillment_service_location/create.json"),
		httpmock.NewStringResponder(200, `{"location":{"id":"loc-001"}}`))

	apiReq := &order.CreateFulfillmentLocationAPIReq{}
	apiResp := &order.CreateFulfillmentLocationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

// ══════════════════════════════════════════════════════════════════════════════
// Carrier Services
// ══════════════════════════════════════════════════════════════════════════════

func TestReceiveAListOfAllCarrierServices(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "carrier_services.json"),
		httpmock.NewStringResponder(200, `{"carrier_services":[{"id":"cs-001","name":"My Carrier"}]}`))

	apiReq := &order.ReceiveAListOfAllCarrierServicesAPIReq{}
	apiResp := &order.ReceiveAListOfAllCarrierServicesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestReceiveASingleCarrierService(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	carrierId := "cs-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("carrier_services/%s.json", carrierId)),
		httpmock.NewStringResponder(200, `{"carrier_service":{"id":"cs-001","name":"My Carrier"}}`))

	apiReq := &order.ReceiveASingleCarrierServiceAPIReq{CarrierServiceId: carrierId}
	apiResp := &order.ReceiveASingleCarrierServiceAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestReceiveASingleCarrierService_MissingCarrierServiceId(t *testing.T) {
	err := (&order.ReceiveASingleCarrierServiceAPIReq{}).Verify()
	assert.EqualError(t, err, "CarrierServiceId is required")
}

func TestCreateANewCarrierService(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "carrier_services.json"),
		httpmock.NewStringResponder(200, `{"carrier_service":{"id":"cs-001","name":"My Carrier"}}`))

	apiReq := &order.CreateANewCarrierServiceAPIReq{
		CallbackUrl: "https://example.com/callback",
		Name:        "My Carrier",
	}
	apiResp := &order.CreateANewCarrierServiceAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCreateANewCarrierService_MissingCallbackUrl(t *testing.T) {
	err := (&order.CreateANewCarrierServiceAPIReq{}).Verify()
	assert.EqualError(t, err, "CallbackUrl is required")
}

func TestCreateANewCarrierService_MissingName(t *testing.T) {
	err := (&order.CreateANewCarrierServiceAPIReq{CallbackUrl: "https://example.com/callback"}).Verify()
	assert.EqualError(t, err, "Name is required")
}

func TestModifyAnExistingCarrierService(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	carrierId := "cs-001"
	httpmock.RegisterResponder("PUT", orderURL(cli, fmt.Sprintf("carrier_services/%s.json", carrierId)),
		httpmock.NewStringResponder(200, `{"carrier_service":{"id":"cs-001"}}`))

	apiReq := &order.ModifyAnExistingCarrierServiceAPIReq{CarrierServiceId: carrierId}
	apiResp := &order.ModifyAnExistingCarrierServiceAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestModifyAnExistingCarrierService_MissingCarrierServiceId(t *testing.T) {
	err := (&order.ModifyAnExistingCarrierServiceAPIReq{}).Verify()
	assert.EqualError(t, err, "CarrierServiceId is required")
}

func TestRemoveAnExistingCarrierService(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	carrierId := "cs-001"
	httpmock.RegisterResponder("DELETE", orderURL(cli, fmt.Sprintf("carrier_services/%s.json", carrierId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.RemoveAnExistingCarrierServiceAPIReq{CarrierServiceId: carrierId}
	apiResp := &order.RemoveAnExistingCarrierServiceAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestRemoveAnExistingCarrierService_MissingCarrierServiceId(t *testing.T) {
	err := (&order.RemoveAnExistingCarrierServiceAPIReq{}).Verify()
	assert.EqualError(t, err, "CarrierServiceId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Abandoned Checkouts
// ══════════════════════════════════════════════════════════════════════════════

func TestGetAbandonedCheckouts(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "orders/abandoned_orders.json"),
		httpmock.NewStringResponder(200, `{"abandoned_orders":[{"id":"ao-001"}]}`))

	apiReq := &order.GetAbandonedCheckoutsAPIReq{}
	apiResp := &order.GetAbandonedCheckoutsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestRetrievesACountOfCheckouts(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "orders/abandoned_orders/count.json"),
		httpmock.NewStringResponder(200, `{"count":7}`))

	apiReq := &order.RetrievesACountOfCheckoutsAPIReq{}
	apiResp := &order.RetrievesACountOfCheckoutsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestArchiveAndUnarchiveAbandonedCheckouts(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "orders/abandoned_order/hide_mark.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.ArchiveAndUnarchiveAbandonedCheckoutsAPIReq{HideMark: "hide"}
	apiResp := &order.ArchiveAndUnarchiveAbandonedCheckoutsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestArchiveAndUnarchiveAbandonedCheckouts_MissingHideMark(t *testing.T) {
	err := (&order.ArchiveAndUnarchiveAbandonedCheckoutsAPIReq{}).Verify()
	assert.EqualError(t, err, "HideMark is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Risk Fraud
// ══════════════════════════════════════════════════════════════════════════════

func TestGetFraudRisksForAnOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("orders/v2/%s/risks.json", orderId)),
		httpmock.NewStringResponder(200, `{"risks":[{"id":"risk-001"}]}`))

	apiReq := &order.GetFraudRisksForAnOrderAPIReq{OrderId: orderId}
	apiResp := &order.GetFraudRisksForAnOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestGetFraudRisksForAnOrder_MissingOrderId(t *testing.T) {
	err := (&order.GetFraudRisksForAnOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestQueryRiskFraudForOrderSpecify(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	riskId := "risk-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("orders/v2/%s/risks/%s.json", orderId, riskId)),
		httpmock.NewStringResponder(200, `{"risk":{"id":"risk-001"}}`))

	apiReq := &order.QueryRiskFraudForOrderSpecifyAPIReq{OrderId: orderId, RiskId: riskId}
	apiResp := &order.QueryRiskFraudForOrderSpecifyAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryRiskFraudForOrderSpecify_MissingOrderId(t *testing.T) {
	err := (&order.QueryRiskFraudForOrderSpecifyAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestQueryRiskFraudForOrderSpecify_MissingRiskId(t *testing.T) {
	err := (&order.QueryRiskFraudForOrderSpecifyAPIReq{OrderId: "order-001"}).Verify()
	assert.EqualError(t, err, "RiskId is required")
}

func TestAddRiskFraud(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("orders/v2/%s/risks.json", orderId)),
		httpmock.NewStringResponder(200, `{"risk":{"id":"risk-001"}}`))

	apiReq := &order.AddRiskFraudAPIReq{OrderId: orderId}
	apiResp := &order.AddRiskFraudAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestAddRiskFraud_MissingOrderId(t *testing.T) {
	err := (&order.AddRiskFraudAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestUpdateOrderRiskFraud(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	riskId := "risk-001"
	httpmock.RegisterResponder("PUT", orderURL(cli, fmt.Sprintf("orders/v2/%s/risks/%s.json", orderId, riskId)),
		httpmock.NewStringResponder(200, `{"risk":{"id":"risk-001"}}`))

	apiReq := &order.UpdateOrderRiskFraudAPIReq{OrderId: orderId, RiskId: riskId}
	apiResp := &order.UpdateOrderRiskFraudAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestUpdateOrderRiskFraud_MissingOrderId(t *testing.T) {
	err := (&order.UpdateOrderRiskFraudAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestUpdateOrderRiskFraud_MissingRiskId(t *testing.T) {
	err := (&order.UpdateOrderRiskFraudAPIReq{OrderId: "order-001"}).Verify()
	assert.EqualError(t, err, "RiskId is required")
}

func TestDeleteAllRiskFraudRelatedToTheOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("DELETE", orderURL(cli, fmt.Sprintf("orders/v2/%s/risks.json", orderId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.DeleteAllRiskFraudRelatedToTheOrderAPIReq{OrderId: orderId}
	apiResp := &order.DeleteAllRiskFraudRelatedToTheOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteAllRiskFraudRelatedToTheOrder_MissingOrderId(t *testing.T) {
	err := (&order.DeleteAllRiskFraudRelatedToTheOrderAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestRemoveTheSpecifyOfTheOrderRiskFraud(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	riskId := "risk-001"
	httpmock.RegisterResponder("DELETE", orderURL(cli, fmt.Sprintf("orders/v2/%s/risks/%s.json", orderId, riskId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.RemoveTheSpecifyOfTheOrderRiskFraudAPIReq{OrderId: orderId, RiskId: riskId}
	apiResp := &order.RemoveTheSpecifyOfTheOrderRiskFraudAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestRemoveTheSpecifyOfTheOrderRiskFraud_MissingOrderId(t *testing.T) {
	err := (&order.RemoveTheSpecifyOfTheOrderRiskFraudAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestRemoveTheSpecifyOfTheOrderRiskFraud_MissingRiskId(t *testing.T) {
	err := (&order.RemoveTheSpecifyOfTheOrderRiskFraudAPIReq{OrderId: "order-001"}).Verify()
	assert.EqualError(t, err, "RiskId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Payments
// ══════════════════════════════════════════════════════════════════════════════

func TestCreatePaymentSlip(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "orders/transactions.json"),
		httpmock.NewStringResponder(200, `{"transaction":{"id":"txn-001"}}`))

	apiReq := &order.CreatePaymentSlipAPIReq{}
	apiResp := &order.CreatePaymentSlipAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestUpdatePaymentSlip(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "orders/update_transactions.json"),
		httpmock.NewStringResponder(200, `{"transaction":{"id":"txn-001"}}`))

	apiReq := &order.UpdatePaymentSlipAPIReq{}
	apiResp := &order.UpdatePaymentSlipAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestGetOrderTransactions(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "transactions/query.json"),
		httpmock.NewStringResponder(200, `{"transactions":[{"id":"txn-001"}]}`))

	apiReq := &order.GetOrderTransactionsAPIReq{OrderId: "order-001"}
	apiResp := &order.GetOrderTransactionsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestGetOrderTransactions_MissingOrderId(t *testing.T) {
	err := (&order.GetOrderTransactionsAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderId is required")
}

func TestQueryStorePaymentChannels(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "orders/pay_channels.json"),
		httpmock.NewStringResponder(200, `{"pay_channels":[{"id":"pc-001"}]}`))

	apiReq := &order.QueryStorePaymentChannelsAPIReq{}
	apiResp := &order.QueryStorePaymentChannelsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryPaymentSettings(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "checkouts/payment_settings.json"),
		httpmock.NewStringResponder(200, `{"payment_settings":[{"id":"ps-001"}]}`))

	apiReq := &order.QueryPaymentSettingsAPIReq{}
	apiResp := &order.QueryPaymentSettingsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

// ══════════════════════════════════════════════════════════════════════════════
// Order Edit
// ══════════════════════════════════════════════════════════════════════════════

func TestStartEditing(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	orderId := "order-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("orders/%s/order_edit_begin.json", orderId)),
		httpmock.NewStringResponder(200, `{"calculated_order":{"id":"calc-001"}}`))

	apiReq := &order.StartEditingAPIReq{Id: orderId}
	apiResp := &order.StartEditingAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestStartEditing_MissingId(t *testing.T) {
	err := (&order.StartEditingAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestAddACustomProduct(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "orders/order_edit_add_customized_item.json"),
		httpmock.NewStringResponder(200, `{"calculated_order":{"id":"calc-001"}}`))

	apiReq := &order.AddACustomProductAPIReq{}
	apiResp := &order.AddACustomProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestAddItemsInOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "orders/order_edit_add_variant.json"),
		httpmock.NewStringResponder(200, `{"calculated_order":{"id":"calc-001"}}`))

	apiReq := &order.AddItemsInOrderAPIReq{}
	apiResp := &order.AddItemsInOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestAddProductDiscount(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "orders/order_edit_add_line_item_discount.json"),
		httpmock.NewStringResponder(200, `{"calculated_order":{"id":"calc-001"}}`))

	apiReq := &order.AddProductDiscountAPIReq{}
	apiResp := &order.AddProductDiscountAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteProductDiscount(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "orders/order_edit_remove_line_item_discount.json"),
		httpmock.NewStringResponder(200, `{"calculated_order":{"id":"calc-001"}}`))

	apiReq := &order.DeleteProductDiscountAPIReq{}
	apiResp := &order.DeleteProductDiscountAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestSetItemQuantity(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "orders/order_edit_set_quantity.json"),
		httpmock.NewStringResponder(200, `{"calculated_order":{"id":"calc-001"}}`))

	apiReq := &order.SetItemQuantityAPIReq{}
	apiResp := &order.SetItemQuantityAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestSubmitEditResults(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "orders/order_edit_commit.json"),
		httpmock.NewStringResponder(200, `{"order":{"id":"order-001"}}`))

	apiReq := &order.SubmitEditResultsAPIReq{}
	apiResp := &order.SubmitEditResultsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

// ══════════════════════════════════════════════════════════════════════════════
// Order Attribution
// ══════════════════════════════════════════════════════════════════════════════

func TestGetOrderAttributionInBulk(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "orders/order_attribution_info.json"),
		httpmock.NewStringResponder(200, `{"data":[{"order_seq":"SEQ-001"}]}`))

	apiReq := &order.GetOrderAttributionInBulkAPIReq{Orders: []string{"SEQ-001"}}
	apiResp := &order.GetOrderAttributionInBulkAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Data, 1)
	assert.Equal(t, "SEQ-001", apiResp.Data[0].OrderSeq)
}

// ══════════════════════════════════════════════════════════════════════════════
// Tax / Countries
// ══════════════════════════════════════════════════════════════════════════════

func TestQueryAllTaxInformationInTheStore(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "countries.json"),
		httpmock.NewStringResponder(200, `{"countries":[{"id":"country-001","name":"United States","code":"US"}]}`))

	apiReq := &order.QueryAllTaxInformationInTheStoreAPIReq{}
	apiResp := &order.QueryAllTaxInformationInTheStoreAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryCountryNumberOfTaxAlreadyConfiguredInStores(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "countries/count.json"),
		httpmock.NewStringResponder(200, `{"count":10}`))

	apiReq := &order.QueryCountryNumberOfTaxAlreadyConfiguredInStoresAPIReq{}
	apiResp := &order.QueryCountryNumberOfTaxAlreadyConfiguredInStoresAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQuerySpecifyTaxInformationFromTheCountry(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	countryId := "country-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("countries/%s.json", countryId)),
		httpmock.NewStringResponder(200, `{"country":{"id":"country-001","name":"United States","code":"US"}}`))

	apiReq := &order.QuerySpecifyTaxInformationFromTheCountryAPIReq{Id: countryId}
	apiResp := &order.QuerySpecifyTaxInformationFromTheCountryAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQuerySpecifyTaxInformationFromTheCountry_MissingId(t *testing.T) {
	err := (&order.QuerySpecifyTaxInformationFromTheCountryAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestQueryQuantityOfAllStateProvincesInACountry(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	countryId := "country-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("provinces/count/%s.json", countryId)),
		httpmock.NewStringResponder(200, `{"count":50}`))

	apiReq := &order.QueryQuantityOfAllStateProvincesInACountryAPIReq{Id: countryId}
	apiResp := &order.QueryQuantityOfAllStateProvincesInACountryAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryQuantityOfAllStateProvincesInACountry_MissingId(t *testing.T) {
	err := (&order.QueryQuantityOfAllStateProvincesInACountryAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestQuerySpecifyAllStateProvinceInformationUnderTheCountry(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	countryId := "country-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("provinces/%s.json", countryId)),
		httpmock.NewStringResponder(200, `{"provinces":[{"id":"province-001","name":"California","code":"CA"}]}`))

	apiReq := &order.QuerySpecifyAllStateProvinceInformationUnderTheCountryAPIReq{Id: countryId}
	apiResp := &order.QuerySpecifyAllStateProvinceInformationUnderTheCountryAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQuerySpecifyAllStateProvinceInformationUnderTheCountry_MissingId(t *testing.T) {
	err := (&order.QuerySpecifyAllStateProvinceInformationUnderTheCountryAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestQueryStateProvinceInformation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	countryId := "country-001"
	provinceId := "province-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("provinces/%s/countries/%s.json", countryId, provinceId)),
		httpmock.NewStringResponder(200, `{"province":{"id":"province-001","name":"California","code":"CA"}}`))

	apiReq := &order.QueryStateProvinceInformationAPIReq{Id: countryId, ProvinceId: provinceId}
	apiResp := &order.QueryStateProvinceInformationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryStateProvinceInformation_MissingId(t *testing.T) {
	err := (&order.QueryStateProvinceInformationAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestQueryStateProvinceInformation_MissingProvinceId(t *testing.T) {
	err := (&order.QueryStateProvinceInformationAPIReq{Id: "country-001"}).Verify()
	assert.EqualError(t, err, "ProvinceId is required")
}

func TestQueryTaxChannel(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "store_tax_service.json"),
		httpmock.NewStringResponder(200, `{"tax_services":[{"id":"ts-001"}]}`))

	apiReq := &order.QueryTaxChannelAPIReq{CountryCode: "US"}
	apiResp := &order.QueryTaxChannelAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryTaxChannel_MissingCountryCode(t *testing.T) {
	err := (&order.QueryTaxChannelAPIReq{}).Verify()
	assert.EqualError(t, err, "CountryCode is required")
}

func TestUpdateTaxChannel(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "store_tax_service/update.json"),
		httpmock.NewStringResponder(200, `{"tax_service":{"id":"ts-001"}}`))

	apiReq := &order.UpdateTaxChannelAPIReq{}
	apiResp := &order.UpdateTaxChannelAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteTaxChannel(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	countryCode := "US"
	httpmock.RegisterResponder("DELETE", orderURL(cli, fmt.Sprintf("store_tax_service/delete/%s.json", countryCode)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.DeleteTaxChannelAPIReq{CountryCode: countryCode}
	apiResp := &order.DeleteTaxChannelAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteTaxChannel_MissingCountryCode(t *testing.T) {
	err := (&order.DeleteTaxChannelAPIReq{}).Verify()
	assert.EqualError(t, err, "CountryCode is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Inventory / Locations / Shipping
// ══════════════════════════════════════════════════════════════════════════════

func TestQueryInventoryAddress(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "locations.json"),
		httpmock.NewStringResponder(200, `{"locations":[{"id":"loc-001","name":"Warehouse A"}]}`))

	apiReq := &order.QueryInventoryAddressAPIReq{}
	apiResp := &order.QueryInventoryAddressAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryAllLocalShippingInformation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "local_delivery/list.json"),
		httpmock.NewStringResponder(200, `{"local_deliveries":[{"id":"ld-001"}]}`))

	apiReq := &order.QueryAllLocalShippingInformationAPIReq{}
	apiResp := &order.QueryAllLocalShippingInformationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryAllPickupInformation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "pickup/list.json"),
		httpmock.NewStringResponder(200, `{"pickups":[{"id":"pickup-001"}]}`))

	apiReq := &order.QueryAllPickupInformationAPIReq{}
	apiResp := &order.QueryAllPickupInformationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestBatchShippingPlan(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "logistics/shipping_zones.json"),
		httpmock.NewStringResponder(200, `{"shipping_zones":[{"id":"sz-001"}]}`))

	apiReq := &order.BatchShippingPlanAPIReq{}
	apiResp := &order.BatchShippingPlanAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestBatchShippingPlanV2(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "logistics/shipping_zones_v2.json"),
		httpmock.NewStringResponder(200, `{"shipping_zones":[{"id":"sz-001"}]}`))

	apiReq := &order.BatchShippingPlanV2APIReq{}
	apiResp := &order.BatchShippingPlanV2APIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

// ══════════════════════════════════════════════════════════════════════════════
// Returns
// ══════════════════════════════════════════════════════════════════════════════

func TestQueryReturns(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "return.json"),
		httpmock.NewStringResponder(200, `{"returns":[{"id":"return-001"}]}`))

	apiReq := &order.QueryReturnsAPIReq{}
	apiResp := &order.QueryReturnsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCreateAReturn(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "return/create.json"),
		httpmock.NewStringResponder(200, `{"return":{"id":"return-001"}}`))

	apiReq := &order.CreateAReturnAPIReq{OrderSeq: "SEQ-001"}
	apiResp := &order.CreateAReturnAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCreateAReturn_MissingOrderSeq(t *testing.T) {
	err := (&order.CreateAReturnAPIReq{}).Verify()
	assert.EqualError(t, err, "OrderSeq is required")
}

func TestCloseReturn(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	returnId := "return-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("return/%s/close.json", returnId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.CloseReturnAPIReq{Id: returnId}
	apiResp := &order.CloseReturnAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCloseReturn_MissingId(t *testing.T) {
	err := (&order.CloseReturnAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestQueryReturnFulfillmentOrder(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "return_fulfillment.json"),
		httpmock.NewStringResponder(200, `{"return_fulfillment_orders":[{"id":"rfo-001"}]}`))

	apiReq := &order.QueryReturnFulfillmentOrderAPIReq{}
	apiResp := &order.QueryReturnFulfillmentOrderAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryReturnFulfillment(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", orderURL(cli, "return_package.json"),
		httpmock.NewStringResponder(200, `{"return_packages":[{"id":"rp-001"}]}`))

	apiReq := &order.QueryReturnFulfillmentAPIReq{ReturnId: "return-001"}
	apiResp := &order.QueryReturnFulfillmentAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQueryReturnFulfillment_MissingReturnId(t *testing.T) {
	err := (&order.QueryReturnFulfillmentAPIReq{}).Verify()
	assert.EqualError(t, err, "ReturnId is required")
}

func TestCreateAReturnFulfillment(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "return/package/create.json"),
		httpmock.NewStringResponder(200, `{"return_package":{"id":"rp-001"}}`))

	apiReq := &order.CreateAReturnFulfillmentAPIReq{}
	apiResp := &order.CreateAReturnFulfillmentAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestUpdateLogisticsInformationForAReturnFulfillment(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "return_express_update.json"),
		httpmock.NewStringResponder(200, `{"return_package":{"id":"rp-001"}}`))

	apiReq := &order.UpdateLogisticsInformationForAReturnFulfillmentAPIReq{}
	apiResp := &order.UpdateLogisticsInformationForAReturnFulfillmentAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

// ══════════════════════════════════════════════════════════════════════════════
// Subscription
// ══════════════════════════════════════════════════════════════════════════════

func TestQuerySubscriptionContractList(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", orderURL(cli, "subscription/contract.json"),
		httpmock.NewStringResponder(200, `{"contracts":[{"id":"contract-001"}]}`))

	apiReq := &order.QuerySubscriptionContractListAPIReq{}
	apiResp := &order.QuerySubscriptionContractListAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQuerySubscriptionContractDetails(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	contractId := "contract-001"
	httpmock.RegisterResponder("GET", orderURL(cli, fmt.Sprintf("subscription/%s/contract.json", contractId)),
		httpmock.NewStringResponder(200, `{"contract":{"id":"contract-001"}}`))

	apiReq := &order.QuerySubscriptionContractDetailsAPIReq{Id: contractId}
	apiResp := &order.QuerySubscriptionContractDetailsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestQuerySubscriptionContractDetails_MissingId(t *testing.T) {
	err := (&order.QuerySubscriptionContractDetailsAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestEditSubscriptionContractInformation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	contractId := "contract-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("subscription/%s/edit_subscription.json", contractId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.EditSubscriptionContractInformationAPIReq{Id: contractId}
	apiResp := &order.EditSubscriptionContractInformationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestEditSubscriptionContractInformation_MissingId(t *testing.T) {
	err := (&order.EditSubscriptionContractInformationAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestCancelSubscriptionContract(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	contractId := "contract-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("subscription/%s/cancel.json", contractId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.CancelSubscriptionContractAPIReq{Id: contractId}
	apiResp := &order.CancelSubscriptionContractAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCancelSubscriptionContract_MissingId(t *testing.T) {
	err := (&order.CancelSubscriptionContractAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestCreateYourOrderAsSoonAsYourSubscriptionContractIsCreated(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	contractId := "contract-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("subscription/%s/create_order.json", contractId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.CreateYourOrderAsSoonAsYourSubscriptionContractIsCreatedAPIReq{Id: contractId}
	apiResp := &order.CreateYourOrderAsSoonAsYourSubscriptionContractIsCreatedAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCreateYourOrderAsSoonAsYourSubscriptionContractIsCreated_MissingId(t *testing.T) {
	err := (&order.CreateYourOrderAsSoonAsYourSubscriptionContractIsCreatedAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestReviseNextBillTime(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	contractId := "contract-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("subscription/%s/edit_next_billing_date.json", contractId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.ReviseNextBillTimeAPIReq{Id: contractId}
	apiResp := &order.ReviseNextBillTimeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestReviseNextBillTime_MissingId(t *testing.T) {
	err := (&order.ReviseNextBillTimeAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestSkipTheNextBill(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	contractId := "contract-001"
	httpmock.RegisterResponder("POST", orderURL(cli, fmt.Sprintf("subscription/%s/skip_next_bill.json", contractId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &order.SkipTheNextBillAPIReq{Id: contractId}
	apiResp := &order.SkipTheNextBillAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestSkipTheNextBill_MissingId(t *testing.T) {
	err := (&order.SkipTheNextBillAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}
