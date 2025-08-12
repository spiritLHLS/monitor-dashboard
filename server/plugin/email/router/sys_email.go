package router

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
	"server/plugin/email/api"
)

type EmailRouter struct{}

func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Use(middleware.OperationRecord())
	EmailApi := api.ApiGroupApp.EmailApi.EmailTest
	SendEmail := api.ApiGroupApp.EmailApi.SendEmail
	CheckEmail := api.ApiGroupApp.EmailApi.CheckEmail
	{
		emailRouter.POST("emailTest", EmailApi)    // 发送测试邮件
		emailRouter.POST("sendEmail", SendEmail)   // 发送邮件
		emailRouter.POST("checkEmail", CheckEmail) // 测试邮件一对一推送
	}
}
