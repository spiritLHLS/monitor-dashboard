package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/client/api"
	"github.com/gin-gonic/gin"
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
