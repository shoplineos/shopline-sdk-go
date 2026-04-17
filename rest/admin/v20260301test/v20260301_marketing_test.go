package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/test"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/marketing"
	"github.com/stretchr/testify/assert"
)

func marketingURL(cli *client.Client, path string) string {
	return fmt.Sprintf("https://%s.myshopline.com/%s/%s/%s",
		cli.StoreHandle, cli.PathPrefix, cli.ApiVersion, path)
}

// ══════════════════════════════════════════════════════════════════════════════
// Price Rules
// ══════════════════════════════════════════════════════════════════════════════

func TestRetrievePriceRulesList(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"price_rules":[{"id":"6018695669979551970","title":"Summer Sale","value":"-20","value_type":"fixed_amount"}]}`
	httpmock.RegisterResponder("GET", marketingURL(cli, "sales/price_rules.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.RetrievePriceRulesListAPIReq{}
	apiResp := &marketing.RetrievePriceRulesListAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.PriceRules, 1)
	assert.Equal(t, "Summer Sale", apiResp.PriceRules[0].Title)
}

func TestRetrieveAllPriceRulesCounts(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", marketingURL(cli, "sales/price_rules/count.json"),
		httpmock.NewStringResponder(200, `{"count":5}`))

	apiReq := &marketing.RetrieveAllPriceRulesCountsAPIReq{}
	apiResp := &marketing.RetrieveAllPriceRulesCountsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 5, apiResp.Count)
}

func TestGetAPriceRuleOfACodeDiscount(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	priceRuleId := "6018695669979551970"
	mockResp := `{"price_rule":{"id":"6018695669979551970","title":"Summer Sale","value":"-20","value_type":"fixed_amount"}}`
	httpmock.RegisterResponder("GET", marketingURL(cli, fmt.Sprintf("sales/price_rules/%s.json", priceRuleId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.GetAPriceRuleOfACodeDiscountAPIReq{PriceRuleId: priceRuleId}
	apiResp := &marketing.GetAPriceRuleOfACodeDiscountAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "6018695669979551970", apiResp.PriceRule.Id)
}

func TestGetAPriceRuleOfACodeDiscount_MissingPriceRuleId(t *testing.T) {
	err := (&marketing.GetAPriceRuleOfACodeDiscountAPIReq{}).Verify()
	assert.EqualError(t, err, "PriceRuleId is required")
}

func TestCreateAPriceRuleOfACodeDiscount(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"price_rule":{"id":"6018695669979551970","title":"New Rule","value":"-10","value_type":"percentage"}}`
	httpmock.RegisterResponder("POST", marketingURL(cli, "sales/price_rules.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.CreateAPriceRuleOfACodeDiscountAPIReq{}
	apiResp := &marketing.CreateAPriceRuleOfACodeDiscountAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "6018695669979551970", apiResp.PriceRule.Id)
}

func TestUpdateThePriceRuleOfACodeDiscount(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	priceRuleId := "6018695669979551970"
	mockResp := `{"price_rule":{"id":"6018695669979551970","title":"Updated Rule","value":"-15","value_type":"fixed_amount"}}`
	httpmock.RegisterResponder("PUT", marketingURL(cli, fmt.Sprintf("sales/price_rules/%s.json", priceRuleId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.UpdateThePriceRuleOfACodeDiscountAPIReq{PriceRuleId: priceRuleId}
	apiResp := &marketing.UpdateThePriceRuleOfACodeDiscountAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "6018695669979551970", apiResp.PriceRule.Id)
	assert.Equal(t, "Updated Rule", apiResp.PriceRule.Title)
}

func TestUpdateThePriceRuleOfACodeDiscount_MissingPriceRuleId(t *testing.T) {
	err := (&marketing.UpdateThePriceRuleOfACodeDiscountAPIReq{}).Verify()
	assert.EqualError(t, err, "PriceRuleId is required")
}

func TestDeleteExistingPriceRule(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	priceRuleId := "6018695669979551970"
	httpmock.RegisterResponder("DELETE", marketingURL(cli, fmt.Sprintf("sales/price_rules/%s.json", priceRuleId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &marketing.DeleteExistingPriceRuleAPIReq{PriceRuleId: priceRuleId}
	apiResp := &marketing.DeleteExistingPriceRuleAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteExistingPriceRule_MissingPriceRuleId(t *testing.T) {
	err := (&marketing.DeleteExistingPriceRuleAPIReq{}).Verify()
	assert.EqualError(t, err, "PriceRuleId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Discount Codes
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateDiscountCode(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	priceRuleId := "6018703930879383908"
	mockResp := `{"discount_code":{"id":"6019053959104174638","code":"WELCOME10","price_rule_id":"6018703930879383908","create_at":"2023-07-12T00:08:02+08:00"}}`
	httpmock.RegisterResponder("POST", marketingURL(cli, fmt.Sprintf("sales/price_rules/%s/discount_codes.json", priceRuleId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.CreateDiscountCodeAPIReq{PriceRuleId: priceRuleId}
	apiResp := &marketing.CreateDiscountCodeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "6019053959104174638", apiResp.DiscountCode.Id)
	assert.Equal(t, "WELCOME10", apiResp.DiscountCode.Code)
}

func TestCreateDiscountCode_MissingPriceRuleId(t *testing.T) {
	err := (&marketing.CreateDiscountCodeAPIReq{}).Verify()
	assert.EqualError(t, err, "PriceRuleId is required")
}

func TestGetADiscountCodeById(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	discountCodeId := "6018708477890794011"
	priceRuleId := "6018695669979551970"
	mockResp := `{"discount_code":{"id":"6018708477890794011","code":"WELCOME_CODE","price_rule_id":"6018695669979551970","create_at":"2023-07-11T18:24:49+08:00"}}`
	httpmock.RegisterResponder("GET", marketingURL(cli, fmt.Sprintf("sales/price_rules/%s/discount_codes/%s.json", discountCodeId, priceRuleId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.GetADiscountCodeByIdAPIReq{DiscountCodeId: discountCodeId, PriceRuleId: priceRuleId}
	apiResp := &marketing.GetADiscountCodeByIdAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "6018708477890794011", apiResp.DiscountCode.Id)
	assert.Equal(t, "WELCOME_CODE", apiResp.DiscountCode.Code)
}

func TestGetADiscountCodeById_MissingDiscountCodeId(t *testing.T) {
	err := (&marketing.GetADiscountCodeByIdAPIReq{}).Verify()
	assert.EqualError(t, err, "DiscountCodeId is required")
}

func TestGetADiscountCodeById_MissingPriceRuleId(t *testing.T) {
	err := (&marketing.GetADiscountCodeByIdAPIReq{DiscountCodeId: "6018708477890794011"}).Verify()
	assert.EqualError(t, err, "PriceRuleId is required")
}

func TestGetADiscountCodeByItsCode(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"discount_code":{"id":"6018708477890794011","code":"WELCOME_CODE","price_rule_id":"6018695669979551970","create_at":"2023-07-11T18:24:49+08:00"}}`
	httpmock.RegisterResponder("GET", marketingURL(cli, "sales/discount_codes/lookup.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.GetADiscountCodeByItsCodeAPIReq{Code: "WELCOME_CODE"}
	apiResp := &marketing.GetADiscountCodeByItsCodeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "6018708477890794011", apiResp.DiscountCode.Id)
	assert.Equal(t, "WELCOME_CODE", apiResp.DiscountCode.Code)
}

func TestUpdateDiscountCode(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	discountCodeId := "6018708477890794011"
	priceRuleId := "6018695669979551970"
	mockResp := `{"discount_code":{"id":"6018708477890794011","code":"NEWCODE","price_rule_id":"6018695669979551970","create_at":"2023-07-11T18:24:49+08:00"}}`
	httpmock.RegisterResponder("PUT", marketingURL(cli, fmt.Sprintf("sales/price_rules/%s/discount_codes/%s.json", discountCodeId, priceRuleId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.UpdateDiscountCodeAPIReq{DiscountCodeId: discountCodeId, PriceRuleId: priceRuleId}
	apiResp := &marketing.UpdateDiscountCodeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "NEWCODE", apiResp.DiscountCode.Code)
}

func TestUpdateDiscountCode_MissingDiscountCodeId(t *testing.T) {
	err := (&marketing.UpdateDiscountCodeAPIReq{}).Verify()
	assert.EqualError(t, err, "DiscountCodeId is required")
}

func TestUpdateDiscountCode_MissingPriceRuleId(t *testing.T) {
	err := (&marketing.UpdateDiscountCodeAPIReq{DiscountCodeId: "6018708477890794011"}).Verify()
	assert.EqualError(t, err, "PriceRuleId is required")
}

func TestDeleteDiscountCode(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	discountCodeId := "6018708477890794011"
	priceRuleId := "6018695669979551970"
	httpmock.RegisterResponder("DELETE", marketingURL(cli, fmt.Sprintf("sales/price_rules/%s/discount_codes/%s.json", discountCodeId, priceRuleId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &marketing.DeleteDiscountCodeAPIReq{DiscountCodeId: discountCodeId, PriceRuleId: priceRuleId}
	apiResp := &marketing.DeleteDiscountCodeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteDiscountCode_MissingDiscountCodeId(t *testing.T) {
	err := (&marketing.DeleteDiscountCodeAPIReq{}).Verify()
	assert.EqualError(t, err, "DiscountCodeId is required")
}

func TestDeleteDiscountCode_MissingPriceRuleId(t *testing.T) {
	err := (&marketing.DeleteDiscountCodeAPIReq{DiscountCodeId: "6018708477890794011"}).Verify()
	assert.EqualError(t, err, "PriceRuleId is required")
}

func TestQueryDiscountCodeByDiscountRulesId(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	priceRuleId := "6018695669979551970"
	mockResp := `{"discount_codes":[{"id":"6018708477890794011","code":"WELCOME_CODE","price_rule_id":"6018695669979551970"}]}`
	httpmock.RegisterResponder("GET", marketingURL(cli, fmt.Sprintf("sales/price_rules/%s/discount_codes.json", priceRuleId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.QueryDiscountCodeByDiscountRulesIdAPIReq{PriceRuleId: priceRuleId}
	apiResp := &marketing.QueryDiscountCodeByDiscountRulesIdAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.DiscountCodes, 1)
	assert.Equal(t, "6018708477890794011", apiResp.DiscountCodes[0].Id)
}

func TestQueryDiscountCodeByDiscountRulesId_MissingPriceRuleId(t *testing.T) {
	err := (&marketing.QueryDiscountCodeByDiscountRulesIdAPIReq{}).Verify()
	assert.EqualError(t, err, "PriceRuleId is required")
}

func TestQueryStoreSTotalNumberOfDiscountCode(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", marketingURL(cli, "sales/discount_codes/count.json"),
		httpmock.NewStringResponder(200, `{"count":10}`))

	apiReq := &marketing.QueryStoreSTotalNumberOfDiscountCodeAPIReq{}
	apiResp := &marketing.QueryStoreSTotalNumberOfDiscountCodeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 10, apiResp.Count)
}

func TestQueryDiscountCode(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"discount_codes":[{"id":"6018708477890794011","code":"WELCOME_CODE","price_rule_id":"6018695669979551970"}],"page_info":{"page_num":1,"page_size":20,"total":1}}`
	httpmock.RegisterResponder("GET", marketingURL(cli, "sales/discount/code"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.QueryDiscountCodeAPIReq{}
	apiResp := &marketing.QueryDiscountCodeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.DiscountCodes, 1)
	assert.Equal(t, 1, apiResp.PageInfo.Total)
}

// ══════════════════════════════════════════════════════════════════════════════
// Bulk Discount Code Tasks
// ══════════════════════════════════════════════════════════════════════════════

func TestCreatePromoCodeCreatedTasks(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	priceRuleId := "5682516987487987235"
	mockResp := `{"discount_code_creation":{"id":"5682520099208570423","status":"queued","price_rule_id":"5682516987487987235","codes_count":1,"failed_count":0,"created_at":"2022-11-21T20:11:40+08:00"}}`
	httpmock.RegisterResponder("POST", marketingURL(cli, fmt.Sprintf("sales/price_rules/%s/batch.json", priceRuleId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.CreatePromoCodeCreatedTasksAPIReq{PriceRuleId: priceRuleId}
	apiResp := &marketing.CreatePromoCodeCreatedTasksAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "5682520099208570423", apiResp.DiscountCodeCreation.Id)
	assert.Equal(t, "queued", apiResp.DiscountCodeCreation.Status)
}

func TestCreatePromoCodeCreatedTasks_MissingPriceRuleId(t *testing.T) {
	err := (&marketing.CreatePromoCodeCreatedTasksAPIReq{}).Verify()
	assert.EqualError(t, err, "PriceRuleId is required")
}

func TestQueryBulkCreatedDiscountCodeList(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	batchId := "6018676942227642851"
	priceRuleId := "6018642574989862445"
	mockResp := `{"discount_codes":[{"id":"6018708477890794011","code":"BATCH_CODE_001"}]}`
	httpmock.RegisterResponder("GET", marketingURL(cli, fmt.Sprintf("sales/price_rules/%s/batch/%s/discount_codes.json", batchId, priceRuleId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.QueryBulkCreatedDiscountCodeListAPIReq{BatchId: batchId, PriceRuleId: priceRuleId}
	apiResp := &marketing.QueryBulkCreatedDiscountCodeListAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.DiscountCodes, 1)
	assert.Equal(t, "BATCH_CODE_001", apiResp.DiscountCodes[0].Code)
}

func TestQueryBulkCreatedDiscountCodeList_MissingBatchId(t *testing.T) {
	err := (&marketing.QueryBulkCreatedDiscountCodeListAPIReq{}).Verify()
	assert.EqualError(t, err, "BatchId is required")
}

func TestQueryBulkCreatedDiscountCodeList_MissingPriceRuleId(t *testing.T) {
	err := (&marketing.QueryBulkCreatedDiscountCodeListAPIReq{BatchId: "6018676942227642851"}).Verify()
	assert.EqualError(t, err, "PriceRuleId is required")
}

func TestBasedOnMissionIdQueryPromoCodeInformation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	batchId := "6018676942227642851"
	priceRuleId := "6018642574989862445"
	mockResp := `{"discount_code_creation":{"id":"6018676942227642851","status":"completed","price_rule_id":"6018642574989862445","codes_count":1,"failed_count":0,"imported_count":1}}`
	httpmock.RegisterResponder("GET", marketingURL(cli, fmt.Sprintf("sales/price_rules/%s/batch/%s.json", batchId, priceRuleId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.BasedOnMissionIdQueryPromoCodeInformationAPIReq{BatchId: batchId, PriceRuleId: priceRuleId}
	apiResp := &marketing.BasedOnMissionIdQueryPromoCodeInformationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "6018676942227642851", apiResp.DiscountCodeCreation.Id)
	assert.Equal(t, "completed", apiResp.DiscountCodeCreation.Status)
}

func TestBasedOnMissionIdQueryPromoCodeInformation_MissingBatchId(t *testing.T) {
	err := (&marketing.BasedOnMissionIdQueryPromoCodeInformationAPIReq{}).Verify()
	assert.EqualError(t, err, "BatchId is required")
}

func TestBasedOnMissionIdQueryPromoCodeInformation_MissingPriceRuleId(t *testing.T) {
	err := (&marketing.BasedOnMissionIdQueryPromoCodeInformationAPIReq{BatchId: "6018676942227642851"}).Verify()
	assert.EqualError(t, err, "PriceRuleId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Automatic Discounts
// ══════════════════════════════════════════════════════════════════════════════

func TestQueryAutomaticDiscountActivity(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"automatic_discount":[{"starts_at":"2023-01-01T00:00:00+08:00","ends_at":"2023-12-31T00:00:00+08:00"}],"page_info":{"total":1,"page_num":1,"page_size":20}}`
	httpmock.RegisterResponder("GET", marketingURL(cli, "sales/discount/automatic"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.QueryAutomaticDiscountActivityAPIReq{}
	apiResp := &marketing.QueryAutomaticDiscountActivityAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.AutomaticDiscount, 1)
	assert.Equal(t, 1, apiResp.PageInfo.Total)
}

// ══════════════════════════════════════════════════════════════════════════════
// Marketing Events
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateMarketingEvents(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"marketing_event":{"id":"4073","marketing_event_id":"MA_6019943243021559089","event_type":"ad","marketing_channel":"search","utm_source":"manong.myshopline.com"}}`
	httpmock.RegisterResponder("POST", marketingURL(cli, "sales/marketing_events.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.CreateMarketingEventsAPIReq{}
	apiResp := &marketing.CreateMarketingEventsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "4073", apiResp.MarketingEvent.Id)
	assert.Equal(t, "MA_6019943243021559089", apiResp.MarketingEvent.MarketingEventId)
	assert.Equal(t, "ad", apiResp.MarketingEvent.EventType)
}

func TestQueryMarketingEventsList(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"marketing_events":[{"id":"4073","marketing_event_id":"MA_6062143268514178052","event_type":"ad","marketing_channel":"search"}]}`
	httpmock.RegisterResponder("GET", marketingURL(cli, "sales/marketing_events.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.QueryMarketingEventsListAPIReq{}
	apiResp := &marketing.QueryMarketingEventsListAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.MarketingEvents, 1)
	assert.Equal(t, "MA_6062143268514178052", apiResp.MarketingEvents[0].MarketingEventId)
}

func TestQueryIndividualMarketingEvents(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	marketingEventId := "4073"
	mockResp := `{"marketing_event":{"id":"4073","marketing_event_id":"MA_6062143268514178052","event_type":"ad","marketing_channel":"search"}}`
	httpmock.RegisterResponder("GET", marketingURL(cli, fmt.Sprintf("sales/marketing_events/%s.json", marketingEventId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.QueryIndividualMarketingEventsAPIReq{MarketingEventId: marketingEventId}
	apiResp := &marketing.QueryIndividualMarketingEventsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "4073", apiResp.MarketingEvent.Id)
}

func TestQueryIndividualMarketingEvents_MissingMarketingEventId(t *testing.T) {
	err := (&marketing.QueryIndividualMarketingEventsAPIReq{}).Verify()
	assert.EqualError(t, err, "MarketingEventId is required")
}

func TestUpdateMarketingEvents(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	marketingEventId := "4073"
	mockResp := `{"marketing_event":{"event_type":"ad","marketing_channel":"search","utm_source":"www.shopline.com","marketing_event_id":"MA_6019932248375043369"}}`
	httpmock.RegisterResponder("POST", marketingURL(cli, fmt.Sprintf("sales/marketing_events/%s.json", marketingEventId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.UpdateMarketingEventsAPIReq{MarketingEventId: marketingEventId}
	apiResp := &marketing.UpdateMarketingEventsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "ad", apiResp.MarketingEvent.EventType)
	assert.Equal(t, "search", apiResp.MarketingEvent.MarketingChannel)
}

func TestUpdateMarketingEvents_MissingMarketingEventId(t *testing.T) {
	err := (&marketing.UpdateMarketingEventsAPIReq{}).Verify()
	assert.EqualError(t, err, "MarketingEventId is required")
}

func TestDeleteMarketingEvents(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	marketingEventId := "4073"
	httpmock.RegisterResponder("DELETE", marketingURL(cli, fmt.Sprintf("sales/marketing_events/%s.json", marketingEventId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &marketing.DeleteMarketingEventsAPIReq{MarketingEventId: marketingEventId}
	apiResp := &marketing.DeleteMarketingEventsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteMarketingEvents_MissingMarketingEventId(t *testing.T) {
	err := (&marketing.DeleteMarketingEventsAPIReq{}).Verify()
	assert.EqualError(t, err, "MarketingEventId is required")
}

func TestQueryNumberOfMarketingEvents(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", marketingURL(cli, "sales/marketing_events/count.json"),
		httpmock.NewStringResponder(200, `{"count":7}`))

	apiReq := &marketing.QueryNumberOfMarketingEventsAPIReq{}
	apiResp := &marketing.QueryNumberOfMarketingEventsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 7, apiResp.Count)
}

func TestMassSyncMarketingEventStatistics(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", marketingURL(cli, "sales/marketing_events/sync_data.json"),
		httpmock.NewStringResponder(200, `{"sucess_count":3}`))

	apiReq := &marketing.MassSyncMarketingEventStatisticsAPIReq{}
	apiResp := &marketing.MassSyncMarketingEventStatisticsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 3, apiResp.SucessCount)
}

// ══════════════════════════════════════════════════════════════════════════════
// Marketing Expansion
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateMarketingExpansion(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"marketing_activity":{"marketing_activity_id":"MA_6019932248291157288","status":"draft","remote_id":"53e7f97e-8df5-4a6b-b672-5cda96b90821","utm_campaign":"google"}}`
	httpmock.RegisterResponder("POST", marketingURL(cli, "sales/marketing_activity_extension.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.CreateMarketingExpansionAPIReq{}
	apiResp := &marketing.CreateMarketingExpansionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "MA_6019932248291157288", apiResp.MarketingActivity.MarketingActivityId)
	assert.Equal(t, "draft", apiResp.MarketingActivity.Status)
}

func TestUpdateMarketingExpansion(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	marketingActivityId := "MA_5797003750150774786"
	mockResp := `{"marketing_activity":{"marketing_activity_id":"MA_5797003750150774786","status":"draft","target_status":"active","remote_id":"123"}}`
	httpmock.RegisterResponder("POST", marketingURL(cli, fmt.Sprintf("sales/marketing_activity_extension/%s.json", marketingActivityId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &marketing.UpdateMarketingExpansionAPIReq{MarketingActivityId: marketingActivityId}
	apiResp := &marketing.UpdateMarketingExpansionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "123", apiResp.MarketingActivity.RemoteId)
	assert.Equal(t, "active", apiResp.MarketingActivity.TargetStatus)
}

func TestUpdateMarketingExpansion_MissingMarketingActivityId(t *testing.T) {
	err := (&marketing.UpdateMarketingExpansionAPIReq{}).Verify()
	assert.EqualError(t, err, "MarketingActivityId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Campaign Summary & Flow Statistics
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateCampaignSummaryData(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", marketingURL(cli, "sales/promotion_marketing.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &marketing.CreateCampaignSummaryDataAPIReq{}
	apiResp := &marketing.CreateCampaignSummaryDataAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestProcessStatisticsSynchronization(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("POST", marketingURL(cli, "sales/flow/batch_sync_flow_data.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &marketing.ProcessStatisticsSynchronizationAPIReq{
		AppKey:     "test_app_key",
		SourceType: "flow",
		StoreId:    "store-001",
	}
	apiResp := &marketing.ProcessStatisticsSynchronizationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestProcessStatisticsSynchronization_MissingAppKey(t *testing.T) {
	err := (&marketing.ProcessStatisticsSynchronizationAPIReq{}).Verify()
	assert.EqualError(t, err, "AppKey is required")
}

func TestProcessStatisticsSynchronization_MissingSourceType(t *testing.T) {
	err := (&marketing.ProcessStatisticsSynchronizationAPIReq{AppKey: "test_app_key"}).Verify()
	assert.EqualError(t, err, "SourceType is required")
}

func TestProcessStatisticsSynchronization_MissingStoreId(t *testing.T) {
	err := (&marketing.ProcessStatisticsSynchronizationAPIReq{AppKey: "test_app_key", SourceType: "flow"}).Verify()
	assert.EqualError(t, err, "StoreId is required")
}
