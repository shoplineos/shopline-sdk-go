package metafield

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// GetMetafieldDefinitionAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/get-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/get-a-metafield-definition?version=v20251201
type GetMetafieldDefinitionAPIReq struct {
	client.BaseAPIRequest
	Id string
}

func (d *GetMetafieldDefinitionAPIReq) Method() string {
	return "GET"
}

func (d *GetMetafieldDefinitionAPIReq) GetQuery() interface{} {
	return d
}

func (d *GetMetafieldDefinitionAPIReq) Verify() error {
	if d.Id == "" {
		return errors.New("MetafieldDefinition Id is empty")
	}
	return nil
}

func (d *GetMetafieldDefinitionAPIReq) Endpoint() string {
	return "metafield_definition.json"
}

type GetMetafieldDefinitionAPIResp struct {
	client.BaseAPIResponse
	MetafieldDefinition MetafieldDefinition `json:"definition,omitempty"`
}
