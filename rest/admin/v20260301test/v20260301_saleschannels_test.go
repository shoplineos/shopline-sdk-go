package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/test"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/saleschannels"
	"github.com/stretchr/testify/assert"
)

func salesChannelsURL(cli *client.Client, path string) string {
	return fmt.Sprintf("https://%s.myshopline.com/%s/%s/%s",
		cli.StoreHandle, cli.PathPrefix, cli.ApiVersion, path)
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection Listings
// ══════════════════════════════════════════════════════════════════════════════

func TestQueryCollectionOfSalesChannels(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"collection_listings":[{"id":"12248721639068681442230282","title":"Summer Collection","collection_type":"smart","published_scope":"web"}]}`
	httpmock.RegisterResponder("GET", salesChannelsURL(cli, "collection_listings.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &saleschannels.QueryCollectionOfSalesChannelsAPIReq{}
	apiResp := &saleschannels.QueryCollectionOfSalesChannelsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.CollectionListings, 1)
	assert.Equal(t, "12248721639068681442230282", apiResp.CollectionListings[0].Id)
	assert.Equal(t, "Summer Collection", apiResp.CollectionListings[0].Title)
}

func TestQuerySpecifyCollectionInSalesChannels(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	collectionListingId := "12248721639068681442230283"
	mockResp := `{"collection_listing":{"id":"12248721639068681442230283","title":"My Collection","collection_type":"smart","published_scope":"web","handle":"my-collection"}}`
	httpmock.RegisterResponder("GET", salesChannelsURL(cli, fmt.Sprintf("collection_listings/%s.json", collectionListingId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &saleschannels.QuerySpecifyCollectionInSalesChannelsAPIReq{CollectionListingId: collectionListingId}
	apiResp := &saleschannels.QuerySpecifyCollectionInSalesChannelsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "12248721639068681442230283", apiResp.CollectionListing.Id)
	assert.Equal(t, "My Collection", apiResp.CollectionListing.Title)
}

func TestQuerySpecifyCollectionInSalesChannels_MissingCollectionListingId(t *testing.T) {
	err := (&saleschannels.QuerySpecifyCollectionInSalesChannelsAPIReq{}).Verify()
	assert.EqualError(t, err, "CollectionListingId is required")
}

func TestAddCollectionToSalesChannels(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	collectionListingId := "12248721639068681442230282"
	mockResp := `{"collection_listing":{"id":"12248721639068681442230282","title":"This is a collection title","collection_type":"smart","published_scope":"web"}}`
	httpmock.RegisterResponder("PUT", salesChannelsURL(cli, fmt.Sprintf("collection_listings/%s.json", collectionListingId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &saleschannels.AddCollectionToSalesChannelsAPIReq{CollectionListingId: collectionListingId}
	apiResp := &saleschannels.AddCollectionToSalesChannelsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "12248721639068681442230282", apiResp.CollectionListing.Id)
	assert.Equal(t, "This is a collection title", apiResp.CollectionListing.Title)
}

func TestAddCollectionToSalesChannels_MissingCollectionListingId(t *testing.T) {
	err := (&saleschannels.AddCollectionToSalesChannelsAPIReq{}).Verify()
	assert.EqualError(t, err, "CollectionListingId is required")
}

func TestRemoveCollectionFromSalesChannels(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	collectionListingId := "12248721639068681442230282"
	httpmock.RegisterResponder("DELETE", salesChannelsURL(cli, fmt.Sprintf("collection_listings/%s.json", collectionListingId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &saleschannels.RemoveCollectionFromSalesChannelsAPIReq{CollectionListingId: collectionListingId}
	apiResp := &saleschannels.RemoveCollectionFromSalesChannelsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestRemoveCollectionFromSalesChannels_MissingCollectionListingId(t *testing.T) {
	err := (&saleschannels.RemoveCollectionFromSalesChannelsAPIReq{}).Verify()
	assert.EqualError(t, err, "CollectionListingId is required")
}

func TestQuerySalesChannelsSpecifyCollectionProductId(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	collectionListingId := "12248721639068681442230282"
	mockResp := `{"product_ids":["16057039432335097907370283","16057039432335097907380282"]}`
	httpmock.RegisterResponder("GET", salesChannelsURL(cli, fmt.Sprintf("collection_listings/%s/product_ids.json", collectionListingId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &saleschannels.QuerySalesChannelsSpecifyCollectionProductIdAPIReq{CollectionListingId: collectionListingId}
	apiResp := &saleschannels.QuerySalesChannelsSpecifyCollectionProductIdAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.ProductIds, 2)
	assert.Equal(t, "16057039432335097907370283", apiResp.ProductIds[0])
}

func TestQuerySalesChannelsSpecifyCollectionProductId_MissingCollectionListingId(t *testing.T) {
	err := (&saleschannels.QuerySalesChannelsSpecifyCollectionProductIdAPIReq{}).Verify()
	assert.EqualError(t, err, "CollectionListingId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Product Listings
// ══════════════════════════════════════════════════════════════════════════════

func TestGetTheProductCountForYourSalesChannel(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("GET", salesChannelsURL(cli, "product_listings/count.json"),
		httpmock.NewStringResponder(200, `{"count":10}`))

	apiReq := &saleschannels.GetTheProductCountForYourSalesChannelAPIReq{}
	apiResp := &saleschannels.GetTheProductCountForYourSalesChannelAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 10, apiResp.Count)
}

func TestQueryProductIdForSalesChannels(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"product_ids":["16057039432335097907370282","16057039432335097907380282"]}`
	httpmock.RegisterResponder("GET", salesChannelsURL(cli, "product_listings/product_ids.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &saleschannels.QueryProductIdForSalesChannelsAPIReq{}
	apiResp := &saleschannels.QueryProductIdForSalesChannelsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.ProductIds, 2)
	assert.Equal(t, "16057039432335097907370282", apiResp.ProductIds[0])
}

func TestQueryProductOfSalesChannels(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"product_listings":[{"product_id":"16057039432335097907370282","title":"Test Product","status":"active","vendor":"Shopline","published_scope":"web"}]}`
	httpmock.RegisterResponder("GET", salesChannelsURL(cli, "product_listings.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &saleschannels.QueryProductOfSalesChannelsAPIReq{}
	apiResp := &saleschannels.QueryProductOfSalesChannelsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.ProductListings, 1)
	assert.Equal(t, "16057039432335097907370282", apiResp.ProductListings[0].ProductId)
	assert.Equal(t, "Test Product", apiResp.ProductListings[0].Title)
}

func TestQuerySpecifyProductForSalesChannels(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	productListingId := "16057039432335097907370282"
	mockResp := `{"product_listing":{"product_id":"16057039432335097907370282","title":"Test Product","status":"active","vendor":"Shopline"}}`
	httpmock.RegisterResponder("GET", salesChannelsURL(cli, fmt.Sprintf("product_listings/%s.json", productListingId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &saleschannels.QuerySpecifyProductForSalesChannelsAPIReq{ProductListingId: productListingId}
	apiResp := &saleschannels.QuerySpecifyProductForSalesChannelsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "16057039432335097907370282", apiResp.ProductListing.ProductId)
	assert.Equal(t, "Test Product", apiResp.ProductListing.Title)
}

func TestQuerySpecifyProductForSalesChannels_MissingProductListingId(t *testing.T) {
	err := (&saleschannels.QuerySpecifyProductForSalesChannelsAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductListingId is required")
}

func TestAddProductToSalesChannels(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	productListingId := "16057039432335097907370282"
	mockResp := `{"product_listing":{"product_id":"16057039432335097907370282","title":"Test Product","status":"active","vendor":"Shopline","published_scope":"web"}}`
	httpmock.RegisterResponder("PUT", salesChannelsURL(cli, fmt.Sprintf("product_listings/%s.json", productListingId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &saleschannels.AddProductToSalesChannelsAPIReq{ProductListingId: productListingId}
	apiResp := &saleschannels.AddProductToSalesChannelsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "16057039432335097907370282", apiResp.ProductListing.ProductId)
	assert.Equal(t, "Test Product", apiResp.ProductListing.Title)
}

func TestAddProductToSalesChannels_MissingProductListingId(t *testing.T) {
	err := (&saleschannels.AddProductToSalesChannelsAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductListingId is required")
}

func TestRemoveFromSalesChannelsProduct(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	productListingId := "16057039432335097907370282"
	httpmock.RegisterResponder("DELETE", salesChannelsURL(cli, fmt.Sprintf("product_listings/%s.json", productListingId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &saleschannels.RemoveFromSalesChannelsProductAPIReq{ProductListingId: productListingId}
	apiResp := &saleschannels.RemoveFromSalesChannelsProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestRemoveFromSalesChannelsProduct_MissingProductListingId(t *testing.T) {
	err := (&saleschannels.RemoveFromSalesChannelsProductAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductListingId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Resource Feedback
// ══════════════════════════════════════════════════════════════════════════════

func TestResourcesFeedback(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	productId := "16057476865495856634490282"
	mockResp := `{"resource_feedback":{"resource_id":"16057476865495856634490282","resource_type":"Product","state":"require_actions","messages":["This is a question"],"created_at":"2023-01-12T15:20:08+08:00"}}`
	httpmock.RegisterResponder("POST", salesChannelsURL(cli, fmt.Sprintf("products/%s/resource_feedback.json", productId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &saleschannels.ResourcesFeedbackAPIReq{ProductId: productId}
	apiResp := &saleschannels.ResourcesFeedbackAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "16057476865495856634490282", apiResp.ResourceFeedback.ResourceId)
	assert.Equal(t, "require_actions", apiResp.ResourceFeedback.State)
	assert.Equal(t, "Product", apiResp.ResourceFeedback.ResourceType)
}

func TestResourcesFeedback_MissingProductId(t *testing.T) {
	err := (&saleschannels.ResourcesFeedbackAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

func TestListOfResourcesFeedback(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	productId := "16057476865495856634490282"
	mockResp := `{"resource_feedback":[{"resource_id":"16057476865495856634490282","resource_type":"Product","state":"require_actions","messages":["This is a question"],"created_at":"2023-01-12T15:20:08+08:00"}]}`
	httpmock.RegisterResponder("GET", salesChannelsURL(cli, fmt.Sprintf("products/%s/resource_feedback.json", productId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &saleschannels.ListOfResourcesFeedbackAPIReq{ProductId: productId}
	apiResp := &saleschannels.ListOfResourcesFeedbackAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.ResourceFeedback, 1)
	assert.Equal(t, "16057476865495856634490282", apiResp.ResourceFeedback[0].ResourceId)
	assert.Equal(t, "require_actions", apiResp.ResourceFeedback[0].State)
}

func TestListOfResourcesFeedback_MissingProductId(t *testing.T) {
	err := (&saleschannels.ListOfResourcesFeedbackAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}
