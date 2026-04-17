package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/test"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/app"
	"github.com/stretchr/testify/assert"
)

func appURL(cli *client.Client, path string) string {
	return fmt.Sprintf("https://%s.myshopline.com/%s/%s/%s",
		cli.StoreHandle, cli.PathPrefix, cli.ApiVersion, path)
}

// ══════════════════════════════════════════════════════════════════════════════
// Size Chart APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateASizeChart(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"size_chart_id":"13266288837293171311580520"}}`
	httpmock.RegisterResponder("POST", appURL(cli, "sizechart.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &app.CreateASizeChartAPIReq{}
	apiResp := &app.CreateASizeChartAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "13266288837293171311580520", apiResp.Data.SizeChartId)
}

func TestUpdateASizeChart(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"size_chart_id":"13266288837293171311580520"}}`
	httpmock.RegisterResponder("PUT", appURL(cli, "sizechart.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &app.UpdateASizeChartAPIReq{}
	apiResp := &app.UpdateASizeChartAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "13266288837293171311580520", apiResp.Data.SizeChartId)
}

func TestQueryProductSizeDataInBatch(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"total":1,"size_charts":[{"size_chart_id":"13266288837293171311580520","title":"My SizeChart","status":false}]}}`
	httpmock.RegisterResponder("GET", appURL(cli, "sizecharts.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &app.QueryProductSizeDataInBatchAPIReq{}
	apiResp := &app.QueryProductSizeDataInBatchAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), apiResp.Data.Total)
}

func TestBatchDeleteSizechart(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("DELETE", appURL(cli, "sizecharts.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &app.BatchDeleteSizechartAPIReq{}
	apiResp := &app.BatchDeleteSizechartAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

// ══════════════════════════════════════════════════════════════════════════════
// Variant Image APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestQueryImagesOfVariant(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	productId := "16049619743346320200470282"
	mockResp := `{"data":{"product_id":"16049619743346320200470282","variant_images":[{"variant_id":"18050375907221969070393099","urls":["https://example.com/img.png"]}]}}`
	httpmock.RegisterResponder("GET", appURL(cli, fmt.Sprintf("variant_images/%s.json", productId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &app.QueryImagesOfVariantAPIReq{ProductId: productId}
	apiResp := &app.QueryImagesOfVariantAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Data.VariantImages, 1)
	assert.Equal(t, "18050375907221969070393099", apiResp.Data.VariantImages[0].VariantId)
}

func TestQueryImagesOfVariant_MissingProductId(t *testing.T) {
	err := (&app.QueryImagesOfVariantAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

func TestBatchUpdateVariantImages(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	productId := "16049619743346320200470282"
	httpmock.RegisterResponder("PUT", appURL(cli, fmt.Sprintf("variant_images/%s.json", productId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &app.BatchUpdateVariantImagesAPIReq{ProductId: productId}
	apiResp := &app.BatchUpdateVariantImagesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestBatchUpdateVariantImages_MissingProductId(t *testing.T) {
	err := (&app.BatchUpdateVariantImagesAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// CDP / Buyer Behavior APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestReportBuyerBehaviorEvents(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"channel_id":"6689609588495885399","events":[{"id":"8725726607153221","event_code":"sales_order","one_id":"0155eb0038b6f96a3b6d5f977ad076fa"}]}}`
	httpmock.RegisterResponder("POST", appURL(cli, "cdp/track.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &app.ReportBuyerBehaviorEventsAPIReq{}
	apiResp := &app.ReportBuyerBehaviorEventsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "6689609588495885399", apiResp.Data.ChannelId)
}

func TestReportBuyerIdentity(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"one_id":"0155eb0038b6f96a3b6d5f977ad076fa","user_traits":{}}}`
	httpmock.RegisterResponder("POST", appURL(cli, "cdp/identify.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &app.ReportBuyerIdentityAPIReq{ChannelId: "6689609588495885399"}
	apiResp := &app.ReportBuyerIdentityAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "0155eb0038b6f96a3b6d5f977ad076fa", apiResp.Data.OneId)
}

func TestReportBuyerIdentity_MissingChannelId(t *testing.T) {
	err := (&app.ReportBuyerIdentityAPIReq{}).Verify()
	assert.EqualError(t, err, "ChannelId is required")
}
