package privmsg

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/privmsg"
    privmsgReq "github.com/flipped-aurora/gin-vue-admin/server/model/privmsg/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PusherConfigApi struct {}



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
		response.FailWithMessage("创建失败:" + err.Error(), c)
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
		response.FailWithMessage("删除失败:" + err.Error(), c)
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
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
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
		response.FailWithMessage("更新失败:" + err.Error(), c)
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
		response.FailWithMessage("查询失败:" + err.Error(), c)
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
