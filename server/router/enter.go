package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/partitions"
	"github.com/flipped-aurora/gin-vue-admin/server/router/products"
	"github.com/flipped-aurora/gin-vue-admin/server/router/shops"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Products   products.RouterGroup
	Shops      shops.RouterGroup
	Partitions partitions.RouterGroup
}
