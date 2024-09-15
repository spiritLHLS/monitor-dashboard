package findallpd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FindallpdRouter struct {}

// InitFindallpdRouter 初始化 findallpd表 路由信息
func (s *FindallpdRouter) InitFindallpdRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	fapdRouter := Router.Group("fapd").Use(middleware.OperationRecord())
	fapdRouterWithoutRecord := Router.Group("fapd")
	fapdRouterWithoutAuth := PublicRouter.Group("fapd")
	{
		fapdRouter.POST("createFindallpd", fapdApi.CreateFindallpd)   // 新建findallpd表
		fapdRouter.DELETE("deleteFindallpd", fapdApi.DeleteFindallpd) // 删除findallpd表
		fapdRouter.DELETE("deleteFindallpdByIds", fapdApi.DeleteFindallpdByIds) // 批量删除findallpd表
		fapdRouter.PUT("updateFindallpd", fapdApi.UpdateFindallpd)    // 更新findallpd表
	}
	{
		fapdRouterWithoutRecord.GET("findFindallpd", fapdApi.FindFindallpd)        // 根据ID获取findallpd表
		fapdRouterWithoutRecord.GET("getFindallpdList", fapdApi.GetFindallpdList)  // 获取findallpd表列表
	}
	{
	    fapdRouterWithoutAuth.GET("getFindallpdPublic", fapdApi.GetFindallpdPublic)  // findallpd表开放接口
	}
}
