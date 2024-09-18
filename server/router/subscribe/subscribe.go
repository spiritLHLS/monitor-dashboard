package subscribe

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SubscribeRouter struct {}

// InitSubscribeRouter 初始化 订阅 路由信息
func (s *SubscribeRouter) InitSubscribeRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	subRouter := Router.Group("sub").Use(middleware.OperationRecord())
	subRouterWithoutRecord := Router.Group("sub")
	subRouterWithoutAuth := PublicRouter.Group("sub")
	{
		subRouter.POST("createSubscribe", subApi.CreateSubscribe)   // 新建订阅
		subRouter.DELETE("deleteSubscribe", subApi.DeleteSubscribe) // 删除订阅
		subRouter.DELETE("deleteSubscribeByIds", subApi.DeleteSubscribeByIds) // 批量删除订阅
		subRouter.PUT("updateSubscribe", subApi.UpdateSubscribe)    // 更新订阅
	}
	{
		subRouterWithoutRecord.GET("findSubscribe", subApi.FindSubscribe)        // 根据ID获取订阅
		subRouterWithoutRecord.GET("getSubscribeList", subApi.GetSubscribeList)  // 获取订阅列表
	}
	{
	    subRouterWithoutAuth.GET("getSubscribePublic", subApi.GetSubscribePublic)  // 订阅开放接口
	}
}
