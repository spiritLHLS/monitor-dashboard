package tgchannel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TgchannelRouter struct {}

// InitTgchannelRouter 初始化 tgchannel表 路由信息
func (s *TgchannelRouter) InitTgchannelRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	tgcRouter := Router.Group("tgc").Use(middleware.OperationRecord())
	tgcRouterWithoutRecord := Router.Group("tgc")
	tgcRouterWithoutAuth := PublicRouter.Group("tgc")
	{
		tgcRouter.POST("createTgchannel", tgcApi.CreateTgchannel)   // 新建tgchannel表
		tgcRouter.DELETE("deleteTgchannel", tgcApi.DeleteTgchannel) // 删除tgchannel表
		tgcRouter.DELETE("deleteTgchannelByIds", tgcApi.DeleteTgchannelByIds) // 批量删除tgchannel表
		tgcRouter.PUT("updateTgchannel", tgcApi.UpdateTgchannel)    // 更新tgchannel表
	}
	{
		tgcRouterWithoutRecord.GET("findTgchannel", tgcApi.FindTgchannel)        // 根据ID获取tgchannel表
		tgcRouterWithoutRecord.GET("getTgchannelList", tgcApi.GetTgchannelList)  // 获取tgchannel表列表
	}
	{
	    tgcRouterWithoutAuth.GET("getTgchannelPublic", tgcApi.GetTgchannelPublic)  // tgchannel表开放接口
	}
}
