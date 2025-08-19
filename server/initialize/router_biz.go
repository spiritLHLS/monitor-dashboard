package initialize

import (
	"github.com/gin-gonic/gin"
	"server/router"
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
	}
	{
		ecsusersRouter := router.RouterGroupApp.Ecsusers
		ecsusersRouter.InitEcsUsersRouter(privateGroup, publicGroup)
	}
	{
		subscribeRouter := router.RouterGroupApp.Subscribe
		subscribeRouter.InitSubscribeRouter(privateGroup, publicGroup)
	}
	{
		privmsgRouter := router.RouterGroupApp.Privmsg
		privmsgRouter.InitPusherConfigRouter(privateGroup, publicGroup)
	}
	{
		invite_codesRouter := router.RouterGroupApp.Invite_codes
		invite_codesRouter.InitInviteCodesRouter(privateGroup, publicGroup)
	}
	{
		digitalproductsRouter := router.RouterGroupApp.Digitalproducts
		digitalproductsRouter.InitDigitalProductsRouter(privateGroup, publicGroup)
	} // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
	{
		claweragentRouter := router.RouterGroupApp.Claweragent
		claweragentRouter.InitAgentIpRecordRouter(privateGroup, publicGroup)
	}
}
