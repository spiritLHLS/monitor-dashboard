package partitions

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

type PartitionspageRouter struct{}

// InitPartitionspageRouter 初始化 partitionspage表 路由信息
func (s *PartitionspageRouter) InitPartitionspageRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	pspRouter := Router.Group("psp").Use(middleware.OperationRecord())
	pspRouterWithoutRecord := Router.Group("psp")
	pspRouterWithoutAuth := PublicRouter.Group("psp")
	{
		pspRouter.POST("createPartitionspage", pspApi.CreatePartitionspage)             // 新建partitionspage表
		pspRouter.DELETE("deletePartitionspage", pspApi.DeletePartitionspage)           // 删除partitionspage表
		pspRouter.DELETE("deletePartitionspageByIds", pspApi.DeletePartitionspageByIds) // 批量删除partitionspage表
		pspRouter.PUT("updatePartitionspage", pspApi.UpdatePartitionspage)              // 更新partitionspage表
	}
	{
		pspRouterWithoutRecord.GET("findPartitionspage", pspApi.FindPartitionspage)       // 根据ID获取partitionspage表
		pspRouterWithoutRecord.GET("getPartitionspageList", pspApi.GetPartitionspageList) // 获取partitionspage表列表
	}
	{
		pspRouterWithoutAuth.GET("getPartitionspagePublic", pspApi.GetPartitionspagePublic) // partitionspage表开放接口
	}
}
