package invite_codes

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InviteCodesRouter struct {}

// InitInviteCodesRouter 初始化 邀请码 路由信息
func (s *InviteCodesRouter) InitInviteCodesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	invcodeRouter := Router.Group("invcode").Use(middleware.OperationRecord())
	invcodeRouterWithoutRecord := Router.Group("invcode")
	invcodeRouterWithoutAuth := PublicRouter.Group("invcode")
	{
		invcodeRouter.POST("createInviteCodes", invcodeApi.CreateInviteCodes)   // 新建邀请码
		invcodeRouter.DELETE("deleteInviteCodes", invcodeApi.DeleteInviteCodes) // 删除邀请码
		invcodeRouter.DELETE("deleteInviteCodesByIds", invcodeApi.DeleteInviteCodesByIds) // 批量删除邀请码
		invcodeRouter.PUT("updateInviteCodes", invcodeApi.UpdateInviteCodes)    // 更新邀请码
	}
	{
		invcodeRouterWithoutRecord.GET("findInviteCodes", invcodeApi.FindInviteCodes)        // 根据ID获取邀请码
		invcodeRouterWithoutRecord.GET("getInviteCodesList", invcodeApi.GetInviteCodesList)  // 获取邀请码列表
	}
	{
	    invcodeRouterWithoutAuth.GET("getInviteCodesPublic", invcodeApi.GetInviteCodesPublic)  // 邀请码开放接口
	}
}
