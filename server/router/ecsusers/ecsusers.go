package ecsusers

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EcsUsersRouter struct {}

// InitEcsUsersRouter 初始化 订阅用户 路由信息
func (s *EcsUsersRouter) InitEcsUsersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	eusrRouter := Router.Group("eusr").Use(middleware.OperationRecord())
	eusrRouterWithoutRecord := Router.Group("eusr")
	eusrRouterWithoutAuth := PublicRouter.Group("eusr")
	{
		eusrRouter.POST("createEcsUsers", eusrApi.CreateEcsUsers)   // 新建订阅用户
		eusrRouter.DELETE("deleteEcsUsers", eusrApi.DeleteEcsUsers) // 删除订阅用户
		eusrRouter.DELETE("deleteEcsUsersByIds", eusrApi.DeleteEcsUsersByIds) // 批量删除订阅用户
		eusrRouter.PUT("updateEcsUsers", eusrApi.UpdateEcsUsers)    // 更新订阅用户
	}
	{
		eusrRouterWithoutRecord.GET("findEcsUsers", eusrApi.FindEcsUsers)        // 根据ID获取订阅用户
		eusrRouterWithoutRecord.GET("getEcsUsersList", eusrApi.GetEcsUsersList)  // 获取订阅用户列表
	}
	{
	    eusrRouterWithoutAuth.GET("getEcsUsersPublic", eusrApi.GetEcsUsersPublic)  // 订阅用户开放接口
	}
}
