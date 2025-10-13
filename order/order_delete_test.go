package order

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"testing"
)

//func TestDeleteOrderCase3(t *testing.T) {
//
//	apiReq := &DeleteOrderAPIReq{
//		OrderId: "21071580496925210798359834",
//	}
//
//	c := manager.GetDefaultClient()
//
//	_, err := DeleteOrder(c, apiReq)
//	if err != nil {
//		t.Errorf("DeleteOrder returned an error %v", err)
//	} else {
//		fmt.Printf("Delete order successful！orderID: %s\n", apiReq.OrderId)
//	}
//
//	a := assert.New(t)
//	a.Nil(err)
//}

func TestOrderDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(client.MethodDelete,
		fmt.Sprintf("https://%s.myshopline.com/%s/%s/orders/123.json", cli.StoreHandle, cli.PathPrefix, cli.ApiVersion),
		httpmock.NewStringResponder(200, ""))

	apiReq := &DeleteOrderAPIReq{
		OrderId: "123",
	}

	_, err := GetOrderService().Delete(context.Background(), apiReq)
	if err != nil {
		t.Errorf("DeleteOrder returned an error %v", err)
	}

}
