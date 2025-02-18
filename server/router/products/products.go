package products

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductsRouter struct{}

func (s *ProductsRouter) InitProductsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	pdRouter := Router.Group("pd").Use(middleware.OperationRecord())
	pdRouterWithoutRecord := Router.Group("pd")
	pdRouterWithoutAuth := PublicRouter.Group("pd")
	{
		pdRouter.POST("createProducts", pdApi.CreateProducts)
		pdRouter.DELETE("deleteProducts", pdApi.DeleteProducts)
		pdRouter.DELETE("deleteProductsByIds", pdApi.DeleteProductsByIds)
		pdRouter.PUT("updateProducts", pdApi.UpdateProducts)
	}
	{
		pdRouterWithoutRecord.GET("findProducts", pdApi.FindProducts)
		pdRouterWithoutRecord.GET("getProductsList", pdApi.GetProductsList)
	}
	{
		pdRouterWithoutAuth.GET("getProductsPublic", pdApi.GetProductsPublic)
	}
}
