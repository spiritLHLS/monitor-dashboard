package claweragent

import (
	
	"server/global"
    "server/model/common/response"
    "server/model/claweragent"
    claweragentReq "server/model/claweragent/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AgentIpRecordApi struct {}



// CreateAgentIpRecord 创建IP记录
// @Tags AgentIpRecord
// @Summary 创建IP记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body claweragent.AgentIpRecord true "创建IP记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /air/createAgentIpRecord [post]
func (airApi *AgentIpRecordApi) CreateAgentIpRecord(c *gin.Context) {
	var air claweragent.AgentIpRecord
	err := c.ShouldBindJSON(&air)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = airService.CreateAgentIpRecord(&air)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAgentIpRecord 删除IP记录
// @Tags AgentIpRecord
// @Summary 删除IP记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body claweragent.AgentIpRecord true "删除IP记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /air/deleteAgentIpRecord [delete]
func (airApi *AgentIpRecordApi) DeleteAgentIpRecord(c *gin.Context) {
	ID := c.Query("ID")
	err := airService.DeleteAgentIpRecord(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAgentIpRecordByIds 批量删除IP记录
// @Tags AgentIpRecord
// @Summary 批量删除IP记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /air/deleteAgentIpRecordByIds [delete]
func (airApi *AgentIpRecordApi) DeleteAgentIpRecordByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := airService.DeleteAgentIpRecordByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAgentIpRecord 更新IP记录
// @Tags AgentIpRecord
// @Summary 更新IP记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body claweragent.AgentIpRecord true "更新IP记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /air/updateAgentIpRecord [put]
func (airApi *AgentIpRecordApi) UpdateAgentIpRecord(c *gin.Context) {
	var air claweragent.AgentIpRecord
	err := c.ShouldBindJSON(&air)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = airService.UpdateAgentIpRecord(air)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAgentIpRecord 用id查询IP记录
// @Tags AgentIpRecord
// @Summary 用id查询IP记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询IP记录"
// @Success 200 {object} response.Response{data=claweragent.AgentIpRecord,msg=string} "查询成功"
// @Router /air/findAgentIpRecord [get]
func (airApi *AgentIpRecordApi) FindAgentIpRecord(c *gin.Context) {
	ID := c.Query("ID")
	reair, err := airService.GetAgentIpRecord(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reair, c)
}
// GetAgentIpRecordList 分页获取IP记录列表
// @Tags AgentIpRecord
// @Summary 分页获取IP记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query claweragentReq.AgentIpRecordSearch true "分页获取IP记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /air/getAgentIpRecordList [get]
func (airApi *AgentIpRecordApi) GetAgentIpRecordList(c *gin.Context) {
	var pageInfo claweragentReq.AgentIpRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := airService.GetAgentIpRecordInfoList(pageInfo)
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

// GetAgentIpRecordPublic 不需要鉴权的IP记录接口
// @Tags AgentIpRecord
// @Summary 不需要鉴权的IP记录接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /air/getAgentIpRecordPublic [get]
func (airApi *AgentIpRecordApi) GetAgentIpRecordPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    airService.GetAgentIpRecordPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的IP记录接口信息",
    }, "获取成功", c)
}
