package order

import (
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/manager"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 21071580496925210798359834
func TestCancelOrder(t *testing.T) {

	apiReq := &CancelOrderAPIReq{
		OrderId: "123",
	}

	c := manager.GetDefaultClient()

	apiResp, err := CancelOrder(c, apiReq)
	if err != nil {
		fmt.Println("Cancel order failed, err:", err)
	} else {
		fmt.Printf("Cancel order successful！orderID: %s\n", apiResp.Order.ID)
	}

	a := assert.New(t)
	a.NotNil(err)

}

// 21071580496925210798359834
func TestCancelOrderCase2(t *testing.T) {

	apiReq := &CancelOrderAPIReq{
		OrderId: "21071580496925210798359811",
	}

	c := manager.GetDefaultClient()

	apiResp, err := CancelOrder(c, apiReq)
	if err != nil {
		fmt.Println("Cancel order failed, err:", err)
	} else {
		fmt.Printf("Cancel order successful！orderID: %s\n", apiResp.Order.ID)
	}

	a := assert.New(t)
	a.NotNil(err)

}

// 21071580496925210798359834
func TestCancelOrderCase3(t *testing.T) {

	apiReq := &CancelOrderAPIReq{
		OrderId: "21071580496925210798359834",
	}

	c := manager.GetDefaultClient()

	apiResp, err := CancelOrder(c, apiReq)
	if err != nil {
		fmt.Println("Cancel order failed, err:", err)
	} else {
		fmt.Printf("Cancel order successful！orderID: %s\n", apiResp.Order.ID)
	}

	a := assert.New(t)
	a.NotNil(err)

}
