package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/product"
	"github.com/stretchr/testify/assert"
)

func productURL(cli *client.Client, path string) string {
	return fmt.Sprintf("https://%s.myshopline.com/%s/%s/%s",
		cli.StoreHandle, cli.PathPrefix, cli.ApiVersion, path)
}

// ══════════════════════════════════════════════════════════════════════════════
// Product APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestGetProducts(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"products":[{"id":"16057850264845250791280282","title":"Test Product","status":"active","vendor":"SHOPLINE"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "products/products.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetProductsAPIReq{}
	apiResp := &product.GetProductsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Products, 1)
	assert.Equal(t, "16057850264845250791280282", apiResp.Products[0].Id)
	assert.Equal(t, "Test Product", apiResp.Products[0].Title)
}

func TestGetAProduct(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16057850264845250791280282"
	mockResp := `{"product":{"id":"16057850264845250791280282","title":"Test Product","status":"active","vendor":"SHOPLINE"}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("products/%s.json", productId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetAProductAPIReq{ProductId: productId}
	apiResp := &product.GetAProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, productId, apiResp.Product.Id)
	assert.Equal(t, "Test Product", apiResp.Product.Title)
}

func TestGetAProduct_MissingProductId(t *testing.T) {
	err := (&product.GetAProductAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

func TestGetProductCount(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("GET", productURL(cli, "products/count.json"),
		httpmock.NewStringResponder(200, `{"count":42}`))

	apiReq := &product.GetProductCountAPIReq{}
	apiResp := &product.GetProductCountAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 42, apiResp.Count)
}

func TestUpdateAProduct(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16057850264845250791280282"
	mockResp := `{"product":{"id":"16057850264845250791280282","title":"Updated Product","status":"active"}}`
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("products/%s.json", productId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.UpdateAProductAPIReq{
		ProductId: productId,
		Product:   product.UpdateAProductAPIReqProduct{Title: "Updated Product"},
	}
	apiResp := &product.UpdateAProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, productId, apiResp.Product.Id)
	assert.Equal(t, "Updated Product", apiResp.Product.Title)
}

func TestUpdateAProduct_MissingProductId(t *testing.T) {
	err := (&product.UpdateAProductAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

func TestDeleteAProduct_New(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16057850264845250791280282"
	httpmock.RegisterResponder("DELETE", productURL(cli, fmt.Sprintf("products/%s.json", productId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.DeleteAProductAPIReq{ProductId: productId}
	apiResp := &product.DeleteAProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteAProduct_MissingProductId(t *testing.T) {
	err := (&product.DeleteAProductAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Product Image APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestGetProductImages(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16057850264845250791280282"
	mockResp := `{"images":[{"id":"5785060242207917075","src":"https://example.com/img.png","product_id":"16057850264845250791280282"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("products/%s/images.json", productId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetProductImagesAPIReq{ProductId: productId, Ids: "5785060242207917075"}
	apiResp := &product.GetProductImagesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Images, 1)
	assert.Equal(t, "5785060242207917075", apiResp.Images[0].Id)
}

func TestGetProductImages_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&product.GetProductImagesAPIReq{}).Verify(), "ProductId is required")
	assert.EqualError(t, (&product.GetProductImagesAPIReq{ProductId: "x"}).Verify(), "Ids is required")
}

func TestGetAProductImage(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	imageId := "5785060242207917075"
	productId := "16057850264845250791280282"
	mockResp := `{"image":{"id":"5785060242207917075","src":"https://example.com/img.png","product_id":"16057850264845250791280282"}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("products/%s/images/%s.json", imageId, productId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetAProductImageAPIReq{ImageId: imageId, ProductId: productId}
	apiResp := &product.GetAProductImageAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, imageId, apiResp.Image.Id)
}

func TestGetAProductImage_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&product.GetAProductImageAPIReq{}).Verify(), "ImageId is required")
	assert.EqualError(t, (&product.GetAProductImageAPIReq{ImageId: "x"}).Verify(), "ProductId is required")
}

func TestGetTheCountOfImagesForAProduct(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16057850264845250791280282"
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("products/%s/images/count.json", productId)),
		httpmock.NewStringResponder(200, `{"count":5}`))

	apiReq := &product.GetTheCountOfImagesForAProductAPIReq{ProductId: productId}
	apiResp := &product.GetTheCountOfImagesForAProductAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 5, apiResp.Count)
}

func TestGetTheCountOfImagesForAProduct_MissingProductId(t *testing.T) {
	err := (&product.GetTheCountOfImagesForAProductAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

func TestCreateAProductImage(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16057850264845250791280282"
	mockResp := `{"image":{"id":"5785060242207917075","src":"https://example.com/img.png","product_id":"16057850264845250791280282","position":1}}`
	httpmock.RegisterResponder("POST", productURL(cli, fmt.Sprintf("products/%s/images.json", productId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.CreateAProductImageAPIReq{
		ProductId: productId,
		Image:     product.CreateAProductImageAPIReqImage{Src: "https://example.com/img.png"},
	}
	apiResp := &product.CreateAProductImageAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "5785060242207917075", apiResp.Image.Id)
	assert.Equal(t, productId, apiResp.Image.ProductId)
}

func TestCreateAProductImage_MissingProductId(t *testing.T) {
	err := (&product.CreateAProductImageAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

func TestUpdateImageInformation(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16057850264845250791280282"
	imageId := "5785060242207917075"
	mockResp := `{"image":{"id":"5785060242207917075","alt":"Updated alt","product_id":"16057850264845250791280282"}}`
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("products/%s/images/%s.json", productId, imageId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.UpdateImageInformationAPIReq{ProductId: productId, ImageId: imageId}
	apiResp := &product.UpdateImageInformationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, imageId, apiResp.Image.Id)
}

func TestUpdateImageInformation_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&product.UpdateImageInformationAPIReq{}).Verify(), "ProductId is required")
	assert.EqualError(t, (&product.UpdateImageInformationAPIReq{ProductId: "x"}).Verify(), "ImageId is required")
}

func TestDeleteAProductImage(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	imageId := "5785060242207917075"
	productId := "16057850264845250791280282"
	httpmock.RegisterResponder("DELETE", productURL(cli, fmt.Sprintf("products/%s/images/%s.json", imageId, productId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.DeleteAProductImageAPIReq{ImageId: imageId, ProductId: productId}
	apiResp := &product.DeleteAProductImageAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteAProductImage_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&product.DeleteAProductImageAPIReq{}).Verify(), "ImageId is required")
	assert.EqualError(t, (&product.DeleteAProductImageAPIReq{ImageId: "x"}).Verify(), "ProductId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Product Variant APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestGetProductVariants(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16057850264845250791280282"
	mockResp := `{"variants":[{"id":"18057039439794751459380282","price":"90.22","sku":"S00000001","product_id":"16057850264845250791280282"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("products/%s/variants.json", productId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetProductVariantsAPIReq{ProductId: productId}
	apiResp := &product.GetProductVariantsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Variants, 1)
	assert.Equal(t, "18057039439794751459380282", apiResp.Variants[0].Id)
}

func TestGetProductVariants_MissingProductId(t *testing.T) {
	err := (&product.GetProductVariantsAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

func TestGetAProductVariant(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	variantId := "18057039439794751459380282"
	mockResp := `{"variant":{"id":"18057039439794751459380282","price":"90.22","sku":"S00000001"}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("products/variants/%s.json", variantId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetAProductVariantAPIReq{VariantId: variantId}
	apiResp := &product.GetAProductVariantAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, variantId, apiResp.Variant.Id)
	assert.Equal(t, "90.22", apiResp.Variant.Price)
}

func TestGetAProductVariant_MissingVariantId(t *testing.T) {
	err := (&product.GetAProductVariantAPIReq{}).Verify()
	assert.EqualError(t, err, "VariantId is required")
}

func TestGetTheTotalCountOfProductVariants(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16057850264845250791280282"
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("products/%s/variants/count.json", productId)),
		httpmock.NewStringResponder(200, `{"count":3}`))

	apiReq := &product.GetTheTotalCountOfProductVariantsAPIReq{ProductId: productId}
	apiResp := &product.GetTheTotalCountOfProductVariantsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 3, apiResp.Count)
}

func TestGetTheTotalCountOfProductVariants_MissingProductId(t *testing.T) {
	err := (&product.GetTheTotalCountOfProductVariantsAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

func TestCreateAVariant(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16057850264845250791280282"
	mockResp := `{"variant":{"id":"18057039439794751459380282","price":"90.22","sku":"S00000001","product_id":"16057850264845250791280282"}}`
	httpmock.RegisterResponder("POST", productURL(cli, fmt.Sprintf("products/%s/variants.json", productId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.CreateAVariantAPIReq{
		ProductId: productId,
		Variant:   product.CreateAVariantAPIReqVariant{Price: "90.22", Sku: "S00000001"},
	}
	apiResp := &product.CreateAVariantAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "18057039439794751459380282", apiResp.Variant.Id)
	assert.Equal(t, "90.22", apiResp.Variant.Price)
}

func TestCreateAVariant_MissingProductId(t *testing.T) {
	err := (&product.CreateAVariantAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductId is required")
}

func TestUpdateAVariant(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	variantId := "18057039439794751459380282"
	mockResp := `{"variant":{"id":"18057039439794751459380282","price":"99.99","sku":"S00000001"}}`
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("products/variants/%s.json", variantId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.UpdateAVariantAPIReq{VariantId: variantId}
	apiResp := &product.UpdateAVariantAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, variantId, apiResp.Variant.Id)
}

func TestUpdateAVariant_MissingVariantId(t *testing.T) {
	err := (&product.UpdateAVariantAPIReq{}).Verify()
	assert.EqualError(t, err, "VariantId is required")
}

func TestDeleteAVariant(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	productId := "16057850264845250791280282"
	variantId := "18057039439794751459380282"
	httpmock.RegisterResponder("DELETE", productURL(cli, fmt.Sprintf("products/%s/variants/%s.json", productId, variantId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.DeleteAVariantAPIReq{ProductId: productId, VariantId: variantId}
	apiResp := &product.DeleteAVariantAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteAVariant_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&product.DeleteAVariantAPIReq{}).Verify(), "ProductId is required")
	assert.EqualError(t, (&product.DeleteAVariantAPIReq{ProductId: "x"}).Verify(), "VariantId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Manual Collection APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestGetManualCollections(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"custom_collections":[{"id":"12257170618007271602093384","title":"spring clothing","handle":"spring-clothing"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "products/custom_collections.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetManualCollectionsAPIReq{}
	apiResp := &product.GetManualCollectionsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.CustomCollections, 1)
	assert.Equal(t, "12257170618007271602093384", apiResp.CustomCollections[0].Id)
}

func TestGetTheCountOfManualCollections(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("GET", productURL(cli, "products/custom_collections/count.json"),
		httpmock.NewStringResponder(200, `{"count":7}`))

	apiReq := &product.GetTheCountOfManualCollectionsAPIReq{}
	apiResp := &product.GetTheCountOfManualCollectionsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 7, apiResp.Count)
}

func TestCreateManualCollection(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"custom_collection":{"id":"12257170618007271602093384","title":"spring clothing","handle":"spring-clothing"}}`
	httpmock.RegisterResponder("POST", productURL(cli, "products/custom_collections.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.CreateManualCollectionAPIReq{}
	apiResp := &product.CreateManualCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "12257170618007271602093384", apiResp.CustomCollection.Id)
	assert.Equal(t, "spring clothing", apiResp.CustomCollection.Title)
}

func TestQueryManualCollectionAttributesById(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectionId := "12257170618007271602093384"
	mockResp := `{"custom_collection":{"id":"12257170618007271602093384","title":"spring clothing","handle":"collection-handle"}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("products/custom_collections/%s.json", collectionId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.QueryManualCollectionAttributesByIdAPIReq{CustomCollectionId: collectionId}
	apiResp := &product.QueryManualCollectionAttributesByIdAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, collectionId, apiResp.CustomCollection.Id)
}

func TestQueryManualCollectionAttributesById_MissingCustomCollectionId(t *testing.T) {
	err := (&product.QueryManualCollectionAttributesByIdAPIReq{}).Verify()
	assert.EqualError(t, err, "CustomCollectionId is required")
}

func TestUpdateManualCollection(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectionId := "12257170618007271602093384"
	mockResp := `{"custom_collection":{"id":"12257170618007271602093384","title":"updated collection"}}`
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("products/custom_collections/%s.json", collectionId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.UpdateManualCollectionAPIReq{CustomCollectionId: collectionId}
	apiResp := &product.UpdateManualCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, collectionId, apiResp.CustomCollection.Id)
}

func TestUpdateManualCollection_MissingCustomCollectionId(t *testing.T) {
	err := (&product.UpdateManualCollectionAPIReq{}).Verify()
	assert.EqualError(t, err, "CustomCollectionId is required")
}

func TestDeleteManualCollection(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectionId := "12257170618007271602093384"
	httpmock.RegisterResponder("DELETE", productURL(cli, fmt.Sprintf("products/custom_collections/%s.json", collectionId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.DeleteManualCollectionAPIReq{CustomCollectionId: collectionId}
	apiResp := &product.DeleteManualCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteManualCollection_MissingCustomCollectionId(t *testing.T) {
	err := (&product.DeleteManualCollectionAPIReq{}).Verify()
	assert.EqualError(t, err, "CustomCollectionId is required")
}

func TestAddProductsToAManualCollection(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectionId := "12257170618007271602093384"
	httpmock.RegisterResponder("POST", productURL(cli, fmt.Sprintf("products/custom_collections/%s/products.json", collectionId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.AddProductsToAManualCollectionAPIReq{CollectionId: collectionId}
	apiResp := &product.AddProductsToAManualCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestAddProductsToAManualCollection_MissingCollectionId(t *testing.T) {
	err := (&product.AddProductsToAManualCollectionAPIReq{}).Verify()
	assert.EqualError(t, err, "CollectionId is required")
}

func TestRemoveProductsFromAManualCollection(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectionId := "12257170618007271602093384"
	httpmock.RegisterResponder("DELETE", productURL(cli, fmt.Sprintf("products/custom_collections/%s/products.json", collectionId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.RemoveProductsFromAManualCollectionAPIReq{CollectionId: collectionId}
	apiResp := &product.RemoveProductsFromAManualCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestRemoveProductsFromAManualCollection_MissingCollectionId(t *testing.T) {
	err := (&product.RemoveProductsFromAManualCollectionAPIReq{}).Verify()
	assert.EqualError(t, err, "CollectionId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Smart Collection APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestGetIntelligentCollections(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"smart_collections":[{"id":"12257170618007271602093384","title":"smart collection","handle":"smart-collection"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "products/smart_collections.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetIntelligentCollectionsAPIReq{}
	apiResp := &product.GetIntelligentCollectionsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.SmartCollections, 1)
	assert.Equal(t, "12257170618007271602093384", apiResp.SmartCollections[0].Id)
}

func TestGetTheCountOfIntelligentCollections(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("GET", productURL(cli, "products/smart_collections/count.json"),
		httpmock.NewStringResponder(200, `{"count":4}`))

	apiReq := &product.GetTheCountOfIntelligentCollectionsAPIReq{}
	apiResp := &product.GetTheCountOfIntelligentCollectionsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 4, apiResp.Count)
}

func TestCreateSmartCollection(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"smart_collection":{"id":"12257170618007271602093384","title":"spring clothing","handle":"spring-clothing"}}`
	httpmock.RegisterResponder("POST", productURL(cli, "products/smart_collections.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.CreateSmartCollectionAPIReq{}
	apiResp := &product.CreateSmartCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "12257170618007271602093384", apiResp.SmartCollection.Id)
}

func TestQueryIntelligentCollectionAttributesById(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectionId := "12257170618007271602093384"
	mockResp := `{"smart_collection":{"id":"12257170618007271602093384","title":"smart collection"}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("products/smart_collections/%s.json", collectionId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.QueryIntelligentCollectionAttributesByIdAPIReq{SmartCollectionId: collectionId}
	apiResp := &product.QueryIntelligentCollectionAttributesByIdAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, collectionId, apiResp.SmartCollection.Id)
}

func TestQueryIntelligentCollectionAttributesById_MissingSmartCollectionId(t *testing.T) {
	err := (&product.QueryIntelligentCollectionAttributesByIdAPIReq{}).Verify()
	assert.EqualError(t, err, "SmartCollectionId is required")
}

func TestUpdateSmartCollection(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectionId := "12257170618007271602093384"
	mockResp := `{"smart_collection":{"id":"12257170618007271602093384","title":"updated smart collection"}}`
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("products/smart_collections/%s.json", collectionId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.UpdateSmartCollectionAPIReq{SmartCollectionId: collectionId}
	apiResp := &product.UpdateSmartCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, collectionId, apiResp.SmartCollection.Id)
}

func TestUpdateSmartCollection_MissingSmartCollectionId(t *testing.T) {
	err := (&product.UpdateSmartCollectionAPIReq{}).Verify()
	assert.EqualError(t, err, "SmartCollectionId is required")
}

func TestUpdateProductSortingInSmartCollections(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectionId := "12257170618007271602093384"
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("products/smart_collections/%s/order.json", collectionId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.UpdateProductSortingInSmartCollectionsAPIReq{CollectionId: collectionId}
	apiResp := &product.UpdateProductSortingInSmartCollectionsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestUpdateProductSortingInSmartCollections_MissingCollectionId(t *testing.T) {
	err := (&product.UpdateProductSortingInSmartCollectionsAPIReq{}).Verify()
	assert.EqualError(t, err, "CollectionId is required")
}

func TestDeleteIntelligentCollection(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectionId := "12257170618007271602093384"
	httpmock.RegisterResponder("DELETE", productURL(cli, fmt.Sprintf("products/smart_collections/%s.json", collectionId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.DeleteIntelligentCollectionAPIReq{SmartCollectionId: collectionId}
	apiResp := &product.DeleteIntelligentCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteIntelligentCollection_MissingSmartCollectionId(t *testing.T) {
	err := (&product.DeleteIntelligentCollectionAPIReq{}).Verify()
	assert.EqualError(t, err, "SmartCollectionId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection Relationship APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestGetProductCollectionRelationships(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"collects":[{"id":5392284410644534000,"product_id":"16057039432335097907370282","collection_id":"12249026592161154694000282"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "products/collects.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetProductCollectionRelationshipsAPIReq{}
	apiResp := &product.GetProductCollectionRelationshipsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Collects, 1)
}

func TestGetTheCountOfProductCollectionRelationships(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("GET", productURL(cli, "products/collects/count.json"),
		httpmock.NewStringResponder(200, `{"count":10}`))

	apiReq := &product.GetTheCountOfProductCollectionRelationshipsAPIReq{}
	apiResp := &product.GetTheCountOfProductCollectionRelationshipsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 10, apiResp.Count)
}

func TestGetAProductCollectionRelationship(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectId := "5392284410644534000"
	mockResp := `{"collect":{"id":5392284410644534000,"product_id":"16057039432335097907370282","collection_id":"12249026592161154694000282"}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("products/collects/%s.json", collectId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetAProductCollectionRelationshipAPIReq{CollectId: collectId}
	apiResp := &product.GetAProductCollectionRelationshipAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, int64(5392284410644534000), apiResp.Collect.Id)
}

func TestGetAProductCollectionRelationship_MissingCollectId(t *testing.T) {
	err := (&product.GetAProductCollectionRelationshipAPIReq{}).Verify()
	assert.EqualError(t, err, "CollectId is required")
}

func TestCreateAProductCollectionRelationship(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"collect":{"id":5704511242397819000,"product_id":"16057039432335097907370282","collection_id":"12249026592161154694000282"}}`
	httpmock.RegisterResponder("POST", productURL(cli, "products/collects.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.CreateAProductCollectionRelationshipAPIReq{}
	apiResp := &product.CreateAProductCollectionRelationshipAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, int64(5704511242397819000), apiResp.Collect.Id)
}

func TestDeleteAProductCollectionRelationship(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectId := "5392284410644534000"
	httpmock.RegisterResponder("DELETE", productURL(cli, fmt.Sprintf("products/collects/%s.json", collectId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.DeleteAProductCollectionRelationshipAPIReq{CollectId: collectId}
	apiResp := &product.DeleteAProductCollectionRelationshipAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteAProductCollectionRelationship_MissingCollectId(t *testing.T) {
	err := (&product.DeleteAProductCollectionRelationshipAPIReq{}).Verify()
	assert.EqualError(t, err, "CollectId is required")
}

func TestQueryCollectionInformationById(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectionId := "12257170618007271602093385"
	mockResp := `{"collection":{"id":"12257170618007271602093385","title":"This is a title","handle":"collection-handle"}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("products/collections/%s.json", collectionId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.QueryCollectionInformationByIdAPIReq{CollectionId: collectionId}
	apiResp := &product.QueryCollectionInformationByIdAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, collectionId, apiResp.Collection.Id)
}

func TestQueryCollectionInformationById_MissingCollectionId(t *testing.T) {
	err := (&product.QueryCollectionInformationByIdAPIReq{}).Verify()
	assert.EqualError(t, err, "CollectionId is required")
}

func TestGetProductsInACollection(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	collectionId := "12257170618007271602093384"
	mockResp := `{"products":[{"id":"16057850264845250791280282","title":"Test Product"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("products/collections/%s/products.json", collectionId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetProductsInACollectionAPIReq{CollectionId: collectionId}
	apiResp := &product.GetProductsInACollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Products, 1)
}

func TestGetProductsInACollection_MissingCollectionId(t *testing.T) {
	err := (&product.GetProductsInACollectionAPIReq{}).Verify()
	assert.EqualError(t, err, "CollectionId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Gift Card APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestGetGiftCards(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"gift_cards":[{"id":"30171274557691301804060045","initial_value":"200","balance":"200","currency":"USD"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "gift_cards.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetGiftCardsAPIReq{}
	apiResp := &product.GetGiftCardsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.GiftCards, 1)
	assert.Equal(t, "30171274557691301804060045", apiResp.GiftCards[0].Id)
}

func TestQueryNumberOfGiftCards(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("GET", productURL(cli, "gift_cards/count.json"),
		httpmock.NewStringResponder(200, `{"count":8}`))

	apiReq := &product.QueryNumberOfGiftCardsAPIReq{}
	apiResp := &product.QueryNumberOfGiftCardsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 8, apiResp.Count)
}

func TestQuerySingleGiftCard(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	giftCardId := "30157043359245339231360282"
	mockResp := `{"gift_card":{"id":"30157043359245339231360282","initial_value":"200","balance":"200","currency":"CNY"}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("gift_cards/%s.json", giftCardId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.QuerySingleGiftCardAPIReq{GiftCardId: giftCardId}
	apiResp := &product.QuerySingleGiftCardAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, giftCardId, apiResp.GiftCard.Id)
}

func TestQuerySingleGiftCard_MissingGiftCardId(t *testing.T) {
	err := (&product.QuerySingleGiftCardAPIReq{}).Verify()
	assert.EqualError(t, err, "GiftCardId is required")
}

func TestCreateGiftCards(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"gift_card":{"id":"30157043359245339231360282","initial_value":"200","balance":"200","currency":"CNY","code":"dbe95e7f02606fc4"}}`
	httpmock.RegisterResponder("POST", productURL(cli, "gift_cards.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.CreateGiftCardsAPIReq{InitialValue: "200"}
	apiResp := &product.CreateGiftCardsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "30157043359245339231360282", apiResp.GiftCard.Id)
}

func TestCreateGiftCards_MissingInitialValue(t *testing.T) {
	err := (&product.CreateGiftCardsAPIReq{}).Verify()
	assert.EqualError(t, err, "InitialValue is required")
}

func TestUpdateGiftCard(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	giftCardId := "30157043359245339231360282"
	mockResp := `{"gift_card":{"id":"30157043359245339231360282","note":"updated note","balance":"200"}}`
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("gift_cards/%s.json", giftCardId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.UpdateGiftCardAPIReq{GiftCardId: giftCardId}
	apiResp := &product.UpdateGiftCardAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, giftCardId, apiResp.GiftCard.Id)
}

func TestUpdateGiftCard_MissingGiftCardId(t *testing.T) {
	err := (&product.UpdateGiftCardAPIReq{}).Verify()
	assert.EqualError(t, err, "GiftCardId is required")
}

func TestDisableGiftCard(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	giftCardId := "30157043359245339231360282"
	mockResp := `{"gift_card":{"id":"30157043359245339231360282","disabled_at":"2023-08-16T23:59:59+08:00"}}`
	httpmock.RegisterResponder("POST", productURL(cli, fmt.Sprintf("gift_cards/%s/disable.json", giftCardId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.DisableGiftCardAPIReq{GiftCardId: giftCardId}
	apiResp := &product.DisableGiftCardAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, giftCardId, apiResp.GiftCard.Id)
}

func TestDisableGiftCard_MissingGiftCardId(t *testing.T) {
	err := (&product.DisableGiftCardAPIReq{}).Verify()
	assert.EqualError(t, err, "GiftCardId is required")
}

func TestGetGiftCardOperationRecords(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	giftCardId := "30157043359245339231360282"
	mockResp := `{"gift_card_records":[{"id":"5504120995916023830","operation_type":"OPERATE_ACTION_DEDUCT","cur_balance":100}]}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("gift_cards/%s/operation_records.json", giftCardId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetGiftCardOperationRecordsAPIReq{GiftCardId: giftCardId}
	apiResp := &product.GetGiftCardOperationRecordsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.GiftCardRecords, 1)
}

func TestGetGiftCardOperationRecords_MissingGiftCardId(t *testing.T) {
	err := (&product.GetGiftCardOperationRecordsAPIReq{}).Verify()
	assert.EqualError(t, err, "GiftCardId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Inventory APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestGetInventoryItems(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"inventory_items":[{"id":"5703943989105270965","sku":"K0000000001","cost":"10.91"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "inventory_items.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetInventoryItemsAPIReq{Ids: "5703943989105270965"}
	apiResp := &product.GetInventoryItemsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.InventoryItems, 1)
	assert.Equal(t, "5703943989105270965", apiResp.InventoryItems[0].Id)
}

func TestGetInventoryItems_MissingIds(t *testing.T) {
	err := (&product.GetInventoryItemsAPIReq{}).Verify()
	assert.EqualError(t, err, "Ids is required")
}

func TestGetAnInventoryItem(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	inventoryItemId := "5703943989105270965"
	mockResp := `{"inventory_item":{"id":"5703943989105270965","sku":"K0000000001","cost":"10.91"}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("inventory_items/%s.json", inventoryItemId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetAnInventoryItemAPIReq{InventoryItemId: inventoryItemId}
	apiResp := &product.GetAnInventoryItemAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, inventoryItemId, apiResp.InventoryItem.Id)
}

func TestGetAnInventoryItem_MissingInventoryItemId(t *testing.T) {
	err := (&product.GetAnInventoryItemAPIReq{}).Verify()
	assert.EqualError(t, err, "InventoryItemId is required")
}

func TestUpdateAnInventoryItem(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	inventoryItemId := "5703943989105270965"
	mockResp := `{"inventory_item":{"id":"5703943989105270965","sku":"K00000001","cost":"10.9"}}`
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("inventory_items/%s.json", inventoryItemId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.UpdateAnInventoryItemAPIReq{InventoryItemId: inventoryItemId}
	apiResp := &product.UpdateAnInventoryItemAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, inventoryItemId, apiResp.InventoryItem.Id)
}

func TestUpdateAnInventoryItem_MissingInventoryItemId(t *testing.T) {
	err := (&product.UpdateAnInventoryItemAPIReq{}).Verify()
	assert.EqualError(t, err, "InventoryItemId is required")
}

func TestGetInventoryQuantities(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"inventory_levels":[{"inventory_item_id":"7177011084762551696","location_id":"5421704248135526901","available":10}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "inventory_levels.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetInventoryQuantitiesAPIReq{InventoryItemIds: "7177011084762551696"}
	apiResp := &product.GetInventoryQuantitiesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.InventoryLevels, 1)
}

func TestGetInventoryQuantities_MissingInventoryItemIds(t *testing.T) {
	err := (&product.GetInventoryQuantitiesAPIReq{}).Verify()
	assert.EqualError(t, err, "InventoryItemIds is required")
}

func TestSetItemInventory(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"inventory_level":{"inventory_item_id":"5703943240925324252","location_id":"5421704248135526901","available":10}}`
	httpmock.RegisterResponder("POST", productURL(cli, "inventory_levels/set.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.SetItemInventoryAPIReq{InventoryItemId: "5703943240925324252", LocationId: "5421704248135526901"}
	apiResp := &product.SetItemInventoryAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "5703943240925324252", apiResp.InventoryLevel.InventoryItemId)
}

func TestSetItemInventory_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&product.SetItemInventoryAPIReq{}).Verify(), "InventoryItemId is required")
	assert.EqualError(t, (&product.SetItemInventoryAPIReq{InventoryItemId: "x"}).Verify(), "LocationId is required")
}

func TestUpdateItemInventory(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"inventory_level":{"inventory_item_id":"7177011084762551696","location_id":"5421703880295066100","available":10}}`
	httpmock.RegisterResponder("POST", productURL(cli, "inventory_levels/adjust.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.UpdateItemInventoryAPIReq{InventoryItemId: "7177011084762551696", LocationId: "5421703880295066100"}
	apiResp := &product.UpdateItemInventoryAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "7177011084762551696", apiResp.InventoryLevel.InventoryItemId)
}

func TestUpdateItemInventory_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&product.UpdateItemInventoryAPIReq{}).Verify(), "InventoryItemId is required")
	assert.EqualError(t, (&product.UpdateItemInventoryAPIReq{InventoryItemId: "x"}).Verify(), "LocationId is required")
}

func TestLinkAnInventoryItemToALocation(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"inventory_level":{"inventory_item_id":"7177011084762551696","location_id":"5421704248135526901","available":10}}`
	httpmock.RegisterResponder("POST", productURL(cli, "inventory_levels/connect.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.LinkAnInventoryItemToALocationAPIReq{InventoryItemId: "7177011084762551696", LocationId: "5421704248135526901"}
	apiResp := &product.LinkAnInventoryItemToALocationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "7177011084762551696", apiResp.InventoryLevel.InventoryItemId)
}

func TestLinkAnInventoryItemToALocation_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&product.LinkAnInventoryItemToALocationAPIReq{}).Verify(), "InventoryItemId is required")
	assert.EqualError(t, (&product.LinkAnInventoryItemToALocationAPIReq{InventoryItemId: "x"}).Verify(), "LocationId is required")
}

func TestDisconnectAnInventoryItemFromALocation(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("DELETE", productURL(cli, "inventory_levels/disconnect.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.DisconnectAnInventoryItemFromALocationAPIReq{InventoryItemId: "7177011084762551696", LocationId: "5421704248135526901"}
	apiResp := &product.DisconnectAnInventoryItemFromALocationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDisconnectAnInventoryItemFromALocation_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&product.DisconnectAnInventoryItemFromALocationAPIReq{}).Verify(), "InventoryItemId is required")
	assert.EqualError(t, (&product.DisconnectAnInventoryItemFromALocationAPIReq{InventoryItemId: "x"}).Verify(), "LocationId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Location APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestGetLocations(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"locations":[{"id":"5668167071852661850","name":"Ottawa Store","active":true}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "locations/list.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetLocationsAPIReq{}
	apiResp := &product.GetLocationsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Locations, 1)
	assert.Equal(t, "5668167071852661850", apiResp.Locations[0].Id)
}

func TestStatisticsNumberOfLocations(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("GET", productURL(cli, "locations/count.json"),
		httpmock.NewStringResponder(200, `{"count":4}`))

	apiReq := &product.StatisticsNumberOfLocationsAPIReq{}
	apiResp := &product.StatisticsNumberOfLocationsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 4, apiResp.Count)
}

func TestBasedOnIdQueryLocation(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	locationId := "5668167071852661850"
	mockResp := `{"location":{"id":"5668167071852661850","name":"Ottawa Store","active":true}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("locations/%s.json", locationId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.BasedOnIdQueryLocationAPIReq{Id: locationId}
	apiResp := &product.BasedOnIdQueryLocationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, locationId, apiResp.Location.Id)
}

func TestBasedOnIdQueryLocation_MissingId(t *testing.T) {
	err := (&product.BasedOnIdQueryLocationAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// File APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestGetFiles(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"files":[{"id":"5824341616684517429","url":"https://example.com/file.png","alt":"image alt"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "files/files.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetFilesAPIReq{}
	apiResp := &product.GetFilesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Files, 1)
	assert.Equal(t, "5824341616684517429", apiResp.Files[0].Id)
}

func TestGetAFile(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	fileId := "5824341616684517429"
	mockResp := `{"id":"5824341616684517429","alt":"image alt","url":"https://example.com/file.png"}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("files/%s.json", fileId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetAFileAPIReq{FileId: fileId}
	apiResp := &product.GetAFileAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestGetAFile_MissingFileId(t *testing.T) {
	err := (&product.GetAFileAPIReq{}).Verify()
	assert.EqualError(t, err, "FileId is required")
}

func TestCreateAFile(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"id":"5824341616684517429","alt":"image alt","url":"https://example.com/file.png"}`
	httpmock.RegisterResponder("POST", productURL(cli, "files/files.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.CreateAFileAPIReq{ContentType: "IMAGE", OriginalSource: "https://example.com/img.png"}
	apiResp := &product.CreateAFileAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestCreateAFile_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&product.CreateAFileAPIReq{}).Verify(), "ContentType is required")
	assert.EqualError(t, (&product.CreateAFileAPIReq{ContentType: "IMAGE"}).Verify(), "OriginalSource is required")
}

func TestUpdateAFile(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	fileId := "5824341616684517429"
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("files/%s.json", fileId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.UpdateAFileAPIReq{FileId: fileId}
	apiResp := &product.UpdateAFileAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestUpdateAFile_MissingFileId(t *testing.T) {
	err := (&product.UpdateAFileAPIReq{}).Verify()
	assert.EqualError(t, err, "FileId is required")
}

func TestDeleteAFile(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	fileId := "5824341616684517429"
	httpmock.RegisterResponder("DELETE", productURL(cli, fmt.Sprintf("files/%s.json", fileId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.DeleteAFileAPIReq{FileId: fileId}
	apiResp := &product.DeleteAFileAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteAFile_MissingFileId(t *testing.T) {
	err := (&product.DeleteAFileAPIReq{}).Verify()
	assert.EqualError(t, err, "FileId is required")
}

func TestInitiateAFileStagedUpload(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"signs":[{"key":"temp/file/store/image.jpg","url":"https://example.com/upload"}]}`
	httpmock.RegisterResponder("POST", productURL(cli, "files/upload.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.InitiateAFileStagedUploadAPIReq{}
	apiResp := &product.InitiateAFileStagedUploadAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Signs, 1)
}

// ══════════════════════════════════════════════════════════════════════════════
// Taxonomy APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestQueryTaxonomyOfTheStore(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"taxonomies":[{"id":"7004220176356349981","collection_id":"12257170618007271602093384","parent_id":"7004220176355548881"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "taxonomies.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.QueryTaxonomyOfTheStoreAPIReq{}
	apiResp := &product.QueryTaxonomyOfTheStoreAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Taxonomies, 1)
	assert.Equal(t, "7004220176356349981", apiResp.Taxonomies[0].Id)
}

func TestGetATaxonomyCollectionNode(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	taxonomyId := "7004220176356349981"
	mockResp := `{"taxonomy":{"id":"7004220176356349981","collection_id":"16050375155238626683133099","is_leaf":false}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("taxonomies/%s.json", taxonomyId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetATaxonomyCollectionNodeAPIReq{TaxonomyId: taxonomyId}
	apiResp := &product.GetATaxonomyCollectionNodeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, taxonomyId, apiResp.Taxonomy.Id)
}

func TestGetATaxonomyCollectionNode_MissingTaxonomyId(t *testing.T) {
	err := (&product.GetATaxonomyCollectionNodeAPIReq{}).Verify()
	assert.EqualError(t, err, "TaxonomyId is required")
}

func TestGetChildTaxonomyCollectionNodes(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"taxonomies":[{"id":"7004220176356349981","collection_id":"16050375155238626683133099"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "taxonomies/children.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetChildTaxonomyCollectionNodesAPIReq{}
	apiResp := &product.GetChildTaxonomyCollectionNodesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Taxonomies, 1)
}

func TestGetTaxonomyCollectionNodes(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"products_taxonomies":[{"product_id":"16057850264845250791280282"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "taxonomies/taxonomy_collections.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetTaxonomyCollectionNodesAPIReq{ProductIds: "16057850264845250791280282"}
	apiResp := &product.GetTaxonomyCollectionNodesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.ProductsTaxonomies, 1)
}

func TestGetTaxonomyCollectionNodes_MissingProductIds(t *testing.T) {
	err := (&product.GetTaxonomyCollectionNodesAPIReq{}).Verify()
	assert.EqualError(t, err, "ProductIds is required")
}

func TestCreateATaxonomyCollectionNode(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"taxonomy":{"id":"7004220176356349981","collection_id":"16050375155238626683133099","is_leaf":false}}`
	httpmock.RegisterResponder("POST", productURL(cli, "taxonomies.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.CreateATaxonomyCollectionNodeAPIReq{CollectionId: "16050375155238626683133099"}
	apiResp := &product.CreateATaxonomyCollectionNodeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "7004220176356349981", apiResp.Taxonomy.Id)
}

func TestCreateATaxonomyCollectionNode_MissingCollectionId(t *testing.T) {
	err := (&product.CreateATaxonomyCollectionNodeAPIReq{}).Verify()
	assert.EqualError(t, err, "CollectionId is required")
}

func TestUpdateATaxonomyCollectionNode(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	taxonomyId := "7004220176356349981"
	mockResp := `{"taxonomy":{"id":"7004220176356349981","collection_id":"16050375155238626683133099"}}`
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("taxonomies/%s.json", taxonomyId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.UpdateATaxonomyCollectionNodeAPIReq{TaxonomyId: taxonomyId}
	apiResp := &product.UpdateATaxonomyCollectionNodeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, taxonomyId, apiResp.Taxonomy.Id)
}

func TestUpdateATaxonomyCollectionNode_MissingTaxonomyId(t *testing.T) {
	err := (&product.UpdateATaxonomyCollectionNodeAPIReq{}).Verify()
	assert.EqualError(t, err, "TaxonomyId is required")
}

func TestDeleteATaxonomyCollectionNode(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	taxonomyId := "7004220176356349981"
	httpmock.RegisterResponder("DELETE", productURL(cli, fmt.Sprintf("taxonomies/%s.json", taxonomyId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.DeleteATaxonomyCollectionNodeAPIReq{TaxonomyId: taxonomyId}
	apiResp := &product.DeleteATaxonomyCollectionNodeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteATaxonomyCollectionNode_MissingTaxonomyId(t *testing.T) {
	err := (&product.DeleteATaxonomyCollectionNodeAPIReq{}).Verify()
	assert.EqualError(t, err, "TaxonomyId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Selling Plan Group APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestQueryMultipleSellingPrograms(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"selling_plan_group_list":[{"id":"14056200245844372441100009","name":"Subscription plan","position":1}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "selling_plan_group/selling_plan_group.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.QueryMultipleSellingProgramsAPIReq{}
	apiResp := &product.QueryMultipleSellingProgramsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.SellingPlanGroupList, 1)
	assert.Equal(t, "14056200245844372441100009", apiResp.SellingPlanGroupList[0].Id)
}

func TestGetTheSalesPlanGroupTotalNumber(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("GET", productURL(cli, "selling_plan_groups/count.json"),
		httpmock.NewStringResponder(200, `{"count":10}`))

	apiReq := &product.GetTheSalesPlanGroupTotalNumberAPIReq{}
	apiResp := &product.GetTheSalesPlanGroupTotalNumberAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 10, apiResp.Count)
}

func TestQuerySalesProgramGroupDetails(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	sellingPlanGroupId := "14056200245844372441100009"
	mockResp := `{"selling_plan_group":{"id":"14056200245844372441100009","name":"Subscription plan","position":1}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("selling_plan_group/%s.json", sellingPlanGroupId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.QuerySalesProgramGroupDetailsAPIReq{SellingPlanGroupId: sellingPlanGroupId}
	apiResp := &product.QuerySalesProgramGroupDetailsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, sellingPlanGroupId, apiResp.SellingPlanGroup.Id)
}

func TestQuerySalesProgramGroupDetails_MissingSellingPlanGroupId(t *testing.T) {
	err := (&product.QuerySalesProgramGroupDetailsAPIReq{}).Verify()
	assert.EqualError(t, err, "SellingPlanGroupId is required")
}

func TestCreateASalesPlanGroup(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"selling_plan_group":{"id":"14156200245844372441120009","name":"Subscription plan","position":1}}`
	httpmock.RegisterResponder("POST", productURL(cli, "selling_plan_groups/selling_plan_group.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.CreateASalesPlanGroupAPIReq{Name: "Subscription plan"}
	apiResp := &product.CreateASalesPlanGroupAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "14156200245844372441120009", apiResp.SellingPlanGroup.Id)
}

func TestCreateASalesPlanGroup_MissingName(t *testing.T) {
	err := (&product.CreateASalesPlanGroupAPIReq{}).Verify()
	assert.EqualError(t, err, "Name is required")
}

func TestUpdateSalesPlanGroup(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"selling_plan_group":{"id":"14056200245844372441100009","name":"Updated plan","position":1}}`
	httpmock.RegisterResponder("PUT", productURL(cli, "selling_plan_groups/selling_plan_group.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.UpdateSalesPlanGroupAPIReq{Id: "14056200245844372441100009", Name: "Updated plan"}
	apiResp := &product.UpdateSalesPlanGroupAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "14056200245844372441100009", apiResp.SellingPlanGroup.Id)
}

func TestUpdateSalesPlanGroup_MissingId(t *testing.T) {
	err := (&product.UpdateSalesPlanGroupAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestDeleteASalesPlanGroup(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	sellingPlanGroupId := "14056200245844372441100009"
	httpmock.RegisterResponder("DELETE", productURL(cli, fmt.Sprintf("selling_plan_group/%s.json", sellingPlanGroupId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.DeleteASalesPlanGroupAPIReq{SellingPlanGroupId: sellingPlanGroupId}
	apiResp := &product.DeleteASalesPlanGroupAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteASalesPlanGroup_MissingSellingPlanGroupId(t *testing.T) {
	err := (&product.DeleteASalesPlanGroupAPIReq{}).Verify()
	assert.EqualError(t, err, "SellingPlanGroupId is required")
}

func TestAddProductsToASalesPlanGroup(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("POST", productURL(cli, "selling_plan_group/binding.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.AddProductsToASalesPlanGroupAPIReq{Id: "14056200245844372441100009"}
	apiResp := &product.AddProductsToASalesPlanGroupAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestAddProductsToASalesPlanGroup_MissingId(t *testing.T) {
	err := (&product.AddProductsToASalesPlanGroupAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestRemoveProductsFromASalesPlanGroup(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("DELETE", productURL(cli, "selling_plan_group/binding.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.RemoveProductsFromASalesPlanGroupAPIReq{Id: "14056200245844372441100009", ProductIds: "16057039432335097907370282"}
	apiResp := &product.RemoveProductsFromASalesPlanGroupAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestRemoveProductsFromASalesPlanGroup_MissingId(t *testing.T) {
	err := (&product.RemoveProductsFromASalesPlanGroupAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Company Location Catalog APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestRetrieveAListOfCompanyLocationCatalogs(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"catalogs":[{"catalog_id":"6598087148213149672","title":"A catalog title","status":1}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "company_location_catalogs.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.RetrieveAListOfCompanyLocationCatalogsAPIReq{}
	apiResp := &product.RetrieveAListOfCompanyLocationCatalogsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Catalogs, 1)
	assert.Equal(t, "6598087148213149672", apiResp.Catalogs[0].CatalogId)
}

func TestRetrieveACompanyLocationCatalog(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	catalogId := "6598087148213149672"
	mockResp := `{"catalog":{"catalog_id":"6598087148213149672","title":"A catalog title","status":1}}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("company_location_catalog/%s.json", catalogId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.RetrieveACompanyLocationCatalogAPIReq{CatalogId: catalogId}
	apiResp := &product.RetrieveACompanyLocationCatalogAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, catalogId, apiResp.Catalog.CatalogId)
}

func TestRetrieveACompanyLocationCatalog_MissingCatalogId(t *testing.T) {
	err := (&product.RetrieveACompanyLocationCatalogAPIReq{}).Verify()
	assert.EqualError(t, err, "CatalogId is required")
}

func TestCreateACompanyLocationCatalog(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"catalog_id":"6598087148213149672"}`
	httpmock.RegisterResponder("POST", productURL(cli, "company_location_catalog/create.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.CreateACompanyLocationCatalogAPIReq{Title: "A catalog title"}
	apiResp := &product.CreateACompanyLocationCatalogAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "6598087148213149672", apiResp.CatalogId)
}

func TestCreateACompanyLocationCatalog_MissingTitle(t *testing.T) {
	err := (&product.CreateACompanyLocationCatalogAPIReq{}).Verify()
	assert.EqualError(t, err, "Title is required")
}

func TestUpdateACompanyLocationCatalog(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	catalogId := "6598087148213149672"
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("company_location_catalog/%s.json", catalogId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.UpdateACompanyLocationCatalogAPIReq{CatalogId: catalogId}
	apiResp := &product.UpdateACompanyLocationCatalogAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestUpdateACompanyLocationCatalog_MissingCatalogId(t *testing.T) {
	err := (&product.UpdateACompanyLocationCatalogAPIReq{}).Verify()
	assert.EqualError(t, err, "CatalogId is required")
}

func TestDeleteCatalog(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	catalogId := "6598087148213149672"
	httpmock.RegisterResponder("DELETE", productURL(cli, fmt.Sprintf("company_location_catalog/%s.json", catalogId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.DeleteCatalogAPIReq{CatalogId: catalogId}
	apiResp := &product.DeleteCatalogAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteCatalog_MissingCatalogId(t *testing.T) {
	err := (&product.DeleteCatalogAPIReq{}).Verify()
	assert.EqualError(t, err, "CatalogId is required")
}

func TestGetCatalogLocations(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	mockResp := `{"catalog_locations":[{"catalog_id":"6598087148213149672","company_location_ids":["6598085940706559169"],"company_locations_count":1}]}`
	httpmock.RegisterResponder("GET", productURL(cli, "company_location_catalog/locations.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetCatalogLocationsAPIReq{}
	apiResp := &product.GetCatalogLocationsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.CatalogLocations, 1)
	assert.Equal(t, "6598087148213149672", apiResp.CatalogLocations[0].CatalogId)
}

func TestGetCatalogLocationsRecommended(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	catalogId := "6598087148213149672"
	mockResp := `{"company_location_ids":["6598085940706559169"]}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("company_location_catalog/%s/locations.json", catalogId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetCatalogLocationsRecommendedAPIReq{CatalogId: catalogId}
	apiResp := &product.GetCatalogLocationsRecommendedAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, []string{"6598085940706559169"}, apiResp.CompanyLocationIds)
}

func TestGetCatalogLocationsRecommended_MissingCatalogId(t *testing.T) {
	err := (&product.GetCatalogLocationsRecommendedAPIReq{}).Verify()
	assert.EqualError(t, err, "CatalogId is required")
}

func TestAddCompanyLocations(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	catalogId := "6598087148213149672"
	mockResp := `{"company_location_ids":["6598085940706559169"]}`
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("company_location_catalog/%s/locations.json", catalogId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.AddCompanyLocationsAPIReq{CatalogId: catalogId}
	apiResp := &product.AddCompanyLocationsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, []string{"6598085940706559169"}, apiResp.CompanyLocationIds)
}

func TestAddCompanyLocations_MissingCatalogId(t *testing.T) {
	err := (&product.AddCompanyLocationsAPIReq{}).Verify()
	assert.EqualError(t, err, "CatalogId is required")
}

func TestRemoveCompanyLocations(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	catalogId := "6598087148213149672"
	httpmock.RegisterResponder("DELETE", productURL(cli, fmt.Sprintf("company_location_catalog/%s/locations.json", catalogId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.RemoveCompanyLocationsAPIReq{CatalogId: catalogId}
	apiResp := &product.RemoveCompanyLocationsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestRemoveCompanyLocations_MissingCatalogId(t *testing.T) {
	err := (&product.RemoveCompanyLocationsAPIReq{}).Verify()
	assert.EqualError(t, err, "CatalogId is required")
}

func TestGetCatalogProducts(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	catalogId := "6598087148213149672"
	mockResp := `{"individual_price_products_count":1,"products":[{"id":"16057850264845250791280282"}]}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("company_location_catalog/%s/products.json", catalogId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.GetCatalogProductsAPIReq{CatalogId: catalogId}
	apiResp := &product.GetCatalogProductsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, 1, apiResp.IndividualPriceProductsCount)
}

func TestGetCatalogProducts_MissingCatalogId(t *testing.T) {
	err := (&product.GetCatalogProductsAPIReq{}).Verify()
	assert.EqualError(t, err, "CatalogId is required")
}

func TestAddCatalogProducts(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	catalogId := "6598087148213149672"
	mockResp := `{"product_ids":["16064649853088531841170520"]}`
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("company_location_catalog/%s/published_products.json", catalogId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.AddCatalogProductsAPIReq{CatalogId: catalogId}
	apiResp := &product.AddCatalogProductsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, []string{"16064649853088531841170520"}, apiResp.ProductIds)
}

func TestAddCatalogProducts_MissingCatalogId(t *testing.T) {
	err := (&product.AddCatalogProductsAPIReq{}).Verify()
	assert.EqualError(t, err, "CatalogId is required")
}

func TestRemoveCatalogProducts(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	catalogId := "6598087148213149672"
	mockResp := `{"product_ids":["16064649853088531841170520"]}`
	httpmock.RegisterResponder("PUT", productURL(cli, fmt.Sprintf("company_location_catalog/%s/unpublished_products.json", catalogId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.RemoveCatalogProductsAPIReq{CatalogId: catalogId}
	apiResp := &product.RemoveCatalogProductsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, []string{"16064649853088531841170520"}, apiResp.ProductIds)
}

func TestRemoveCatalogProducts_MissingCatalogId(t *testing.T) {
	err := (&product.RemoveCatalogProductsAPIReq{}).Verify()
	assert.EqualError(t, err, "CatalogId is required")
}

func TestSeparatePricingForCatalogVariants(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	httpmock.RegisterResponder("PUT", productURL(cli, "company_location_catalog/variants_pricing.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &product.SeparatePricingForCatalogVariantsAPIReq{CatalogId: "6598087148213149672", ProductId: "16057850264845250791280282"}
	apiResp := &product.SeparatePricingForCatalogVariantsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestSeparatePricingForCatalogVariants_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&product.SeparatePricingForCatalogVariantsAPIReq{}).Verify(), "CatalogId is required")
	assert.EqualError(t, (&product.SeparatePricingForCatalogVariantsAPIReq{CatalogId: "x"}).Verify(), "ProductId is required")
}

func TestSpecifyVariantForPricingRules(t *testing.T) {
	client.SetupWithVersion(ApiVersion)
	defer client.Teardown()
	cli := client.GetClient()

	catalogId := "6598087148213149672"
	variantId := "18064649853096752677070520"
	mockResp := `{"variant_id":"18064649853096752677070520"}`
	httpmock.RegisterResponder("GET", productURL(cli, fmt.Sprintf("company_location_catalog/%s/%s.json", catalogId, variantId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &product.SpecifyVariantForPricingRulesAPIReq{CatalogId: catalogId, VariantId: variantId}
	apiResp := &product.SpecifyVariantForPricingRulesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, variantId, apiResp.VariantId)
}

func TestSpecifyVariantForPricingRules_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&product.SpecifyVariantForPricingRulesAPIReq{}).Verify(), "CatalogId is required")
	assert.EqualError(t, (&product.SpecifyVariantForPricingRulesAPIReq{CatalogId: "x"}).Verify(), "VariantId is required")
}
