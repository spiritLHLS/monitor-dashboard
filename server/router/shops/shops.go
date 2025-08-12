package shops

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

type ShopsRouter struct{}

// InitShopsRouter 初始化 shops表 路由信息
func (s *ShopsRouter) InitShopsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	spRouter := Router.Group("sp").Use(middleware.OperationRecord())
	spRouterWithoutRecord := Router.Group("sp")
	spRouterWithoutAuth := PublicRouter.Group("sp")
	{
		spRouter.POST("createShops", spApi.CreateShops)             // 新建shops表
		spRouter.DELETE("deleteShops", spApi.DeleteShops)           // 删除shops表
		spRouter.DELETE("deleteShopsByIds", spApi.DeleteShopsByIds) // 批量删除shops表
		spRouter.PUT("updateShops", spApi.UpdateShops)              // 更新shops表
	}
	{
		spRouterWithoutRecord.GET("findShops", spApi.FindShops)       // 根据ID获取shops表
		spRouterWithoutRecord.GET("getShopsList", spApi.GetShopsList) // 获取shops表列表
	}
	{
		spRouterWithoutAuth.GET("getShopsPublic", spApi.GetShopsPublic) // shops表开放接口
	}
}
