package users

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {}

// InitUserRouter 初始化 用户 路由信息
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	usrRouter := Router.Group("usr").Use(middleware.OperationRecord())
	usrRouterWithoutRecord := Router.Group("usr")
	usrRouterWithoutAuth := PublicRouter.Group("usr")
	{
		usrRouter.POST("createUser", usrApi.CreateUser)   // 新建用户
		usrRouter.DELETE("deleteUser", usrApi.DeleteUser) // 删除用户
		usrRouter.DELETE("deleteUserByIds", usrApi.DeleteUserByIds) // 批量删除用户
		usrRouter.PUT("updateUser", usrApi.UpdateUser)    // 更新用户
	}
	{
		usrRouterWithoutRecord.GET("findUser", usrApi.FindUser)        // 根据ID获取用户
		usrRouterWithoutRecord.GET("getUserList", usrApi.GetUserList)  // 获取用户列表
	}
	{
	    usrRouterWithoutAuth.GET("getUserPublic", usrApi.GetUserPublic)  // 用户开放接口
	}
}
