package invite_codes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/invite_codes"
	invite_codesReq "server/model/invite_codes/request"
)

type InviteCodesApi struct{}

// CreateInviteCodes 创建邀请码
// @Tags InviteCodes
// @Summary 创建邀请码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body invite_codes.InviteCodes true "创建邀请码"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /invcode/createInviteCodes [post]
func (invcodeApi *InviteCodesApi) CreateInviteCodes(c *gin.Context) {
	var invcode invite_codes.InviteCodes
	err := c.ShouldBindJSON(&invcode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = invcodeService.CreateInviteCodes(&invcode)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteInviteCodes 删除邀请码
// @Tags InviteCodes
// @Summary 删除邀请码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body invite_codes.InviteCodes true "删除邀请码"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /invcode/deleteInviteCodes [delete]
func (invcodeApi *InviteCodesApi) DeleteInviteCodes(c *gin.Context) {
	ID := c.Query("ID")
	err := invcodeService.DeleteInviteCodes(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteInviteCodesByIds 批量删除邀请码
// @Tags InviteCodes
// @Summary 批量删除邀请码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /invcode/deleteInviteCodesByIds [delete]
func (invcodeApi *InviteCodesApi) DeleteInviteCodesByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := invcodeService.DeleteInviteCodesByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateInviteCodes 更新邀请码
// @Tags InviteCodes
// @Summary 更新邀请码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body invite_codes.InviteCodes true "更新邀请码"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /invcode/updateInviteCodes [put]
func (invcodeApi *InviteCodesApi) UpdateInviteCodes(c *gin.Context) {
	var invcode invite_codes.InviteCodes
	err := c.ShouldBindJSON(&invcode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = invcodeService.UpdateInviteCodes(invcode)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindInviteCodes 用id查询邀请码
// @Tags InviteCodes
// @Summary 用id查询邀请码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query invite_codes.InviteCodes true "用id查询邀请码"
// @Success 200 {object} response.Response{data=invite_codes.InviteCodes,msg=string} "查询成功"
// @Router /invcode/findInviteCodes [get]
func (invcodeApi *InviteCodesApi) FindInviteCodes(c *gin.Context) {
	ID := c.Query("ID")
	reinvcode, err := invcodeService.GetInviteCodes(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reinvcode, c)
}

// GetInviteCodesList 分页获取邀请码列表
// @Tags InviteCodes
// @Summary 分页获取邀请码列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query invite_codesReq.InviteCodesSearch true "分页获取邀请码列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /invcode/getInviteCodesList [get]
func (invcodeApi *InviteCodesApi) GetInviteCodesList(c *gin.Context) {
	var pageInfo invite_codesReq.InviteCodesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := invcodeService.GetInviteCodesInfoList(pageInfo)
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

// GetInviteCodesPublic 不需要鉴权的邀请码接口
// @Tags InviteCodes
// @Summary 不需要鉴权的邀请码接口
// @accept application/json
// @Produce application/json
// @Param data query invite_codesReq.InviteCodesSearch true "分页获取邀请码列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /invcode/getInviteCodesPublic [get]
func (invcodeApi *InviteCodesApi) GetInviteCodesPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	invcodeService.GetInviteCodesPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的邀请码接口信息",
	}, "获取成功", c)
}

// BatchBuildCodes 批量生成邀请码
// @Tags InviteCodes
// @Summary 批量生成邀请码
// @accept application/json
// @Produce application/json
// @Param data query invite_codesReq.InviteCodesSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /invcode/batchBuildCodes [POST]
func (invcodeApi *InviteCodesApi) BatchBuildCodes(c *gin.Context) {
	var invcode invite_codes.BatchBuildCodes
	err := c.ShouldBindJSON(&invcode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	content, err2 := invcodeService.BatchBuildCodes(*invcode.Count)
	if err2 != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err2))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithDetailed(gin.H{
		"Codes": content,
	}, "获取成功", c)
}

// BatchExportCodes 批量导出邀请码
// @Tags InviteCodes
// @Summary 批量导出邀请码
// @accept application/json
// @Produce application/json
// @Param data query invite_codesReq.InviteCodesSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /invcode/batchExportCodes [POST]
func (invcodeApi *InviteCodesApi) BatchExportCodes(c *gin.Context) {
	var invcodes invite_codes.BatchExportCodes
	err := c.ShouldBindJSON(&invcodes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	content, err2 := invcodeService.BatchExportCodes(invcodes.IDs)
	if err2 != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err2))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithDetailed(gin.H{
		"Codes": content,
	}, "获取成功", c)
}

// GetPublicInviteStatus 获取是否启用邀请码注册
// @Tags InviteCodes
// @Summary 获取是否启用邀请码注册
// @accept application/json
// @Produce application/json
// @Param data query invite_codesReq.InviteCodesSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /invcode/getPublicInviteStatus [GET]
func (invcodeApi *InviteCodesApi) GetPublicInviteStatus(c *gin.Context) {
	response.OkWithDetailed(invite_codes.PublicInviteCodesStatus, "查询成功", c)
}

// ControlPublicInvite 修改是否使用邀请码注册
// @Tags InviteCodes
// @Summary 修改是否使用邀请码注册
// @accept application/json
// @Produce application/json
// @Param data query invite_codesReq.InviteCodesSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /invcode/controlPublicInvite [POST]
func (invcodeApi *InviteCodesApi) ControlPublicInvite(c *gin.Context) {
	var control invite_codes.InviteControl
	if err := c.ShouldBindJSON(&control); err != nil {
		response.FailWithMessage("入参失败", c)
	}
	if invite_codes.PublicInviteCodesStatus != control.EnablePublicInvite {
		invite_codes.PublicInviteCodesStatus = control.EnablePublicInvite
		if err := invcodeService.ControlPublicInvite(control.EnablePublicInvite); err != nil {
			response.FailWithMessage(fmt.Sprintf("后台任务启动失败: %s", err.Error()), c)
		}
	}
	response.OkWithDetailed(invite_codes.PublicInviteCodesStatus, "修改成功", c)
}
