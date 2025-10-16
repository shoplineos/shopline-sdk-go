package metafield

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// DeleteMetafieldDefinitionAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/delete-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/delete-a-metafield-definition?version=v20251201
type DeleteMetafieldDefinitionAPIReq struct {
	client.BaseAPIRequest
	Id                            string
	DeleteAllAssociatedMetafields bool `json:"delete_all_associated_metafields,omitempty"`
}

func (d *DeleteMetafieldDefinitionAPIReq) Method() string {
	return "DELETE"
}

func (c *DeleteMetafieldDefinitionAPIReq) GetData() interface{} {
	return c
}

func (d *DeleteMetafieldDefinitionAPIReq) Verify() error {
	if d.Id == "" {
		return errors.New("MetafieldDefinition Id is empty")
	}
	return nil
}

func (d *DeleteMetafieldDefinitionAPIReq) Endpoint() string {
	return "metafield_definition.json"
}

type DeleteMetafieldDefinitionAPIResp struct {
	client.BaseAPIResponse
}
