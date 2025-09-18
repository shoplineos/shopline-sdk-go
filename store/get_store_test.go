package store

import (
	"log"
	"testing"
)

func TestGetStoreInfo(t *testing.T) {

	apiReq := &GetStoreAPIReq{}
	storeInfo, err := GetStoreInfo(apiReq)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("storeInfo:%v\n", storeInfo)
}
