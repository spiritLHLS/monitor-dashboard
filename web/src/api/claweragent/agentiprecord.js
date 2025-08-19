import service from '@/utils/request'
// @Tags AgentIpRecord
// @Summary 创建IP记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentIpRecord true "创建IP记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /air/createAgentIpRecord [post]
export const createAgentIpRecord = (data) => {
  return service({
    url: '/air/createAgentIpRecord',
    method: 'post',
    data
  })
}

// @Tags AgentIpRecord
// @Summary 删除IP记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentIpRecord true "删除IP记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /air/deleteAgentIpRecord [delete]
export const deleteAgentIpRecord = (params) => {
  return service({
    url: '/air/deleteAgentIpRecord',
    method: 'delete',
    params
  })
}

// @Tags AgentIpRecord
// @Summary 批量删除IP记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IP记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /air/deleteAgentIpRecord [delete]
export const deleteAgentIpRecordByIds = (params) => {
  return service({
    url: '/air/deleteAgentIpRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags AgentIpRecord
// @Summary 更新IP记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentIpRecord true "更新IP记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /air/updateAgentIpRecord [put]
export const updateAgentIpRecord = (data) => {
  return service({
    url: '/air/updateAgentIpRecord',
    method: 'put',
    data
  })
}

// @Tags AgentIpRecord
// @Summary 用id查询IP记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AgentIpRecord true "用id查询IP记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /air/findAgentIpRecord [get]
export const findAgentIpRecord = (params) => {
  return service({
    url: '/air/findAgentIpRecord',
    method: 'get',
    params
  })
}

// @Tags AgentIpRecord
// @Summary 分页获取IP记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IP记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /air/getAgentIpRecordList [get]
export const getAgentIpRecordList = (params) => {
  return service({
    url: '/air/getAgentIpRecordList',
    method: 'get',
    params
  })
}

// @Tags AgentIpRecord
// @Summary 不需要鉴权的IP记录接口
// @Accept application/json
// @Produce application/json
// @Param data query claweragentReq.AgentIpRecordSearch true "分页获取IP记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /air/getAgentIpRecordPublic [get]
export const getAgentIpRecordPublic = () => {
  return service({
    url: '/air/getAgentIpRecordPublic',
    method: 'get',
  })
}
