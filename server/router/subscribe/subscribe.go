package subscribe

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SubscribeRouter struct{}

func (s *SubscribeRouter) InitSubscribeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	subRouter := Router.Group("sub").Use(middleware.OperationRecord())
	subRouterWithoutRecord := Router.Group("sub")
	subRouterWithoutAuth := PublicRouter.Group("sub")
	{
		subRouter.POST("createSubscribe", subApi.CreateSubscribe)
		subRouter.DELETE("deleteSubscribe", subApi.DeleteSubscribe)
		subRouter.DELETE("deleteSubscribeByIds", subApi.DeleteSubscribeByIds)
		subRouter.PUT("updateSubscribe", subApi.UpdateSubscribe)
	}
	{
		subRouter.GET("selfGetSub", subApi.SelfGetSub)
		subRouter.POST("selfCreateSub", subApi.SelfCreateSub)
		subRouter.POST("selfDeleteSub", subApi.SelfDeleteSub)
		subRouter.POST("selfUpdateSub", subApi.SelfUpdateSub)
		subRouter.POST("selfBatchUpdateStatus", subApi.SelfBatchUpdateStatus)
	}
	{
		subRouterWithoutRecord.GET("findSubscribe", subApi.FindSubscribe)
		subRouterWithoutRecord.GET("getSubscribeList", subApi.GetSubscribeList)
	}
	{
		subRouterWithoutAuth.GET("selfGetAllPd", subApi.SelfGetAllPd)
	}
}
