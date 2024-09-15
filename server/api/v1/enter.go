package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/partitions"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/products"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/shops"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/tgchannel"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	ProductsApiGroup   products.ApiGroup
	ShopsApiGroup      shops.ApiGroup
	PartitionsApiGroup partitions.ApiGroup
	TgchannelApiGroup  tgchannel.ApiGroup
}
