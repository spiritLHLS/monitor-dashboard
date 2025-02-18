package privmsg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PusherConfigRouter struct{}

func (s *PusherConfigRouter) InitPusherConfigRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	pcRouter := Router.Group("pc").Use(middleware.OperationRecord())
	pcRouterWithoutRecord := Router.Group("pc")
	pcRouterWithoutAuth := PublicRouter.Group("pc")
	{
		pcRouter.POST("createPusherConfig", pcApi.CreatePusherConfig)
		pcRouter.DELETE("deletePusherConfig", pcApi.DeletePusherConfig)
		pcRouter.DELETE("deletePusherConfigByIds", pcApi.DeletePusherConfigByIds)
		pcRouter.PUT("updatePusherConfig", pcApi.UpdatePusherConfig)

		pcRouter.GET("getPublicPushStatus", pcApi.GetPublicPushStatus)
		pcRouter.POST("controlPublicPushStatus", pcApi.ControlPublicPushStatus)
		pcRouter.GET("getTelegramBotPushStatus", pcApi.GetTelegramBotPushStatus)
		pcRouter.POST("controlTelegramBotPushStatus", pcApi.ControlTelegramBotPushStatus)
		pcRouter.GET("getEmailPushStatus", pcApi.GetEmailPushStatus)
		pcRouter.POST("controlEmailPushStatus", pcApi.ControlEmailPushStatus)
	}
	{
		pcRouterWithoutRecord.GET("findPusherConfig", pcApi.FindPusherConfig)
		pcRouterWithoutRecord.GET("getPusherConfigList", pcApi.GetPusherConfigList)
	}
	{
		pcRouterWithoutAuth.GET("getPusherConfigPublic", pcApi.GetPusherConfigPublic)
	}
}
