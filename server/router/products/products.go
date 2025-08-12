package products

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
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
