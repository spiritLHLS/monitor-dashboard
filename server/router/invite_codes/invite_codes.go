package invite_codes

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InviteCodesRouter struct{}

func (s *InviteCodesRouter) InitInviteCodesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	invcodeRouter := Router.Group("invcode").Use(middleware.OperationRecord())
	invcodeRouterWithoutRecord := Router.Group("invcode")
	invcodeRouterWithoutAuth := PublicRouter.Group("invcode")
	{
		invcodeRouter.POST("createInviteCodes", invcodeApi.CreateInviteCodes)
		invcodeRouter.DELETE("deleteInviteCodes", invcodeApi.DeleteInviteCodes)
		invcodeRouter.DELETE("deleteInviteCodesByIds", invcodeApi.DeleteInviteCodesByIds)
		invcodeRouter.PUT("updateInviteCodes", invcodeApi.UpdateInviteCodes)
	}
	{
		invcodeRouterWithoutRecord.GET("findInviteCodes", invcodeApi.FindInviteCodes)
		invcodeRouterWithoutRecord.GET("getInviteCodesList", invcodeApi.GetInviteCodesList)
	}
	{
		invcodeRouterWithoutAuth.GET("getInviteCodesPublic", invcodeApi.GetInviteCodesPublic)
		invcodeRouterWithoutAuth.POST("batchBuildCodes", invcodeApi.BatchBuildCodes)
		invcodeRouterWithoutAuth.POST("batchExportCodes", invcodeApi.BatchExportCodes)
	}
}
