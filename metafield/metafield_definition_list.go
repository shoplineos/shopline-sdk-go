package metafield

import (
	"errors"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListMetafieldDefinitionAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/get-metafield-definitions?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/get-metafield-definitions?version=v20251201
type ListMetafieldDefinitionAPIReq struct {
	client.BaseAPIRequest
	UpdatedAtMin string `url:"updated_at_min,omitempty"` // Updated time,eg: 2021-04-25T16:16:47+04:00
	UpdatedAtMax string `url:"updated_at_max,omitempty"` // Updated time
	CreatedAtMin string `url:"created_at_min,omitempty"` // Updated time
	CreatedAtMax string `url:"created_at_max,omitempty"` // Updated time

	AccessAdmin     string `url:"access_admin,omitempty"`
	DefinitionState string `url:"definition_state,omitempty"`
	Key             string `url:"key,omitempty"`
	OwnerResource   string `url:"owner_resource,omitempty"` // required
	Type            string `url:"type,omitempty"`

	Limit    string `url:"limit,omitempty"`
	SinceId  string `url:"since_id,omitempty"`
	PageInfo string `url:"page_info,omitempty"`
}

func (l *ListMetafieldDefinitionAPIReq) GetMethod() string {
	return "GET"
}

func (l *ListMetafieldDefinitionAPIReq) GetQuery() interface{} {
	return l
}

func (l *ListMetafieldDefinitionAPIReq) Verify() error {
	if l.OwnerResource == "" {
		return errors.New("owner_resource is required")
	}
	return nil
}

func (l *ListMetafieldDefinitionAPIReq) GetEndpoint() string {
	return "metafield_definitions.json"
}

type ListMetafieldDefinitionAPIResp struct {
	client.BaseAPIResponse
	Data ListMetafieldDefinitionAPIRespData `json:"data,omitempty"`
}

type ListMetafieldDefinitionAPIRespData struct {
	MetafieldDefinitions []MetafieldDefinition `json:"metafield_definitions,omitempty"`
}
