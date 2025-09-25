package product

import (
	"github.com/shoplineos/shopline-sdk-go/manager"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestDeleteProduct(t *testing.T) {

	productId := "16070828389930044390453380"
	apiReq := &DeleteProductAPIReq{
		ProductId: productId,
	}

	c := manager.GetDefaultClient()

	apiResp, err := DeleteProduct(c, apiReq)
	if err != nil {
		log.Printf("delete product error, apiResp: %v, err:%v", apiResp, err)
	} else {
		log.Printf("delete product success")
	}

	assert.Nil(t, err, "err should be nil")
}

// when delete a product again, it's return success
func TestDeleteProductCase2(t *testing.T) {

	productId := "16070828389930044390453380"
	apiReq := &DeleteProductAPIReq{
		ProductId: productId,
	}

	c := manager.GetDefaultClient()

	apiResp, err := DeleteProduct(c, apiReq)
	if err != nil {
		log.Printf("delete product error, apiResp: %v, err:%v", apiResp, err)
	} else {
		log.Printf("delete product success")
	}

	assert.Nil(t, err)
}
