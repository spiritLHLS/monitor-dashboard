package router

import (
	"github.com/gin-gonic/gin"
	"server/plugin/client/api"
)

type RegisterRouter struct {
}

func (s *RegisterRouter) InitRegisterRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.RegisterApi
	{
		plugRouter.POST("code", plugApi.Code)
		plugRouter.POST("register", plugApi.Register)
		plugRouter.POST("registerWithInvite", plugApi.RegisterWithInvite)
		plugRouter.POST("changePassword", plugApi.ChangePassword)
		plugRouter.POST("login", plugApi.Login)
	}
}
