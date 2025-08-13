package digitalproducts

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

type DigitalProductsRouter struct{}

// InitDigitalProductsRouter 初始化 数字商品表 路由信息
func (s *DigitalProductsRouter) InitDigitalProductsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	dpdRouter := Router.Group("dpd").Use(middleware.OperationRecord())
	dpdRouterWithoutRecord := Router.Group("dpd")
	dpdRouterWithoutAuth := PublicRouter.Group("dpd")
	{
		dpdRouter.POST("createDigitalProducts", dpdApi.CreateDigitalProducts)             // 新建数字商品表
		dpdRouter.DELETE("deleteDigitalProducts", dpdApi.DeleteDigitalProducts)           // 删除数字商品表
		dpdRouter.DELETE("deleteDigitalProductsByIds", dpdApi.DeleteDigitalProductsByIds) // 批量删除数字商品表
		dpdRouter.PUT("updateDigitalProducts", dpdApi.UpdateDigitalProducts)              // 更新数字商品表
	}
	{
		dpdRouterWithoutRecord.GET("findDigitalProducts", dpdApi.FindDigitalProducts)           // 根据ID获取数字商品表
		dpdRouterWithoutRecord.GET("getDigitalProductsList", dpdApi.GetDigitalProductsList)     // 获取数字商品表列表
		dpdRouterWithoutRecord.GET("convertProductsToDigital", dpdApi.ConvertProductsToDigital) // 将products表数据转换为数字商品表
	}
	{
		dpdRouterWithoutAuth.GET("getDigitalProductsPublic", dpdApi.GetDigitalProductsPublic) // 数字商品表开放接口
	}
}
