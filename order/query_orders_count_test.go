package order

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestQueryOrdersCount(t *testing.T) {

	apiReq := &GetOrdersCountAPIReq{
		Status: "any",
	}

	apiResp, err := QueryOrdersCount(apiReq)

	fmt.Printf("Count: %v\n", apiResp)
	if err != nil {
		log.Printf("Request failed, error: %v", err)
	}

	a := assert.New(t)
	a.Greater(apiResp.Count, 0)
}
