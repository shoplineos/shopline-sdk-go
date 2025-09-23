package order

import (
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/manager"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestQueryOrdersCount(t *testing.T) {

	apiReq := &GetOrdersCountAPIReq{
		Status: "any",
	}

	c := manager.GetDefaultClient()
	apiResp, err := QueryOrdersCount(c, apiReq)

	fmt.Printf("Count: %v\n", apiResp)
	if err != nil {
		log.Printf("Request failed, error: %v", err)
	}

	a := assert.New(t)
	a.Greater(apiResp.Count, 0)
}
