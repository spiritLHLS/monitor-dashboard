package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var EncryptedLink = new(EL)

type EL struct {}

// Init 初始化 加密链接 路由信息
func (r *EL) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("EL").Use(middleware.OperationRecord())
		group.POST("createEncryptedLink", apiEncryptedLink.CreateEncryptedLink)   // 新建加密链接
		group.DELETE("deleteEncryptedLink", apiEncryptedLink.DeleteEncryptedLink) // 删除加密链接
		group.DELETE("deleteEncryptedLinkByIds", apiEncryptedLink.DeleteEncryptedLinkByIds) // 批量删除加密链接
		group.PUT("updateEncryptedLink", apiEncryptedLink.UpdateEncryptedLink)    // 更新加密链接
	}
	{
	    group := private.Group("EL")
		group.GET("findEncryptedLink", apiEncryptedLink.FindEncryptedLink)        // 根据ID获取加密链接
		group.GET("getEncryptedLinkList", apiEncryptedLink.GetEncryptedLinkList)  // 获取加密链接列表
	}
	{
	    group := public.Group("EL")
	    group.GET("getEncryptedLinkPublic", apiEncryptedLink.GetEncryptedLinkPublic)  // 加密链接开放接口
	}
}
