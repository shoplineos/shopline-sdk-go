package metafield

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// DetailMetafieldDefinitionAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/get-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/get-a-metafield-definition?version=v20251201
type DetailMetafieldDefinitionAPIReq struct {
	ID string
}

func (d DetailMetafieldDefinitionAPIReq) Verify() error {
	if d.ID == "" {
		return errors.New("MetafieldDefinition ID is empty")
	}
	return nil
}

func (d DetailMetafieldDefinitionAPIReq) Endpoint() string {
	return "metafield_definition.json"
}

type DetailMetafieldDefinitionAPIResp struct {
	MetafieldDefinition MetafieldDefinition `json:"definition,omitempty"`
	client.BaseAPIResponse
}
