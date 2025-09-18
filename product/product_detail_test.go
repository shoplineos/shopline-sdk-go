package product

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestProductDetail(t *testing.T) {
	productId := "16070822412102455208483380"
	apiReq := &GetProductDetailAPIReq{
		ProductId: productId,
	}

	apiResp, err := GetProductDetail(apiReq)
	log.Printf("apiResp: %v, err:%v", apiResp, err)
	assert.Nil(t, err, "err should be nil")
}
