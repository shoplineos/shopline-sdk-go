package store

import (
	"github.com/shoplineos/shopline-sdk-go/manager"
	"log"
	"testing"
)

func TestGetStoreInfo(t *testing.T) {

	c := manager.GetDefaultClient()

	apiReq := &GetStoreAPIReq{}
	storeInfo, err := GetStoreInfo(c, apiReq)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("storeInfo:%v\n", storeInfo)
}
