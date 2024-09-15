package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/partitions"
	"github.com/flipped-aurora/gin-vue-admin/server/service/products"
	"github.com/flipped-aurora/gin-vue-admin/server/service/shops"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	ProductsServiceGroup   products.ServiceGroup
	ShopsServiceGroup      shops.ServiceGroup
	PartitionsServiceGroup partitions.ServiceGroup
}
