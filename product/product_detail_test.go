package product

import (
	"github.com/shoplineos/shopline-sdk-go/manager"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestProductDetail(t *testing.T) {
	//productId := "16070822412102455208483380"
	productId := "16070828389930044390453380"
	apiReq := &GetProductDetailAPIReq{
		ProductId: productId,
	}

	c := manager.GetDefaultClient()

	apiResp, err := GetProductDetail(c, apiReq)
	log.Printf("apiResp: %v, err:%v", apiResp, err)
	assert.Nil(t, err, "err should be nil")
}

// product DATA_NOT_EXIST
func TestGetProductDetailCase2(t *testing.T) {
	//productId := "16070822412102455208483380"
	productId := "16070828389930044390453380"
	apiReq := &GetProductDetailAPIReq{
		ProductId: productId,
	}

	c := manager.GetDefaultClient()

	_, err := GetProductDetail(c, apiReq)
	//log.Printf("apiResp: %v, err:%v", apiResp, err)
	if err != nil {
		log.Printf("get product detail error, apiResp: %v, err:%v", apiReq, err)
	}
	assert.NotNil(t, err)
}
