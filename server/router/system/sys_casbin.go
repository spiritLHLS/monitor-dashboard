package system

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinRouterWithoutRecord := Router.Group("casbin")
	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
	}
	{
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
