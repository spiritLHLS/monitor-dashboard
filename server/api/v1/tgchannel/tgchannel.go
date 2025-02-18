package tgchannel

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/tgchannel"
    tgchannelReq "github.com/flipped-aurora/gin-vue-admin/server/model/tgchannel/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TgchannelApi struct {}



// CreateTgchannel 创建tgchannel表
// @Tags Tgchannel
// @Summary 创建tgchannel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tgchannel.Tgchannel true "创建tgchannel表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /tgc/createTgchannel [post]
func (tgcApi *TgchannelApi) CreateTgchannel(c *gin.Context) {
	var tgc tgchannel.Tgchannel
	err := c.ShouldBindJSON(&tgc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = tgcService.CreateTgchannel(&tgc)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteTgchannel 删除tgchannel表
// @Tags Tgchannel
// @Summary 删除tgchannel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tgchannel.Tgchannel true "删除tgchannel表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /tgc/deleteTgchannel [delete]
func (tgcApi *TgchannelApi) DeleteTgchannel(c *gin.Context) {
	ID := c.Query("ID")
	err := tgcService.DeleteTgchannel(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTgchannelByIds 批量删除tgchannel表
// @Tags Tgchannel
// @Summary 批量删除tgchannel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /tgc/deleteTgchannelByIds [delete]
func (tgcApi *TgchannelApi) DeleteTgchannelByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := tgcService.DeleteTgchannelByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTgchannel 更新tgchannel表
// @Tags Tgchannel
// @Summary 更新tgchannel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tgchannel.Tgchannel true "更新tgchannel表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /tgc/updateTgchannel [put]
func (tgcApi *TgchannelApi) UpdateTgchannel(c *gin.Context) {
	var tgc tgchannel.Tgchannel
	err := c.ShouldBindJSON(&tgc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = tgcService.UpdateTgchannel(tgc)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTgchannel 用id查询tgchannel表
// @Tags Tgchannel
// @Summary 用id查询tgchannel表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tgchannel.Tgchannel true "用id查询tgchannel表"
// @Success 200 {object} response.Response{data=tgchannel.Tgchannel,msg=string} "查询成功"
// @Router /tgc/findTgchannel [get]
func (tgcApi *TgchannelApi) FindTgchannel(c *gin.Context) {
	ID := c.Query("ID")
	retgc, err := tgcService.GetTgchannel(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(retgc, c)
}

// GetTgchannelList 分页获取tgchannel表列表
// @Tags Tgchannel
// @Summary 分页获取tgchannel表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tgchannelReq.TgchannelSearch true "分页获取tgchannel表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tgc/getTgchannelList [get]
func (tgcApi *TgchannelApi) GetTgchannelList(c *gin.Context) {
	var pageInfo tgchannelReq.TgchannelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := tgcService.GetTgchannelInfoList(pageInfo)
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

// GetTgchannelPublic 不需要鉴权的tgchannel表接口
// @Tags Tgchannel
// @Summary 不需要鉴权的tgchannel表接口
// @accept application/json
// @Produce application/json
// @Param data query tgchannelReq.TgchannelSearch true "分页获取tgchannel表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tgc/getTgchannelPublic [get]
func (tgcApi *TgchannelApi) GetTgchannelPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    tgcService.GetTgchannelPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的tgchannel表接口信息",
    }, "获取成功", c)
}
