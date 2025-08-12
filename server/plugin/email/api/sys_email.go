package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	email_response "server/plugin/email/model/response"
	"server/plugin/email/service"
	"server/utils"
)

type EmailApi struct{}

// EmailTest
// @Tags      System
// @Summary   发送测试邮件
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {string}  string  "{"success":true,"data":{},"msg":"发送成功"}"
// @Router    /email/emailTest [post]
func (s *EmailApi) EmailTest(c *gin.Context) {
	err := service.ServiceGroupApp.EmailTest()
	if err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
		return
	}
	response.OkWithMessage("发送成功", c)
}

// SendEmail
// @Tags      System
// @Summary   发送邮件
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      email_response.Email  true  "发送邮件必须的参数"
// @Success   200   {string}  string                "{"success":true,"data":{},"msg":"发送成功"}"
// @Router    /email/sendEmail [post]
func (s *EmailApi) SendEmail(c *gin.Context) {
	var email email_response.Email
	err := c.ShouldBindJSON(&email)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.SendEmailWithSMTPDetails("", 0, "", "",
		email.To, email.Subject, email.Body)
	if err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
		return
	}
	response.OkWithMessage("发送成功", c)
}

// CheckEmail 检查并发送邮件的接口
// @Tags      System
// @Summary   测试邮件推送是否成功
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      email_response.Email  true  "发送邮件必须的参数"
// @Success   200   {string}  string                "{"success":true,"data":{},"msg":"发送成功"}"
// @Router    /email/checkEmail [post]
func (s *EmailApi) CheckEmail(c *gin.Context) {
	clientIP := c.ClientIP()
	if clientIP == "" {
		response.FailWithMessage("检测失败，检测不到用户IP", c)
	}
	uuid := utils.GetUserUuid(c)
	result := service.ServiceGroupApp.CheckEmailLogic(clientIP, uuid)
	if result.Success {
		response.OkWithMessage(result.Message, c)
	} else {
		response.FailWithMessage(result.Message, c)
	}
}
