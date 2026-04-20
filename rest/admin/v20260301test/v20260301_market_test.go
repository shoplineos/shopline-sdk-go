package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/market"
	"github.com/stretchr/testify/assert"
)

func marketURL(cli *client.Client, path string) string {
	return fmt.Sprintf("https://%s.myshopline.com/%s/%s/%s",
		cli.StoreHandle, cli.PathPrefix, cli.ApiVersion, path)
}

func TestQueryMarketList(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"markets":[{"market_id":"2805657769119556579334","market_name":"US-MARKET","enabled":true,"primary":true,"regions":[{"code":"US","currency":{"currency_code":"USD"}}]}]}`
	httpmock.RegisterResponder("GET", marketURL(cli, "markets.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &market.QueryMarketListAPIReq{}
	apiResp := &market.QueryMarketListAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Markets, 1)
	assert.Equal(t, "2805657769119556579334", apiResp.Markets[0].MarketId)
	assert.Equal(t, "US-MARKET", apiResp.Markets[0].MarketName)
	assert.True(t, apiResp.Markets[0].Enabled)
	assert.True(t, apiResp.Markets[0].Primary)
}

func TestReturnsAMarketResourceById(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	marketId := "2805657769119556579334"
	mockResp := `{"market":{"market_id":"2805657769119556579334","market_name":"US-MARKET","enabled":true,"primary":true,"currency_settings":{"base_currency":{"currency_code":"USD","currency_name":"USD","auto_exchange_rate":true},"local_currencies":true}}}`
	httpmock.RegisterResponder("GET", marketURL(cli, fmt.Sprintf("markets/%s.json", marketId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &market.ReturnsAMarketResourceByIdAPIReq{Id: marketId}
	apiResp := &market.ReturnsAMarketResourceByIdAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "2805657769119556579334", apiResp.Market.MarketId)
	assert.Equal(t, "US-MARKET", apiResp.Market.MarketName)
	assert.Equal(t, "USD", apiResp.Market.CurrencySettings.BaseCurrency.CurrencyCode)
}

func TestReturnsAMarketResourceById_MissingId(t *testing.T) {
	err := (&market.ReturnsAMarketResourceByIdAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestAddProductsToTheMarket(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	marketsId := "2805657769119556579334"
	mockResp := `{"product_ids":["16055431542830326121993183"]}`
	httpmock.RegisterResponder("POST", marketURL(cli, fmt.Sprintf("markets/%s/published_products.json", marketsId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &market.AddProductsToTheMarketAPIReq{
		MarketsId:  marketsId,
		ProductIds: []string{"16055431542830326121993183"},
	}
	apiResp := &market.AddProductsToTheMarketAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, []string{"16055431542830326121993183"}, apiResp.ProductIds)
}

func TestAddProductsToTheMarket_MissingMarketsId(t *testing.T) {
	err := (&market.AddProductsToTheMarketAPIReq{}).Verify()
	assert.EqualError(t, err, "MarketsId is required")
}

func TestGetMarketsPublishedProducts(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	marketsId := "2805657769119556579335"
	mockResp := `{"products_count":2,"product_ids":["16054492295522822738072519","16054492295522822738072520"]}`
	httpmock.RegisterResponder("GET", marketURL(cli, fmt.Sprintf("markets/%s/products.json", marketsId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &market.GetMarketsPublishedProductsAPIReq{MarketsId: marketsId}
	apiResp := &market.GetMarketsPublishedProductsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 2, apiResp.ProductsCount)
	assert.Len(t, apiResp.ProductIds, 2)
}

func TestGetMarketsPublishedProducts_MissingMarketsId(t *testing.T) {
	err := (&market.GetMarketsPublishedProductsAPIReq{}).Verify()
	assert.EqualError(t, err, "MarketsId is required")
}

func TestRemoveMarketProducts(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	marketsId := "2805657769119556579335"
	mockResp := `{"product_ids":["16054492295522822738072519"]}`
	httpmock.RegisterResponder("POST", marketURL(cli, fmt.Sprintf("markets/%s/unpublished_products.json", marketsId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &market.RemoveMarketProductsAPIReq{
		MarketsId:  marketsId,
		ProductIds: []string{"16054492295522822738072519"},
	}
	apiResp := &market.RemoveMarketProductsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, []string{"16054492295522822738072519"}, apiResp.ProductIds)
}

func TestRemoveMarketProducts_MissingMarketsId(t *testing.T) {
	err := (&market.RemoveMarketProductsAPIReq{}).Verify()
	assert.EqualError(t, err, "MarketsId is required")
}

func TestMarketCurrencyUpdate(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	marketId := "2805657769119556579334"
	httpmock.RegisterResponder("PUT", marketURL(cli, fmt.Sprintf("markets/%s/currency_settings.json", marketId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &market.MarketCurrencyUpdateAPIReq{
		Id:           marketId,
		ExchangeRate: "1.01234567892",
	}
	apiResp := &market.MarketCurrencyUpdateAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestMarketCurrencyUpdate_MissingId(t *testing.T) {
	err := (&market.MarketCurrencyUpdateAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestMarketCurrencyUpdate_MissingExchangeRate(t *testing.T) {
	err := (&market.MarketCurrencyUpdateAPIReq{Id: "2805657769119556579334"}).Verify()
	assert.EqualError(t, err, "ExchangeRate is required")
}

func TestQueryTheMarketPriceOfTheVariant(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	variantId := "18060895302863070249743144"
	mockResp := `{"market_prices":[{"country_code":"CN","market_price":{"amount":"10","currency_code":"CNY"}},{"country_code":"US","market_price":{"amount":"1.5","currency_code":"USD"}}]}`
	httpmock.RegisterResponder("GET", marketURL(cli, fmt.Sprintf("markets/%s/price.json", variantId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &market.QueryTheMarketPriceOfTheVariantAPIReq{
		VariantId:    variantId,
		CountryCodes: "CN,US",
	}
	apiResp := &market.QueryTheMarketPriceOfTheVariantAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.MarketPrices, 2)
	assert.Equal(t, "CN", apiResp.MarketPrices[0].CountryCode)
	assert.Equal(t, "10", apiResp.MarketPrices[0].MarketPrice.Amount)
	assert.Equal(t, "CNY", apiResp.MarketPrices[0].MarketPrice.CurrencyCode)
}

func TestQueryTheMarketPriceOfTheVariant_MissingVariantId(t *testing.T) {
	err := (&market.QueryTheMarketPriceOfTheVariantAPIReq{}).Verify()
	assert.EqualError(t, err, "VariantId is required")
}

func TestQueryTheMarketPriceOfTheVariant_MissingCountryCodes(t *testing.T) {
	err := (&market.QueryTheMarketPriceOfTheVariantAPIReq{VariantId: "18060895302863070249743144"}).Verify()
	assert.EqualError(t, err, "CountryCodes is required")
}

func TestQueryTheMarketPriceRangesOfTheProduct(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16060895302856191591063144"
	mockResp := `{"market_prices":[{"country_code":"CN","min_market_price":{"amount":"5","currency_code":"CNY"},"max_market_price":{"amount":"10","currency_code":"CNY"}}]}`
	httpmock.RegisterResponder("GET", marketURL(cli, fmt.Sprintf("markets/%s/price/range.json", productId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &market.QueryTheMarketPriceRangesOfTheProductAPIReq{
		ProductId:    productId,
		CountryCodes: "CN",
	}
	apiResp := &market.QueryTheMarketPriceRangesOfTheProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.MarketPrices, 1)
	assert.Equal(t, "CN", apiResp.MarketPrices[0].CountryCode)
	assert.Equal(t, "5", apiResp.MarketPrices[0].MinMarketPrice.Amount)
	assert.Equal(t, "10", apiResp.MarketPrices[0].MaxMarketPrice.Amount)
}

func TestQueryTheMarketPriceRangesOfTheProduct_MissingProductId(t *testing.T) {
	err := (&market.QueryTheMarketPriceRangesOfTheProductAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

func TestQueryTheMarketPriceRangesOfTheProduct_MissingCountryCodes(t *testing.T) {
	err := (&market.QueryTheMarketPriceRangesOfTheProductAPIReq{ProductId: "16060895302856191591063144"}).Verify()
	assert.EqualError(t, err, "CountryCodes is required")
}

func TestQueryTheMarketPublishedOfTheProduct(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16060895302856191591063144"
	mockResp := `{"markets":[{"country_code":"US","published":true},{"country_code":"CN","published":false}]}`
	httpmock.RegisterResponder("GET", marketURL(cli, fmt.Sprintf("markets/%s/published.json", productId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &market.QueryTheMarketPublishedOfTheProductAPIReq{
		ProductId:    productId,
		CountryCodes: "US,CN",
	}
	apiResp := &market.QueryTheMarketPublishedOfTheProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Markets, 2)
	assert.Equal(t, "US", apiResp.Markets[0].CountryCode)
	assert.True(t, apiResp.Markets[0].Published)
	assert.Equal(t, "CN", apiResp.Markets[1].CountryCode)
	assert.False(t, apiResp.Markets[1].Published)
}

func TestQueryTheMarketPublishedOfTheProduct_MissingProductId(t *testing.T) {
	err := (&market.QueryTheMarketPublishedOfTheProductAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

func TestQueryTheMarketPublishedOfTheProduct_MissingCountryCodes(t *testing.T) {
	err := (&market.QueryTheMarketPublishedOfTheProductAPIReq{ProductId: "16060895302856191591063144"}).Verify()
	assert.EqualError(t, err, "CountryCodes is required")
}
