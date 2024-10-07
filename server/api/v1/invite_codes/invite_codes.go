package invite_codes

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/invite_codes"
    invite_codesReq "github.com/flipped-aurora/gin-vue-admin/server/model/invite_codes/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type InviteCodesApi struct {}



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
		response.FailWithMessage("创建失败:" + err.Error(), c)
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
		response.FailWithMessage("删除失败:" + err.Error(), c)
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
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
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
		response.FailWithMessage("更新失败:" + err.Error(), c)
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
		response.FailWithMessage("查询失败:" + err.Error(), c)
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
        response.FailWithMessage("获取失败:" + err.Error(), c)
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
