package privmsg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PusherConfigRouter struct {}

// InitPusherConfigRouter 初始化 推送配置 路由信息
func (s *PusherConfigRouter) InitPusherConfigRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	pcRouter := Router.Group("pc").Use(middleware.OperationRecord())
	pcRouterWithoutRecord := Router.Group("pc")
	pcRouterWithoutAuth := PublicRouter.Group("pc")
	{
		pcRouter.POST("createPusherConfig", pcApi.CreatePusherConfig)   // 新建推送配置
		pcRouter.DELETE("deletePusherConfig", pcApi.DeletePusherConfig) // 删除推送配置
		pcRouter.DELETE("deletePusherConfigByIds", pcApi.DeletePusherConfigByIds) // 批量删除推送配置
		pcRouter.PUT("updatePusherConfig", pcApi.UpdatePusherConfig)    // 更新推送配置
	}
	{
		pcRouterWithoutRecord.GET("findPusherConfig", pcApi.FindPusherConfig)        // 根据ID获取推送配置
		pcRouterWithoutRecord.GET("getPusherConfigList", pcApi.GetPusherConfigList)  // 获取推送配置列表
	}
	{
	    pcRouterWithoutAuth.GET("getPusherConfigPublic", pcApi.GetPusherConfigPublic)  // 推送配置开放接口
	}
}
