package metafield

import (
	"context"
	"github.com/shoplineos/shopline-sdk-go/client"
)

// 中文：https://developer.shopline.com/zh-hans-cn/docs/admin-rest-api/shopline-metafields/metafield-definition/create-a-metafield-definition?version=v20251201
// En：https://developer.shopline.com/docs/admin-rest-api/shopline-metafields/metafield-definition/create-a-metafield-definition?version=v20251201
type IMetafieldService interface {
	// todo
	List(context.Context)
	ListAll(context.Context)
	Detail(context.Context)
	Update(context.Context)
	Delete(context.Context)
	Count(context.Context)
	Create(context.Context, *CreateMetafieldDefinitionAPIReq) (*CreateMetafieldDefinitionAPIResp, error)
}

var metafieldServiceInst = &MetafieldService{}

type MetafieldService struct {
	client.BaseService
}

func GetMetafieldService() *MetafieldService {
	return metafieldServiceInst
}
