package products

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductsRouter struct {}

// InitProductsRouter 初始化 products表 路由信息
func (s *ProductsRouter) InitProductsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	pdRouter := Router.Group("pd").Use(middleware.OperationRecord())
	pdRouterWithoutRecord := Router.Group("pd")
	pdRouterWithoutAuth := PublicRouter.Group("pd")
	{
		pdRouter.POST("createProducts", pdApi.CreateProducts)   // 新建products表
		pdRouter.DELETE("deleteProducts", pdApi.DeleteProducts) // 删除products表
		pdRouter.DELETE("deleteProductsByIds", pdApi.DeleteProductsByIds) // 批量删除products表
		pdRouter.PUT("updateProducts", pdApi.UpdateProducts)    // 更新products表
	}
	{
		pdRouterWithoutRecord.GET("findProducts", pdApi.FindProducts)        // 根据ID获取products表
		pdRouterWithoutRecord.GET("getProductsList", pdApi.GetProductsList)  // 获取products表列表
	}
	{
	    pdRouterWithoutAuth.GET("getProductsPublic", pdApi.GetProductsPublic)  // products表开放接口
	}
}
