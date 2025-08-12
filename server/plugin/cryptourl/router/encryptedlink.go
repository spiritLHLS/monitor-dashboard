package router

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

var EncryptedLink = new(EL)

type EL struct{}

func (r *EL) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("EL").Use(middleware.OperationRecord())
		group.POST("createEncryptedLink", apiEncryptedLink.CreateEncryptedLink)
		group.DELETE("deleteEncryptedLink", apiEncryptedLink.DeleteEncryptedLink)
		group.DELETE("deleteEncryptedLinkByIds", apiEncryptedLink.DeleteEncryptedLinkByIds)
		group.PUT("updateEncryptedLink", apiEncryptedLink.UpdateEncryptedLink)
	}
	{
		group := private.Group("EL")
		group.GET("findEncryptedLink", apiEncryptedLink.FindEncryptedLink)
		group.GET("getEncryptedLinkList", apiEncryptedLink.GetEncryptedLinkList)
		group.POST("buildEncryptedUrl", apiEncryptedLink.BuildEncryptedUrl)
	}
	{
		group := public.Group("EL")
		group.GET("handleRedirect", apiEncryptedLink.HandleRedirect)
		group.GET("getEncryptedLinkPublic", apiEncryptedLink.GetEncryptedLinkPublic) // TODO 未实现
	}
}
