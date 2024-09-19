package ecsusers

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EcsUsersRouter struct{}

func (s *EcsUsersRouter) InitEcsUsersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	eusrRouter := Router.Group("eusr").Use(middleware.OperationRecord())
	eusrRouterWithoutRecord := Router.Group("eusr")
	eusrRouterWithoutAuth := PublicRouter.Group("eusr")
	{
		eusrRouter.POST("createEcsUsers", eusrApi.CreateEcsUsers)
		eusrRouter.DELETE("deleteEcsUsers", eusrApi.DeleteEcsUsers)
		eusrRouter.DELETE("deleteEcsUsersByIds", eusrApi.DeleteEcsUsersByIds)
		eusrRouter.PUT("updateEcsUsers", eusrApi.UpdateEcsUsers)
	}
	{
		eusrRouterWithoutRecord.GET("findEcsUsers", eusrApi.FindEcsUsers)
		eusrRouterWithoutRecord.GET("getEcsUsersList", eusrApi.GetEcsUsersList)
	}
	{
		eusrRouterWithoutAuth.GET("getEcsUsersPublic", eusrApi.GetEcsUsersPublic)
		eusrRouterWithoutAuth.PUT("adminChangePassword", eusrApi.AdminChangePassword)
		eusrRouterWithoutAuth.GET("selfGetUserInfo", eusrApi.SelfGetUserInfo)
		eusrRouterWithoutAuth.POST("selfModifyInfo", eusrApi.SelfModifyInfo)
	}
}
