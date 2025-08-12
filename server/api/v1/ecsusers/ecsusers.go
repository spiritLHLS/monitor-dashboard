package ecsusers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/ecsusers"
	ecsusersReq "server/model/ecsusers/request"
	"server/plugin/client/model"
	"server/utils"
)

type EcsUsersApi struct{}

// CreateEcsUsers 创建订阅用户
// @Tags EcsUsers
// @Summary 创建订阅用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ecsusers.EcsUsers true "创建订阅用户"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /eusr/createEcsUsers [post]
func (eusrApi *EcsUsersApi) CreateEcsUsers(c *gin.Context) {
	var eusr ecsusers.EcsUsers
	err := c.ShouldBindJSON(&eusr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = eusrService.CreateEcsUsers(&eusr)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEcsUsers 删除订阅用户
// @Tags EcsUsers
// @Summary 删除订阅用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ecsusers.EcsUsers true "删除订阅用户"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /eusr/deleteEcsUsers [delete]
func (eusrApi *EcsUsersApi) DeleteEcsUsers(c *gin.Context) {
	ID := c.Query("ID")
	err := eusrService.DeleteEcsUsers(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEcsUsersByIds 批量删除订阅用户
// @Tags EcsUsers
// @Summary 批量删除订阅用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /eusr/deleteEcsUsersByIds [delete]
func (eusrApi *EcsUsersApi) DeleteEcsUsersByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := eusrService.DeleteEcsUsersByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEcsUsers 更新订阅用户
// @Tags EcsUsers
// @Summary 更新订阅用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ecsusers.EcsUsers true "更新订阅用户"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /eusr/updateEcsUsers [put]
func (eusrApi *EcsUsersApi) UpdateEcsUsers(c *gin.Context) {
	var eusr ecsusers.EcsUsers
	err := c.ShouldBindJSON(&eusr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	eusr.Username = ""
	eusr.Password = ""
	err = eusrService.UpdateEcsUsers(eusr)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEcsUsers 用id查询订阅用户
// @Tags EcsUsers
// @Summary 用id查询订阅用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ecsusers.EcsUsers true "用id查询订阅用户"
// @Success 200 {object} response.Response{data=ecsusers.EcsUsers,msg=string} "查询成功"
// @Router /eusr/findEcsUsers [get]
func (eusrApi *EcsUsersApi) FindEcsUsers(c *gin.Context) {
	ID := c.Query("ID")
	reeusr, err := eusrService.GetEcsUsers(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reeusr, c)
}

// GetEcsUsersList 分页获取订阅用户列表
// @Tags EcsUsers
// @Summary 分页获取订阅用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ecsusersReq.EcsUsersSearch true "分页获取订阅用户列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /eusr/getEcsUsersList [get]
func (eusrApi *EcsUsersApi) GetEcsUsersList(c *gin.Context) {
	var pageInfo ecsusersReq.EcsUsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := eusrService.GetEcsUsersInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetEcsUsersPublic 不需要鉴权的订阅用户接口
// @Tags EcsUsers
// @Summary 不需要鉴权的订阅用户接口
// @accept application/json
// @Produce application/json
// @Param data query ecsusersReq.EcsUsersSearch true "分页获取订阅用户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /eusr/getEcsUsersPublic [get]
func (eusrApi *EcsUsersApi) GetEcsUsersPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	eusrService.GetEcsUsersPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的订阅用户接口信息",
	}, "获取成功", c)
}

// AdminChangePassword 管理员修改用户密码
// @Tags EcsUsers
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Param data body ecsusersReq.AdminChangePasswordReq true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/adminChangePassword [PUT]
func (eusrApi *EcsUsersApi) AdminChangePassword(c *gin.Context) {
	var req ecsusersReq.AdminChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 请添加自己的业务逻辑
	err = eusrService.AdminChangePassword(req)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}

// SelfGetUserInfo 前端用户信息查询
// @Tags EcsUsers
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Param data query ecsusersReq.EcsUsersSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/getUserInfo [GET]
func (eusrApi *EcsUsersApi) SelfGetUserInfo(c *gin.Context) {
	id := utils.GetUserID(c)
	if id == 0 {
		response.FailWithMessage("失败", c)
		return
	}
	// 正常进行信息获取
	user, err := eusrService.GetUserInfo(id)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData(user, c)
}

// SelfModifyInfo 仅认证用户修改自己的信息
// @Tags EcsUsers
// @Summary 仅认证用户修改自己的信息
// @accept application/json
// @Produce application/json
// @Param data query ecsusersReq.EcsUsersSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/selfModifyInfo [POST]
func (eusrApi *EcsUsersApi) SelfModifyInfo(c *gin.Context) {
	var eusr ecsusers.EcsSubUsers
	err := c.ShouldBindJSON(&eusr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 通过cookie中的token获取uuid比对，如果不是同一用户，则不予修改更新
	currentUUID := utils.GetUserUuid(c)
	currentUUIDStr := currentUUID.String()
	if eusr.UUID.String() != currentUUIDStr {
		response.FailWithMessage("无权修改其他用户的信息", c)
		return
	}
	// 进行修改更新
	err = eusrService.SelfModifyInfo(&eusr)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetPublicRegisterStatus 获取是否公开注册
// @Tags EcsUsers
// @Summary 获取是否公开注册
// @accept application/json
// @Produce application/json
// @Param data query ecsusersReq.EcsUsersSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/getPublicRegisterStatus [GET]
func (eusrApi *EcsUsersApi) GetPublicRegisterStatus(c *gin.Context) {
	response.OkWithDetailed(ecsusers.EcsUserPublicRegisterStatus, "查询成功", c)
}

// ControlPublicRegister 启用关闭公开注册
// @Tags EcsUsers
// @Summary 启用关闭公开注册
// @accept application/json
// @Produce application/json
// @Param data query ecsusersReq.EcsUsersSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/controlPublicRegister [POST]
func (eusrApi *EcsUsersApi) ControlPublicRegister(c *gin.Context) {
	var control ecsusers.RegisterControl
	if err := c.ShouldBindJSON(&control); err != nil {
		response.FailWithMessage("修改失败", c)
	}
	if ecsusers.EcsUserPublicRegisterStatus != control.EnablePublicRegister {
		ecsusers.EcsUserPublicRegisterStatus = control.EnablePublicRegister
	}
	response.OkWithDetailed(ecsusers.EcsUserPublicRegisterStatus, "修改成功", c)
}

// GetTGRegisterStatus 获取是否公开TG注册
// @Tags EcsUsers
// @Summary 获取是否公开TG注册
// @accept application/json
// @Produce application/json
// @Param data query ecsusersReq.EcsUsersSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/getTGRegisterStatus [GET]
func (eusrApi *EcsUsersApi) GetTGRegisterStatus(c *gin.Context) {
	response.OkWithDetailed(model.EnableTGRegister, "查询成功", c)
}

// ControlTGRegister 启用关闭公开TG注册
// @Tags EcsUsers
// @Summary 启用关闭公开TG注册
// @accept application/json
// @Produce application/json
// @Param data query ecsusersReq.EcsUsersSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/controlTGRegister [POST]
func (eusrApi *EcsUsersApi) ControlTGRegister(c *gin.Context) {
	var control model.TGRegisterControl
	if err := c.ShouldBindJSON(&control); err != nil {
		response.FailWithMessage("修改失败", c)
	}
	if model.EnableTGRegister != control.EnableTGRegister {
		model.EnableTGRegister = control.EnableTGRegister
	}
	response.OkWithDetailed(model.EnableTGRegister, "修改成功", c)
}

// GetTGLoginStatus 获取是否启用TG登录
// @Tags EcsUsers
// @Summary 获取是否启用TG登录
// @accept application/json
// @Produce application/json
// @Param data query ecsusersReq.EcsUsersSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/getTGLoginStatus [GET]
func (eusrApi *EcsUsersApi) GetTGLoginStatus(c *gin.Context) {
	response.OkWithDetailed(model.EnableTGLogin, "查询成功", c)
}

// ControlTGLogin 启用关闭TG登录
// @Tags EcsUsers
// @Summary 启用关闭TG登录
// @accept application/json
// @Produce application/json
// @Param data query ecsusersReq.EcsUsersSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /eusr/controlTGLogin [POST]
func (eusrApi *EcsUsersApi) ControlTGLogin(c *gin.Context) {
	var control model.TGLoginControl
	if err := c.ShouldBindJSON(&control); err != nil {
		response.FailWithMessage("修改失败", c)
		return
	}
	if model.EnableTGLogin != control.EnableTGLogin {
		model.EnableTGLogin = control.EnableTGLogin
	}
	response.OkWithDetailed(model.EnableTGLogin, "修改成功", c)
}
