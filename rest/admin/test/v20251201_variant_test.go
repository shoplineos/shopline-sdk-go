package test

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	product2 "github.com/shoplineos/shopline-sdk-go/rest/admin/v20251201/product"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestDeleteVariant(t *testing.T) {
	client.Setup()
	defer client.Teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://%s.myshopline.com/%s/%s/products/1/variants/1.json", client.GetClient().StoreHandle, client.GetClient().PathPrefix, client.GetClient().ApiVersion),
		httpmock.NewStringResponder(200, ""))

	productId := "1"
	apiReq := &product2.DeleteAVariantAPIReq{
		ProductId: productId,
		VariantId: "1",
	}

	apiResp := apiReq.NewAPIResp()
	err := client.GetClient().Call(context.Background(), apiReq, apiResp)

	if err != nil {
		log.Printf("delete variant error, apiResp: %v, err:%v", apiResp, err)
	} else {
		log.Printf("delete variant success")
	}

	assert.Nil(t, err, "err should be nil")
}
