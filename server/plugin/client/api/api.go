package api

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/client/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/client/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type RegisterApi struct{}

// @Tags code
// @Summary 发送code
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /client/code [post]
func (p *RegisterApi) Code(c *gin.Context) {
	var req model.CodeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("获取tg_id失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.ServiceGroupApp.Code(req.TgId); err != nil {
		global.GVA_LOG.Error("发送Code失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed("", "发送Code成功", c)
	}
}

// @Tags Register
// @Summary 注册用户
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /client/client [post]
func (p *RegisterApi) Register(c *gin.Context) {
	var req model.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("绑定JSON失败", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("绑定JSON失败: %s", err.Error()), c)
		return
	}
	if token, claims, user, err := service.ServiceGroupApp.Register(req); err != nil {
		global.GVA_LOG.Error("注册用户失败", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("注册用户失败: %s", err.Error()), c)
	} else {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

// @Tags RegisterWithInvite
// @Summary 通过邀请码注册用户
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /client/client [post]
func (p *RegisterApi) RegisterWithInvite(c *gin.Context) {
	var req model.RegisterWithInviteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("绑定JSON失败", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("绑定JSON失败: %s", err.Error()), c)
		return
	}
	if token, claims, user, err := service.ServiceGroupApp.RegisterWithInvite(req); err != nil {
		global.GVA_LOG.Error("注册用户失败", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("注册用户失败: %s", err.Error()), c)
	} else {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

// @Tags ChangePassword
// @Summary 修改密码
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /client/changePassword [post]
func (p *RegisterApi) ChangePassword(c *gin.Context) {
	var req model.ChangePasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("绑定JSON失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		global.GVA_LOG.Error("密码输入为空", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	}
	if err := service.ServiceGroupApp.ChangePassword(req); err != nil {
		global.GVA_LOG.Error("修改密码失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("修改密码成功，请回到登录页面登录", c)
	}
}

// @Tags Login
// @Summary 用户登录
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /client/login [post]
func (p *RegisterApi) Login(c *gin.Context) {
	var req model.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("绑定JSON失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	key := c.ClientIP()
	if token, claims, user, err := service.ServiceGroupApp.Login(req, key); err != nil {
		global.GVA_LOG.Error("用户登录失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		maxAge := int(claims.RegisteredClaims.ExpiresAt.Unix() - time.Now().Unix())
		utils.SetToken(c, token, maxAge)
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}
