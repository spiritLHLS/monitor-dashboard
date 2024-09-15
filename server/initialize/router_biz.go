package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		productsRouter := router.RouterGroupApp.Products
		productsRouter.InitProductsRouter(privateGroup, publicGroup)
	}
	{
		shopsRouter := router.RouterGroupApp.Shops
		shopsRouter.InitShopsRouter(privateGroup, publicGroup)
	}
	{
		partitionsRouter := router.RouterGroupApp.Partitions
		partitionsRouter.InitPartitionspageRouter(privateGroup, publicGroup)
	}
	{
		tgchannelRouter := router.RouterGroupApp.Tgchannel
		tgchannelRouter.InitTgchannelRouter(privateGroup, publicGroup)
	}
	{
		findallpdRouter := router.RouterGroupApp.Findallpd
		findallpdRouter.InitFindallpdRouter(privateGroup, publicGroup)
	} // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
	{
		ecsusersRouter := router.RouterGroupApp.Ecsusers
		ecsusersRouter.InitEcsUsersRouter(privateGroup, publicGroup)
	}
}
