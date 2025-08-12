package privmsg

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/privmsg"
	privmsgReq "server/model/privmsg/request"
	"server/plugin/obopush/config"
)

type PusherConfigApi struct{}

// CreatePusherConfig 创建推送配置
// @Tags PusherConfig
// @Summary 创建推送配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body privmsg.PusherConfig true "创建推送配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /pc/createPusherConfig [post]
func (pcApi *PusherConfigApi) CreatePusherConfig(c *gin.Context) {
	var pc privmsg.PusherConfig
	err := c.ShouldBindJSON(&pc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = pcService.CreatePusherConfig(&pc)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePusherConfig 删除推送配置
// @Tags PusherConfig
// @Summary 删除推送配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body privmsg.PusherConfig true "删除推送配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /pc/deletePusherConfig [delete]
func (pcApi *PusherConfigApi) DeletePusherConfig(c *gin.Context) {
	ID := c.Query("ID")
	err := pcService.DeletePusherConfig(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePusherConfigByIds 批量删除推送配置
// @Tags PusherConfig
// @Summary 批量删除推送配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /pc/deletePusherConfigByIds [delete]
func (pcApi *PusherConfigApi) DeletePusherConfigByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := pcService.DeletePusherConfigByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePusherConfig 更新推送配置
// @Tags PusherConfig
// @Summary 更新推送配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body privmsg.PusherConfig true "更新推送配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /pc/updatePusherConfig [put]
func (pcApi *PusherConfigApi) UpdatePusherConfig(c *gin.Context) {
	var pc privmsg.PusherConfig
	err := c.ShouldBindJSON(&pc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = pcService.UpdatePusherConfig(pc)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPusherConfig 用id查询推送配置
// @Tags PusherConfig
// @Summary 用id查询推送配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query privmsg.PusherConfig true "用id查询推送配置"
// @Success 200 {object} response.Response{data=privmsg.PusherConfig,msg=string} "查询成功"
// @Router /pc/findPusherConfig [get]
func (pcApi *PusherConfigApi) FindPusherConfig(c *gin.Context) {
	ID := c.Query("ID")
	repc, err := pcService.GetPusherConfig(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(repc, c)
}

// GetPusherConfigList 分页获取推送配置列表
// @Tags PusherConfig
// @Summary 分页获取推送配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query privmsgReq.PusherConfigSearch true "分页获取推送配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /pc/getPusherConfigList [get]
func (pcApi *PusherConfigApi) GetPusherConfigList(c *gin.Context) {
	var pageInfo privmsgReq.PusherConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := pcService.GetPusherConfigInfoList(pageInfo)
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

// GetPusherConfigPublic 不需要鉴权的推送配置接口
// @Tags PusherConfig
// @Summary 不需要鉴权的推送配置接口
// @accept application/json
// @Produce application/json
// @Param data query privmsgReq.PusherConfigSearch true "分页获取推送配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pc/getPusherConfigPublic [get]
func (pcApi *PusherConfigApi) GetPusherConfigPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	pcService.GetPusherConfigPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的推送配置接口信息",
	}, "获取成功", c)
}

// GetPublicPushStatus 获取整体推送是否启用
// @Tags PusherConfig
// @Summary 获取整体推送是否启用
// @accept application/json
// @Produce application/json
// @Param data query privmsgReq.PusherConfigSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /pc/getPublicPushStatus [GET]
func (pcApi *PusherConfigApi) GetPublicPushStatus(c *gin.Context) {
	response.OkWithDetailed(config.Enabelobopush, "查询成功", c)
}

// ControlPublicPushStatus 修改整体推送是否启用
// @Tags PusherConfig
// @Summary 修改整体推送是否启用
// @accept application/json
// @Produce application/json
// @Param data query privmsgReq.PusherConfigSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /pc/controlPublicPushStatus [POST]
func (pcApi *PusherConfigApi) ControlPublicPushStatus(c *gin.Context) {
	var control config.OboPushControl
	if err := c.ShouldBindJSON(&control); err != nil {
		response.FailWithMessage("修改失败", c)
	}
	if config.Enabelobopush != control.EnablePublicOboPush {
		config.Enabelobopush = control.EnablePublicOboPush
	}
	response.OkWithDetailed(config.Enabelobopush, "修改成功", c)
}

// GetTelegramBotPushStatus 获取是否启用TG的BOT进行推送
// @Tags PusherConfig
// @Summary 获取是否启用TG的BOT进行推送
// @accept application/json
// @Produce application/json
// @Param data query privmsgReq.PusherConfigSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /pc/getTelegramBotPushStatus [GET]
func (pcApi *PusherConfigApi) GetTelegramBotPushStatus(c *gin.Context) {
	response.OkWithDetailed(config.EnabelTelegramBotPush, "查询成功", c)
}

// ControlTelegramBotPushStatus 修改是否启用TG的BOT进行推送
// @Tags PusherConfig
// @Summary 修改是否启用TG的BOT进行推送
// @accept application/json
// @Produce application/json
// @Param data query privmsgReq.PusherConfigSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /pc/controlTelegramBotPushStatus [POST]
func (pcApi *PusherConfigApi) ControlTelegramBotPushStatus(c *gin.Context) {
	var control config.TelegramBotPushControl
	if err := c.ShouldBindJSON(&control); err != nil {
		response.FailWithMessage("修改失败", c)
	}
	if config.EnabelTelegramBotPush != control.EnableTelegramBotPush {
		config.EnabelTelegramBotPush = control.EnableTelegramBotPush
	}
	response.OkWithDetailed(config.EnabelTelegramBotPush, "修改成功", c)
}

// GetEmailPushStatus 获取邮件推送状态
// @Tags PusherConfig
// @Summary 获取邮件推送状态
// @accept application/json
// @Produce application/json
// @Param data query privmsgReq.PusherConfigSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /pc/getEmailPushStatus [GET]
func (pcApi *PusherConfigApi) GetEmailPushStatus(c *gin.Context) {
	response.OkWithDetailed(config.EnableEmailPush, "查询成功", c)
}

// ControlEmailPushStatus 控制启用关闭邮件推送
// @Tags PusherConfig
// @Summary 控制启用关闭邮件推送
// @accept application/json
// @Produce application/json
// @Param data query privmsgReq.PusherConfigSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /pc/controlEmailPushStatus [POST]
func (pcApi *PusherConfigApi) ControlEmailPushStatus(c *gin.Context) {
	var control config.EmailPushControl
	if err := c.ShouldBindJSON(&control); err != nil {
		response.FailWithMessage("修改失败", c)
	}
	if config.EnableEmailPush != control.EnableEmailPush {
		config.EnableEmailPush = control.EnableEmailPush
	}
	response.OkWithDetailed(config.EnableEmailPush, "修改成功", c)
}
