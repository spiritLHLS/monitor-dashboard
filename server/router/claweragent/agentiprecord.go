package claweragent

import (
	"server/middleware"
	"github.com/gin-gonic/gin"
)

type AgentIpRecordRouter struct {}

// InitAgentIpRecordRouter 初始化 IP记录 路由信息
func (s *AgentIpRecordRouter) InitAgentIpRecordRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	airRouter := Router.Group("air").Use(middleware.OperationRecord())
	airRouterWithoutRecord := Router.Group("air")
	airRouterWithoutAuth := PublicRouter.Group("air")
	{
		airRouter.POST("createAgentIpRecord", airApi.CreateAgentIpRecord)   // 新建IP记录
		airRouter.DELETE("deleteAgentIpRecord", airApi.DeleteAgentIpRecord) // 删除IP记录
		airRouter.DELETE("deleteAgentIpRecordByIds", airApi.DeleteAgentIpRecordByIds) // 批量删除IP记录
		airRouter.PUT("updateAgentIpRecord", airApi.UpdateAgentIpRecord)    // 更新IP记录
	}
	{
		airRouterWithoutRecord.GET("findAgentIpRecord", airApi.FindAgentIpRecord)        // 根据ID获取IP记录
		airRouterWithoutRecord.GET("getAgentIpRecordList", airApi.GetAgentIpRecordList)  // 获取IP记录列表
	}
	{
	    airRouterWithoutAuth.GET("getAgentIpRecordPublic", airApi.GetAgentIpRecordPublic)  // IP记录开放接口
	}
}
