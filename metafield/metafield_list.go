package metafield

import (
	"errors"
	"fmt"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// ListMetafieldAPIReq
// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafields/resource-metafields/get-a-metafield-for-a-resource?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafields/resource-metafields/get-a-metafield-for-a-resource?version=v20251201
type ListMetafieldAPIReq struct {
	client.BaseAPIRequest
	// enum：product、order、customer、collection、
	// shop、variant、draft_order
	// eg："product"
	OwnerResource string
	OwnerId       string

	UpdatedAtMin string `url:"updated_at_min,omitempty"` // Updated time,eg: 2021-04-25T16:16:47+04:00
	UpdatedAtMax string `url:"updated_at_max,omitempty"` // Updated time
	CreatedAtMin string `url:"created_at_min,omitempty"` // Updated time
	CreatedAtMax string `url:"created_at_max,omitempty"` // Updated time

	Key  string `url:"key,omitempty"`
	Type string `url:"type,omitempty"`

	Fields   string `url:"fields,omitempty"`
	Limit    string `url:"limit,omitempty"`
	SinceId  string `url:"since_id,omitempty"`
	PageInfo string `url:"page_info,omitempty"`
}

func (l *ListMetafieldAPIReq) Method() string {
	return "GET"
}

func (l *ListMetafieldAPIReq) GetQuery() interface{} {
	return l
}

func (l *ListMetafieldAPIReq) Verify() error {
	if l.OwnerResource == "" {
		return errors.New("owner_resource is required")
	}
	if l.OwnerId == "" {
		return errors.New("owner_id is required")
	}
	return nil
}

func (l *ListMetafieldAPIReq) Endpoint() string {
	return fmt.Sprintf("/%s/%s/metafields.json", l.OwnerResource, l.OwnerId)
}

type ListMetafieldAPIResp struct {
	client.BaseAPIResponse
	Metafields []Metafield `json:"metafields,omitempty"`
}
