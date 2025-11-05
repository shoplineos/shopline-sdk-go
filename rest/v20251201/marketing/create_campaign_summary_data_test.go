package marketing

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"testing"
)

func TestCreateCampaignSummaryDataAPIReq(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(client.MethodPost,
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/sales/promotion_marketing.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	apiReq := &CreateCampaignSummaryDataAPIReq{
		SalesAmount:     100,
		TradeOrderCount: 1,
		BenefitAmount:   5,
	}

	apiResp := &CreateCampaignSummaryDataAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)
	if err != nil {
		t.Errorf("DeleteOrder returned an error %v", err)
	}

}
